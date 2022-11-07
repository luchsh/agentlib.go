// Copyright 2020 chuanshenglu@gmail.com
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

// +build ignore

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var (
	javaHome string
)

func init() {
	if javaHome = os.Getenv("JAVA_HOME"); javaHome == "" {
		panic("JAVA_HOME env not found!")
	}
}

type jvmtiPipe struct {
	out    io.Writer
	in     io.Reader
	consts [][]string // constatn definition, converted from enum
	fns    []fn
}

// parsed function info
type fn struct {
	name  string
	ret   string
	ptype []string // params types
	pname []string // params arg names
}

// e.g.
//  jvmtiError (JNICALL *SetEventNotificationMode) (jvmtiEnv* env,
//    jvmtiEventMode mode,
//    jvmtiEvent event_type,
//    jthread event_thread,
//     ...);
func parseJvmtiFuncDecl(s string) fn {
	var f fn
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "(JNICALL *", "", 1)
	s = strings.Replace(s, ") (", " (", 1)
	ilb := strings.Index(s, "(")
	irb := strings.Index(s, ")")
	l := s[0:ilb]
	params := s[ilb+1 : irb]
	fds := strings.Fields(l)
	f.ret = fds[0]
	f.name = fds[1]

	for _, p := range strings.Split(params, ",") {
		fds := strings.Fields(p)
		if len(fds) == 1 {
			// SKIP ...
			//f.ptype = append(f.ptype, fds[0])
			//f.pname = append(f.pname, "")
		} else if len(fds) >= 2 {
			f.ptype = append(f.ptype, strings.Join(fds[0:len(fds)-1], " "))
			f.pname = append(f.pname, fds[len(fds)-1])
		}
	}
	return f
}

func (f *fn) cdecl() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s %s(", f.ret, f.name))
	for i := 0; i < len(f.ptype); i++ {
		sb.WriteString(fmt.Sprintf("%s %s", f.ptype[i], f.pname[i]))
		if i < len(f.ptype)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")")
	return sb.String()
}

// converted cgo
func (f *fn) cgoimpl(lnpfx string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%sstatic ", lnpfx))
	sb.WriteString(f.cdecl())
	sb.WriteString(" {\n")
	sb.WriteString(fmt.Sprintf("%s  return (*env)->%s(", lnpfx, f.name))
	for i := 0; i < len(f.ptype); i++ {
		if len(f.pname[i]) > 0 {
			sb.WriteString(f.pname[i])
		} else {
			sb.WriteString(f.ptype[i])
		}
		if i < len(f.ptype)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(fmt.Sprintf(");\n%s}", lnpfx))
	return sb.String()
}

// convert C type decl to Go
func goTypeOfC(ct string) string {
	ct = strings.Replace(ct, "const ", "", 1)
	famousTypes := map[string]string {
		"char*": "*C.char",
		"char**": "**C.char",
		"unsigned char**": "**C.uchar",
		"unsigned char*": "*C.uchar",
		"void*": "unsafe.Pointer",
		"void**": "*unsafe.Pointer",
		"...":"[]interface{}",
	}
	if v,ok := famousTypes[ct]; ok {
		return v
	}
	fds := strings.Fields(ct)
	if len(fds) == 1 {
		ct = fmt.Sprintf("C.%s", fds[0])
	}
	for strings.HasSuffix(ct, "*") {
		ct = ct[0:len(ct)-1]
		ct = "*"+ct
	}
	return ct
}

func (f *fn) goimpl() string {
	var sb strings.Builder
	//sb.WriteString("//TODO: manual adjustment needed here\n")
	sb.WriteString(fmt.Sprintf("func (jvmti JvmtiEnv) %s(", f.name))
	for i := 1; i < len(f.ptype); i++ {
		if len(f.pname[i]) > 0 {
			sb.WriteString(f.pname[i])
		} else {
			sb.WriteString("args")
		}
		sb.WriteString(fmt.Sprintf(" %s", goTypeOfC(f.ptype[i])))
		if i < len(f.ptype)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(fmt.Sprintf(") C.%s {\n", f.ret))
	sb.WriteString(fmt.Sprintf("  return C.%s(", f.name))
	for i := 0; i < len(f.ptype); i++ {
		if i == 0 {
			sb.WriteString("jvmti.raw(), ")
			continue
		}
		if len(f.pname[i]) > 0 {
			sb.WriteString(f.pname[i])
		} else {
			sb.WriteString("args")
		}
		if i < len(f.ptype)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString(")\n}\n");
	return sb.String()
}

func (jp *jvmtiPipe) parse() {
	rd := bufio.NewReader(jp.in)
	// load enum consts
	for {
		bs, _, e := rd.ReadLine()
		if e == io.EOF {
			return
		}
		ln := string(bs)
		if strings.Contains(ln, "enum {") {
			var cconsts []string
			for {
				bs, _, e := rd.ReadLine()
				ln := string(bs)
				if e == io.EOF || strings.Contains(ln, "}") {
					break
				}
				cconsts = append(cconsts, ln)
			}
			jp.consts = append(jp.consts, cconsts)
		} else if strings.HasPrefix(ln, "typedef struct jvmtiInterface_1_ {") {
			for {
				bs, _, e := rd.ReadLine()
				ln := string(bs)
				if e == io.EOF || strings.Contains(ln, "}") {
					break
				}

				if strings.Contains(ln, "(JNICALL") && strings.Contains(ln, "(jvmtiEnv* env,") {
					var sb strings.Builder
					sb.WriteString(ln)
					for {
						bs, _, e := rd.ReadLine()
						if e != nil {
							panic("e")
						}
						ln := string(bs)
						sb.WriteString(ln)
						if strings.Contains(ln, ")") {
							break
						}
					}
					jp.fns = append(jp.fns, parseJvmtiFuncDecl(sb.String()))
				}
			}
		}
	}
}

func (jp *jvmtiPipe) printCgoWrapper() {
	for _, fn := range jp.fns {
		fmt.Fprintf(jp.out, "%s\n", fn.cgoimpl("// "))
	}
}

func (jp *jvmtiPipe) printGoWrapper() {
	for _, fn := range jp.fns {
		fmt.Fprintf(jp.out, "%s\n", fn.goimpl())
	}
}

func (jp *jvmtiPipe) printJvmtiDef() {
	fmt.Fprintf(jp.out, "%s\n",
`// JvmtiEnv corresponds to jvmtiEnv*
type JvmtiEnv uintptr

func (jvmti JvmtiEnv) raw() *C.jvmtiEnv {
	return (*C.jvmtiEnv)(unsafe.Pointer(jvmti))
}

func (jvmti JvmtiEnv) asPointer() unsafe.Pointer {
	return unsafe.Pointer(jvmti)
}`)
}

func (jp *jvmtiPipe) print() {
	jp.printHeader()
	jp.printCgoWrapper()
	fmt.Fprintf(jp.out, "import \"C\"\n\n")
	fmt.Fprintf(jp.out, "import \"unsafe\"\n\n")

	// consts
	for _, cc := range jp.consts {
		fmt.Fprintf(jp.out, "const (\n")
		for _, c := range cc {
			if strings.HasSuffix(c, ",") {
				c = c[0 : len(c)-1]
			}
			fmt.Fprintf(jp.out, "%s\n", c)
		}
		fmt.Fprintf(jp.out, ")\n\n")
	}

	fmt.Fprintf(jp.out, "\n\n")
	jp.printJvmtiDef()
	fmt.Fprintf(jp.out, "\n\n")
	jp.printGoWrapper()
}

func (jp *jvmtiPipe) Pump() {
	jp.parse()
	jp.print()
}

func (jp *jvmtiPipe) printHeader() {
	fmt.Fprintf(jp.out, "// Copyright %d chuanshenglu@gmail.com\n", time.Now().Year())
	fmt.Fprintf(jp.out, `//
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
// limitations under the License.\n\n

package jpprof

// #include <jvmti.h>
//
`)
}

func main() {
	jvmtiH := fmt.Sprintf("%s/include/jvmti.h", javaHome)
	fin, err := os.Open(jvmtiH)
	if err != nil {
		panic(err)
	}

	jp := &jvmtiPipe{
		out: os.Stdout,
		in:  fin,
	}
	jp.Pump()
}
