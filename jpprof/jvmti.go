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

// #include "wrapper.h"
// #include "jvmti_wrapper.h"
// #include <jvmti.h>
import "C"

import (
	"unsafe"
)

// JvmtiEnv corresponds to jvmtiEnv*
type JvmtiEnv uintptr

func (jvmti JvmtiEnv) raw() *C.jvmtiEnv {
	return (*C.jvmtiEnv)(unsafe.Pointer(jvmti))
}

func (jvmti JvmtiEnv) asPointer() unsafe.Pointer {
	return unsafe.Pointer(jvmti)
}

func (jvmti JvmtiEnv) GetClassSignature(clazz C.jclass) (res string) {
	var sigp,genp *C.char
	C.GetClassSignature(jvmti.raw(), clazz, &sigp, &genp)
	if sigp != nil {
		defer jvmti.Deallocate((*C.uchar)(unsafe.Pointer(sigp)))
		res = C.GoString(sigp);
	}
	if genp != nil {
		defer jvmti.Deallocate((*C.uchar)(unsafe.Pointer(genp)))
		tg := C.GoString(genp);
		res = res + "<" + tg + ">"
	}
	return res
}

//////////////////////// Memory management ////////////////////
func (jvmti JvmtiEnv) Allocate(sz int64) (res *C.uchar, err C.jvmtiError) {
	err = C.Allocate(jvmti.raw(), C.jlong(sz), &res)
	return res,err
}

func (jvmti JvmtiEnv) Deallocate(mem *C.uchar) C.jvmtiError {
	return C.Deallocate(jvmti.raw(), mem)
}

//////////////////////// Thread ////////////////////
func (jvmti JvmtiEnv) GetThreadState(thrd C.jthread) (stat C.jint) {
	C.GetThreadState(jvmti.raw(), thrd, &stat)
	return stat
}

func (jvmti JvmtiEnv) GetCurrentThread() (thrd C.jthread) {
	C.GetCurrentThread(jvmti.raw(), &thrd)
	return thrd
}

func (jvmti JvmtiEnv) GetAllThreads() (threads []C.jthread) {
	var n C.jint
	var ts *C.jthread
	if er := C.GetAllThreads(jvmti.raw(), &n, &ts); er == JVMTI_ERROR_NONE {
		for i := 0; i < int(n); i++ {
			t := (C.jthread)(uintptr(unsafe.Pointer(ts)) + uintptr(i) * ptrSize)
			threads = append(threads, t)
		}
	}
	return threads
}

type ThreadInfo struct {
	Name string
	Priority int
	IsDaemon bool
	ThreadGroup uintptr
	ContextClassLoader uintptr
};

func (jvmti JvmtiEnv) GetThreadInfo(thrd C.jthread) (info *ThreadInfo) {
	ps := &C.struct__jvmtiThreadInfo { }
	C.GetThreadInfo(jvmti.raw(), thrd, ps)
	info = &ThreadInfo {
		Name: C.GoString(ps.name),
		Priority: int(ps.priority),
		IsDaemon: (ps.is_daemon != 0),
		ThreadGroup: uintptr(ps.thread_group),
		ContextClassLoader: uintptr(ps.context_class_loader),
	}
	return info
}
