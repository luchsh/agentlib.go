//
// Copyright 2020 Lu Chuan Sheng
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

// if should print debug output
static jint debug_output = JNI_TRUE;

static JavaVM* java_vm;

#define DEBUG(code) do { \
  if (debug_output == JNI_TRUE) { \
    code; \
  }   \
} while(0)

// Just to lable that this call is to Go funcs
#define GO_CALL(code) do {\
  code; \
} while(0)

// exported CGO functions writen in Go lang
extern void MainForwardLoop();
extern void OnJvmtiVmInit(uintptr_t, uintptr_t, uintptr_t);

// entry of the newly create thread
static void
go_bridge_thread_entry(jvmtiEnv* jvmti_env,
                       JNIEnv* jni_env,
                       void* arg) {
  DEBUG(printf("Bridge thread started\n"));

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

// global JVMTI callback table
static jvmtiEventCallbacks *_callbacks =NULL;

// Here we define C wrappers of all JVMTI callbacks, and if corresponding hook
// is enabled in Go end, we will invoke.
// The calling convention is currently purely based on exported func name
static void on_Breakpoint(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jmethodID method,
            jlocation location) { }
static void on_SingleStep(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jmethodID method,
            jlocation location) { }
static void on_FieldAccess(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jmethodID method,
            jlocation location,
            jclass field_klass,
            jobject object,
            jfieldID field) { }
static void on_FieldModification(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jmethodID method,
            jlocation location,
            jclass field_klass,
            jobject object,
            jfieldID field,
            char signature_type,
            jvalue new_value) {}
static void on_FramePop(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jmethodID method,
            jboolean was_popped_by_exception) {}
static void on_MethodEntry(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jmethodID method) {}
static void on_MethodExit(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jmethodID method,
            jboolean was_popped_by_exception,
            jvalue return_value) {}
static void on_NativeMethodBind(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jmethodID method,
            void* address,
            void** new_address_ptr) { }
static void on_Exception(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jmethodID method,
            jlocation location,
            jobject exception,
            jmethodID catch_method,
            jlocation catch_location) { }
static void on_ExceptionCatch(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jmethodID method,
            jlocation location,
            jobject exception) { }
static void on_ThreadStart(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread) { }
static void on_ThreadEnd(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread) { }
static void on_ClassLoad(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jclass klass) { }
static void on_ClassPrepare(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jclass klass) { }
static void on_ClassFileLoadHook(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jclass class_being_redefined,
            jobject loader,
            const char* name,
            jobject protection_domain,
            jint class_data_len,
            const unsigned char* class_data,
            jint* new_class_data_len,
            unsigned char** new_class_data) {}
static void on_VMStart(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env) { }
static void on_VMInit(jvmtiEnv *jvmti,
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
  OnJvmtiVmInit((uintptr_t)jvmti, (uintptr_t)jni, (uintptr_t)thread);
}

static void on_VMDeath(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env) { }
static void on_CompiledMethodLoad(jvmtiEnv *jvmti_env,
            jmethodID method,
            jint code_size,
            const void* code_addr,
            jint map_length,
            const jvmtiAddrLocationMap* map,
            const void* compile_info) { }
static void on_CompiledMethodUnload(jvmtiEnv *jvmti_env,
            jmethodID method,
            const void* code_addr) { }
static void on_DynamicCodeGenerated(jvmtiEnv *jvmti_env,
            const char* name,
            const void* address,
            jint length) { }
static void on_DataDumpRequest(jvmtiEnv *jvmti_env) { }
static void on_MonitorContendedEnter(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jobject object) { }
static void on_MonitorContendedEntered(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jobject object) { }
static void on_MonitorWait(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jobject object,
            jlong timeout) {}
static void on_MonitorWaited(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jobject object,
            jboolean timed_out) {}
static void on_ResourceExhausted(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jint flags,
            const void* reserved,
            const char* description) {}
static void on_VMObjectAlloc(jvmtiEnv *jvmti_env,
            JNIEnv* jni_env,
            jthread thread,
            jobject object,
            jclass object_klass,
            jlong size) { }
static void on_ObjectFree(jvmtiEnv *jvmti_env,
            jlong tag) {}
static void on_GarbageCollectionStart(jvmtiEnv *jvmti_env) {}
static void on_GarbageCollectionFinish(jvmtiEnv *jvmti_env) {}

// Setup JVMTI callbacks
static void linkLocalJvmtiCallback(int event_id) {
  // link event id
  switch (event_id) {
	  case JVMTI_EVENT_VM_INIT:
    _callbacks->VMInit = &on_VMInit;
    break;
	case JVMTI_EVENT_VM_DEATH:
    _callbacks->VMDeath = &on_VMDeath;
    break;
	case JVMTI_EVENT_THREAD_START:
    _callbacks->ThreadEnd = &on_ThreadStart;
    break;
	case JVMTI_EVENT_THREAD_END:
    _callbacks->ThreadEnd = &on_ThreadEnd;
    break;
	case JVMTI_EVENT_CLASS_FILE_LOAD_HOOK:
    _callbacks->ClassFileLoadHook = &on_ClassFileLoadHook;
    break;
	case JVMTI_EVENT_CLASS_LOAD:
    _callbacks->ClassLoad = &on_ClassLoad;
    break;
	case JVMTI_EVENT_CLASS_PREPARE:
    _callbacks->ClassPrepare = &on_ClassPrepare;
    break;
	case JVMTI_EVENT_VM_START:
    _callbacks->VMStart = &on_VMStart;
    break;
	case JVMTI_EVENT_EXCEPTION:
    _callbacks->Exception = &on_Exception;
    break;
	case JVMTI_EVENT_EXCEPTION_CATCH:
    _callbacks->ExceptionCatch = &on_ExceptionCatch;
    break;
	case JVMTI_EVENT_SINGLE_STEP:
    _callbacks->SingleStep = &on_SingleStep;
    break;
	case JVMTI_EVENT_FRAME_POP:
    _callbacks->FramePop = &on_FramePop;
    break;
	case JVMTI_EVENT_BREAKPOINT:
    _callbacks->Breakpoint = &on_Breakpoint;
    break;
	case JVMTI_EVENT_FIELD_ACCESS:
    _callbacks->FieldAccess = &on_FieldAccess;
    break;
	case JVMTI_EVENT_FIELD_MODIFICATION:
    _callbacks->FieldModification = &on_FieldModification;
    break;
	case JVMTI_EVENT_METHOD_ENTRY:
    _callbacks->MethodEntry = &on_MethodEntry;
    break;
	case JVMTI_EVENT_METHOD_EXIT:
    _callbacks->MethodExit = &on_MethodExit;
    break;
	case JVMTI_EVENT_NATIVE_METHOD_BIND:
    _callbacks->NativeMethodBind = &on_NativeMethodBind;
    break;
	case JVMTI_EVENT_COMPILED_METHOD_LOAD:
    _callbacks->CompiledMethodLoad = &on_CompiledMethodLoad;
    break;
	case JVMTI_EVENT_COMPILED_METHOD_UNLOAD:
    _callbacks->CompiledMethodUnload = &on_CompiledMethodUnload;
    break;
	case JVMTI_EVENT_DYNAMIC_CODE_GENERATED:
    _callbacks->DynamicCodeGenerated = &on_DynamicCodeGenerated;
    break;
	case JVMTI_EVENT_DATA_DUMP_REQUEST:
    _callbacks->DataDumpRequest = &on_DataDumpRequest;
    break;
	case JVMTI_EVENT_MONITOR_WAIT:
    _callbacks->MonitorWait = &on_MonitorWait;
    break;
	case JVMTI_EVENT_MONITOR_WAITED:
    _callbacks->MonitorWaited = &on_MonitorWaited;
    break;
	case JVMTI_EVENT_MONITOR_CONTENDED_ENTER:
    _callbacks->MonitorContendedEnter = &on_MonitorContendedEnter;
    break;
	case JVMTI_EVENT_MONITOR_CONTENDED_ENTERED:
    _callbacks->MonitorContendedEntered = &on_MonitorContendedEntered;
    break;
	case JVMTI_EVENT_RESOURCE_EXHAUSTED:
    _callbacks->ResourceExhausted = &on_ResourceExhausted;
    break;
	case JVMTI_EVENT_GARBAGE_COLLECTION_START:
    _callbacks->GarbageCollectionStart = &on_GarbageCollectionStart;
    break;
	case JVMTI_EVENT_GARBAGE_COLLECTION_FINISH:
    _callbacks->GarbageCollectionFinish = &on_GarbageCollectionFinish;
    break;
	case JVMTI_EVENT_OBJECT_FREE:
    _callbacks->ObjectFree = &on_ObjectFree;
    break;
	case JVMTI_EVENT_VM_OBJECT_ALLOC:
    _callbacks->VMObjectAlloc = &on_VMObjectAlloc;
    break;
  default:
    printf("Should not reach here, invalid event!");
  }
}

void EnableJvmtiCallback(jvmtiEnv* jvmti, int event_id) {
  if (event_id < JVMTI_MIN_EVENT_TYPE_VAL || event_id > JVMTI_MAX_EVENT_TYPE_VAL) {
    printf("Invalid jvmti event to enable: %d\n", event_id);
    return;
  }
  if (jvmti == NULL) {
    printf("NULL jvmtiEnv pointer\n");
    return;
  }

  linkLocalJvmtiCallback(event_id);

  jvmtiError jvmti_res = (*jvmti)->SetEventCallbacks(jvmti, _callbacks, sizeof(jvmtiEventCallbacks));
  if (jvmti_res != JVMTI_ERROR_NONE) {
    printf("Failed to set jvmti callbacks\n");
    return;
  }
  jvmti_res = (*jvmti)->SetEventNotificationMode(jvmti, JVMTI_ENABLE, event_id, NULL);
  if (jvmti_res != JVMTI_ERROR_NONE) {
    printf("Failed to set jvmti callbacks\n");
    return;
  }
}

// JVMTI start-up point
jint Agent_OnLoad(JavaVM* javaVM, char* options, void* reserved) {
  DEBUG(printf("agent.go loaded\n"));
  java_vm = javaVM;
  if (java_vm == NULL) {
    printf("Cannot find a valid Java VM");
    return JNI_ERR;
  }

  GO_CALL(OnAgentLoad((void*)javaVM, options));

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

  EnableJvmtiCallback(jvmti, JVMTI_EVENT_VM_INIT);

  return JNI_OK;
}

// destroy resources;
void Agent_OnUnload(JavaVM* javaVM) {
  jvmtiEnv* jvmti = NULL;
  jvmtiError jvmti_res = (*javaVM)->GetEnv(javaVM, (void**)&jvmti, JVMTI_VERSION_1_1);
  if (jvmti_res != JVMTI_ERROR_NONE || jvmti == NULL) {
    printf("Failed to get jvmtiEnv from JavaVM");
    return;
  }
  // clear all events
  memset(_callbacks, 0, sizeof(_callbacks));
  jvmti_res = (*jvmti)->SetEventCallbacks(jvmti, _callbacks, sizeof(jvmtiEventCallbacks));
  if (jvmti_res != JVMTI_ERROR_NONE) {
    printf("Failed to set jvmti callbacks\n");
    return;
  }
  for (int e = JVMTI_MIN_EVENT_TYPE_VAL; e <= JVMTI_MAX_EVENT_TYPE_VAL; ++e) {
    jvmti_res = (*jvmti)->SetEventNotificationMode(jvmti, JVMTI_DISABLE, e, NULL);
    if (jvmti_res != JVMTI_ERROR_NONE) {
      printf("Failed to set jvmti callbacks\n");
      return;
    }
  }

  jvmtiEventCallbacks* cb = _callbacks;
  _callbacks = NULL;
  free(cb);
}