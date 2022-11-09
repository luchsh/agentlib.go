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

func TestGetProperties(t *testing.T) {
	wantedPrpos := []string {
		"java.vm.vendor",
		"java.vm.version",
		"java.vm.name",
		"java.vm.info",
		"java.library.path",
		"java.class.path",
	}
	props, e := jvm.GetProperties()
	assert.Nil(t, e)
	assert.GreaterOrEqual(t, len(props), len(wantedPrpos))
	for _,wp := range wantedPrpos {
		t.Logf("%s=%s", wp, props[wp])
		if _,ok := props[wp]; !ok {
			t.Fatalf("mandatory property %s not found", wp)
		}
	}
}
