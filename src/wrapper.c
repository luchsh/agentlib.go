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

static void JNICALL
on_vm_init(jvmtiEnv *jvmti,
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

  jvmtiEventCallbacks callbacks;
  memset(&callbacks, 0, sizeof(callbacks));
  callbacks.VMInit = &on_vm_init;
  jvmti_res = (*jvmti)->SetEventCallbacks(jvmti, &callbacks, sizeof(callbacks));
  if (jvmti_res != JVMTI_ERROR_NONE) {
    printf("Failed to set jvmti callbacks\n");
    return JNI_ERR;
  }
  jvmti_res = (*jvmti)->SetEventNotificationMode(jvmti, JVMTI_ENABLE, JVMTI_EVENT_VM_INIT, NULL);
  if (jvmti_res != JVMTI_ERROR_NONE) {
    printf("Failed to set jvmti callbacks\n");
    return JNI_ERR;
  }
  return JNI_OK;
}
