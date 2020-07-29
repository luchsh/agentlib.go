package main

import (
	"fmt"
)

// This file contains the code that user has to write

// AgentGoOnLoad is the mandatory global hook provided by user code
func AgentGoOnLoad(lib *AgentLib) {
	fmt.Println("GO: AgentGoOnload")
	fmt.Println("GO: Agent command line options:", lib.options)
	lib.GetCallbacks().SetVmInitCallback(func(jvmti JvmtiEnv, jni JniEnv, thread uintptr) {
		fmt.Println("GO: OnJvmtiVmInit(): triggered on Go level")
	})
	lib.GetCallbacks().SetClassLoadCallback(func(jvmti JvmtiEnv, jni JniEnv, thread, clazz uintptr) {
		fmt.Println("GO: ClassLoad event")
	})
	lib.GetCallbacks().SetClassPrepareCallback(func(jvmti JvmtiEnv, jni JniEnv, thread, clazz uintptr) {
		fmt.Println("GO: ClassPrepare event")
	})
}

func AgentGoOnUnload() {
	fmt.Println("GO: AgentGoOnUnload")
}

func main() {}
