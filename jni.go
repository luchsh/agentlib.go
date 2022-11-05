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
package jpprof

// #include <jni.h>
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/ClarkGuan/jni"
)

const (
	ptrSize = 8
)

type JniEnv jni.Env
type JVM jni.VM

func JniGetCreatedJavaVMs() (vms []jni.VM) {
	l := 128
	buf := C.malloc(C.size_t(l * ptrSize))
	defer C.free(buf)
	var n C.jsize
	if 0 == C.JNI_GetCreatedJavaVMs((**C.JavaVM)(buf), C.jsize(l), &n) {
		for i := C.jsize(0); i < n; i++ {
			p := uintptr(buf)+uintptr(i)*ptrSize
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

	vmargs := C.malloc(C.sizeof_JavaVMOption)
	defer C.free(vmargs)
	jva := (*C.JavaVMInitArgs)(vmargs)
	jva.version = jni.JNI_VERSION_1_6
	jva.nOptions = 0
	jva.ignoreUnrecognized = jni.JNI_TRUE
	if len(args) > 0 {
		fds := strings.Fields(args)
		opts := C.malloc(C.size_t(C.sizeof_JavaVMOption * len(fds)))
		defer C.free(opts)
		jva.nOptions = C.jint(len(fds))
		jva.options = (*C.JavaVMOption)(opts)
		for i,a := range fds {
			o := (*C.JavaVMOption)(unsafe.Pointer(uintptr(opts)+uintptr(i)*C.sizeof_JavaVMOption))
			o.optionString = C.CString(a)
		}
	}
	vmargs = unsafe.Pointer(jva)

	e := C.JNI_CreateJavaVM((**C.JavaVM)(unsafe.Pointer(&vmp)), (*unsafe.Pointer)(unsafe.Pointer(&envp)), vmargs)
	if e != jni.JNI_OK {
		fmt.Printf("Failed to create Java VM, error=%d (%s)\n", e, DescribeJNIError(e))
		return 0,0
	}

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

var jniErrorTextMap = map[int]string {
	jni.JNI_ERR: "JNI_ERR",
	jni.JNI_EDETACHED: "JNI_EDETACHED",
	jni.JNI_EVERSION: "JNI_EVERSION",
	jni.JNI_ENOMEM: "JNI_ENOMEM",
	jni.JNI_EEXIST: "JNI_EEXIST",
	jni.JNI_EINVAL: "JNI_EINVAL",
}

func DescribeJNIError(ev C.int) string {
	if tv,ok := jniErrorTextMap[int(ev)]; ok {
		return tv
	}
	return "Unknown error"
}
