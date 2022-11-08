package main

import "github.com/luchsh/agentlib.go/jpprof"

func main() {
	println("agentlib.go")
	jpprof.Onload = func(lib *jpprof.AgentLib) {
		println(lib.Options)
	}
}
