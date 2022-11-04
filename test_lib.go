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

// +build ignore

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

// This file contains the code that user has to write
var classesLoaded int64

// AgentGoOnLoad is the mandatory global hook provided by user code
func AgentGoOnLoad(lib *AgentLib) {
	fmt.Println("GO: AgentGoOnload")
	fmt.Println("GO: Agent command line options:", lib.options)
	lib.GetCallbacks().SetCallback(JVMTI_EVENT_VM_INIT, func(jvmti JvmtiEnv, args ...JvmtiArg) {
		p := jvmti.Allocate(int64(4096))
		defer jvmti.Deallocate(p)
		fmt.Printf("GO: OnJvmtiVmInit(): triggered on Go level\np=%v\n", p)

		ch := make(chan bool)
		TestMain(lib.jvmti, ch)
		//<-ch
	})
	lib.GetCallbacks().SetCallback(JVMTI_EVENT_CLASS_LOAD, func(jvmti JvmtiEnv, args ...JvmtiArg) {
		name := jvmti.GetClassSignature(uintptr(args[2]))
		fmt.Printf("GO: ClassLoad [%d] event: %s\n", classesLoaded, name)
		atomic.AddInt64(&classesLoaded, int64(1))
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

	go func() {
		http.HandleFunc("/classesLoaded", func(w http.ResponseWriter, _ *http.Request) {
			io.WriteString(w, fmt.Sprintf("Total classes loaded: %d\n", classesLoaded))
		})
		println("Server loaded...")
		log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
	}()
}

func AgentGoOnUnload() {
	fmt.Println("GO: AgentGoOnUnload")
}

func testThreads(jvmti JvmtiEnv) {
	allThreads := jvmti.GetAllThreads()
	if len(allThreads) <= 0 {
		panic("Failed to get all threads")
	}
	for _,t := range allThreads {
		info := jvmti.GetThreadInfo(t)
		fmt.Printf("Thread %v %v\n", t, info)
	}
}

func TestMain(jvmti JvmtiEnv, ch chan bool) {
	time.Sleep(5 * time.Second)
	testThreads(jvmti)
	//ch <- true
}
