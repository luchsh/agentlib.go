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

// AgentLib defines the information of a agent library
type AgentLib struct {
	// The unique instance of JavaVM
	javaVM uintptr 
	// command line options to this agent
	options string
	// current jvmti env
	jvmti	JvmtiEnv
	// Callbacks
	callbacks JvmtiCallbacks
}

// GetCallbacks returns the callbacks of this lib
func (agent *AgentLib) GetCallbacks() *JvmtiCallbacks {
	return &agent.callbacks
}

// The global instance of this agent lib
var _lib *AgentLib

//export OnAgentLoad
func OnAgentLoad(javaVM, jvmti uintptr, options *C.char) {
	_lib = new(AgentLib)
	_lib.javaVM = javaVM
	_lib.options = C.GoString(options)
	_lib.jvmti = JvmtiEnv(jvmti)
	AgentGoOnLoad(_lib)
}

//export OnAgentUnload
func OnAgentUnload() int32 {
	if _lib.callbacks.onAgentUnload != nil {
		_lib.callbacks.onAgentUnload()
	}
	return 0
}

//export MainForwardLoop
func MainForwardLoop() {
	// TODO: cross-runtime forwarding mechanism
}
