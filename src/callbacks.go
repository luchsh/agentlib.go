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

import "C"

import (
	"fmt"
	"unsafe"
)

//export OnJvmtiEvent
func OnJvmtiEvent(eventId int, jvmti, jni, params unsafe.Pointer, paramsLen int) {

}

/*
static void on_Breakpoint(jvmtiEnv *jvmti_env,
	JNIEnv* jni_env,
	jthread thread,
	jmethodID method,
	jlocation location) { }
*/
//export OnJvmtiBreakpoint
func OnJvmtiBreakpoint(jvmti, jni, thread, method, loc uintptr) {
	fmt.Println("OnJvmtiBreakpoint")
}

/*
static void on_SingleStep(jvmtiEnv *jvmti_env,
	JNIEnv* jni_env,
	jthread thread,
	jmethodID method,
	jlocation location) { }
*/
//export OnJvmtiSingleStep
func OnJvmtiSingleStep(jvmti, jni, thread, method, loc uintptr) {
	fmt.Println("OnJvmtiSingleStep")
}

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
//export OnJvmtiFieldAccess
func OnJvmtiFieldAccess(jvmti, jni, method, loc, fclazz, obj, field uintptr) {
	fmt.Println("OnJvmtiFieldAccess")
}

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
//export OnJvmtiFieldModification
func OnJvmtiFieldModification(jvmti, jni, thread, method, loc, fklazz, obj, field, sig, newval uintptr) {
	fmt.Println("OnJvmtiFieldModification")
}

/*
static void on_FramePop(jvmtiEnv *jvmti_env,
	JNIEnv* jni_env,
	jthread thread,
	jmethodID method,
	jboolean was_popped_by_exception) {}
*/
//export OnJvmtiFramePop
func OnJvmtiFramePop(jvmti, jni, thread, method uintptr, by_excep bool) {
	fmt.Println("OnJvmtiFramePop")
}

/*
static void on_MethodEntry(jvmtiEnv *jvmti_env,
	JNIEnv* jni_env,
	jthread thread,
	jmethodID method) {}
*/
//export OnJvmtiMethodEntry
func OnJvmtiMethodEntry(jvmi, jni, thread, method uintptr) {
	fmt.Println("OnJvmtiMethodEntry")
}

/*
static void on_MethodExit(jvmtiEnv *jvmti_env,
	JNIEnv* jni_env,
	jthread thread,
	jmethodID method,
	jboolean was_popped_by_exception,
	jvalue return_value) {}
*/
//export OnJvmtiMethodExit
func OnJvmtiMethodExit(jvmti, jni, thread, method uintptr, by_excep bool, ret_val uintptr) {
	fmt.Println("OnJvmtiMethodExit")
}

/*
static void on_NativeMethodBind(jvmtiEnv *jvmti_env,
	JNIEnv* jni_env,
	jthread thread,
	jmethodID method,
	void* address,
	void** new_address_ptr) { }
*/
//export OnJvmtiNativeMethodBind
func OnJvmtiNativeMethodBind(jvmti, jni, thread, method, addr, new_addr_ptr uintptr) {
	fmt.Println("OnJvmtiNativeMethodBind")
}

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
//export OnJvmtiException
func OnJvmtiException(jvmti, jni, thread, method, location, exception, catch_m, catch_l uintptr) {
	fmt.Println("OnJvmtiException")
}

/*
static void on_ExceptionCatch(jvmtiEnv *jvmti_env,
	JNIEnv* jni_env,
	jthread thread,
	jmethodID method,
	jlocation location,
	jobject exception) { }
*/
//export OnJvmtiExceptionCatch
func OnJvmtiExceptionCatch(jvmti, jni, thread, method, location, exception uintptr) {
	fmt.Println("OnJvmtiExceptionCatch")
}

/*
static void on_ThreadStart(jvmtiEnv *jvmti_env,
	JNIEnv* jni_env,
	jthread thread) { }
*/
//export OnJvmtiThreadStart
func OnJvmtiThreadStart(jvmti, jni, thread uintptr) {
	fmt.Println("OnJvmtiThreadStart")
}

/*
static void on_ThreadEnd(jvmtiEnv *jvmti_env,
	JNIEnv* jni_env,
	jthread thread) { }
*/
//export OnJvmtiThreadEnd
func OnJvmtiThreadEnd(jvmti, jni, thread uintptr) {
	fmt.Println("OnJvmtiThreadEnd")
}

/*
static void on_ClassLoad(jvmtiEnv *jvmti_env,
	JNIEnv* jni_env,
	jthread thread,
	jclass klass) { }
*/
//export OnJvmtiClassLoad
func OnJvmtiClassLoad(jvmti, jni, thread, klass uintptr) {
	fmt.Println("OnJvmtiClassLoad")
}

/*
static void on_ClassPrepare(jvmtiEnv *jvmti_env,
	JNIEnv* jni_env,
	jthread thread,
	jclass klass) { }
*/
//export OnJvmtiClassPrepare
func OnJvmtiClassPrepare(jvmti, jni, thread, klass uintptr) {
	fmt.Println("OnJvmtiClassPrepare")
}

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
//export OnJvmtiClassFileLoadHook
func OnJvmtiClassFileLoadHook(jvmti, jni, cls_defined, loader, name, pd uintptr,
	clsd_len int32,
	class_data, new_data_len, new_data uintptr) {
	fmt.Println("OnJvmtiClassFileLoadHook")
}

/*
static void on_VMStart(jvmtiEnv *jvmti_env,
	JNIEnv* jni_env) { }
*/
//export OnJvmtiVmStart
func OnJvmtiVmStart(jvmti, jni uintptr) {
	fmt.Println("OnJvmtiVmInit()\n")
}

/*
static void on_VMInit(jvmtiEnv *jvmti,
	JNIEnv* jni,
	jthread thread)*/
//export OnJvmtiVmInit
func OnJvmtiVmInit(jvmti, jni, thread uintptr) {
	fmt.Println("OnJvmtiVmInit()\n")
}
