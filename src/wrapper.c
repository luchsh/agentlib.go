//
// Copyright 2020 chuanshenglu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

#include "wrapper.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdarg.h>

// if should print debug output
static jint debug_output = JNI_TRUE;

#define DEBUG(code) do { \
  if (debug_output == JNI_TRUE) { \
    code; \
  }   \
} while(0)

#define ARRAY_LEN(arr)  (sizeof(arr) / sizeof(arr[0]))

// Just to lable that this call is to Go funcs
#define GO_CALL(code) do {\
  code; \
} while(0)

// Main loop of the forward thread to make JNI, JVMTI calls
// in a correct Java thread context
extern void MainForwardLoop();
// primary interface to transport JVMTI event to GO level
extern void OnJvmtiEvent(int event_id,
                         uintptr_t jvmti,
                         uintptr_t params_addr, // aruments are packed into uint64 array
                         int params_len);

// entry of the newly create thread
static void
go_bridge_thread_entry(jvmtiEnv* jvmti_env,
                       JNIEnv* jni_env,
                       void* arg) {
  DEBUG(printf("C: Bridge thread started\n"));

  // enter Go realm
  GO_CALL(MainForwardLoop());
}

// Allocate and initialize a new java.lang.Thread object
static jthread
allocate_thread_obj(JNIEnv* env) {
  jclass clazz = (*env)->FindClass(env, "Ljava/lang/Thread;");
  if ((*env)->ExceptionOccurred(env) != NULL) {
    printf("Cannot find class java.lang.Thread\n");
    (*env)->ExceptionClear(env);
    return NULL;
  }
  jmethodID constructor = (*env)->GetMethodID(env, clazz, "<init>", "()V");
  if ((*env)->ExceptionOccurred(env) != NULL) {
    printf("Cannot get default constructor of thread class\n");
    (*env)->ExceptionClear(env);
    return NULL;
  }
  jthread thread_obj = (*env)->NewObject(env, clazz, constructor);
  if ((*env)->ExceptionOccurred(env) != NULL) {
    printf("Cannot create a new thread object\n");
    (*env)->ExceptionClear(env);
    return NULL;
  }
  return thread_obj;
}

// extra internal initialization work
static void __internalOnVMInit(jvmtiEnv *jvmti,
            JNIEnv* jni,
            jthread thread) {
  DEBUG(printf("VMInit event triggered!\n"));
  // create a agent thread to forward commands from Go runtime to Java runtime
  // because neither JNIEnv or jvmtiEnv exists on Go's 'm'
  jthread thread_obj = allocate_thread_obj(jni);
  jvmtiError jvmti_res = (*jvmti)->RunAgentThread(jvmti,
                                       thread_obj,
                                       &go_bridge_thread_entry,
                                       NULL,
                                       JVMTI_THREAD_NORM_PRIORITY);
  if (jvmti_res != JVMTI_ERROR_NONE) {
    printf("Failed to start agent thread\n");
  }
}

// global JVMTI callback table
static jvmtiEventCallbacks *_callbacks =NULL;

// number of args (include jvmti and jni) for each event callback
static size_t nof_args_for_event[] = {
    3, /* JVMTI_EVENT_VM_INIT = 50, */
    3, /* JVMTI_EVENT_VM_DEATH = 51, */
    3, /* JVMTI_EVENT_THREAD_START = 52, */
    3, /* JVMTI_EVENT_THREAD_END = 53, */
    10, /* JVMTI_EVENT_CLASS_FILE_LOAD_HOOK = 54, */
    4, /* JVMTI_EVENT_CLASS_LOAD = 55, */
    4, /* JVMTI_EVENT_CLASS_PREPARE = 56, */
    2, /* JVMTI_EVENT_VM_START = 57, */
    8, /* JVMTI_EVENT_EXCEPTION = 58, */
    6, /* JVMTI_EVENT_EXCEPTION_CATCH = 59, */
    5, /* JVMTI_EVENT_SINGLE_STEP = 60, */
    5, /* JVMTI_EVENT_FRAME_POP = 61, */
    5, /* JVMTI_EVENT_BREAKPOINT = 62, */
    8, /* JVMTI_EVENT_FIELD_ACCESS = 63, */
    10, /* JVMTI_EVENT_FIELD_MODIFICATION = 64, */
    4, /* JVMTI_EVENT_METHOD_ENTRY = 65, */
    6, /* JVMTI_EVENT_METHOD_EXIT = 66, */
    6, /* JVMTI_EVENT_NATIVE_METHOD_BIND = 67, */
    7, /* JVMTI_EVENT_COMPILED_METHOD_LOAD = 68, */
    3, /* JVMTI_EVENT_COMPILED_METHOD_UNLOAD = 69, */
    4, /* JVMTI_EVENT_DYNAMIC_CODE_GENERATED = 70, */
    1, /* JVMTI_EVENT_DATA_DUMP_REQUEST = 71, */
    5, /* JVMTI_EVENT_MONITOR_WAIT = 73, */
    5, /* JVMTI_EVENT_MONITOR_WAITED = 74, */
    4, /* JVMTI_EVENT_MONITOR_CONTENDED_ENTER = 75, */
    4, /* JVMTI_EVENT_MONITOR_CONTENDED_ENTERED = 76, */
    5, /* JVMTI_EVENT_RESOURCE_EXHAUSTED = 80, */
    1, /* JVMTI_EVENT_GARBAGE_COLLECTION_START = 81, */
    1, /* JVMTI_EVENT_GARBAGE_COLLECTION_FINISH = 82, */
    2, /* JVMTI_EVENT_OBJECT_FREE = 83, */
    6, /* JVMTI_EVENT_VM_OBJECT_ALLOC = 84, */
    // Below are fake
    0, /* JVMTI_EVENT_AGENT_UNLOAD = 85, */
};

// Handler generator MACRO.
// Can be simplified using C++ template function
//   `template<int EventId> void onJvmtiEvent(jvmtiEnv* jvmti, ...);`
// but it may involve SWIG, a little complex, do it later.
#define GEN_JVMTI_EVENT_HANDLER(EVENT)                                    \
  static void onJvmtiEvent_ ## EVENT(jvmtiEnv* jvmti, ...) {              \
    DEBUG(printf("C: event " #EVENT " triggerred\n"));                    \
    size_t idx = EVENT - JVMTI_MIN_EVENT_TYPE_VAL;                        \
    size_t nof_args = nof_args_for_event[idx] - 1 /* rm jvmti */;         \
    uintptr_t args[nof_args];                                             \
    va_list ap;                                                           \
    va_start(ap, jvmti);                                                  \
    for (int i = 0; i < nof_args; ++i) {                                  \
      args[i] = va_arg(ap, uintptr_t);                                    \
    }                                                                     \
    va_end(ap);                                                           \
    if (EVENT == JVMTI_EVENT_VM_INIT) {                                   \
      __internalOnVMInit(jvmti, (JNIEnv*)(args[0]), (jthread)(args[1]));  \
    }                                                                     \
    OnJvmtiEvent(EVENT, jvmti, args, ARRAY_LEN(args));                    \
  }

// Generate event handler methods
FOR_EACH_JVMTI_EVENT(GEN_JVMTI_EVENT_HANDLER)

// a table to quickly map event index to corresponding handler method
#define NOF_JVMTI_EVENTS  (JVMTI_MAX_EVENT_TYPE_VAL - JVMTI_MIN_EVENT_TYPE_VAL + 1)
#define EXPAND_JVMTI_EVENT_HANDLER(EVENT) (&onJvmtiEvent_##EVENT),
static void* builtin_handler_table[NOF_JVMTI_EVENTS] = {
  FOR_EACH_JVMTI_EVENT(EXPAND_JVMTI_EVENT_HANDLER)
  /* fake events has separated entry */
  // TODO: wrong index
  &Agent_OnUnload, // JVMTI_EVENT_AGENT_UNLOAD
};

// a table containing all jvmti event names
#define JVMTI_EVENT_NAME(EVENT)  (#EVENT),
static const char* jvmti_event_names[NOF_JVMTI_EVENTS] = {
  FOR_EACH_JVMTI_EVENT(JVMTI_EVENT_NAME)
};

#define NOF_FAKE_JVMTI_EVENTS   (JVMTI_MAX_FAKE_EVENT_TYPE_VAL - JVMTI_MIN_EVENT_TYPE_VAL + 1)
static const char* jvmti_fake_event_names[NOF_FAKE_JVMTI_EVENTS] = {
  FOR_EACH_FAKE_JVMTI_EVENT(JVMTI_EVENT_NAME)
};

static const char* get_jvmti_event_name(int event_id) {
  if (event_id <= JVMTI_MAX_EVENT_TYPE_VAL && event_id >= JVMTI_MIN_EVENT_TYPE_VAL) {
    int idx = event_id - JVMTI_MIN_EVENT_TYPE_VAL;
    return jvmti_event_names[idx];
  } else if (event_id <= JVMTI_MAX_FAKE_EVENT_TYPE_VAL && event_id >= JVMTI_MIN_FAKE_EVENT_TYPE_VAL) {
    int idx = event_id - JVMTI_MIN_FAKE_EVENT_TYPE_VAL;
    return jvmti_fake_event_names[idx];
  }
  DEBUG(printf("C: no name for event id %d\n", event_id));
  return NULL;
}

// Setup JVMTI callbacks
// NOTE: this method highly depends on the sequence and layout defined in jvmti.h
static void linkLocalJvmtiCallback(int event_id) {
  void** ftable = (void**)&(_callbacks->VMInit); // VMInit is the first entry of the struct
  int idx = event_id - JVMTI_MIN_EVENT_TYPE_VAL;
  ftable[idx] = builtin_handler_table[idx];
}

void EnableJvmtiCallback(void* p, int event_id) {
  jvmtiEnv* jvmti = (jvmtiEnv*)p;
  //printf("%d-%d-%d", JVMTI_MIN_EVENT_TYPE_VAL, event_id, JVMTI_MAX_EVENT_TYPE_VAL);
  if (event_id >= JVMTI_MIN_EVENT_TYPE_VAL && event_id <= JVMTI_MAX_EVENT_TYPE_VAL) {
    if (jvmti == NULL) {
      printf("NULL jvmtiEnv pointer\n");
      return;
    }
 
    linkLocalJvmtiCallback(event_id);

    jvmtiError jvmti_res = (*jvmti)->SetEventCallbacks(jvmti, _callbacks, sizeof(jvmtiEventCallbacks));
    if (jvmti_res != JVMTI_ERROR_NONE) {
      printf("Failed to set jvmti callbacks: %s\n", get_jvmti_event_name(event_id));
      return;
    }
    jvmti_res = (*jvmti)->SetEventNotificationMode(jvmti, JVMTI_ENABLE, event_id, NULL);
    if (jvmti_res != JVMTI_ERROR_NONE) {
      printf("Failed to set jvmti callback notification mode: %s\n", get_jvmti_event_name(event_id));
      return;
    }
    DEBUG(printf("C: enabled event id=%d, name=%s\n", event_id, get_jvmti_event_name(event_id)));
  } else if (event_id >= JVMTI_MIN_FAKE_EVENT_TYPE_VAL && event_id <= JVMTI_MAX_FAKE_EVENT_TYPE_VAL) {
    // nothing to do
    DEBUG(printf("C: enabled event id=%d, name=%s\n", event_id, get_jvmti_event_name(event_id)));
  } else {
    DEBUG(printf("C: invalid event id=%d, name=%s\n", event_id, get_jvmti_event_name(event_id)));
  }
}

//export OnAgentLoad
#ifdef __cplusplus
extern "C" {
#endif
extern void OnAgentLoad(uintptr_t, uintptr_t, const char*);
extern void OnAgentUnload();
#ifdef __cplusplus
}
#endif

// JVMTI start-up point
jint Agent_OnLoad(JavaVM* javaVM, char* options, void* reserved) {
  DEBUG(printf("agent.go loaded\n"));

  jvmtiEnv* jvmti = NULL;
  jvmtiError jvmti_res = (*javaVM)->GetEnv(javaVM, (void**)&jvmti, JVMTI_VERSION_1_1);
  if (jvmti_res != JVMTI_ERROR_NONE || jvmti == NULL) {
    printf("Failed to get jvmtiEnv from JavaVM");
    return JNI_ERR;
  }

  if (_callbacks == NULL) {
    _callbacks = (jvmtiEventCallbacks*)malloc(sizeof(jvmtiEventCallbacks));
    memset(_callbacks, 0, sizeof(jvmtiEventCallbacks));
  }

  GO_CALL(OnAgentLoad((uintptr_t)javaVM, jvmti, options));

  return JNI_OK;
}

// destroy resources;
void Agent_OnUnload(JavaVM* javaVM) {
  GO_CALL(OnAgentUnload());

  jvmtiEventCallbacks* cb = _callbacks;
  _callbacks = NULL;
  free(cb);
}
