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

// generate wrappers from jni headers

//+build ignore

package main

import (
	"fmt"
	"os"
)

var javaHome string
//var javaVersion []int // 14.0.1

func init() {
	if javaHome = os.Getenv("JAVA_HOME"); javaHome == "" {
		panic("Cannot find a valid JAVA_HOME")
	}
	if fi,e := os.Lstat(javaHome); e == nil {
		if !fi.IsDir() {
			panic(fmt.Errorf("JAVA_HOME=%s is not a directory", javaHome))
		}
	} else {
		panic(e)
	}
}

func main() {
}
