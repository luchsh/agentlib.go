
#ifndef __JVMTI_WRAPPER_H__
#define __JVMTI_WRAPPER_H__

#include "jvmti.h"

#ifdef __cplusplus
extern "C" {
#endif

int GetClassSignature(void* jvmti, void* clazz, void* sigptr, void* genptr);
int Deallocate(void* jvmti, void* mem);

#ifdef __cplusplus
}
#endif

#endif // #ifndef __JVMTI_WRAPPER_H__