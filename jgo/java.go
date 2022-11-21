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

package jgo

//#include<jni.h>
//#include<stdlib.h>
//#include<jvmti.h>
import "C"

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	. "github.com/ClarkGuan/jni"
)

type JClass uintptr
type JObject uintptr

// Unified interface for both jni and jvmti
type JavaVM struct {
	jvmti jvmtiEnv
	jni   Env
	jvm   VM
}

// The JavaVM instance for main thread
// TODO: maybe we should remove curVM
var theVM *JavaVM

// create and launch a new Java virtual machine
// with current thread attached as the main thread
func Exec(args []string) (*JavaVM, error) {
	if theVM != nil {
		return nil, fmt.Errorf("Cannot create multiple JVM in the same process")
	}
	jvm, jni := jniCreateJavaVM(args)
	if jvm == 0 || jni == 0 {
		return nil, fmt.Errorf("Failed to create JavaVM with args: %s", strings.Join(args, " "))
	}
	jvmti, err := jvm.GetEnv(JVMTI_VERSION_1_1)
	if err != JNI_OK {
		return nil, fmt.Errorf("GetEnv error=%d", err)
	}

	theVM = &JavaVM{
		jni:   Env(jni),
		jvm:   VM(jvm),
		jvmti: jvmtiEnv(jvmti),
	}
	return theVM, nil
}

// create a usable *JavaVM instance from raw jni.VM
func createJavaVM(vm VM) (*JavaVM, error) {
	if vm == 0 {
		return nil, fmt.Errorf("NULL VM")
	}
	env, i := vm.AttachCurrentThread()
	if i != JNI_OK {
		return nil, fmt.Errorf("Cannot attach current thread, error=%s", describeJNIError(i))
	}
	je, e := vm.GetEnv(JVMTI_VERSION_1_1)
	if e != JNI_OK {
		return nil, fmt.Errorf("Cannot get JVMTI env, error=%s", describeJNIError(e))
	}
	return &JavaVM{
		jni:   env,
		jvmti: jvmtiEnv(je),
		jvm:   vm,
	}, nil
}

// Get the context VM for a goroutine to use
// must be called after setting up the global unique VM instance
// e.g. calling Exec(...)
func contextVM() (*JavaVM, error) {
	if theVM == nil {
		return nil, fmt.Errorf("JavaVM instance not found for this process")
	}
	return createJavaVM(theVM.jvm)
}

// Retrieve the unique VM instance
// Current spec only allows one JVM in each process
func CurrentVM() *JavaVM {
	v, e := contextVM()
	if e != nil {
		panic(e)
	}
	env, _ := v.jvm.AttachCurrentThread()
	return &JavaVM{
		jni:   env,
		jvm:   v.jvm,
		jvmti: v.jvmti,
	}
}

func (jvm *JavaVM) FullVersion() string {
	iv := int(jvm.jni.GetVersion())
	primVer := (iv & 0xFFFF0000) >> 16
	minorVer := (iv & 0xFFFF)
	return fmt.Sprintf("Java version %d.%d", primVer, minorVer)
}

// Retrieve all the properties from the Java VM
func (jvm *JavaVM) GetSystemProperties() (map[string]string, error) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	var n C.jint
	var p **C.char
	if e := jvm.jvmti.getSystemProperties(&n, &p); e != JVMTI_ERROR_NONE {
		return nil, fmt.Errorf("Failed to get system properties: %s", describeJvmtiError(int(e)))
	}
	if n <= 0 {
		return nil, fmt.Errorf("no properties found!")
	}

	defer deallocate(p)

	res := make(map[string]string)
	for i := 0; i < int(n); i++ {
		addr := addrAt(p, i)
		ks := *(**C.char)(unsafe.Pointer(addr))
		var vs *C.char
		if e := jvm.jvmti.getSystemProperty(ks, &vs); e != JVMTI_ERROR_NONE {
			return nil, fmt.Errorf("failed get prop %s, error=%s", C.GoString(ks), describeJvmtiError(int(e)))
		}
		defer deallocate(vs)
		res[C.GoString(ks)] = C.GoString(vs)
	}
	return res, nil
}

// Retrieve property value of given key
// returns "" if not found or error
func (jvm *JavaVM) GetSystemProperty(key string) string {
	var vs *C.char
	ks := C.CString(key)
	defer C.free(unsafe.Pointer(ks))
	if e := jvm.jvmti.getSystemProperty(ks, &vs); e != JVMTI_ERROR_NONE {
		return ""
	}
	defer deallocate(vs)
	return C.GoString(vs)
}

// Set the value of target system property
// only allowed to be called at agent OnLoad phase
// According to https://docs.oracle.com/javase/8/docs/platform/jvmti/jvmti.html#GetSystemProperty
func (jvm *JavaVM) SetSystemProperty(key, value string) error {
	ck := C.CString(key)
	cv := C.CString(value)
	defer func() {
		C.free(unsafe.Pointer(ck))
		C.free(unsafe.Pointer(cv))
	}()
	if e := jvm.jvmti.setSystemProperty(ck, cv); e != JVMTI_ERROR_NONE {
		return fmt.Errorf("Failed to set property, error=%s\n", describeJvmtiError(int(e)))
	}
	return nil
}

type Frame struct {
	PC     uintptr
	Func   string
	Source string
}

func (f *Frame) String() string {
	return fmt.Sprintf("%s@%p(%s)", f.Func, unsafe.Pointer(f.PC), f.Source)
}

type Thread struct {
	jt          C.jobject
	state       int
	IsDaemon    bool
	Name        string
	StackTraces []Frame
}

func (t *Thread) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Thread@%p: %s, is_daemon=%v\n", unsafe.Pointer(t.jt), t.Name, t.IsDaemon))
	for i, fr := range t.StackTraces {
		sb.WriteString(fmt.Sprintf("\t[%d] %s\n", i, fr.String()))
	}
	return sb.String()
}

func (jvm *JavaVM) fillThread(jt C.jobject) (rt *Thread, err error) {
	rt = &Thread{
		jt: jt,
	}
	ti := &C.struct__jvmtiThreadInfo{}
	if e := jvm.jvmti.getThreadInfo(rt.jt, ti); e != JVMTI_ERROR_NONE {
		return nil, fmt.Errorf("JVMTI GetThreadInfo returns %s", describeJvmtiError(int(e)))
	}
	rt.IsDaemon = (ti.is_daemon == C.JNI_TRUE)
	rt.Name = C.GoString(ti.name)

	var tst C.jint
	if e := jvm.jvmti.getThreadState(rt.jt, &tst); e != JVMTI_ERROR_NONE {
		return nil, fmt.Errorf("JVMTI GetThreadState returns %s", describeJvmtiError(int(e)))
	}
	rt.state = int(tst)
	rt.StackTraces, err = jvm.stackTraceOf(jt)
	return
}

func (jvm *JavaVM) stackTraceOf(jt C.jobject) (frames []Frame, err error) {
	var nfrs C.jint
	var jfrs *C.struct__jvmtiFrameInfo
	flmt := C.jint(1024)
	if e := jvm.jvmti.allocate(C.jlong(C.sizeof_struct__jvmtiFrameInfo)*C.jlong(flmt), (**C.uchar)(unsafe.Pointer(&jfrs))); e != JVMTI_ERROR_NONE {
		defer deallocate(jfrs)
	}
	if e := jvm.jvmti.getStackTrace(jt, C.jint(0), flmt, jfrs, &nfrs); e != JVMTI_ERROR_NONE {
		return nil, fmt.Errorf("JVMTI GetStackTrace failed with %s", describeJvmtiError(int(e)))
	}
	for i := 0; i < int(nfrs); i++ {
		p := elemAt(jfrs, i)
		//jloc := p.location
		var name *C.char
		var sig *C.char
		var gen *C.char
		if e := jvm.jvmti.getMethodName(p.method, &name, &sig, &gen); e != JVMTI_ERROR_NONE {
			return nil, fmt.Errorf("JVMTI GetMethodName failed with %s", describeJvmtiError(int(e)))
		}
		defer func() {
			deallocate(name)
			deallocate(sig)
			deallocate(gen)
		}()

		fr := Frame{
			PC:   uintptr(p.location),
			Func: fmt.Sprintf("%s:%s:%s", C.GoString(name), C.GoString(sig), C.GoString(gen)),
		}
		frames = append(frames, fr)
	}
	return
}

// retrieve current thread, caller goroutine must pin thread before invocation
func (jvm *JavaVM) CurrentThread() (rt *Thread, err error) {
	var jt C.jobject
	if e := jvm.jvmti.getCurrentThread(&jt); e != JVMTI_ERROR_NONE {
		return nil, fmt.Errorf("JVMTI GetCurrentThread returns %s", describeJvmtiError(int(e)))
	}
	return jvm.fillThread(jt)
}

// Dump all the threads
func (jvm *JavaVM) DumpThreads() (thrds []*Thread, err error) {
	var jts *C.jobject
	var nt C.jint
	if e := jvm.jvmti.getAllThreads(&nt, &jts); e != JVMTI_ERROR_NONE {
		return nil, fmt.Errorf("JVMTI GetAllThreads returns %s", describeJvmtiError(int(e)))
	}
	defer deallocate(jts)
	for i := 0; i < int(nt); i++ {
		p := elemAt(jts, i)
		if t, e := jvm.fillThread(*p); e != nil {
			return nil, e
		} else {
			thrds = append(thrds, t)
		}
	}
	return
}


// get all the loaded classes
func (jvm *JavaVM) GetLoadedClasses() (classes []string, err error) {
	var cls *C.jclass
	var n C.jint
	if e := jvm.jvmti.getLoadedClasses(&n, &cls); e == JVMTI_ERROR_NONE {
		defer deallocate(cls)
		for i := 0; i < int(n); i++ {
			c := elemAt(cls, i)
			var sig *C.char
			var gen *C.char
			if e := jvm.jvmti.getClassSignature(c, &sig, &gen); e == JVMTI_ERROR_NONE {
				defer func() {
					deallocate(sig)
					deallocate(gen)
				}()
				s := fmt.Sprintf("Class: %s (%s)", C.GoString(sig), C.GoString(gen))
				classes = append(classes, s)
			}
		}
	}
	return
}

// helper to simplify the deallcoation of jvmti resource
func deallocate[T any](p *T) {
	jvm,e := contextVM()
	if e != nil {
		panic(e)
	}
	jvm.jvmti.deallocate((*C.uchar)(unsafe.Pointer(p)))
}

// helper to simplify the address manipulation of C memory
func elemAt[T any](base *T, idx int) T {
	return *(*T)(unsafe.Pointer((uintptr(unsafe.Pointer(base)) + uintptr(idx)*ptrSize)))
}

func addrAt[T any](base *T, idx int) *T {
	return (*T)(unsafe.Pointer((uintptr(unsafe.Pointer(base)) + uintptr(idx)*ptrSize)))
}
