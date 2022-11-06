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
package jpprof

// #include "wrapper.h"
// #include "jvmti_wrapper.h"
// #include <jvmti.h>
import "C"

import (
	"unsafe"
)

// JvmtiEnv corresponds to jvmtiEnv*
type JvmtiEnv uintptr

func (jvmti JvmtiEnv) raw() *C.jvmtiEnv {
	return (*C.jvmtiEnv)(unsafe.Pointer(jvmti))
}

func (jvmti JvmtiEnv) asPointer() unsafe.Pointer {
	return unsafe.Pointer(jvmti)
}


/*   2 : Set Event Notification Mode */
// jvmtiError (JNICALL *SetEventNotificationMode) (jvmtiEnv* env,
//  jvmtiEventMode mode,
//  jvmtiEvent event_type,
//  jthread event_thread,
//   ...);
//func (jvmti JvmtiEnv) SetEventNotificationMode(mode C.jvmtiEventMode, event_type C.jvmtiEvent, event_thread C.jthread, args []interface{}) C.jvmtiError {
//	return C.SetEventNotificationMode((*C.jvmtiEnv)(jvmti), mode, event_type, event_thread)
//}

//  /*   3 : Get All Modules */
//  jvmtiError (JNICALL *GetAllModules) (jvmtiEnv* env,
//    jint* module_count_ptr,
//    jobject** modules_ptr);
//func (jvmti JvmtiEnv) GetAllModules(module_counter_ptr *C.jint, modules_ptr **C.jobject) C.jvmtiError {
//	return C.GetAllModules((*C.jvmtiEnv)(jvmti), module_counter_ptr, modules_ptr)
//}

//  /*   4 : Get All Threads */
//  jvmtiError (JNICALL *GetAllThreads) (jvmtiEnv* env,
//    jint* threads_count_ptr,
//    jthread** threads_ptr);
func (jvmti JvmtiEnv) getAllThreads(thread_count_ptr *C.jint, threads_ptr **C.jthread) C.jvmtiError {
	return C.jvmtiError(C.GetAllThreads(jvmti.raw(), thread_count_ptr, threads_ptr))
}

//  /*   5 : Suspend Thread */
//  jvmtiError (JNICALL *SuspendThread) (jvmtiEnv* env,
//    jthread thread);
//
//  /*   6 : Resume Thread */
//  jvmtiError (JNICALL *ResumeThread) (jvmtiEnv* env,
//    jthread thread);
//
//  /*   7 : Stop Thread */
//  jvmtiError (JNICALL *StopThread) (jvmtiEnv* env,
//    jthread thread,
//    jobject exception);
//
//  /*   8 : Interrupt Thread */
//  jvmtiError (JNICALL *InterruptThread) (jvmtiEnv* env,
//    jthread thread);
//
//  /*   9 : Get Thread Info */
//  jvmtiError (JNICALL *GetThreadInfo) (jvmtiEnv* env,
//    jthread thread,
//    jvmtiThreadInfo* info_ptr);
//
//  /*   10 : Get Owned Monitor Info */
//  jvmtiError (JNICALL *GetOwnedMonitorInfo) (jvmtiEnv* env,
//    jthread thread,
//    jint* owned_monitor_count_ptr,
//    jobject** owned_monitors_ptr);
//
//  /*   11 : Get Current Contended Monitor */
//  jvmtiError (JNICALL *GetCurrentContendedMonitor) (jvmtiEnv* env,
//    jthread thread,
//    jobject* monitor_ptr);
//
//  /*   12 : Run Agent Thread */
//  jvmtiError (JNICALL *RunAgentThread) (jvmtiEnv* env,
//    jthread thread,
//    jvmtiStartFunction proc,
//    const void* arg,
//    jint priority);
//
//  /*   13 : Get Top Thread Groups */
//  jvmtiError (JNICALL *GetTopThreadGroups) (jvmtiEnv* env,
//    jint* group_count_ptr,
//    jthreadGroup** groups_ptr);
//
//  /*   14 : Get Thread Group Info */
//  jvmtiError (JNICALL *GetThreadGroupInfo) (jvmtiEnv* env,
//    jthreadGroup group,
//    jvmtiThreadGroupInfo* info_ptr);
//
//  /*   15 : Get Thread Group Children */
//  jvmtiError (JNICALL *GetThreadGroupChildren) (jvmtiEnv* env,
//    jthreadGroup group,
//    jint* thread_count_ptr,
//    jthread** threads_ptr,
//    jint* group_count_ptr,
//    jthreadGroup** groups_ptr);
//
//  /*   16 : Get Frame Count */
//  jvmtiError (JNICALL *GetFrameCount) (jvmtiEnv* env,
//    jthread thread,
//    jint* count_ptr);
//
//  /*   17 : Get Thread State */
//  jvmtiError (JNICALL *GetThreadState) (jvmtiEnv* env,
//    jthread thread,
//    jint* thread_state_ptr);
//
//  /*   18 : Get Current Thread */
//  jvmtiError (JNICALL *GetCurrentThread) (jvmtiEnv* env,
//    jthread* thread_ptr);
//
//  /*   19 : Get Frame Location */
//  jvmtiError (JNICALL *GetFrameLocation) (jvmtiEnv* env,
//    jthread thread,
//    jint depth,
//    jmethodID* method_ptr,
//    jlocation* location_ptr);
//
//  /*   20 : Notify Frame Pop */
//  jvmtiError (JNICALL *NotifyFramePop) (jvmtiEnv* env,
//    jthread thread,
//    jint depth);
//
//  /*   21 : Get Local Variable - Object */
//  jvmtiError (JNICALL *GetLocalObject) (jvmtiEnv* env,
//    jthread thread,
//    jint depth,
//    jint slot,
//    jobject* value_ptr);
//
//  /*   22 : Get Local Variable - Int */
//  jvmtiError (JNICALL *GetLocalInt) (jvmtiEnv* env,
//    jthread thread,
//    jint depth,
//    jint slot,
//    jint* value_ptr);
//
//  /*   23 : Get Local Variable - Long */
//  jvmtiError (JNICALL *GetLocalLong) (jvmtiEnv* env,
//    jthread thread,
//    jint depth,
//    jint slot,
//    jlong* value_ptr);
//
//  /*   24 : Get Local Variable - Float */
//  jvmtiError (JNICALL *GetLocalFloat) (jvmtiEnv* env,
//    jthread thread,
//    jint depth,
//    jint slot,
//    jfloat* value_ptr);
//
//  /*   25 : Get Local Variable - Double */
//  jvmtiError (JNICALL *GetLocalDouble) (jvmtiEnv* env,
//    jthread thread,
//    jint depth,
//    jint slot,
//    jdouble* value_ptr);
//
//  /*   26 : Set Local Variable - Object */
//  jvmtiError (JNICALL *SetLocalObject) (jvmtiEnv* env,
//    jthread thread,
//    jint depth,
//    jint slot,
//    jobject value);
//
//  /*   27 : Set Local Variable - Int */
//  jvmtiError (JNICALL *SetLocalInt) (jvmtiEnv* env,
//    jthread thread,
//    jint depth,
//    jint slot,
//    jint value);
//
//  /*   28 : Set Local Variable - Long */
//  jvmtiError (JNICALL *SetLocalLong) (jvmtiEnv* env,
//    jthread thread,
//    jint depth,
//    jint slot,
//    jlong value);
//
//  /*   29 : Set Local Variable - Float */
//  jvmtiError (JNICALL *SetLocalFloat) (jvmtiEnv* env,
//    jthread thread,
//    jint depth,
//    jint slot,
//    jfloat value);
//
//  /*   30 : Set Local Variable - Double */
//  jvmtiError (JNICALL *SetLocalDouble) (jvmtiEnv* env,
//    jthread thread,
//    jint depth,
//    jint slot,
//    jdouble value);
//
//  /*   31 : Create Raw Monitor */
//  jvmtiError (JNICALL *CreateRawMonitor) (jvmtiEnv* env,
//    const char* name,
//    jrawMonitorID* monitor_ptr);
//
//  /*   32 : Destroy Raw Monitor */
//  jvmtiError (JNICALL *DestroyRawMonitor) (jvmtiEnv* env,
//    jrawMonitorID monitor);
//
//  /*   33 : Raw Monitor Enter */
//  jvmtiError (JNICALL *RawMonitorEnter) (jvmtiEnv* env,
//    jrawMonitorID monitor);
//
//  /*   34 : Raw Monitor Exit */
//  jvmtiError (JNICALL *RawMonitorExit) (jvmtiEnv* env,
//    jrawMonitorID monitor);
//
//  /*   35 : Raw Monitor Wait */
//  jvmtiError (JNICALL *RawMonitorWait) (jvmtiEnv* env,
//    jrawMonitorID monitor,
//    jlong millis);
//
//  /*   36 : Raw Monitor Notify */
//  jvmtiError (JNICALL *RawMonitorNotify) (jvmtiEnv* env,
//    jrawMonitorID monitor);
//
//  /*   37 : Raw Monitor Notify All */
//  jvmtiError (JNICALL *RawMonitorNotifyAll) (jvmtiEnv* env,
//    jrawMonitorID monitor);
//
//  /*   38 : Set Breakpoint */
//  jvmtiError (JNICALL *SetBreakpoint) (jvmtiEnv* env,
//    jmethodID method,
//    jlocation location);
//
//  /*   39 : Clear Breakpoint */
//  jvmtiError (JNICALL *ClearBreakpoint) (jvmtiEnv* env,
//    jmethodID method,
//    jlocation location);
//
//  /*   40 : Get Named Module */
//  jvmtiError (JNICALL *GetNamedModule) (jvmtiEnv* env,
//    jobject class_loader,
//    const char* package_name,
//    jobject* module_ptr);
//
//  /*   41 : Set Field Access Watch */
//  jvmtiError (JNICALL *SetFieldAccessWatch) (jvmtiEnv* env,
//    jclass klass,
//    jfieldID field);
//
//  /*   42 : Clear Field Access Watch */
//  jvmtiError (JNICALL *ClearFieldAccessWatch) (jvmtiEnv* env,
//    jclass klass,
//    jfieldID field);
//
//  /*   43 : Set Field Modification Watch */
//  jvmtiError (JNICALL *SetFieldModificationWatch) (jvmtiEnv* env,
//    jclass klass,
//    jfieldID field);
//
//  /*   44 : Clear Field Modification Watch */
//  jvmtiError (JNICALL *ClearFieldModificationWatch) (jvmtiEnv* env,
//    jclass klass,
//    jfieldID field);
//
//  /*   45 : Is Modifiable Class */
//  jvmtiError (JNICALL *IsModifiableClass) (jvmtiEnv* env,
//    jclass klass,
//    jboolean* is_modifiable_class_ptr);
//
//  /*   46 : Allocate */
//  jvmtiError (JNICALL *Allocate) (jvmtiEnv* env,
//    jlong size,
//    unsigned char** mem_ptr);
//
//  /*   47 : Deallocate */
//  jvmtiError (JNICALL *Deallocate) (jvmtiEnv* env,
//    unsigned char* mem);
//
//  /*   48 : Get Class Signature */
//  jvmtiError (JNICALL *GetClassSignature) (jvmtiEnv* env,
//    jclass klass,
//    char** signature_ptr,
//    char** generic_ptr);
//
//  /*   49 : Get Class Status */
//  jvmtiError (JNICALL *GetClassStatus) (jvmtiEnv* env,
//    jclass klass,
//    jint* status_ptr);
//
//  /*   50 : Get Source File Name */
//  jvmtiError (JNICALL *GetSourceFileName) (jvmtiEnv* env,
//    jclass klass,
//    char** source_name_ptr);
//
//  /*   51 : Get Class Modifiers */
//  jvmtiError (JNICALL *GetClassModifiers) (jvmtiEnv* env,
//    jclass klass,
//    jint* modifiers_ptr);
//
//  /*   52 : Get Class Methods */
//  jvmtiError (JNICALL *GetClassMethods) (jvmtiEnv* env,
//    jclass klass,
//    jint* method_count_ptr,
//    jmethodID** methods_ptr);
//
//  /*   53 : Get Class Fields */
//  jvmtiError (JNICALL *GetClassFields) (jvmtiEnv* env,
//    jclass klass,
//    jint* field_count_ptr,
//    jfieldID** fields_ptr);
//
//  /*   54 : Get Implemented Interfaces */
//  jvmtiError (JNICALL *GetImplementedInterfaces) (jvmtiEnv* env,
//    jclass klass,
//    jint* interface_count_ptr,
//    jclass** interfaces_ptr);
//
//  /*   55 : Is Interface */
//  jvmtiError (JNICALL *IsInterface) (jvmtiEnv* env,
//    jclass klass,
//    jboolean* is_interface_ptr);
//
//  /*   56 : Is Array Class */
//  jvmtiError (JNICALL *IsArrayClass) (jvmtiEnv* env,
//    jclass klass,
//    jboolean* is_array_class_ptr);
//
//  /*   57 : Get Class Loader */
//  jvmtiError (JNICALL *GetClassLoader) (jvmtiEnv* env,
//    jclass klass,
//    jobject* classloader_ptr);
//
//  /*   58 : Get Object Hash Code */
//  jvmtiError (JNICALL *GetObjectHashCode) (jvmtiEnv* env,
//    jobject object,
//    jint* hash_code_ptr);
//
//  /*   59 : Get Object Monitor Usage */
//  jvmtiError (JNICALL *GetObjectMonitorUsage) (jvmtiEnv* env,
//    jobject object,
//    jvmtiMonitorUsage* info_ptr);
//
//  /*   60 : Get Field Name (and Signature) */
//  jvmtiError (JNICALL *GetFieldName) (jvmtiEnv* env,
//    jclass klass,
//    jfieldID field,
//    char** name_ptr,
//    char** signature_ptr,
//    char** generic_ptr);
//
//  /*   61 : Get Field Declaring Class */
//  jvmtiError (JNICALL *GetFieldDeclaringClass) (jvmtiEnv* env,
//    jclass klass,
//    jfieldID field,
//    jclass* declaring_class_ptr);
//
//  /*   62 : Get Field Modifiers */
//  jvmtiError (JNICALL *GetFieldModifiers) (jvmtiEnv* env,
//    jclass klass,
//    jfieldID field,
//    jint* modifiers_ptr);
//
//  /*   63 : Is Field Synthetic */
//  jvmtiError (JNICALL *IsFieldSynthetic) (jvmtiEnv* env,
//    jclass klass,
//    jfieldID field,
//    jboolean* is_synthetic_ptr);
//
//  /*   64 : Get Method Name (and Signature) */
//  jvmtiError (JNICALL *GetMethodName) (jvmtiEnv* env,
//    jmethodID method,
//    char** name_ptr,
//    char** signature_ptr,
//    char** generic_ptr);
//
//  /*   65 : Get Method Declaring Class */
//  jvmtiError (JNICALL *GetMethodDeclaringClass) (jvmtiEnv* env,
//    jmethodID method,
//    jclass* declaring_class_ptr);
//
//  /*   66 : Get Method Modifiers */
//  jvmtiError (JNICALL *GetMethodModifiers) (jvmtiEnv* env,
//    jmethodID method,
//    jint* modifiers_ptr);
//
//  /*   67 :  RESERVED */
//  void *reserved67;
//
//  /*   68 : Get Max Locals */
//  jvmtiError (JNICALL *GetMaxLocals) (jvmtiEnv* env,
//    jmethodID method,
//    jint* max_ptr);
//
//  /*   69 : Get Arguments Size */
//  jvmtiError (JNICALL *GetArgumentsSize) (jvmtiEnv* env,
//    jmethodID method,
//    jint* size_ptr);
//
//  /*   70 : Get Line Number Table */
//  jvmtiError (JNICALL *GetLineNumberTable) (jvmtiEnv* env,
//    jmethodID method,
//    jint* entry_count_ptr,
//    jvmtiLineNumberEntry** table_ptr);
//
//  /*   71 : Get Method Location */
//  jvmtiError (JNICALL *GetMethodLocation) (jvmtiEnv* env,
//    jmethodID method,
//    jlocation* start_location_ptr,
//    jlocation* end_location_ptr);
//
//  /*   72 : Get Local Variable Table */
//  jvmtiError (JNICALL *GetLocalVariableTable) (jvmtiEnv* env,
//    jmethodID method,
//    jint* entry_count_ptr,
//    jvmtiLocalVariableEntry** table_ptr);
//
//  /*   73 : Set Native Method Prefix */
//  jvmtiError (JNICALL *SetNativeMethodPrefix) (jvmtiEnv* env,
//    const char* prefix);
//
//  /*   74 : Set Native Method Prefixes */
//  jvmtiError (JNICALL *SetNativeMethodPrefixes) (jvmtiEnv* env,
//    jint prefix_count,
//    char** prefixes);
//
//  /*   75 : Get Bytecodes */
//  jvmtiError (JNICALL *GetBytecodes) (jvmtiEnv* env,
//    jmethodID method,
//    jint* bytecode_count_ptr,
//    unsigned char** bytecodes_ptr);
//
//  /*   76 : Is Method Native */
//  jvmtiError (JNICALL *IsMethodNative) (jvmtiEnv* env,
//    jmethodID method,
//    jboolean* is_native_ptr);
//
//  /*   77 : Is Method Synthetic */
//  jvmtiError (JNICALL *IsMethodSynthetic) (jvmtiEnv* env,
//    jmethodID method,
//    jboolean* is_synthetic_ptr);
//
//  /*   78 : Get Loaded Classes */
//  jvmtiError (JNICALL *GetLoadedClasses) (jvmtiEnv* env,
//    jint* class_count_ptr,
//    jclass** classes_ptr);
//
//  /*   79 : Get Classloader Classes */
//  jvmtiError (JNICALL *GetClassLoaderClasses) (jvmtiEnv* env,
//    jobject initiating_loader,
//    jint* class_count_ptr,
//    jclass** classes_ptr);
//
//  /*   80 : Pop Frame */
//  jvmtiError (JNICALL *PopFrame) (jvmtiEnv* env,
//    jthread thread);
//
//  /*   81 : Force Early Return - Object */
//  jvmtiError (JNICALL *ForceEarlyReturnObject) (jvmtiEnv* env,
//    jthread thread,
//    jobject value);
//
//  /*   82 : Force Early Return - Int */
//  jvmtiError (JNICALL *ForceEarlyReturnInt) (jvmtiEnv* env,
//    jthread thread,
//    jint value);
//
//  /*   83 : Force Early Return - Long */
//  jvmtiError (JNICALL *ForceEarlyReturnLong) (jvmtiEnv* env,
//    jthread thread,
//    jlong value);
//
//  /*   84 : Force Early Return - Float */
//  jvmtiError (JNICALL *ForceEarlyReturnFloat) (jvmtiEnv* env,
//    jthread thread,
//    jfloat value);
//
//  /*   85 : Force Early Return - Double */
//  jvmtiError (JNICALL *ForceEarlyReturnDouble) (jvmtiEnv* env,
//    jthread thread,
//    jdouble value);
//
//  /*   86 : Force Early Return - Void */
//  jvmtiError (JNICALL *ForceEarlyReturnVoid) (jvmtiEnv* env,
//    jthread thread);
//
//  /*   87 : Redefine Classes */
//  jvmtiError (JNICALL *RedefineClasses) (jvmtiEnv* env,
//    jint class_count,
//    const jvmtiClassDefinition* class_definitions);
//
//  /*   88 : Get Version Number */
//  jvmtiError (JNICALL *GetVersionNumber) (jvmtiEnv* env,
//    jint* version_ptr);
//
//  /*   89 : Get Capabilities */
//  jvmtiError (JNICALL *GetCapabilities) (jvmtiEnv* env,
//    jvmtiCapabilities* capabilities_ptr);
//
//  /*   90 : Get Source Debug Extension */
//  jvmtiError (JNICALL *GetSourceDebugExtension) (jvmtiEnv* env,
//    jclass klass,
//    char** source_debug_extension_ptr);
//
//  /*   91 : Is Method Obsolete */
//  jvmtiError (JNICALL *IsMethodObsolete) (jvmtiEnv* env,
//    jmethodID method,
//    jboolean* is_obsolete_ptr);
//
//  /*   92 : Suspend Thread List */
//  jvmtiError (JNICALL *SuspendThreadList) (jvmtiEnv* env,
//    jint request_count,
//    const jthread* request_list,
//    jvmtiError* results);
//
//  /*   93 : Resume Thread List */
//  jvmtiError (JNICALL *ResumeThreadList) (jvmtiEnv* env,
//    jint request_count,
//    const jthread* request_list,
//    jvmtiError* results);
//
//  /*   94 : Add Module Reads */
//  jvmtiError (JNICALL *AddModuleReads) (jvmtiEnv* env,
//    jobject module,
//    jobject to_module);
//
//  /*   95 : Add Module Exports */
//  jvmtiError (JNICALL *AddModuleExports) (jvmtiEnv* env,
//    jobject module,
//    const char* pkg_name,
//    jobject to_module);
//
//  /*   96 : Add Module Opens */
//  jvmtiError (JNICALL *AddModuleOpens) (jvmtiEnv* env,
//    jobject module,
//    const char* pkg_name,
//    jobject to_module);
//
//  /*   97 : Add Module Uses */
//  jvmtiError (JNICALL *AddModuleUses) (jvmtiEnv* env,
//    jobject module,
//    jclass service);
//
//  /*   98 : Add Module Provides */
//  jvmtiError (JNICALL *AddModuleProvides) (jvmtiEnv* env,
//    jobject module,
//    jclass service,
//    jclass impl_class);
//
//  /*   99 : Is Modifiable Module */
//  jvmtiError (JNICALL *IsModifiableModule) (jvmtiEnv* env,
//    jobject module,
//    jboolean* is_modifiable_module_ptr);
//
//  /*   100 : Get All Stack Traces */
//  jvmtiError (JNICALL *GetAllStackTraces) (jvmtiEnv* env,
//    jint max_frame_count,
//    jvmtiStackInfo** stack_info_ptr,
//    jint* thread_count_ptr);
//
//  /*   101 : Get Thread List Stack Traces */
//  jvmtiError (JNICALL *GetThreadListStackTraces) (jvmtiEnv* env,
//    jint thread_count,
//    const jthread* thread_list,
//    jint max_frame_count,
//    jvmtiStackInfo** stack_info_ptr);
//
//  /*   102 : Get Thread Local Storage */
//  jvmtiError (JNICALL *GetThreadLocalStorage) (jvmtiEnv* env,
//    jthread thread,
//    void** data_ptr);
//
//  /*   103 : Set Thread Local Storage */
//  jvmtiError (JNICALL *SetThreadLocalStorage) (jvmtiEnv* env,
//    jthread thread,
//    const void* data);
//
//  /*   104 : Get Stack Trace */
//  jvmtiError (JNICALL *GetStackTrace) (jvmtiEnv* env,
//    jthread thread,
//    jint start_depth,
//    jint max_frame_count,
//    jvmtiFrameInfo* frame_buffer,
//    jint* count_ptr);
//
//  /*   105 :  RESERVED */
//  void *reserved105;
//
//  /*   106 : Get Tag */
//  jvmtiError (JNICALL *GetTag) (jvmtiEnv* env,
//    jobject object,
//    jlong* tag_ptr);
//
//  /*   107 : Set Tag */
//  jvmtiError (JNICALL *SetTag) (jvmtiEnv* env,
//    jobject object,
//    jlong tag);
//
//  /*   108 : Force Garbage Collection */
//  jvmtiError (JNICALL *ForceGarbageCollection) (jvmtiEnv* env);
//
//  /*   109 : Iterate Over Objects Reachable From Object */
//  jvmtiError (JNICALL *IterateOverObjectsReachableFromObject) (jvmtiEnv* env,
//    jobject object,
//    jvmtiObjectReferenceCallback object_reference_callback,
//    const void* user_data);
//
//  /*   110 : Iterate Over Reachable Objects */
//  jvmtiError (JNICALL *IterateOverReachableObjects) (jvmtiEnv* env,
//    jvmtiHeapRootCallback heap_root_callback,
//    jvmtiStackReferenceCallback stack_ref_callback,
//    jvmtiObjectReferenceCallback object_ref_callback,
//    const void* user_data);
//
//  /*   111 : Iterate Over Heap */
//  jvmtiError (JNICALL *IterateOverHeap) (jvmtiEnv* env,
//    jvmtiHeapObjectFilter object_filter,
//    jvmtiHeapObjectCallback heap_object_callback,
//    const void* user_data);
//
//  /*   112 : Iterate Over Instances Of Class */
//  jvmtiError (JNICALL *IterateOverInstancesOfClass) (jvmtiEnv* env,
//    jclass klass,
//    jvmtiHeapObjectFilter object_filter,
//    jvmtiHeapObjectCallback heap_object_callback,
//    const void* user_data);
//
//  /*   113 :  RESERVED */
//  void *reserved113;
//
//  /*   114 : Get Objects With Tags */
//  jvmtiError (JNICALL *GetObjectsWithTags) (jvmtiEnv* env,
//    jint tag_count,
//    const jlong* tags,
//    jint* count_ptr,
//    jobject** object_result_ptr,
//    jlong** tag_result_ptr);
//
//  /*   115 : Follow References */
//  jvmtiError (JNICALL *FollowReferences) (jvmtiEnv* env,
//    jint heap_filter,
//    jclass klass,
//    jobject initial_object,
//    const jvmtiHeapCallbacks* callbacks,
//    const void* user_data);
//
//  /*   116 : Iterate Through Heap */
//  jvmtiError (JNICALL *IterateThroughHeap) (jvmtiEnv* env,
//    jint heap_filter,
//    jclass klass,
//    const jvmtiHeapCallbacks* callbacks,
//    const void* user_data);
//
//  /*   117 :  RESERVED */
//  void *reserved117;
//
//  /*   118 :  RESERVED */
//  void *reserved118;
//
//  /*   119 :  RESERVED */
//  void *reserved119;
//
//  /*   120 : Set JNI Function Table */
//  jvmtiError (JNICALL *SetJNIFunctionTable) (jvmtiEnv* env,
//    const jniNativeInterface* function_table);
//
//  /*   121 : Get JNI Function Table */
//  jvmtiError (JNICALL *GetJNIFunctionTable) (jvmtiEnv* env,
//    jniNativeInterface** function_table);
//
//  /*   122 : Set Event Callbacks */
//  jvmtiError (JNICALL *SetEventCallbacks) (jvmtiEnv* env,
//    const jvmtiEventCallbacks* callbacks,
//    jint size_of_callbacks);
//
//  /*   123 : Generate Events */
//  jvmtiError (JNICALL *GenerateEvents) (jvmtiEnv* env,
//    jvmtiEvent event_type);
//
//  /*   124 : Get Extension Functions */
//  jvmtiError (JNICALL *GetExtensionFunctions) (jvmtiEnv* env,
//    jint* extension_count_ptr,
//    jvmtiExtensionFunctionInfo** extensions);
//
//  /*   125 : Get Extension Events */
//  jvmtiError (JNICALL *GetExtensionEvents) (jvmtiEnv* env,
//    jint* extension_count_ptr,
//    jvmtiExtensionEventInfo** extensions);
//
//  /*   126 : Set Extension Event Callback */
//  jvmtiError (JNICALL *SetExtensionEventCallback) (jvmtiEnv* env,
//    jint extension_event_index,
//    jvmtiExtensionEvent callback);
//
//  /*   127 : Dispose Environment */
//  jvmtiError (JNICALL *DisposeEnvironment) (jvmtiEnv* env);
//
//  /*   128 : Get Error Name */
//  jvmtiError (JNICALL *GetErrorName) (jvmtiEnv* env,
//    jvmtiError error,
//    char** name_ptr);
//
//  /*   129 : Get JLocation Format */
//  jvmtiError (JNICALL *GetJLocationFormat) (jvmtiEnv* env,
//    jvmtiJlocationFormat* format_ptr);
//
//  /*   130 : Get System Properties */
//  jvmtiError (JNICALL *GetSystemProperties) (jvmtiEnv* env,
//    jint* count_ptr,
//    char*** property_ptr);
//
//  /*   131 : Get System Property */
//  jvmtiError (JNICALL *GetSystemProperty) (jvmtiEnv* env,
//    const char* property,
//    char** value_ptr);
//
//  /*   132 : Set System Property */
//  jvmtiError (JNICALL *SetSystemProperty) (jvmtiEnv* env,
//    const char* property,
//    const char* value_ptr);
//
//  /*   133 : Get Phase */
//  jvmtiError (JNICALL *GetPhase) (jvmtiEnv* env,
//    jvmtiPhase* phase_ptr);
//
//  /*   134 : Get Current Thread CPU Timer Information */
//  jvmtiError (JNICALL *GetCurrentThreadCpuTimerInfo) (jvmtiEnv* env,
//    jvmtiTimerInfo* info_ptr);
//
//  /*   135 : Get Current Thread CPU Time */
//  jvmtiError (JNICALL *GetCurrentThreadCpuTime) (jvmtiEnv* env,
//    jlong* nanos_ptr);
//
//  /*   136 : Get Thread CPU Timer Information */
//  jvmtiError (JNICALL *GetThreadCpuTimerInfo) (jvmtiEnv* env,
//    jvmtiTimerInfo* info_ptr);
//
//  /*   137 : Get Thread CPU Time */
//  jvmtiError (JNICALL *GetThreadCpuTime) (jvmtiEnv* env,
//    jthread thread,
//    jlong* nanos_ptr);
//
//  /*   138 : Get Timer Information */
//  jvmtiError (JNICALL *GetTimerInfo) (jvmtiEnv* env,
//    jvmtiTimerInfo* info_ptr);
//
//  /*   139 : Get Time */
//  jvmtiError (JNICALL *GetTime) (jvmtiEnv* env,
//    jlong* nanos_ptr);
//
//  /*   140 : Get Potential Capabilities */
//  jvmtiError (JNICALL *GetPotentialCapabilities) (jvmtiEnv* env,
//    jvmtiCapabilities* capabilities_ptr);
//
//  /*   141 :  RESERVED */
//  void *reserved141;
//
//  /*   142 : Add Capabilities */
//  jvmtiError (JNICALL *AddCapabilities) (jvmtiEnv* env,
//    const jvmtiCapabilities* capabilities_ptr);
//
//  /*   143 : Relinquish Capabilities */
//  jvmtiError (JNICALL *RelinquishCapabilities) (jvmtiEnv* env,
//    const jvmtiCapabilities* capabilities_ptr);
//
//  /*   144 : Get Available Processors */
//  jvmtiError (JNICALL *GetAvailableProcessors) (jvmtiEnv* env,
//    jint* processor_count_ptr);
//
//  /*   145 : Get Class Version Numbers */
//  jvmtiError (JNICALL *GetClassVersionNumbers) (jvmtiEnv* env,
//    jclass klass,
//    jint* minor_version_ptr,
//    jint* major_version_ptr);
//
//  /*   146 : Get Constant Pool */
//  jvmtiError (JNICALL *GetConstantPool) (jvmtiEnv* env,
//    jclass klass,
//    jint* constant_pool_count_ptr,
//    jint* constant_pool_byte_count_ptr,
//    unsigned char** constant_pool_bytes_ptr);
//
//  /*   147 : Get Environment Local Storage */
//  jvmtiError (JNICALL *GetEnvironmentLocalStorage) (jvmtiEnv* env,
//    void** data_ptr);
//
//  /*   148 : Set Environment Local Storage */
//  jvmtiError (JNICALL *SetEnvironmentLocalStorage) (jvmtiEnv* env,
//    const void* data);
//
//  /*   149 : Add To Bootstrap Class Loader Search */
//  jvmtiError (JNICALL *AddToBootstrapClassLoaderSearch) (jvmtiEnv* env,
//    const char* segment);
//
//  /*   150 : Set Verbose Flag */
//  jvmtiError (JNICALL *SetVerboseFlag) (jvmtiEnv* env,
//    jvmtiVerboseFlag flag,
//    jboolean value);
//
//  /*   151 : Add To System Class Loader Search */
//  jvmtiError (JNICALL *AddToSystemClassLoaderSearch) (jvmtiEnv* env,
//    const char* segment);
//
//  /*   152 : Retransform Classes */
//  jvmtiError (JNICALL *RetransformClasses) (jvmtiEnv* env,
//    jint class_count,
//    const jclass* classes);
//
//  /*   153 : Get Owned Monitor Stack Depth Info */
//  jvmtiError (JNICALL *GetOwnedMonitorStackDepthInfo) (jvmtiEnv* env,
//    jthread thread,
//    jint* monitor_info_count_ptr,
//    jvmtiMonitorStackDepthInfo** monitor_info_ptr);
//
//  /*   154 : Get Object Size */
//  jvmtiError (JNICALL *GetObjectSize) (jvmtiEnv* env,
//    jobject object,
//    jlong* size_ptr);
//
//  /*   155 : Get Local Instance */
//  jvmtiError (JNICALL *GetLocalInstance) (jvmtiEnv* env,
//    jthread thread,
//    jint depth,
//    jobject* value_ptr);
//
//  /*   156 : Set Heap Sampling Interval */
//  jvmtiError (JNICALL *SetHeapSamplingInterval) (jvmtiEnv* env,
//    jint sampling_interval);


func (jvmti JvmtiEnv) GetClassSignature(clazz uintptr) (res string) {
	var sigp,genp unsafe.Pointer
	C.GetClassSignature(jvmti.raw(), unsafe.Pointer(clazz), unsafe.Pointer(&sigp), unsafe.Pointer(&genp))
	if sigp != nil {
		defer jvmti.Deallocate(sigp)
		res = C.GoString((*C.char)(sigp));
	}
	if genp != nil {
		defer jvmti.Deallocate(genp)
		tg := C.GoString((*C.char)(genp));
		res = res + "<" + tg + ">"
	}
	return res
}

//////////////////////// Memory management ////////////////////
func (jvmti JvmtiEnv) Allocate(sz int64) (res unsafe.Pointer) {
	C.Allocate(jvmti.asPointer(), C.longlong(sz), unsafe.Pointer(&res))
	return res
}

func (jvmti JvmtiEnv) Deallocate(mem unsafe.Pointer) int {
	return int(C.Deallocate(jvmti.asPointer(), mem))
}

//////////////////////// Thread ////////////////////
func (jvmti JvmtiEnv) GetThreadState(thrd uintptr) (stat int) {
	C.GetThreadState(jvmti.asPointer(), unsafe.Pointer(thrd), unsafe.Pointer(&stat))
	return stat
}

func (jvmti JvmtiEnv) GetCurrentThread() (thrd uintptr) {
	C.GetCurrentThread(jvmti.asPointer(), unsafe.Pointer(&thrd))
	return thrd
}

func (jvmti JvmtiEnv) GetAllThreads() (threads []C.jthread) {
	var n C.jint
	var ts *C.jthread
	if er := jvmti.getAllThreads(&n, &ts); er == JVMTI_ERROR_NONE {
		for i := 0; i < n; i++ {
			t := (C.jthread)(uintptr(unsafe.Pointer(ts)) + uintptr(i) * ptrSize)
			threads = append(threads, t)
		}
	}
	return threads
}

type ThreadInfo struct {
	Name string
	Priority int
	IsDaemon bool
	ThreadGroup uintptr
	ContextClassLoader uintptr
};

func (jvmti JvmtiEnv) GetThreadInfo(thrd uintptr) (info *ThreadInfo) {
	p := jvmti.Allocate(C.sizeof_struct__jvmtiThreadInfo)
	defer jvmti.Deallocate(p)
	C.GetThreadInfo(jvmti.asPointer(), unsafe.Pointer(thrd), p)
	ps := (*C.struct__jvmtiThreadInfo)(p)
	info = &ThreadInfo {
		Name: C.GoString(ps.name),
		Priority: int(ps.priority),
		IsDaemon: (ps.is_daemon != 0),
		ThreadGroup: uintptr(ps.thread_group),
		ContextClassLoader: uintptr(ps.context_class_loader),
	}
	return info
}
