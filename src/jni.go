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
package main

// #include "wrapper.h"
import "C"

const (
	JNI_VERSION_1_1 = 0x00010001
	JNI_VERSION_1_2 = 0x00010002
)

// JniEnv conrresponds to JNIEnv*
type JniEnv uintptr

// GetVersion corresponding to
// jint GetVersion(JNIEnv *env);
func (jni JniEnv) GetVersion() int32 {
	return JNI_VERSION_1_2
}
