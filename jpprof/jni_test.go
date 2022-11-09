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
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	vm  JVM
	env JniEnv // only for the creating thread
)

//  only create the VM once, spec does not support creating multiple VMs
func init() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	vm, env = JniCreateJavaVM("-verbose:gc -check:jni -Djava.class.path=./testdata")
}

func TestGetCreatedJavaVMs(t *testing.T) {
	vms := JniGetCreatedJavaVMs()
	assert.Equal(t, len(vms), 1)
}

func TestCreateJavaVM(t *testing.T) {
	assert.NotEqual(t, int(vm), 0)
	assert.NotEqual(t, int(env), 0)
}

func TestFindClass(t *testing.T) {
	vm.JniRun(func(env JniEnv) {
		jni := env.raw()
		clazz := jni.FindClass("Ljava/lang/Object;")
		assert.NotZero(t, uintptr(clazz))
	})
}
