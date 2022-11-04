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
package jpprof

// #include <jni.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"

	"github.com/ClarkGuan/jni"
)

const (
	ptrSize = 8
)

type JniEnv jni.Env

func JniGetCreatedJavaVMs() (vms []jni.VM) {
	l := 128
	buf := C.malloc(C.size_t(l * ptrSize))
	defer C.free(buf)
	if 0 == C.JNI_GetCreatedJavaVMs(buf, l, nil) {
		for p := uintptr(buf); p < uintptr(buf)+128*ptrSize; p+= ptrSize {
			addr := *(*uintptr)(unsafe.Pointer(p))
			vms = append(vms, jni.VM(addr))
		}
	}
	return
}

// Create a Java VM
func JniCreateJavaVM(env []string, args string) jni.VM {
	var vmp uintptr
	var envp uintptr

	if len(env) > 0 {
		buf := C.malloc(len(env)*ptrSize)
		defer C.free(buf)
		for i,e := range env {
			ce := C.CString(e)
			defer C.free(ce)
			slotAddr := uintptr(buf)+uintptr(i)*ptrSize
			*(**C.char)(unsafe.Pointer(slotAddr)) = ce
		}
	}

	cargs := C.CString(args)
	defer C.free(cargs)
	C.JNI_CreateJavaVM(unsafe.Pointer(&vmp),  unsafe.Pointer(envp), unsafe.Pointer(cargs))

	return jni.VM(vmp)
}

// Retrieve current virtual machine of this process if exists
func CurrentVM() jni.VM {
	vms := JniGetCreatedJavaVMs()
	if len(vms) > 0 {
		return vms[0]
	}
	return 0
}
