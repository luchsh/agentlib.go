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
package main

//#include "wrapper.h"
import "C"

import (
	"fmt"
	"unsafe"
)

// JVMTI event ID definitions
const (
	JVMTI_MIN_EVENT_TYPE_VAL              = 50
	JVMTI_EVENT_VM_INIT                   = 50
	JVMTI_EVENT_VM_DEATH                  = 51
	JVMTI_EVENT_THREAD_START              = 52
	JVMTI_EVENT_THREAD_END                = 53
	JVMTI_EVENT_CLASS_FILE_LOAD_HOOK      = 54
	JVMTI_EVENT_CLASS_LOAD                = 55
	JVMTI_EVENT_CLASS_PREPARE             = 56
	JVMTI_EVENT_VM_START                  = 57
	JVMTI_EVENT_EXCEPTION                 = 58
	JVMTI_EVENT_EXCEPTION_CATCH           = 59
	JVMTI_EVENT_SINGLE_STEP               = 60
	JVMTI_EVENT_FRAME_POP                 = 61
	JVMTI_EVENT_BREAKPOINT                = 62
	JVMTI_EVENT_FIELD_ACCESS              = 63
	JVMTI_EVENT_FIELD_MODIFICATION        = 64
	JVMTI_EVENT_METHOD_ENTRY              = 65
	JVMTI_EVENT_METHOD_EXIT               = 66
	JVMTI_EVENT_NATIVE_METHOD_BIND        = 67
	JVMTI_EVENT_COMPILED_METHOD_LOAD      = 68
	JVMTI_EVENT_COMPILED_METHOD_UNLOAD    = 69
	JVMTI_EVENT_DYNAMIC_CODE_GENERATED    = 70
	JVMTI_EVENT_DATA_DUMP_REQUEST         = 71
	JVMTI_EVENT_MONITOR_WAIT              = 73
	JVMTI_EVENT_MONITOR_WAITED            = 74
	JVMTI_EVENT_MONITOR_CONTENDED_ENTER   = 75
	JVMTI_EVENT_MONITOR_CONTENDED_ENTERED = 76
	JVMTI_EVENT_RESOURCE_EXHAUSTED        = 80
	JVMTI_EVENT_GARBAGE_COLLECTION_START  = 81
	JVMTI_EVENT_GARBAGE_COLLECTION_FINISH = 82
	JVMTI_EVENT_OBJECT_FREE               = 83
	JVMTI_EVENT_VM_OBJECT_ALLOC           = 84
	JVMTI_MAX_EVENT_TYPE_VAL              = 84
)

// JvmtiCallbacks holds the JVMTI event callbacks
type JvmtiCallbacks struct {
	/*
	   static void on_Breakpoint(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env,
	   	jthread thread,
	   	jmethodID method,
	   	jlocation location) { }
	*/
	onJvmtiBreakpoint func(jvmti, jni, thread, method, loc uintptr)

	/*
	   static void on_SingleStep(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env,
	   	jthread thread,
	   	jmethodID method,
	   	jlocation location) { }
	*/
	onJvmtiSingleStep func(jvmti, jni, thread, method, loc uintptr)

	/*
	   static void on_FieldAccess(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env,
	   	jthread thread,
	   	jmethodID method,
	   	jlocation location,
	   	jclass field_klass,
	   	jobject object,
	   	jfieldID field) { }
	*/
	onJvmtiFieldAccess func(jvmti, jni, method, loc, fclazz, obj, field uintptr)

	/*
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
	*/
	onJvmtiFieldModification func(jvmti, jni, thread, method, loc, fklazz, obj, field, sig, newval uintptr)

	/*
	   static void on_FramePop(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env,
	   	jthread thread,
	   	jmethodID method,
	   	jboolean was_popped_by_exception) {}
	*/
	onJvmtiFramePop func(jvmti, jni, thread, method uintptr, by_excep bool)

	/*
	   static void on_MethodEntry(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env,
	   	jthread thread,
	   	jmethodID method) {}
	*/
	onJvmtiMethodEntry func(jvmi, jni, thread, method uintptr)

	/*
	   static void on_MethodExit(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env,
	   	jthread thread,
	   	jmethodID method,
	   	jboolean was_popped_by_exception,
	   	jvalue return_value) {}
	*/
	onJvmtiMethodExit func(jvmti, jni, thread, method uintptr, by_excep bool, ret_val uintptr)

	/*
	   static void on_NativeMethodBind(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env,
	   	jthread thread,
	   	jmethodID method,
	   	void* address,
	   	void** new_address_ptr) { }
	*/
	onJvmtiNativeMethodBind func(jvmti, jni, thread, method, addr, new_addr_ptr uintptr)

	/*
	   static void on_Exception(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env,
	   	jthread thread,
	   	jmethodID method,
	   	jlocation location,
	   	jobject exception,
	   	jmethodID catch_method,
	   	jlocation catch_location) { }
	*/
	onJvmtiException func(jvmti, jni, thread, method, location, exception, catch_m, catch_l uintptr)

	/*
	   static void on_ExceptionCatch(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env,
	   	jthread thread,
	   	jmethodID method,
	   	jlocation location,
	   	jobject exception) { }
	*/
	onJvmtiExceptionCatch func(jvmti, jni, thread, method, location, exception uintptr)

	/*
	   static void on_ThreadStart(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env,
	   	jthread thread) { }
	*/
	onJvmtiThreadStart func(jvmti JvmtiEnv, jni JniEnv, thread uintptr)

	/*
	   static void on_ThreadEnd(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env,
	   	jthread thread) { }
	*/
	onJvmtiThreadEnd func(jvmti, jni, thread uintptr)

	/*
	   static void on_ClassLoad(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env,
	   	jthread thread,
	   	jclass klass) { }
	*/
	OnJvmtiClassLoad func(jvmti, jni, thread, klass uintptr)

	/*
	   static void on_ClassPrepare(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env,
	   	jthread thread,
	   	jclass klass) { }
	*/
	onJvmtiClassPrepare func(jvmti, jni, thread, klass uintptr)

	/*
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
	*/
	onJvmtiClassFileLoadHook func(jvmti, jni, cls_defined, loader, name, pd uintptr,
		clsd_len int32,
		class_data, new_data_len, new_data uintptr)

	/*
	   static void on_VMStart(jvmtiEnv *jvmti_env,
	   	JNIEnv* jni_env) { }
	*/
	onJvmtiVmStart func(jvmti, jni uintptr)

	/*
	   static void on_VMInit(jvmtiEnv *jvmti,
	   	JNIEnv* jni,
	   	jthread thread)*/
	onJvmtiVmInit func(jvmti JvmtiEnv, jni JniEnv, thread uintptr)
}

// SetVmInitCallback sets callback function for VMInit event
func (callbacks *JvmtiCallbacks) SetVmInitCallback(fn func(JvmtiEnv, JniEnv, uintptr)) {
	callbacks.onJvmtiVmInit = fn
	jvmti := unsafe.Pointer(_lib.jvmti)
	C.EnableJvmtiCallback(jvmti, JVMTI_EVENT_VM_INIT)
}

//export OnJvmtiEvent
func OnJvmtiEvent(eventId int, jvmti, jni, params uintptr, paramsLen int) {
	if _lib == nil {
		return
	}
	callbacks := _lib.GetCallbacks()
	jvmtiEnv := JvmtiEnv(jvmti)
	jniEnv := JniEnv(jni)

	switch eventId {
	case JVMTI_EVENT_VM_INIT:
		if callbacks.onJvmtiVmInit != nil {
			callbacks.onJvmtiVmInit(jvmtiEnv, jniEnv, 0)
		}

	case JVMTI_EVENT_VM_DEATH:

	case JVMTI_EVENT_THREAD_START:
		if callbacks.onJvmtiThreadStart != nil {
			callbacks.onJvmtiThreadStart(jvmtiEnv, jniEnv, 0)
		}

	case JVMTI_EVENT_THREAD_END:

	case JVMTI_EVENT_CLASS_FILE_LOAD_HOOK:

	case JVMTI_EVENT_CLASS_LOAD:

	case JVMTI_EVENT_CLASS_PREPARE:

	case JVMTI_EVENT_VM_START:

	case JVMTI_EVENT_EXCEPTION:

	case JVMTI_EVENT_EXCEPTION_CATCH:

	case JVMTI_EVENT_SINGLE_STEP:

	case JVMTI_EVENT_FRAME_POP:

	case JVMTI_EVENT_BREAKPOINT:

	case JVMTI_EVENT_FIELD_ACCESS:

	case JVMTI_EVENT_FIELD_MODIFICATION:

	case JVMTI_EVENT_METHOD_ENTRY:

	case JVMTI_EVENT_METHOD_EXIT:

	case JVMTI_EVENT_NATIVE_METHOD_BIND:

	case JVMTI_EVENT_COMPILED_METHOD_LOAD:

	case JVMTI_EVENT_COMPILED_METHOD_UNLOAD:

	case JVMTI_EVENT_DYNAMIC_CODE_GENERATED:

	case JVMTI_EVENT_DATA_DUMP_REQUEST:

	case JVMTI_EVENT_MONITOR_WAIT:

	case JVMTI_EVENT_MONITOR_WAITED:

	case JVMTI_EVENT_MONITOR_CONTENDED_ENTER:

	case JVMTI_EVENT_MONITOR_CONTENDED_ENTERED:

	case JVMTI_EVENT_RESOURCE_EXHAUSTED:

	case JVMTI_EVENT_GARBAGE_COLLECTION_START:

	case JVMTI_EVENT_GARBAGE_COLLECTION_FINISH:

	case JVMTI_EVENT_OBJECT_FREE:

	case JVMTI_EVENT_VM_OBJECT_ALLOC:

	default:
		fmt.Println("Unkown event ID!")
	}
}
