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

import (
	"runtime/debug"
	"testing"
)

func TestParseMethod(t *testing.T) {
	in := "jvmtiError (JNICALL *GetThreadCpuTimerInfo) (jvmtiEnv* env, jvmtiTimerInfo* info_ptr);"
	m := parseJvmtiApiDecl(in)
	assertEQ(m.name, "GetThreadCpuTimerInfo", t)
	assertEQ(m.ret, "jvmtiError", t)
	assertEQ(len(m.pname), 2, t)
	assertEQ(len(m.ptype), 2, t)
	assertEQ(m.pname[0], "env", t)
	assertEQ(m.pname[1], "info_ptr", t)
	assertEQ(m.ptype[0], "jvmtiEnv*", t)
	assertEQ(m.ptype[1], "jvmtiTimerInfo*", t)

	in = "jvmtiError (JNICALL *SetEventNotificationMode) (jvmtiEnv* env,jvmtiEventMode mode,jvmtiEvent event_type,const jthread event_thread,...);"
	m = parseJvmtiApiDecl(in)
	assertEQ(m.name, "SetEventNotificationMode", t)
	assertEQ(m.ret, "jvmtiError", t)
	assertEQ(len(m.pname), 4, t)
	assertEQ(m.ptype[3], "const jthread", t)
}

// A lonely traveller from the Java continent
func assertEQ(a, b interface{}, t *testing.T) {
	if a != b {
		debug.PrintStack()
		t.Fatal("assertEQ failed: ", a, "!=", b, "\n")
	}
}
func assertTrue(b bool, t *testing.T) {
	if !b {
		debug.PrintStack()
		t.Fatal("Assertion failed")
	}
}
