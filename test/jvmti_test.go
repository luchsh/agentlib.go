//
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
//

package main_test

import(
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"strconv"
	"testing"
)

const dataDir = "testdata"

var (
	javaHome string
	javaExec string
	javaMajorVer int
)

func TestBasic(t *testing.T) {
	dir,err := ioutil.TempDir(os.TempDir(), "_jvmti_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory")
	} else {
		t.Logf("temp directory created: %s", dir)
	}
	defer os.RemoveAll(dir)

	setup(t)

	c,_ := ioutil.ReadFile(dataDir+"/Loop.java")
	ioutil.WriteFile(dir+"/Loop.java", c, 0644)
	cmd := exec.Command("javac", dir+"/Loop.java")
	if out,err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to compile Loop.java\n%s\n",string(out))
	}

	cwd,_ := os.Getwd()
	defer os.Chdir(cwd)
	os.Chdir(path.Dir(cwd)+"/src")
	cmd = exec.Command("java", "-agentlib:jnigo=hello_options", "-cp", dir, "Loop")
	if out,err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to run java with agent\n%s\n",string(out))
	} else {
		t.Logf("passed:\n%s\n", string(out))
	}
}

// various setup for the testing
func setup(t *testing.T) {
	// detect if java command presents
	s := os.Getenv("JAVA_HOME")
	if fi, err := os.Lstat("some-filename"); err == nil {
		if m := fi.Mode(); m.IsDir() {
			javaHome = s
			javaExec = javaHome+"/bin/java"
		}
	}
	/*
	javaHome = path.Dir(javaExec)
	cmd := exec.Command("which", "java")
	if out,err := cmd.Output(); err == nil {
		javaExec = string(out)
	} else {
		t.Fatalf("Failed to execute 'which java'")
	}
	*/

	// check if it is a valid javaexec
	cmd := exec.Command("java", "-version")
	if out,err := cmd.CombinedOutput(); err == nil {
		for _,l := range strings.Split(string(out), "\n") {
			if strings.HasPrefix(l, "java version") {
				s := strings.Split(l, " ")[2]
				s = strings.ReplaceAll(s, "\"", "")
				s = strings.Split(s, ".")[2]
				if v,e := strconv.ParseInt(s, 10, 32); e != nil {
					javaMajorVer = int(v)
				}
			}
		}
	} else {
		t.Fatalf("Cannot run 'java -version'\n%s\n", string(out))
	}

	// copy testlib to 'src'
	var projRoot string
	cwd,err := os.Getwd()
	if err == nil {
		projRoot = path.Dir(cwd)
		srcDir := projRoot+"/src"
		if c,err := ioutil.ReadFile(dataDir+"/test_lib.go"); err == nil {
			err = ioutil.WriteFile(srcDir+"/test_lib.go", c, 0644)
			if err != nil {
				t.Fatalf("Failed to write to %s/test_lib.go", srcDir)
			}
		} else {
			t.Fatalf("failed to read content of %s/test_lib.go",dataDir)
		}
	} else {
		t.Fatalf("Failed to get cwd")
	}

	// build the dynamic library
	os.Chdir(projRoot)
	defer os.Chdir(cwd)
	cmd = exec.Command("sh", "-x", "build.sh")
	if out,err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to compile the library:\n%s\n", string(out))
	} else {
		t.Logf("%s", string(out))
	}
}
