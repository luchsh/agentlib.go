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
package main

// #include "wrapper.h"
// #include "jvmti_wrapper.h"
import "C"

import(
	"unsafe"
)

// JvmtiEnv corresponds to jvmtiEnv*
type JvmtiEnv uintptr

func (jvmti JvmtiEnv) asPointer() unsafe.Pointer {
	return unsafe.Pointer(jvmti)
}

func (jvmti JvmtiEnv) GetClassSignature(clazz uintptr) (res string) {
	var sigp,genp unsafe.Pointer
	C.GetClassSignature(jvmti.asPointer(), unsafe.Pointer(clazz), unsafe.Pointer(&sigp), unsafe.Pointer(&genp))
	if sigp != nil {
		defer jvmti.Deallocate(sigp)
		res = C.GoString((*C.char)(sigp));
	}
	if genp != nil {
		defer jvmti.Deallocate(genp)
		tg := C.GoString((*C.char)(genp));
		res = res + "<" + tg + ">"
	}
	return res
}

//////////////////////// Memory management ////////////////////
func (jvmti JvmtiEnv) Allocate(sz int64) (res unsafe.Pointer) {
	C.Allocate(jvmti.asPointer(), C.longlong(sz), unsafe.Pointer(&res))
	return res
}

func (jvmti JvmtiEnv) Deallocate(mem unsafe.Pointer) int {
	return int(C.Deallocate(jvmti.asPointer(), mem))
}

//////////////////////// Thread ////////////////////
func (jvmti JvmtiEnv) GetThreadState(thrd uintptr) (stat int) {
	C.GetThreadState(jvmti.asPointer(), unsafe.Pointer(thrd), unsafe.Pointer(&stat))
	return stat
}

func (jvmti JvmtiEnv) GetCurrentThread() (thrd uintptr) {
	C.GetCurrentThread(jvmti.asPointer(), unsafe.Pointer(&thrd))
	return thrd
}

func (jvmti JvmtiEnv) GetAllThreads() (threads []uintptr) {
	var count int
	var p unsafe.Pointer
	C.GetAllThreads(jvmti.asPointer(), unsafe.Pointer(&count), unsafe.Pointer(&p))
	defer jvmti.Deallocate(p)
	if count > 0 {
		threads = make([]uintptr, count)
		for i := 0; i < count; i++ {
			threads[i] = *(*uintptr)(unsafe.Pointer((uintptr(p)+uintptr(i*8))))
		}
	}
	return threads
}
