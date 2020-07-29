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

	// added in jdk11
	// JVMTI_EVENT_SAMPLED_OBJECT_ALLOC      = 85

	// Fake events
	JVMTI_MIN_FAKE_EVENT_TYPE_VAL = 100
	JVMTI_EVENT_AGENT_UNLOAD      = 100
	JVMTI_MAX_FAKE_EVENT_TYPE_VAL = 100
	JVMTI_EVENT_TYPE_LIMIT        = 100
)

// JvmtiCallbacks holds the JVMTI event callbacks
type JvmtiCallbacks struct {
	// to hold callbacks
	cbs []func(JvmtiEnv, ...JvmtiArg)
}

func (callbacks *JvmtiCallbacks) init() {
	callbacks.cbs = make([]func(JvmtiEnv, ...JvmtiArg), JVMTI_EVENT_TYPE_LIMIT-JVMTI_MIN_EVENT_TYPE_VAL+1)
}

type JvmtiArg uintptr

// SetCallback links a go method to process a specific JVMTI event
func (callbacks *JvmtiCallbacks) SetCallback(eventId int, fn func(JvmtiEnv, ...JvmtiArg)) {
	if eventId <= JVMTI_MAX_EVENT_TYPE_VAL && eventId >= JVMTI_MIN_EVENT_TYPE_VAL {
		callbacks.cbs[eventId-JVMTI_MIN_EVENT_TYPE_VAL] = fn
		C.EnableJvmtiCallback(unsafe.Pointer(_lib.jvmti), C.int(eventId))
	} else if eventId <= JVMTI_MAX_FAKE_EVENT_TYPE_VAL && eventId >= JVMTI_MIN_FAKE_EVENT_TYPE_VAL {
		C.EnableJvmtiCallback(unsafe.Pointer(_lib.jvmti), C.int(eventId))
	} else {
		fmt.Println("GO: Bad event id ", eventId)
	}
}

func (callbacks *JvmtiCallbacks) dispatch(eventId int, jvmti JvmtiEnv, args ...JvmtiArg) {
	fn := callbacks.cbs[eventId-JVMTI_MIN_EVENT_TYPE_VAL]
	if fn != nil {
		fn(jvmti, args...)
	}
}

type rawArray struct {
	base  uintptr
	len   int
	eSize int
	idx   int
}

func newRawArray(b uintptr, l int) *rawArray {
	ar := new(rawArray)
	ar.base = b
	ar.len = l
	ar.eSize = int(unsafe.Sizeof(ar.base))
	return ar
}

func (arr *rawArray) ptrAt(idx int) uintptr {
	p := arr.base + uintptr(arr.eSize*idx)
	return *(*uintptr)(unsafe.Pointer(p))
}

func (arr *rawArray) next() uintptr {
	if arr.idx >= arr.len {
		return uintptr(0)
	}
	p := arr.base + uintptr(arr.eSize*arr.idx)
	return *(*uintptr)(unsafe.Pointer(p))
}

// OnJvmtiEvent dispatches all the event to corresponding Go handlers
// runs on a JavaThread
//export OnJvmtiEvent
func OnJvmtiEvent(eventId int, jvmti, params uintptr, paramsLen int) {
	if _lib == nil {
		return
	}
	callbacks := _lib.GetCallbacks()
	jvmtiEnv := JvmtiEnv(jvmti)
	ra := newRawArray(params, paramsLen)

	args := make([]JvmtiArg, paramsLen)
	for i := 0; i < paramsLen; i++ {
		a := JvmtiArg(ra.ptrAt(i))
		args = append(args, a)
	}

	callbacks.dispatch(eventId, jvmtiEnv, args...)
}
