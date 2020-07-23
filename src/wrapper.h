//
// Copyright 2020 Lu Chuan Sheng
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

//
// The C code here provides a thin bridge between Go runtime
// and the JNI, JVMTI interface.
//
// We need this wrapper layer for two reasons:
// 1. To work with cgo's inability of calling function pointer fields of C struct.
//    So far (by Go 1.14) this still cannot be done easily, might be fixed in the future.
//    We may consider removing this layer then.
// 2. To create JNI compatible thread state
//    JNIEnv is threadlocal object, and expected to be accessed only from 'JavaThread',
//    but for a shared library writen in Go, which currently ships a full copy Go runtime by default,
//    its carrier threads, AKA 'm', are not 'JavaThread'.
//

#ifndef __WRAPPER_H__
#define __WRAPPER_H__

#include <jni.h>
#include <jvmti.h>

#ifdef __cplusplus
extern "C" {
#endif

/*
 * Entry point of the library following the JVMTI specification
 */
jint Agent_OnLoad(JavaVM* javaVM, char* options, void* reserved);


#ifdef __cplusplus
}
#endif

#endif // #ifndef __WRAPPER_H__