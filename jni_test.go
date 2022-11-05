package jpprof

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCreatedJavaVMs(t *testing.T) {
}

func TestCreateJavaVM(t *testing.T) {
	vm,jni := JniCreateJavaVM("")
	assert.NotEqual(t, uintptr(vm), 0)
	assert.NotEqual(t, uintptr(jni), 0)

	vm,jni = JniCreateJavaVM("-verbose:gc -Djava.class.path=testdata Loop")
	assert.NotEqual(t, uintptr(vm), 0)
	assert.NotEqual(t, uintptr(jni), 0)
}


