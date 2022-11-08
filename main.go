package main

//#include <jni.h>
//#include <jvmti.h>
import "C"

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	_ "net/http/pprof"

	. "github.com/luchsh/agentlib.go/jpprof"
)

var (
	vm JVM
)

func runJava() {
	var jni JniEnv
	vm,jni = JniCreateJavaVM("-Djava.class.path=./testdata Loop")
	jni.Env().FindClass("Ljava/lang/String;")
}

func jniHandleVersion(w http.ResponseWriter, req *http.Request) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	env,_ := vm.VM().AttachCurrentThread()
	defer vm.VM().DetachCurrentThread()
	v := env.GetVersion()
	w.Write([]byte(fmt.Sprintf("JNI version: %d\n", v)))
}

/*
func getLoadedClasses(w http.ResponseWriter, req *http.Request) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	env := JvmtiEnv(vm.GetEnv(JVMTI_VERSION_1_1))

	var nofCls C.jint
	var clss *C.jclass
	if e := env.GetLoadedClasses(&nofCls, &clss); e != JVMTI_ERROR_NONE {
		panic(e)
	}
}
*/

func main() {
	go runJava()
	http.HandleFunc("/debug/jni/version", jniHandleVersion) 
	//http.HandleFunc("/debug/jvmti/loadedclasses", getLoadedClasses) 
	log.Fatal(http.ListenAndServe(":8081", nil))
}
