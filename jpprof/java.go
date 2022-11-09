//
// Copyright 2022 chuanshenglu@gmail.com
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

// The goal is to decouple the exposed interface from CGO, JNI or JVMTI

package jpprof

//#include<jni.h>
//#include<stdlib.h>
import "C"

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	. "github.com/ClarkGuan/jni"
)

type JClass uintptr
type JObject uintptr

// Unified interface for both jni and jvmti
type JavaVM struct {
	jvmti jvmtiEnv
	jni   Env
	jvm   VM
}

// The JavaVM instance for main thread
// TODO: maybe we should remove curVM
var theVM *JavaVM

// create and launch a new Java virtual machine
// with current thread attached as the main thread
func Exec(args []string) (*JavaVM, error) {
	if theVM != nil {
		return nil, fmt.Errorf("Cannot create multiple JVM in the same process")
	}
	jvm, jni := jniCreateJavaVM(args)
	if jvm == 0 || jni == 0 {
		return nil, fmt.Errorf("Failed to create JavaVM with args: %s", strings.Join(args, " "))
	}
	jvmti, err := jvm.GetEnv(JVMTI_VERSION_1_1)
	if err != JNI_OK {
		return nil, fmt.Errorf("GetEnv error=%d",err)
	}

	theVM = &JavaVM{
		jni: Env(jni),
		jvm: VM(jvm),
		jvmti: jvmtiEnv(jvmti),
	}
	return theVM, nil
}

// create a usable *JavaVM instance from raw jni.VM
func createJavaVM(vm VM) (*JavaVM, error) {
	if vm == 0 {
		return nil, fmt.Errorf("NULL VM")
	}
	env, i := vm.AttachCurrentThread()
	if i != JNI_OK {
		return nil, fmt.Errorf("Cannot attach current thread, error=%s", describeJNIError(i))
	}
	je, e := vm.GetEnv(JVMTI_VERSION_1_1)
	if e != JNI_OK {
		return nil, fmt.Errorf("Cannot get JVMTI env, error=%s", describeJNIError(e))
	}
	return &JavaVM {
		jni: env,
		jvmti: jvmtiEnv(je),
		jvm: vm,
	}, nil
}

// Get the context VM for a goroutine to use
// must be called after setting up the global unique VM instance
// e.g. calling Exec(...)
func contextVM() (*JavaVM, error) {
	if theVM == nil {
		return nil, fmt.Errorf("JavaVM instance not found for this process")
	}
	return createJavaVM(theVM.jvm)
}

// Retrieve the unique VM instance
// Current spec only allows one JVM in each process
func CurrentVM() *JavaVM {
	v,e := contextVM()
	if e != nil {
		panic(e)
	}
	return v
}

// Retrieve all the properties from the Java VM
func (jvm *JavaVM) GetSystemProperties() (map[string]string, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	var n C.jint
	var p **C.char
	if e := jvm.jvmti.getSystemProperties(&n, &p); e != JVMTI_ERROR_NONE {
		return nil, fmt.Errorf("Failed to get system properties: %d", int(e))
	}
	if n <= 0 {
		return nil, fmt.Errorf("no properties found!")
	}

	defer jvm.jvmti.deallocate((*C.uchar)(unsafe.Pointer(p)))

	res := make(map[string]string)
	for i := 0; i < int(n); i++ {
		addr := uintptr(unsafe.Pointer(p)) + uintptr(i) * ptrSize
		ks := *(**C.char)(unsafe.Pointer(addr))
		var vs *C.char
		if e := jvm.jvmti.getSystemProperty(ks, &vs); e != JVMTI_ERROR_NONE {
			return nil, fmt.Errorf("failed get prop %s", C.GoString(ks))
		}
		defer jvm.jvmti.deallocate((*C.uchar)(unsafe.Pointer(vs)))
		res[C.GoString(ks)] = C.GoString(vs)
	}
	return res, nil
}

// Retrieve property value of given key
// returns "" if not found or error
func (jvm *JavaVM) GetSystemProperty(key string) string {
	var vs *C.char
	ks := C.CString(key)
	defer C.free(unsafe.Pointer(ks))
	if e := jvm.jvmti.getSystemProperty(ks, &vs); e != JVMTI_ERROR_NONE {
		return ""
	}
	defer jvm.jvmti.deallocate((*C.uchar)(unsafe.Pointer(vs)))
	return C.GoString(vs)
}

// Set the value of target system property
// only allowed to be called at agent OnLoad phase
// According to https://docs.oracle.com/javase/8/docs/platform/jvmti/jvmti.html#GetSystemProperty
func (jvm *JavaVM) SetSystemProperty(key, value string) error {
	ck := C.CString(key)
	cv := C.CString(value)
	defer func() {
		C.free(unsafe.Pointer(ck))
		C.free(unsafe.Pointer(cv))
	}()
	if e := jvm.jvmti.setSystemProperty(ck, cv); e != JVMTI_ERROR_NONE {
		return fmt.Errorf("Failed to set property, error=%v\n", e)
	}
	return nil
}
