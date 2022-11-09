//
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
// limitations under the License.
//

// The goal is to decouple the exposed interface from CGO, JNI or JVMTI

package jpprof

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateJavaVM(t *testing.T) {
	vm,e := Exec([]string{}) // cannot create more than one VMs
	assert.Nil(t, vm)
	assert.NotNil(t, e)
}

func TestGetProperties(t *testing.T) {
	wantedPrpos := []string {
		"java.vm.vendor",
		"java.vm.version",
		"java.vm.name",
		"java.vm.info",
		"java.library.path",
		"java.class.path",
	}
	props, e := jvm.GetSystemProperties()
	assert.Nil(t, e)
	assert.GreaterOrEqual(t, len(props), len(wantedPrpos))
	for _,wp := range wantedPrpos {
		//t.Logf("%s=%s", wp, props[wp])
		if _,ok := props[wp]; !ok {
			t.Fatalf("mandatory property %s not found", wp)
		}
	}

	for _,wp := range wantedPrpos {
		vs := jvm.GetSystemProperty(wp)
		assert.Greater(t, len(vs), 0)
	}
}

/*
// TODO: run this test at agent OnLoad phase
func TestGetSetProperty(t *testing.T) {
	//key := fmt.Sprintf("java.test%d", time.Now().Nanosecond())
	key := "java.vm.vendor"
	v := jvm.GetSystemProperty(key)
	assert.Equal(t, len(v), 0) // should not exist at first

	value := fmt.Sprintf("value%d", time.Now().Nanosecond())
	e := jvm.SetSystemProperty(key, value)
	assert.Nil(t, e)

	nv := jvm.GetSystemProperty(key)
	assert.NotEqual(t, len(nv), 0)
	assert.Equal(t, nv, value)
}
*/
