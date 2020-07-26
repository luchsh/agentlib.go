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

// AgentLib defines the information of a agent library
type AgentLib struct {
	// The unique instance of JavaVM
	javaVM unsafe.Pointer
	// command line options to this agent
	options string
}

// The global instance of this agent lib
var _lib *AgentLib

//export OnAgentLoad
func OnAgentLoad(javaVM unsafe.Pointer, options *C.char) int32 {
	_lib = new(AgentLib)
	_lib.javaVM = javaVM
	_lib.options = C.GoString(options)
	fmt.Println("Agent command line options:", _lib.options)
	return 0
}

//export OnAgentUnload
func OnAgentUnload() int32 {
	// TODO
	return 0
}

//export OnJvmInit
func OnJvmInit(jvmti unsafe.Pointer, jni unsafe.Pointer) {
}

//export MainForwardLoop
func MainForwardLoop() {
}
