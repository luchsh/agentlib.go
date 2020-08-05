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

//
// The C code here provides a thin bridge between Go runtime
// and the JNI, JVMTI interface.
//
// We need this wrapper layer for two reasons:
// 1. To work with cgo's inability of calling function pointer fields of C struct.
//    So far (by Go 1.14) this still cannot be done easily, might be fixed in the future.
//    We may consider removing this layer then.
// 2. To create JNI compatible thread state
//    JNIEnv is threadlocal object, and expected to be accessed only from 'JavaThread',
//    but for a shared library writen in Go, which currently ships a full copy Go runtime by default,
//    its carrier threads, AKA 'm', are not 'JavaThread'.
//

#ifndef __WRAPPER_H__
#define __WRAPPER_H__

#include <jni.h>
#include <jvmti.h>

#ifdef __cplusplus
extern "C" {
#endif

// Macro to generate operation on all the jvmti events
#define FOR_EACH_JVMTI_EVENT(f) 		     \
	f(JVMTI_EVENT_VM_INIT)                   \
	f(JVMTI_EVENT_VM_DEATH)                  \
	f(JVMTI_EVENT_THREAD_START)              \
	f(JVMTI_EVENT_THREAD_END)                \
	f(JVMTI_EVENT_CLASS_FILE_LOAD_HOOK)      \
	f(JVMTI_EVENT_CLASS_LOAD)                \
	f(JVMTI_EVENT_CLASS_PREPARE)             \
	f(JVMTI_EVENT_VM_START)                  \
	f(JVMTI_EVENT_EXCEPTION)                 \
	f(JVMTI_EVENT_EXCEPTION_CATCH)           \
	f(JVMTI_EVENT_SINGLE_STEP)               \
	f(JVMTI_EVENT_FRAME_POP)                 \
	f(JVMTI_EVENT_BREAKPOINT)                \
	f(JVMTI_EVENT_FIELD_ACCESS)              \
	f(JVMTI_EVENT_FIELD_MODIFICATION)        \
	f(JVMTI_EVENT_METHOD_ENTRY)              \
	f(JVMTI_EVENT_METHOD_EXIT)               \
	f(JVMTI_EVENT_NATIVE_METHOD_BIND)        \
	f(JVMTI_EVENT_COMPILED_METHOD_LOAD)      \
	f(JVMTI_EVENT_COMPILED_METHOD_UNLOAD)    \
	f(JVMTI_EVENT_DYNAMIC_CODE_GENERATED)    \
	f(JVMTI_EVENT_DATA_DUMP_REQUEST)         \
	f(JVMTI_EVENT_MONITOR_WAIT)              \
	f(JVMTI_EVENT_MONITOR_WAITED)            \
	f(JVMTI_EVENT_MONITOR_CONTENDED_ENTER)   \
	f(JVMTI_EVENT_MONITOR_CONTENDED_ENTERED) \
	f(JVMTI_EVENT_RESOURCE_EXHAUSTED)        \
	f(JVMTI_EVENT_GARBAGE_COLLECTION_START)  \
	f(JVMTI_EVENT_GARBAGE_COLLECTION_FINISH) \
	f(JVMTI_EVENT_OBJECT_FREE)               \
	f(JVMTI_EVENT_VM_OBJECT_ALLOC)           \

// Fake event definition
enum {
  JVMTI_MIN_FAKE_EVENT_TYPE_VAL   = 100,
  JVMTI_EVENT_AGENT_UNLOAD        = JVMTI_MIN_FAKE_EVENT_TYPE_VAL,
  JVMTI_MAX_FAKE_EVENT_TYPE_VAL   = JVMTI_EVENT_AGENT_UNLOAD,
};

#define FOR_EACH_FAKE_JVMTI_EVENT(f)    \
  f(JVMTI_EVENT_AGENT_UNLOAD)           \

// Enable a JVMTI event notification
// @arg p         Pointer to jvmtiEnv*, using uintptr_t to make it easier to be called from Go
// @arg event_id  ID of the JVMTI event to be enabled
void EnableJvmtiCallback(void* p, int event_id);

#ifdef __cplusplus
}
#endif

#endif // #ifndef __WRAPPER_H__