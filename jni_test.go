package jpprof

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCreatedJavaVMs(t *testing.T) {
	vms := JniGetCreatedJavaVMs()
	assert.Equal(t, len(vms), 0)
}

func TestCreateJavaVM(t *testing.T) {
	vm,jni := JniCreateJavaVM("")
	assert.NotEqual(t, int(vm), 0)
	assert.NotEqual(t, int(jni), 0)
	vm.DestroyJavaVM()

	vm,jni = JniCreateJavaVM("-verbose:gc")
	println(vm)
	assert.NotEqual(t, int(vm), 0)
	assert.NotEqual(t, int(jni), 0)
	vm.DestroyJavaVM()
}


