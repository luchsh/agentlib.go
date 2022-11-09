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

// create and launch a new Java virtual machine
// with current thread attached as the main thread
func Exec(args []string) (*JavaVM, error) {
	jvm, jni := jniCreateJavaVM(args)
	if jvm == 0 || jni == 0 {
		return nil, fmt.Errorf("Failed to create JavaVM with args: %s", strings.Join(args, " "))
	}
	jvmti, err := jvm.GetEnv(JVMTI_VERSION_1_1)
	if err != JNI_OK {
		return nil, fmt.Errorf("GetEnv error=%d",err)
	}

	vm := &JavaVM{
		jni: Env(jni),
		jvm: VM(jvm),
		jvmti: jvmtiEnv(jvmti),
	}
	return vm, nil
}


func (jvm *JavaVM) GetProperties() (map[string]string, error) {
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
		ks := (*C.char)(unsafe.Pointer(addr))
		var vs *C.char
		if e := jvm.jvmti.getSystemProperty(ks, &vs); e != JVMTI_ERROR_NONE {
			return nil, fmt.Errorf("failed get prop %s", C.GoString(ks))
		}
		defer jvm.jvmti.deallocate((*C.uchar)(unsafe.Pointer(vs)))
		res[C.GoString(ks)] = C.GoString(vs)
	}
	return res, nil
}

func (jvm *JavaVM) GetProperty(key string) string {
	return ""
}

