
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