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

// #include "wrapper.h"
// #include "jvmti_wrapper.h"
import "C"

import(
	"unsafe"
)

// JvmtiEnv corresponds to jvmtiEnv*
type JvmtiEnv uintptr

func (jvmti JvmtiEnv) GetClassSignature(clazz uintptr) (res string) {
	var sigp uintptr
	var genp uintptr
	env := uintptr(jvmti)
	C.GetClassSignature(unsafe.Pointer(env), unsafe.Pointer(clazz), unsafe.Pointer(&sigp), unsafe.Pointer(&genp))
	if sigp != uintptr(0) {
		defer jvmti.Deallocate(sigp)
		res = C.GoString((*C.char)(unsafe.Pointer(sigp)));
	}
	if genp != uintptr(0) {
		defer jvmti.Deallocate(genp)
		tg := C.GoString((*C.char)(unsafe.Pointer(genp)));
		res = res + "<" + tg + ">"
	}
	return res
}

func (jvmti JvmtiEnv) Deallocate(mem uintptr) int {
	env := uintptr(jvmti)
	return int(C.Deallocate(unsafe.Pointer(env), unsafe.Pointer(mem)))
}