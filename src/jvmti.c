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

#include "jvmti.h"
#include <stdio.h>

// jvmtiError GetClassSignature(jvmtiEnv* env, jclass klass, char** signature_ptr, char** generic_ptr)
int GetClassSignature(void* jvmti, void* clazz, void* sigptr, void* genptr) {
  jvmtiEnv* env = (jvmtiEnv*)jvmti;
  jclass cls = (jclass)clazz;
  char** sig_ptr = (char**)sigptr;
  char** gen_ptr = (char**)genptr;
  return (int)(*env)->GetClassSignature(env, cls, sig_ptr, gen_ptr);
}

// jvmtiError Allocate(jvmtiEnv* env, jlong size, unsigned char** mem_ptr)
int Allocate(void* jvmti, long long size, void* ptr) {
  jvmtiEnv* env = (jvmtiEnv*)jvmti;
  jlong jsize = (jlong)size;
  return (int)(*env)->Allocate(env, jsize, (unsigned char**)ptr);
}

// jvmtiError Deallocate(jvmtiEnv* env, unsigned char* mem)
int Deallocate(void* jvmti, void* mem) {
  jvmtiEnv* env = (jvmtiEnv*)jvmti;
  unsigned char* p = (unsigned char*)mem;
  return(int)(*env)->Deallocate(env, p);
}