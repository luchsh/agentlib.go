package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	_ "net/http/pprof"

	"github.com/luchsh/agentlib.go/jgo"
)

func main() {
	var err error
	_,err = jgo.Exec([]string{"-Djava.class.path=./testdata", "loop"})
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/debug/java/version", func(w http.ResponseWriter, req *http.Request) {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		vm := jgo.CurrentVM()
		w.Write([]byte(vm.FullVersion()))
	})

	http.HandleFunc("/debug/java/thread", func(w http.ResponseWriter, req *http.Request) {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		vm := jgo.CurrentVM()

		thrds,e := vm.DumpThreads()
		if e != nil {
			s := fmt.Sprintf("Failed to DumpThreads: %v", e)
			w.Write([]byte(s))
			return
		}
		for i,th := range thrds {
			s := fmt.Sprintf("Thread-%d %s\n", i, th.String())
			w.Write([]byte(s))
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
