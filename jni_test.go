package jpprof

import (
	"runtime"
	"testing"

	"github.com/ClarkGuan/jni"
	"github.com/stretchr/testify/assert"
)

var (
	vm jni.VM
	env JniEnv // only for the creating thread
)

//  only create the VM once, spec does not support creating multiple VMs
func init() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	vm,env = JniCreateJavaVM("-verbose:gc -check:jni -Djava.class.path=./testdata")
}

func TestGetCreatedJavaVMs(t *testing.T) {
	vms := JniGetCreatedJavaVMs()
	assert.Equal(t, len(vms), 1)
}

func TestCreateJavaVM(t *testing.T) {
	assert.NotEqual(t, int(vm), 0)
	assert.NotEqual(t, int(env), 0)
}
