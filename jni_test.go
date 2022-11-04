package jpprof

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCreatedJavaVMs(t *testing.T) {
}

func TestCreateJavaVM(t *testing.T) {
	vm := JniCreateJavaVM(nil, "")
	assert.NotEquals(vm, 0, "should not")
}


