// Copyright 2022 chuanshenglu@gmail.com
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
// limitations under the License.\n\n

package jpprof

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/ClarkGuan/jni"
	"github.com/stretchr/testify/assert"
)

var (
	vm  jni.VM
	env jni.Env // only for the creating thread
	jvm *JavaVM
)

//  only create the VM once, spec does not support creating multiple VMs
func init() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	//vm, env = jniCreateJavaVM([]string{"-verbose:gc", "-check:jni", "-Djava.class.path=./testdata"})
	var err error
	jvm, err = Exec([]string{"-verbose:gc", "-check:jni", "-Djava.class.path=./testdata"})
	if err != nil {
		panic(fmt.Errorf("Failed to create JavaVM, err=%v", err))
	}
	vm = jvm.jvm
	env = jvm.jni
}

func TestGetCreatedJavaVMs(t *testing.T) {
	vms := jniGetCreatedJavaVMs()
	assert.Equal(t, len(vms), 1)
}

func TestFindClass(t *testing.T) {
	JVM(vm).jniRun(func(env JniEnv) {
		jni := env.raw()
		clazz := jni.FindClass("Ljava/lang/Object;")
		assert.NotZero(t, uintptr(clazz))
	})
}
