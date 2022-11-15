package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"runtime"

	_ "net/http/pprof"

	"github.com/luchsh/agentlib.go/jgo"
)

func handleIndex(w http.ResponseWriter, req *http.Request) {
	var b bytes.Buffer
	b.WriteString(`<html>
 <head>
 <title>/debug/java/</title>
 </head>
 <body>
 /debug/java/<br>
 <br>
 <li><a href="/debug/java/version">Java version</a></li>
 <li><a href="/debug/java/threads">All Java threads</a></li>
 </body>
 </html>`)
	w.Write(b.Bytes())
}

func init() {
	var err error
	_, err = jgo.Exec([]string{"-Djava.class.path=./testdata", "loop"})
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/debug/java/version", func(w http.ResponseWriter, req *http.Request) {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		vm := jgo.CurrentVM()
		w.Write([]byte(vm.FullVersion()))
	})

	http.HandleFunc("/debug/java/threads", func(w http.ResponseWriter, req *http.Request) {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		vm := jgo.CurrentVM()

		thrds, e := vm.DumpThreads()
		if e != nil {
			s := fmt.Sprintf("Failed to DumpThreads: %v", e)
			w.Write([]byte(s))
			return
		}
		for i, th := range thrds {
			s := fmt.Sprintf("Thread-%d %s\n", i, th.String())
			w.Write([]byte(s))
		}
	})
	http.HandleFunc("/debug/java", handleIndex)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
