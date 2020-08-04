package main

import (
	"fmt"
)

// This file contains the code that user has to write

// AgentGoOnLoad is the mandatory global hook provided by user code
func AgentGoOnLoad(lib *AgentLib) {
	fmt.Println("GO: AgentGoOnload")
	fmt.Println("GO: Agent command line options:", lib.options)
	lib.GetCallbacks().SetCallback(JVMTI_EVENT_VM_INIT, func(jvmti JvmtiEnv, args ...JvmtiArg) {
		fmt.Println("GO: OnJvmtiVmInit(): triggered on Go level")
	})
	lib.GetCallbacks().SetCallback(JVMTI_EVENT_CLASS_LOAD, func(jvmti JvmtiEnv, args ...JvmtiArg) {
		name := jvmti.GetClassSignature(uintptr(args[2]))
		fmt.Println("GO: ClassLoad event: ", name)
	})
	lib.GetCallbacks().SetCallback(JVMTI_EVENT_CLASS_PREPARE, func(jvmti JvmtiEnv, args ...JvmtiArg) {
		name := jvmti.GetClassSignature(uintptr(args[2]))
		fmt.Println("GO: ClassPrepare event: ", name)
	})
	lib.GetCallbacks().SetCallback(JVMTI_EVENT_AGENT_UNLOAD, func(jvmti JvmtiEnv, args ...JvmtiArg) {
		fmt.Println("GO: AgentUnloaded")
	})
	lib.GetCallbacks().SetCallback(JVMTI_EVENT_METHOD_ENTRY, func(jvmti JvmtiEnv, args ...JvmtiArg) {
		fmt.Println("GO: method entry")
	})
	lib.GetCallbacks().SetCallback(JVMTI_EVENT_METHOD_EXIT, func(jvmti JvmtiEnv, args ...JvmtiArg) {
		fmt.Println("GO: method exit")
	})
}

func AgentGoOnUnload() {
	fmt.Println("GO: AgentGoOnUnload")
}

func main() {}
