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

//+build ignore

package jgo

//#include "wrapper.h"
import "C"

import (
	"fmt"
	"unsafe"
)

// JVMTI event ID definitions
// TODO: need shorter GO name
const (
	// Fake events
	JVMTI_MIN_FAKE_EVENT_TYPE_VAL = 100
	JVMTI_EVENT_AGENT_UNLOAD      = 100
	JVMTI_MAX_FAKE_EVENT_TYPE_VAL = 100
	JVMTI_EVENT_TYPE_LIMIT        = 100
)

// JvmtiCallbacks holds the JVMTI event callbacks
type JvmtiCallbacks struct {
	// to hold callbacks
	cbs []func(jvmtiEnv, ...JvmtiArg)
}

func (callbacks *JvmtiCallbacks) init() {
	callbacks.cbs = make([]func(jvmtiEnv, ...JvmtiArg), JVMTI_EVENT_TYPE_LIMIT-JVMTI_MIN_EVENT_TYPE_VAL+1)
}

type JvmtiArg uintptr

// SetCallback links a go method to process a specific JVMTI event
func (callbacks *JvmtiCallbacks) SetCallback(eventId int, fn func(jvmtiEnv, ...JvmtiArg)) {
	if eventId <= JVMTI_MAX_EVENT_TYPE_VAL && eventId >= JVMTI_MIN_EVENT_TYPE_VAL {
		callbacks.cbs[eventId-JVMTI_MIN_EVENT_TYPE_VAL] = fn
		C.EnableJvmtiCallback(unsafe.Pointer(_lib.jvmti), C.int(eventId))
	} else if eventId <= JVMTI_MAX_FAKE_EVENT_TYPE_VAL && eventId >= JVMTI_MIN_FAKE_EVENT_TYPE_VAL {
		C.EnableJvmtiCallback(unsafe.Pointer(_lib.jvmti), C.int(eventId))
	} else {
		fmt.Println("GO: Bad event id ", eventId)
	}
}

func (callbacks *JvmtiCallbacks) dispatch(eventId int, jvmti jvmtiEnv, args ...JvmtiArg) {
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
	jvmtiEnv := jvmtiEnv(jvmti)
	ra := newRawArray(params, paramsLen)

	args := make([]JvmtiArg, paramsLen)
	for i := 0; i < paramsLen; i++ {
		args[i] = JvmtiArg(ra.ptrAt(i))
	}

	callbacks.dispatch(eventId, jvmtiEnv, args...)
}
