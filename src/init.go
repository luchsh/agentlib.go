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
