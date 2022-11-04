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
	"strings"
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
	var n_vm C.jsize
	if 0 == C.JNI_GetCreatedJavaVMs((**C.JavaVM)(buf), C.jsize(l), &n_vm) {
		if int(n_vm) > l {
			panic("insufficient buffer space")
		}
		for i := 0; i < int(n_vm); i++ {
			p := uintptr(buf) + uintptr(i) * ptrSize
			addr := *(*uintptr)(unsafe.Pointer(p))
			vms = append(vms, jni.VM(addr))
		}
	}
	return
}

// Create a Java VM
func JniCreateJavaVM(args string) (jni.VM, JniEnv) {
	var vmp jni.VM
	var envp JniEnv

	var vmargs unsafe.Pointer
	jva := (*C.JavaVMInitArgs)(C.malloc(C.sizeof_JavaVMInitArgs))
	defer C.free(unsafe.Pointer(jva))
	jva.version = jni.JNI_VERSION_1_6
	jva.nOptions = 0
	if len(args) > 0 {
		fds := strings.Fields(args)
		jva.nOptions = C.jint(len(fds))
		opts := C.malloc(C.size_t(C.sizeof_JavaVMOption * len(fds)))
		defer C.free(opts)
		for i,a := range fds {
			println(a)
			o := (*C.JavaVMOption)(unsafe.Pointer(uintptr(opts)+uintptr(i)*C.sizeof_JavaVMOption))
			ca := C.CString(a)
			defer C.free(unsafe.Pointer(ca))
			o.optionString = ca
		}
		jva.options = (*C.JavaVMOption)(opts)
	}
	vmargs = unsafe.Pointer(jva)

	C.JNI_CreateJavaVM((**C.JavaVM)(unsafe.Pointer(&vmp)), (*unsafe.Pointer)(unsafe.Pointer(&envp)), vmargs)

	return vmp, envp
}

// Retrieve current virtual machine of this process if exists
func CurrentVM() jni.VM {
	vms := JniGetCreatedJavaVMs()
	if len(vms) > 0 {
		return vms[0]
	}
	return 0
}
