// Copyright 2022 chuanshenglu@gmail.com
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
// limitations under the License.\n\n

package jgo

// #include <jvmti.h>
//
// static jvmtiError jvmtiSetEventNotificationMode(jvmtiEnv* env, jvmtiEventMode mode, jvmtiEvent event_type, jthread event_thread) {
//   return (*env)->SetEventNotificationMode(env, mode, event_type, event_thread);
// }
// static jvmtiError jvmtiGetAllModules(jvmtiEnv* env, jint* module_count_ptr, jobject** modules_ptr) {
//   return (*env)->GetAllModules(env, module_count_ptr, modules_ptr);
// }
// static jvmtiError jvmtiGetAllThreads(jvmtiEnv* env, jint* threads_count_ptr, jthread** threads_ptr) {
//   return (*env)->GetAllThreads(env, threads_count_ptr, threads_ptr);
// }
// static jvmtiError jvmtiSuspendThread(jvmtiEnv* env, jthread thread) {
//   return (*env)->SuspendThread(env, thread);
// }
// static jvmtiError jvmtiResumeThread(jvmtiEnv* env, jthread thread) {
//   return (*env)->ResumeThread(env, thread);
// }
// static jvmtiError jvmtiStopThread(jvmtiEnv* env, jthread thread, jobject exception) {
//   return (*env)->StopThread(env, thread, exception);
// }
// static jvmtiError jvmtiInterruptThread(jvmtiEnv* env, jthread thread) {
//   return (*env)->InterruptThread(env, thread);
// }
// static jvmtiError jvmtiGetThreadInfo(jvmtiEnv* env, jthread thread, jvmtiThreadInfo* info_ptr) {
//   return (*env)->GetThreadInfo(env, thread, info_ptr);
// }
// static jvmtiError jvmtiGetOwnedMonitorInfo(jvmtiEnv* env, jthread thread, jint* owned_monitor_count_ptr, jobject** owned_monitors_ptr) {
//   return (*env)->GetOwnedMonitorInfo(env, thread, owned_monitor_count_ptr, owned_monitors_ptr);
// }
// static jvmtiError jvmtiGetCurrentContendedMonitor(jvmtiEnv* env, jthread thread, jobject* monitor_ptr) {
//   return (*env)->GetCurrentContendedMonitor(env, thread, monitor_ptr);
// }
// static jvmtiError jvmtiRunAgentThread(jvmtiEnv* env, jthread thread, jvmtiStartFunction proc, const void* arg, jint priority) {
//   return (*env)->RunAgentThread(env, thread, proc, arg, priority);
// }
// static jvmtiError jvmtiGetTopThreadGroups(jvmtiEnv* env, jint* group_count_ptr, jthreadGroup** groups_ptr) {
//   return (*env)->GetTopThreadGroups(env, group_count_ptr, groups_ptr);
// }
// static jvmtiError jvmtiGetThreadGroupInfo(jvmtiEnv* env, jthreadGroup group, jvmtiThreadGroupInfo* info_ptr) {
//   return (*env)->GetThreadGroupInfo(env, group, info_ptr);
// }
// static jvmtiError jvmtiGetThreadGroupChildren(jvmtiEnv* env, jthreadGroup group, jint* thread_count_ptr, jthread** threads_ptr, jint* group_count_ptr, jthreadGroup** groups_ptr) {
//   return (*env)->GetThreadGroupChildren(env, group, thread_count_ptr, threads_ptr, group_count_ptr, groups_ptr);
// }
// static jvmtiError jvmtiGetFrameCount(jvmtiEnv* env, jthread thread, jint* count_ptr) {
//   return (*env)->GetFrameCount(env, thread, count_ptr);
// }
// static jvmtiError jvmtiGetThreadState(jvmtiEnv* env, jthread thread, jint* thread_state_ptr) {
//   return (*env)->GetThreadState(env, thread, thread_state_ptr);
// }
// static jvmtiError jvmtiGetCurrentThread(jvmtiEnv* env, jthread* thread_ptr) {
//   return (*env)->GetCurrentThread(env, thread_ptr);
// }
// static jvmtiError jvmtiGetFrameLocation(jvmtiEnv* env, jthread thread, jint depth, jmethodID* method_ptr, jlocation* location_ptr) {
//   return (*env)->GetFrameLocation(env, thread, depth, method_ptr, location_ptr);
// }
// static jvmtiError jvmtiNotifyFramePop(jvmtiEnv* env, jthread thread, jint depth) {
//   return (*env)->NotifyFramePop(env, thread, depth);
// }
// static jvmtiError jvmtiGetLocalObject(jvmtiEnv* env, jthread thread, jint depth, jint slot, jobject* value_ptr) {
//   return (*env)->GetLocalObject(env, thread, depth, slot, value_ptr);
// }
// static jvmtiError jvmtiGetLocalInt(jvmtiEnv* env, jthread thread, jint depth, jint slot, jint* value_ptr) {
//   return (*env)->GetLocalInt(env, thread, depth, slot, value_ptr);
// }
// static jvmtiError jvmtiGetLocalLong(jvmtiEnv* env, jthread thread, jint depth, jint slot, jlong* value_ptr) {
//   return (*env)->GetLocalLong(env, thread, depth, slot, value_ptr);
// }
// static jvmtiError jvmtiGetLocalFloat(jvmtiEnv* env, jthread thread, jint depth, jint slot, jfloat* value_ptr) {
//   return (*env)->GetLocalFloat(env, thread, depth, slot, value_ptr);
// }
// static jvmtiError jvmtiGetLocalDouble(jvmtiEnv* env, jthread thread, jint depth, jint slot, jdouble* value_ptr) {
//   return (*env)->GetLocalDouble(env, thread, depth, slot, value_ptr);
// }
// static jvmtiError jvmtiSetLocalObject(jvmtiEnv* env, jthread thread, jint depth, jint slot, jobject value) {
//   return (*env)->SetLocalObject(env, thread, depth, slot, value);
// }
// static jvmtiError jvmtiSetLocalInt(jvmtiEnv* env, jthread thread, jint depth, jint slot, jint value) {
//   return (*env)->SetLocalInt(env, thread, depth, slot, value);
// }
// static jvmtiError jvmtiSetLocalLong(jvmtiEnv* env, jthread thread, jint depth, jint slot, jlong value) {
//   return (*env)->SetLocalLong(env, thread, depth, slot, value);
// }
// static jvmtiError jvmtiSetLocalFloat(jvmtiEnv* env, jthread thread, jint depth, jint slot, jfloat value) {
//   return (*env)->SetLocalFloat(env, thread, depth, slot, value);
// }
// static jvmtiError jvmtiSetLocalDouble(jvmtiEnv* env, jthread thread, jint depth, jint slot, jdouble value) {
//   return (*env)->SetLocalDouble(env, thread, depth, slot, value);
// }
// static jvmtiError jvmtiCreateRawMonitor(jvmtiEnv* env, const char* name, jrawMonitorID* monitor_ptr) {
//   return (*env)->CreateRawMonitor(env, name, monitor_ptr);
// }
// static jvmtiError jvmtiDestroyRawMonitor(jvmtiEnv* env, jrawMonitorID monitor) {
//   return (*env)->DestroyRawMonitor(env, monitor);
// }
// static jvmtiError jvmtiRawMonitorEnter(jvmtiEnv* env, jrawMonitorID monitor) {
//   return (*env)->RawMonitorEnter(env, monitor);
// }
// static jvmtiError jvmtiRawMonitorExit(jvmtiEnv* env, jrawMonitorID monitor) {
//   return (*env)->RawMonitorExit(env, monitor);
// }
// static jvmtiError jvmtiRawMonitorWait(jvmtiEnv* env, jrawMonitorID monitor, jlong millis) {
//   return (*env)->RawMonitorWait(env, monitor, millis);
// }
// static jvmtiError jvmtiRawMonitorNotify(jvmtiEnv* env, jrawMonitorID monitor) {
//   return (*env)->RawMonitorNotify(env, monitor);
// }
// static jvmtiError jvmtiRawMonitorNotifyAll(jvmtiEnv* env, jrawMonitorID monitor) {
//   return (*env)->RawMonitorNotifyAll(env, monitor);
// }
// static jvmtiError jvmtiSetBreakpoint(jvmtiEnv* env, jmethodID method, jlocation location) {
//   return (*env)->SetBreakpoint(env, method, location);
// }
// static jvmtiError jvmtiClearBreakpoint(jvmtiEnv* env, jmethodID method, jlocation location) {
//   return (*env)->ClearBreakpoint(env, method, location);
// }
// static jvmtiError jvmtiGetNamedModule(jvmtiEnv* env, jobject class_loader, const char* package_name, jobject* module_ptr) {
//   return (*env)->GetNamedModule(env, class_loader, package_name, module_ptr);
// }
// static jvmtiError jvmtiSetFieldAccessWatch(jvmtiEnv* env, jclass klass, jfieldID field) {
//   return (*env)->SetFieldAccessWatch(env, klass, field);
// }
// static jvmtiError jvmtiClearFieldAccessWatch(jvmtiEnv* env, jclass klass, jfieldID field) {
//   return (*env)->ClearFieldAccessWatch(env, klass, field);
// }
// static jvmtiError jvmtiSetFieldModificationWatch(jvmtiEnv* env, jclass klass, jfieldID field) {
//   return (*env)->SetFieldModificationWatch(env, klass, field);
// }
// static jvmtiError jvmtiClearFieldModificationWatch(jvmtiEnv* env, jclass klass, jfieldID field) {
//   return (*env)->ClearFieldModificationWatch(env, klass, field);
// }
// static jvmtiError jvmtiIsModifiableClass(jvmtiEnv* env, jclass klass, jboolean* is_modifiable_class_ptr) {
//   return (*env)->IsModifiableClass(env, klass, is_modifiable_class_ptr);
// }
// static jvmtiError jvmtiAllocate(jvmtiEnv* env, jlong size, unsigned char** mem_ptr) {
//   return (*env)->Allocate(env, size, mem_ptr);
// }
// static jvmtiError jvmtiDeallocate(jvmtiEnv* env, unsigned char* mem) {
//   return (*env)->Deallocate(env, mem);
// }
// static jvmtiError jvmtiGetClassSignature(jvmtiEnv* env, jclass klass, char** signature_ptr, char** generic_ptr) {
//   return (*env)->GetClassSignature(env, klass, signature_ptr, generic_ptr);
// }
// static jvmtiError jvmtiGetClassStatus(jvmtiEnv* env, jclass klass, jint* status_ptr) {
//   return (*env)->GetClassStatus(env, klass, status_ptr);
// }
// static jvmtiError jvmtiGetSourceFileName(jvmtiEnv* env, jclass klass, char** source_name_ptr) {
//   return (*env)->GetSourceFileName(env, klass, source_name_ptr);
// }
// static jvmtiError jvmtiGetClassModifiers(jvmtiEnv* env, jclass klass, jint* modifiers_ptr) {
//   return (*env)->GetClassModifiers(env, klass, modifiers_ptr);
// }
// static jvmtiError jvmtiGetClassMethods(jvmtiEnv* env, jclass klass, jint* method_count_ptr, jmethodID** methods_ptr) {
//   return (*env)->GetClassMethods(env, klass, method_count_ptr, methods_ptr);
// }
// static jvmtiError jvmtiGetClassFields(jvmtiEnv* env, jclass klass, jint* field_count_ptr, jfieldID** fields_ptr) {
//   return (*env)->GetClassFields(env, klass, field_count_ptr, fields_ptr);
// }
// static jvmtiError jvmtiGetImplementedInterfaces(jvmtiEnv* env, jclass klass, jint* interface_count_ptr, jclass** interfaces_ptr) {
//   return (*env)->GetImplementedInterfaces(env, klass, interface_count_ptr, interfaces_ptr);
// }
// static jvmtiError jvmtiIsInterface(jvmtiEnv* env, jclass klass, jboolean* is_interface_ptr) {
//   return (*env)->IsInterface(env, klass, is_interface_ptr);
// }
// static jvmtiError jvmtiIsArrayClass(jvmtiEnv* env, jclass klass, jboolean* is_array_class_ptr) {
//   return (*env)->IsArrayClass(env, klass, is_array_class_ptr);
// }
// static jvmtiError jvmtiGetClassLoader(jvmtiEnv* env, jclass klass, jobject* classloader_ptr) {
//   return (*env)->GetClassLoader(env, klass, classloader_ptr);
// }
// static jvmtiError jvmtiGetObjectHashCode(jvmtiEnv* env, jobject object, jint* hash_code_ptr) {
//   return (*env)->GetObjectHashCode(env, object, hash_code_ptr);
// }
// static jvmtiError jvmtiGetObjectMonitorUsage(jvmtiEnv* env, jobject object, jvmtiMonitorUsage* info_ptr) {
//   return (*env)->GetObjectMonitorUsage(env, object, info_ptr);
// }
// static jvmtiError jvmtiGetFieldName(jvmtiEnv* env, jclass klass, jfieldID field, char** name_ptr, char** signature_ptr, char** generic_ptr) {
//   return (*env)->GetFieldName(env, klass, field, name_ptr, signature_ptr, generic_ptr);
// }
// static jvmtiError jvmtiGetFieldDeclaringClass(jvmtiEnv* env, jclass klass, jfieldID field, jclass* declaring_class_ptr) {
//   return (*env)->GetFieldDeclaringClass(env, klass, field, declaring_class_ptr);
// }
// static jvmtiError jvmtiGetFieldModifiers(jvmtiEnv* env, jclass klass, jfieldID field, jint* modifiers_ptr) {
//   return (*env)->GetFieldModifiers(env, klass, field, modifiers_ptr);
// }
// static jvmtiError jvmtiIsFieldSynthetic(jvmtiEnv* env, jclass klass, jfieldID field, jboolean* is_synthetic_ptr) {
//   return (*env)->IsFieldSynthetic(env, klass, field, is_synthetic_ptr);
// }
// static jvmtiError jvmtiGetMethodName(jvmtiEnv* env, jmethodID method, char** name_ptr, char** signature_ptr, char** generic_ptr) {
//   return (*env)->GetMethodName(env, method, name_ptr, signature_ptr, generic_ptr);
// }
// static jvmtiError jvmtiGetMethodDeclaringClass(jvmtiEnv* env, jmethodID method, jclass* declaring_class_ptr) {
//   return (*env)->GetMethodDeclaringClass(env, method, declaring_class_ptr);
// }
// static jvmtiError jvmtiGetMethodModifiers(jvmtiEnv* env, jmethodID method, jint* modifiers_ptr) {
//   return (*env)->GetMethodModifiers(env, method, modifiers_ptr);
// }
// static jvmtiError jvmtiGetMaxLocals(jvmtiEnv* env, jmethodID method, jint* max_ptr) {
//   return (*env)->GetMaxLocals(env, method, max_ptr);
// }
// static jvmtiError jvmtiGetArgumentsSize(jvmtiEnv* env, jmethodID method, jint* size_ptr) {
//   return (*env)->GetArgumentsSize(env, method, size_ptr);
// }
// static jvmtiError jvmtiGetLineNumberTable(jvmtiEnv* env, jmethodID method, jint* entry_count_ptr, jvmtiLineNumberEntry** table_ptr) {
//   return (*env)->GetLineNumberTable(env, method, entry_count_ptr, table_ptr);
// }
// static jvmtiError jvmtiGetMethodLocation(jvmtiEnv* env, jmethodID method, jlocation* start_location_ptr, jlocation* end_location_ptr) {
//   return (*env)->GetMethodLocation(env, method, start_location_ptr, end_location_ptr);
// }
// static jvmtiError jvmtiGetLocalVariableTable(jvmtiEnv* env, jmethodID method, jint* entry_count_ptr, jvmtiLocalVariableEntry** table_ptr) {
//   return (*env)->GetLocalVariableTable(env, method, entry_count_ptr, table_ptr);
// }
// static jvmtiError jvmtiSetNativeMethodPrefix(jvmtiEnv* env, const char* prefix) {
//   return (*env)->SetNativeMethodPrefix(env, prefix);
// }
// static jvmtiError jvmtiSetNativeMethodPrefixes(jvmtiEnv* env, jint prefix_count, char** prefixes) {
//   return (*env)->SetNativeMethodPrefixes(env, prefix_count, prefixes);
// }
// static jvmtiError jvmtiGetBytecodes(jvmtiEnv* env, jmethodID method, jint* bytecode_count_ptr, unsigned char** bytecodes_ptr) {
//   return (*env)->GetBytecodes(env, method, bytecode_count_ptr, bytecodes_ptr);
// }
// static jvmtiError jvmtiIsMethodNative(jvmtiEnv* env, jmethodID method, jboolean* is_native_ptr) {
//   return (*env)->IsMethodNative(env, method, is_native_ptr);
// }
// static jvmtiError jvmtiIsMethodSynthetic(jvmtiEnv* env, jmethodID method, jboolean* is_synthetic_ptr) {
//   return (*env)->IsMethodSynthetic(env, method, is_synthetic_ptr);
// }
// static jvmtiError jvmtiGetLoadedClasses(jvmtiEnv* env, jint* class_count_ptr, jclass** classes_ptr) {
//   return (*env)->GetLoadedClasses(env, class_count_ptr, classes_ptr);
// }
// static jvmtiError jvmtiGetClassLoaderClasses(jvmtiEnv* env, jobject initiating_loader, jint* class_count_ptr, jclass** classes_ptr) {
//   return (*env)->GetClassLoaderClasses(env, initiating_loader, class_count_ptr, classes_ptr);
// }
// static jvmtiError jvmtiPopFrame(jvmtiEnv* env, jthread thread) {
//   return (*env)->PopFrame(env, thread);
// }
// static jvmtiError jvmtiForceEarlyReturnObject(jvmtiEnv* env, jthread thread, jobject value) {
//   return (*env)->ForceEarlyReturnObject(env, thread, value);
// }
// static jvmtiError jvmtiForceEarlyReturnInt(jvmtiEnv* env, jthread thread, jint value) {
//   return (*env)->ForceEarlyReturnInt(env, thread, value);
// }
// static jvmtiError jvmtiForceEarlyReturnLong(jvmtiEnv* env, jthread thread, jlong value) {
//   return (*env)->ForceEarlyReturnLong(env, thread, value);
// }
// static jvmtiError jvmtiForceEarlyReturnFloat(jvmtiEnv* env, jthread thread, jfloat value) {
//   return (*env)->ForceEarlyReturnFloat(env, thread, value);
// }
// static jvmtiError jvmtiForceEarlyReturnDouble(jvmtiEnv* env, jthread thread, jdouble value) {
//   return (*env)->ForceEarlyReturnDouble(env, thread, value);
// }
// static jvmtiError jvmtiForceEarlyReturnVoid(jvmtiEnv* env, jthread thread) {
//   return (*env)->ForceEarlyReturnVoid(env, thread);
// }
// static jvmtiError jvmtiRedefineClasses(jvmtiEnv* env, jint class_count, const jvmtiClassDefinition* class_definitions) {
//   return (*env)->RedefineClasses(env, class_count, class_definitions);
// }
// static jvmtiError jvmtiGetVersionNumber(jvmtiEnv* env, jint* version_ptr) {
//   return (*env)->GetVersionNumber(env, version_ptr);
// }
// static jvmtiError jvmtiGetCapabilities(jvmtiEnv* env, jvmtiCapabilities* capabilities_ptr) {
//   return (*env)->GetCapabilities(env, capabilities_ptr);
// }
// static jvmtiError jvmtiGetSourceDebugExtension(jvmtiEnv* env, jclass klass, char** source_debug_extension_ptr) {
//   return (*env)->GetSourceDebugExtension(env, klass, source_debug_extension_ptr);
// }
// static jvmtiError jvmtiIsMethodObsolete(jvmtiEnv* env, jmethodID method, jboolean* is_obsolete_ptr) {
//   return (*env)->IsMethodObsolete(env, method, is_obsolete_ptr);
// }
// static jvmtiError jvmtiSuspendThreadList(jvmtiEnv* env, jint request_count, const jthread* request_list, jvmtiError* results) {
//   return (*env)->SuspendThreadList(env, request_count, request_list, results);
// }
// static jvmtiError jvmtiResumeThreadList(jvmtiEnv* env, jint request_count, const jthread* request_list, jvmtiError* results) {
//   return (*env)->ResumeThreadList(env, request_count, request_list, results);
// }
// static jvmtiError jvmtiAddModuleReads(jvmtiEnv* env, jobject module, jobject to_module) {
//   return (*env)->AddModuleReads(env, module, to_module);
// }
// static jvmtiError jvmtiAddModuleExports(jvmtiEnv* env, jobject module, const char* pkg_name, jobject to_module) {
//   return (*env)->AddModuleExports(env, module, pkg_name, to_module);
// }
// static jvmtiError jvmtiAddModuleOpens(jvmtiEnv* env, jobject module, const char* pkg_name, jobject to_module) {
//   return (*env)->AddModuleOpens(env, module, pkg_name, to_module);
// }
// static jvmtiError jvmtiAddModuleUses(jvmtiEnv* env, jobject module, jclass service) {
//   return (*env)->AddModuleUses(env, module, service);
// }
// static jvmtiError jvmtiAddModuleProvides(jvmtiEnv* env, jobject module, jclass service, jclass impl_class) {
//   return (*env)->AddModuleProvides(env, module, service, impl_class);
// }
// static jvmtiError jvmtiIsModifiableModule(jvmtiEnv* env, jobject module, jboolean* is_modifiable_module_ptr) {
//   return (*env)->IsModifiableModule(env, module, is_modifiable_module_ptr);
// }
// static jvmtiError jvmtiGetAllStackTraces(jvmtiEnv* env, jint max_frame_count, jvmtiStackInfo** stack_info_ptr, jint* thread_count_ptr) {
//   return (*env)->GetAllStackTraces(env, max_frame_count, stack_info_ptr, thread_count_ptr);
// }
// static jvmtiError jvmtiGetThreadListStackTraces(jvmtiEnv* env, jint thread_count, const jthread* thread_list, jint max_frame_count, jvmtiStackInfo** stack_info_ptr) {
//   return (*env)->GetThreadListStackTraces(env, thread_count, thread_list, max_frame_count, stack_info_ptr);
// }
// static jvmtiError jvmtiGetThreadLocalStorage(jvmtiEnv* env, jthread thread, void** data_ptr) {
//   return (*env)->GetThreadLocalStorage(env, thread, data_ptr);
// }
// static jvmtiError jvmtiSetThreadLocalStorage(jvmtiEnv* env, jthread thread, const void* data) {
//   return (*env)->SetThreadLocalStorage(env, thread, data);
// }
// static jvmtiError jvmtiGetStackTrace(jvmtiEnv* env, jthread thread, jint start_depth, jint max_frame_count, jvmtiFrameInfo* frame_buffer, jint* count_ptr) {
//   return (*env)->GetStackTrace(env, thread, start_depth, max_frame_count, frame_buffer, count_ptr);
// }
// static jvmtiError jvmtiGetTag(jvmtiEnv* env, jobject object, jlong* tag_ptr) {
//   return (*env)->GetTag(env, object, tag_ptr);
// }
// static jvmtiError jvmtiSetTag(jvmtiEnv* env, jobject object, jlong tag) {
//   return (*env)->SetTag(env, object, tag);
// }
// static jvmtiError jvmtiIterateOverObjectsReachableFromObject(jvmtiEnv* env, jobject object, jvmtiObjectReferenceCallback object_reference_callback, const void* user_data) {
//   return (*env)->IterateOverObjectsReachableFromObject(env, object, object_reference_callback, user_data);
// }
// static jvmtiError jvmtiIterateOverReachableObjects(jvmtiEnv* env, jvmtiHeapRootCallback heap_root_callback, jvmtiStackReferenceCallback stack_ref_callback, jvmtiObjectReferenceCallback object_ref_callback, const void* user_data) {
//   return (*env)->IterateOverReachableObjects(env, heap_root_callback, stack_ref_callback, object_ref_callback, user_data);
// }
// static jvmtiError jvmtiIterateOverHeap(jvmtiEnv* env, jvmtiHeapObjectFilter object_filter, jvmtiHeapObjectCallback heap_object_callback, const void* user_data) {
//   return (*env)->IterateOverHeap(env, object_filter, heap_object_callback, user_data);
// }
// static jvmtiError jvmtiIterateOverInstancesOfClass(jvmtiEnv* env, jclass klass, jvmtiHeapObjectFilter object_filter, jvmtiHeapObjectCallback heap_object_callback, const void* user_data) {
//   return (*env)->IterateOverInstancesOfClass(env, klass, object_filter, heap_object_callback, user_data);
// }
// static jvmtiError jvmtiGetObjectsWithTags(jvmtiEnv* env, jint tag_count, const jlong* tags, jint* count_ptr, jobject** object_result_ptr, jlong** tag_result_ptr) {
//   return (*env)->GetObjectsWithTags(env, tag_count, tags, count_ptr, object_result_ptr, tag_result_ptr);
// }
// static jvmtiError jvmtiFollowReferences(jvmtiEnv* env, jint heap_filter, jclass klass, jobject initial_object, const jvmtiHeapCallbacks* callbacks, const void* user_data) {
//   return (*env)->FollowReferences(env, heap_filter, klass, initial_object, callbacks, user_data);
// }
// static jvmtiError jvmtiIterateThroughHeap(jvmtiEnv* env, jint heap_filter, jclass klass, const jvmtiHeapCallbacks* callbacks, const void* user_data) {
//   return (*env)->IterateThroughHeap(env, heap_filter, klass, callbacks, user_data);
// }
// static jvmtiError jvmtiSetJNIFunctionTable(jvmtiEnv* env, const jniNativeInterface* function_table) {
//   return (*env)->SetJNIFunctionTable(env, function_table);
// }
// static jvmtiError jvmtiGetJNIFunctionTable(jvmtiEnv* env, jniNativeInterface** function_table) {
//   return (*env)->GetJNIFunctionTable(env, function_table);
// }
// static jvmtiError jvmtiSetEventCallbacks(jvmtiEnv* env, const jvmtiEventCallbacks* callbacks, jint size_of_callbacks) {
//   return (*env)->SetEventCallbacks(env, callbacks, size_of_callbacks);
// }
// static jvmtiError jvmtiGenerateEvents(jvmtiEnv* env, jvmtiEvent event_type) {
//   return (*env)->GenerateEvents(env, event_type);
// }
// static jvmtiError jvmtiGetExtensionFunctions(jvmtiEnv* env, jint* extension_count_ptr, jvmtiExtensionFunctionInfo** extensions) {
//   return (*env)->GetExtensionFunctions(env, extension_count_ptr, extensions);
// }
// static jvmtiError jvmtiGetExtensionEvents(jvmtiEnv* env, jint* extension_count_ptr, jvmtiExtensionEventInfo** extensions) {
//   return (*env)->GetExtensionEvents(env, extension_count_ptr, extensions);
// }
// static jvmtiError jvmtiSetExtensionEventCallback(jvmtiEnv* env, jint extension_event_index, jvmtiExtensionEvent callback) {
//   return (*env)->SetExtensionEventCallback(env, extension_event_index, callback);
// }
// static jvmtiError jvmtiGetErrorName(jvmtiEnv* env, jvmtiError error, char** name_ptr) {
//   return (*env)->GetErrorName(env, error, name_ptr);
// }
// static jvmtiError jvmtiGetJLocationFormat(jvmtiEnv* env, jvmtiJlocationFormat* format_ptr) {
//   return (*env)->GetJLocationFormat(env, format_ptr);
// }
// static jvmtiError jvmtiGetSystemProperties(jvmtiEnv* env, jint* count_ptr, char*** property_ptr) {
//   return (*env)->GetSystemProperties(env, count_ptr, property_ptr);
// }
// static jvmtiError jvmtiGetSystemProperty(jvmtiEnv* env, const char* property, char** value_ptr) {
//   return (*env)->GetSystemProperty(env, property, value_ptr);
// }
// static jvmtiError jvmtiSetSystemProperty(jvmtiEnv* env, const char* property, const char* value_ptr) {
//   return (*env)->SetSystemProperty(env, property, value_ptr);
// }
// static jvmtiError jvmtiGetPhase(jvmtiEnv* env, jvmtiPhase* phase_ptr) {
//   return (*env)->GetPhase(env, phase_ptr);
// }
// static jvmtiError jvmtiGetCurrentThreadCpuTimerInfo(jvmtiEnv* env, jvmtiTimerInfo* info_ptr) {
//   return (*env)->GetCurrentThreadCpuTimerInfo(env, info_ptr);
// }
// static jvmtiError jvmtiGetCurrentThreadCpuTime(jvmtiEnv* env, jlong* nanos_ptr) {
//   return (*env)->GetCurrentThreadCpuTime(env, nanos_ptr);
// }
// static jvmtiError jvmtiGetThreadCpuTimerInfo(jvmtiEnv* env, jvmtiTimerInfo* info_ptr) {
//   return (*env)->GetThreadCpuTimerInfo(env, info_ptr);
// }
// static jvmtiError jvmtiGetThreadCpuTime(jvmtiEnv* env, jthread thread, jlong* nanos_ptr) {
//   return (*env)->GetThreadCpuTime(env, thread, nanos_ptr);
// }
// static jvmtiError jvmtiGetTimerInfo(jvmtiEnv* env, jvmtiTimerInfo* info_ptr) {
//   return (*env)->GetTimerInfo(env, info_ptr);
// }
// static jvmtiError jvmtiGetTime(jvmtiEnv* env, jlong* nanos_ptr) {
//   return (*env)->GetTime(env, nanos_ptr);
// }
// static jvmtiError jvmtiGetPotentialCapabilities(jvmtiEnv* env, jvmtiCapabilities* capabilities_ptr) {
//   return (*env)->GetPotentialCapabilities(env, capabilities_ptr);
// }
// static jvmtiError jvmtiAddCapabilities(jvmtiEnv* env, const jvmtiCapabilities* capabilities_ptr) {
//   return (*env)->AddCapabilities(env, capabilities_ptr);
// }
// static jvmtiError jvmtiRelinquishCapabilities(jvmtiEnv* env, const jvmtiCapabilities* capabilities_ptr) {
//   return (*env)->RelinquishCapabilities(env, capabilities_ptr);
// }
// static jvmtiError jvmtiGetAvailableProcessors(jvmtiEnv* env, jint* processor_count_ptr) {
//   return (*env)->GetAvailableProcessors(env, processor_count_ptr);
// }
// static jvmtiError jvmtiGetClassVersionNumbers(jvmtiEnv* env, jclass klass, jint* minor_version_ptr, jint* major_version_ptr) {
//   return (*env)->GetClassVersionNumbers(env, klass, minor_version_ptr, major_version_ptr);
// }
// static jvmtiError jvmtiGetConstantPool(jvmtiEnv* env, jclass klass, jint* constant_pool_count_ptr, jint* constant_pool_byte_count_ptr, unsigned char** constant_pool_bytes_ptr) {
//   return (*env)->GetConstantPool(env, klass, constant_pool_count_ptr, constant_pool_byte_count_ptr, constant_pool_bytes_ptr);
// }
// static jvmtiError jvmtiGetEnvironmentLocalStorage(jvmtiEnv* env, void** data_ptr) {
//   return (*env)->GetEnvironmentLocalStorage(env, data_ptr);
// }
// static jvmtiError jvmtiSetEnvironmentLocalStorage(jvmtiEnv* env, const void* data) {
//   return (*env)->SetEnvironmentLocalStorage(env, data);
// }
// static jvmtiError jvmtiAddToBootstrapClassLoaderSearch(jvmtiEnv* env, const char* segment) {
//   return (*env)->AddToBootstrapClassLoaderSearch(env, segment);
// }
// static jvmtiError jvmtiSetVerboseFlag(jvmtiEnv* env, jvmtiVerboseFlag flag, jboolean value) {
//   return (*env)->SetVerboseFlag(env, flag, value);
// }
// static jvmtiError jvmtiAddToSystemClassLoaderSearch(jvmtiEnv* env, const char* segment) {
//   return (*env)->AddToSystemClassLoaderSearch(env, segment);
// }
// static jvmtiError jvmtiRetransformClasses(jvmtiEnv* env, jint class_count, const jclass* classes) {
//   return (*env)->RetransformClasses(env, class_count, classes);
// }
// static jvmtiError jvmtiGetOwnedMonitorStackDepthInfo(jvmtiEnv* env, jthread thread, jint* monitor_info_count_ptr, jvmtiMonitorStackDepthInfo** monitor_info_ptr) {
//   return (*env)->GetOwnedMonitorStackDepthInfo(env, thread, monitor_info_count_ptr, monitor_info_ptr);
// }
// static jvmtiError jvmtiGetObjectSize(jvmtiEnv* env, jobject object, jlong* size_ptr) {
//   return (*env)->GetObjectSize(env, object, size_ptr);
// }
// static jvmtiError jvmtiGetLocalInstance(jvmtiEnv* env, jthread thread, jint depth, jobject* value_ptr) {
//   return (*env)->GetLocalInstance(env, thread, depth, value_ptr);
// }
// static jvmtiError jvmtiSetHeapSamplingInterval(jvmtiEnv* env, jint sampling_interval) {
//   return (*env)->SetHeapSamplingInterval(env, sampling_interval);
// }
import "C"

import "unsafe"

import "fmt"

const (
    JVMTI_VERSION_1   = 0x30010000
    JVMTI_VERSION_1_0 = 0x30010000
    JVMTI_VERSION_1_1 = 0x30010100
    JVMTI_VERSION_1_2 = 0x30010200
    JVMTI_VERSION_9   = 0x30090000
    JVMTI_VERSION_11  = 0x300B0000

    JVMTI_VERSION = 0x30000000 + (14 * 0x10000) + ( 0 * 0x100) + 0  /* version: 14.0.0 */
)

const (
    JVMTI_THREAD_STATE_ALIVE = 0x0001
    JVMTI_THREAD_STATE_TERMINATED = 0x0002
    JVMTI_THREAD_STATE_RUNNABLE = 0x0004
    JVMTI_THREAD_STATE_BLOCKED_ON_MONITOR_ENTER = 0x0400
    JVMTI_THREAD_STATE_WAITING = 0x0080
    JVMTI_THREAD_STATE_WAITING_INDEFINITELY = 0x0010
    JVMTI_THREAD_STATE_WAITING_WITH_TIMEOUT = 0x0020
    JVMTI_THREAD_STATE_SLEEPING = 0x0040
    JVMTI_THREAD_STATE_IN_OBJECT_WAIT = 0x0100
    JVMTI_THREAD_STATE_PARKED = 0x0200
    JVMTI_THREAD_STATE_SUSPENDED = 0x100000
    JVMTI_THREAD_STATE_INTERRUPTED = 0x200000
    JVMTI_THREAD_STATE_IN_NATIVE = 0x400000
    JVMTI_THREAD_STATE_VENDOR_1 = 0x10000000
    JVMTI_THREAD_STATE_VENDOR_2 = 0x20000000
    JVMTI_THREAD_STATE_VENDOR_3 = 0x40000000
)

const (
    JVMTI_JAVA_LANG_THREAD_STATE_MASK = JVMTI_THREAD_STATE_TERMINATED | JVMTI_THREAD_STATE_ALIVE | JVMTI_THREAD_STATE_RUNNABLE | JVMTI_THREAD_STATE_BLOCKED_ON_MONITOR_ENTER | JVMTI_THREAD_STATE_WAITING | JVMTI_THREAD_STATE_WAITING_INDEFINITELY | JVMTI_THREAD_STATE_WAITING_WITH_TIMEOUT
    JVMTI_JAVA_LANG_THREAD_STATE_NEW = 0
    JVMTI_JAVA_LANG_THREAD_STATE_TERMINATED = JVMTI_THREAD_STATE_TERMINATED
    JVMTI_JAVA_LANG_THREAD_STATE_RUNNABLE = JVMTI_THREAD_STATE_ALIVE | JVMTI_THREAD_STATE_RUNNABLE
    JVMTI_JAVA_LANG_THREAD_STATE_BLOCKED = JVMTI_THREAD_STATE_ALIVE | JVMTI_THREAD_STATE_BLOCKED_ON_MONITOR_ENTER
    JVMTI_JAVA_LANG_THREAD_STATE_WAITING = JVMTI_THREAD_STATE_ALIVE | JVMTI_THREAD_STATE_WAITING | JVMTI_THREAD_STATE_WAITING_INDEFINITELY
    JVMTI_JAVA_LANG_THREAD_STATE_TIMED_WAITING = JVMTI_THREAD_STATE_ALIVE | JVMTI_THREAD_STATE_WAITING | JVMTI_THREAD_STATE_WAITING_WITH_TIMEOUT
)

const (
    JVMTI_THREAD_MIN_PRIORITY = 1
    JVMTI_THREAD_NORM_PRIORITY = 5
    JVMTI_THREAD_MAX_PRIORITY = 10
)

const (
    JVMTI_HEAP_FILTER_TAGGED = 0x4
    JVMTI_HEAP_FILTER_UNTAGGED = 0x8
    JVMTI_HEAP_FILTER_CLASS_TAGGED = 0x10
    JVMTI_HEAP_FILTER_CLASS_UNTAGGED = 0x20
)

const (
    JVMTI_VISIT_OBJECTS = 0x100
    JVMTI_VISIT_ABORT = 0x8000
)

const (
    JVMTI_HEAP_REFERENCE_CLASS = 1
    JVMTI_HEAP_REFERENCE_FIELD = 2
    JVMTI_HEAP_REFERENCE_ARRAY_ELEMENT = 3
    JVMTI_HEAP_REFERENCE_CLASS_LOADER = 4
    JVMTI_HEAP_REFERENCE_SIGNERS = 5
    JVMTI_HEAP_REFERENCE_PROTECTION_DOMAIN = 6
    JVMTI_HEAP_REFERENCE_INTERFACE = 7
    JVMTI_HEAP_REFERENCE_STATIC_FIELD = 8
    JVMTI_HEAP_REFERENCE_CONSTANT_POOL = 9
    JVMTI_HEAP_REFERENCE_SUPERCLASS = 10
    JVMTI_HEAP_REFERENCE_JNI_GLOBAL = 21
    JVMTI_HEAP_REFERENCE_SYSTEM_CLASS = 22
    JVMTI_HEAP_REFERENCE_MONITOR = 23
    JVMTI_HEAP_REFERENCE_STACK_LOCAL = 24
    JVMTI_HEAP_REFERENCE_JNI_LOCAL = 25
    JVMTI_HEAP_REFERENCE_THREAD = 26
    JVMTI_HEAP_REFERENCE_OTHER = 27
)

const (
    JVMTI_PRIMITIVE_TYPE_BOOLEAN = 90
    JVMTI_PRIMITIVE_TYPE_BYTE = 66
    JVMTI_PRIMITIVE_TYPE_CHAR = 67
    JVMTI_PRIMITIVE_TYPE_SHORT = 83
    JVMTI_PRIMITIVE_TYPE_INT = 73
    JVMTI_PRIMITIVE_TYPE_LONG = 74
    JVMTI_PRIMITIVE_TYPE_FLOAT = 70
    JVMTI_PRIMITIVE_TYPE_DOUBLE = 68
)

const (
    JVMTI_HEAP_OBJECT_TAGGED = 1
    JVMTI_HEAP_OBJECT_UNTAGGED = 2
    JVMTI_HEAP_OBJECT_EITHER = 3
)

const (
    JVMTI_HEAP_ROOT_JNI_GLOBAL = 1
    JVMTI_HEAP_ROOT_SYSTEM_CLASS = 2
    JVMTI_HEAP_ROOT_MONITOR = 3
    JVMTI_HEAP_ROOT_STACK_LOCAL = 4
    JVMTI_HEAP_ROOT_JNI_LOCAL = 5
    JVMTI_HEAP_ROOT_THREAD = 6
    JVMTI_HEAP_ROOT_OTHER = 7
)

const (
    JVMTI_REFERENCE_CLASS = 1
    JVMTI_REFERENCE_FIELD = 2
    JVMTI_REFERENCE_ARRAY_ELEMENT = 3
    JVMTI_REFERENCE_CLASS_LOADER = 4
    JVMTI_REFERENCE_SIGNERS = 5
    JVMTI_REFERENCE_PROTECTION_DOMAIN = 6
    JVMTI_REFERENCE_INTERFACE = 7
    JVMTI_REFERENCE_STATIC_FIELD = 8
    JVMTI_REFERENCE_CONSTANT_POOL = 9
)

const (
    JVMTI_ITERATION_CONTINUE = 1
    JVMTI_ITERATION_IGNORE = 2
    JVMTI_ITERATION_ABORT = 0
)

const (
    JVMTI_CLASS_STATUS_VERIFIED = 1
    JVMTI_CLASS_STATUS_PREPARED = 2
    JVMTI_CLASS_STATUS_INITIALIZED = 4
    JVMTI_CLASS_STATUS_ERROR = 8
    JVMTI_CLASS_STATUS_ARRAY = 16
    JVMTI_CLASS_STATUS_PRIMITIVE = 32
)

const (
    JVMTI_ENABLE = 1
    JVMTI_DISABLE = 0
)

const (
    JVMTI_TYPE_JBYTE = 101
    JVMTI_TYPE_JCHAR = 102
    JVMTI_TYPE_JSHORT = 103
    JVMTI_TYPE_JINT = 104
    JVMTI_TYPE_JLONG = 105
    JVMTI_TYPE_JFLOAT = 106
    JVMTI_TYPE_JDOUBLE = 107
    JVMTI_TYPE_JBOOLEAN = 108
    JVMTI_TYPE_JOBJECT = 109
    JVMTI_TYPE_JTHREAD = 110
    JVMTI_TYPE_JCLASS = 111
    JVMTI_TYPE_JVALUE = 112
    JVMTI_TYPE_JFIELDID = 113
    JVMTI_TYPE_JMETHODID = 114
    JVMTI_TYPE_CCHAR = 115
    JVMTI_TYPE_CVOID = 116
    JVMTI_TYPE_JNIENV = 117
)

const (
    JVMTI_KIND_IN = 91
    JVMTI_KIND_IN_PTR = 92
    JVMTI_KIND_IN_BUF = 93
    JVMTI_KIND_ALLOC_BUF = 94
    JVMTI_KIND_ALLOC_ALLOC_BUF = 95
    JVMTI_KIND_OUT = 96
    JVMTI_KIND_OUT_BUF = 97
)

const (
    JVMTI_TIMER_USER_CPU = 30
    JVMTI_TIMER_TOTAL_CPU = 31
    JVMTI_TIMER_ELAPSED = 32
)

const (
    JVMTI_PHASE_ONLOAD = 1
    JVMTI_PHASE_PRIMORDIAL = 2
    JVMTI_PHASE_START = 6
    JVMTI_PHASE_LIVE = 4
    JVMTI_PHASE_DEAD = 8
)

const (
    JVMTI_VERSION_INTERFACE_JNI = 0x00000000
    JVMTI_VERSION_INTERFACE_JVMTI = 0x30000000
)

const (
    JVMTI_VERSION_MASK_INTERFACE_TYPE = 0x70000000
    JVMTI_VERSION_MASK_MAJOR = 0x0FFF0000
    JVMTI_VERSION_MASK_MINOR = 0x0000FF00
    JVMTI_VERSION_MASK_MICRO = 0x000000FF
)

const (
    JVMTI_VERSION_SHIFT_MAJOR = 16
    JVMTI_VERSION_SHIFT_MINOR = 8
    JVMTI_VERSION_SHIFT_MICRO = 0
)

const (
    JVMTI_VERBOSE_OTHER = 0
    JVMTI_VERBOSE_GC = 1
    JVMTI_VERBOSE_CLASS = 2
    JVMTI_VERBOSE_JNI = 4
)

const (
    JVMTI_JLOCATION_JVMBCI = 1
    JVMTI_JLOCATION_MACHINEPC = 2
    JVMTI_JLOCATION_OTHER = 0
)

const (
    JVMTI_RESOURCE_EXHAUSTED_OOM_ERROR = 0x0001
    JVMTI_RESOURCE_EXHAUSTED_JAVA_HEAP = 0x0002
    JVMTI_RESOURCE_EXHAUSTED_THREADS = 0x0004
)

const (
    JVMTI_ERROR_NONE = 0
    JVMTI_ERROR_INVALID_THREAD = 10
    JVMTI_ERROR_INVALID_THREAD_GROUP = 11
    JVMTI_ERROR_INVALID_PRIORITY = 12
    JVMTI_ERROR_THREAD_NOT_SUSPENDED = 13
    JVMTI_ERROR_THREAD_SUSPENDED = 14
    JVMTI_ERROR_THREAD_NOT_ALIVE = 15
    JVMTI_ERROR_INVALID_OBJECT = 20
    JVMTI_ERROR_INVALID_CLASS = 21
    JVMTI_ERROR_CLASS_NOT_PREPARED = 22
    JVMTI_ERROR_INVALID_METHODID = 23
    JVMTI_ERROR_INVALID_LOCATION = 24
    JVMTI_ERROR_INVALID_FIELDID = 25
    JVMTI_ERROR_INVALID_MODULE = 26
    JVMTI_ERROR_NO_MORE_FRAMES = 31
    JVMTI_ERROR_OPAQUE_FRAME = 32
    JVMTI_ERROR_TYPE_MISMATCH = 34
    JVMTI_ERROR_INVALID_SLOT = 35
    JVMTI_ERROR_DUPLICATE = 40
    JVMTI_ERROR_NOT_FOUND = 41
    JVMTI_ERROR_INVALID_MONITOR = 50
    JVMTI_ERROR_NOT_MONITOR_OWNER = 51
    JVMTI_ERROR_INTERRUPT = 52
    JVMTI_ERROR_INVALID_CLASS_FORMAT = 60
    JVMTI_ERROR_CIRCULAR_CLASS_DEFINITION = 61
    JVMTI_ERROR_FAILS_VERIFICATION = 62
    JVMTI_ERROR_UNSUPPORTED_REDEFINITION_METHOD_ADDED = 63
    JVMTI_ERROR_UNSUPPORTED_REDEFINITION_SCHEMA_CHANGED = 64
    JVMTI_ERROR_INVALID_TYPESTATE = 65
    JVMTI_ERROR_UNSUPPORTED_REDEFINITION_HIERARCHY_CHANGED = 66
    JVMTI_ERROR_UNSUPPORTED_REDEFINITION_METHOD_DELETED = 67
    JVMTI_ERROR_UNSUPPORTED_VERSION = 68
    JVMTI_ERROR_NAMES_DONT_MATCH = 69
    JVMTI_ERROR_UNSUPPORTED_REDEFINITION_CLASS_MODIFIERS_CHANGED = 70
    JVMTI_ERROR_UNSUPPORTED_REDEFINITION_METHOD_MODIFIERS_CHANGED = 71
    JVMTI_ERROR_UNSUPPORTED_REDEFINITION_CLASS_ATTRIBUTE_CHANGED = 72
    JVMTI_ERROR_UNMODIFIABLE_CLASS = 79
    JVMTI_ERROR_UNMODIFIABLE_MODULE = 80
    JVMTI_ERROR_NOT_AVAILABLE = 98
    JVMTI_ERROR_MUST_POSSESS_CAPABILITY = 99
    JVMTI_ERROR_NULL_POINTER = 100
    JVMTI_ERROR_ABSENT_INFORMATION = 101
    JVMTI_ERROR_INVALID_EVENT_TYPE = 102
    JVMTI_ERROR_ILLEGAL_ARGUMENT = 103
    JVMTI_ERROR_NATIVE_METHOD = 104
    JVMTI_ERROR_CLASS_LOADER_UNSUPPORTED = 106
    JVMTI_ERROR_OUT_OF_MEMORY = 110
    JVMTI_ERROR_ACCESS_DENIED = 111
    JVMTI_ERROR_WRONG_PHASE = 112
    JVMTI_ERROR_INTERNAL = 113
    JVMTI_ERROR_UNATTACHED_THREAD = 115
    JVMTI_ERROR_INVALID_ENVIRONMENT = 116
    JVMTI_ERROR_MAX = 116
)

const (
    JVMTI_MIN_EVENT_TYPE_VAL = 50
    JVMTI_EVENT_VM_INIT = 50
    JVMTI_EVENT_VM_DEATH = 51
    JVMTI_EVENT_THREAD_START = 52
    JVMTI_EVENT_THREAD_END = 53
    JVMTI_EVENT_CLASS_FILE_LOAD_HOOK = 54
    JVMTI_EVENT_CLASS_LOAD = 55
    JVMTI_EVENT_CLASS_PREPARE = 56
    JVMTI_EVENT_VM_START = 57
    JVMTI_EVENT_EXCEPTION = 58
    JVMTI_EVENT_EXCEPTION_CATCH = 59
    JVMTI_EVENT_SINGLE_STEP = 60
    JVMTI_EVENT_FRAME_POP = 61
    JVMTI_EVENT_BREAKPOINT = 62
    JVMTI_EVENT_FIELD_ACCESS = 63
    JVMTI_EVENT_FIELD_MODIFICATION = 64
    JVMTI_EVENT_METHOD_ENTRY = 65
    JVMTI_EVENT_METHOD_EXIT = 66
    JVMTI_EVENT_NATIVE_METHOD_BIND = 67
    JVMTI_EVENT_COMPILED_METHOD_LOAD = 68
    JVMTI_EVENT_COMPILED_METHOD_UNLOAD = 69
    JVMTI_EVENT_DYNAMIC_CODE_GENERATED = 70
    JVMTI_EVENT_DATA_DUMP_REQUEST = 71
    JVMTI_EVENT_MONITOR_WAIT = 73
    JVMTI_EVENT_MONITOR_WAITED = 74
    JVMTI_EVENT_MONITOR_CONTENDED_ENTER = 75
    JVMTI_EVENT_MONITOR_CONTENDED_ENTERED = 76
    JVMTI_EVENT_RESOURCE_EXHAUSTED = 80
    JVMTI_EVENT_GARBAGE_COLLECTION_START = 81
    JVMTI_EVENT_GARBAGE_COLLECTION_FINISH = 82
    JVMTI_EVENT_OBJECT_FREE = 83
    JVMTI_EVENT_VM_OBJECT_ALLOC = 84
    JVMTI_EVENT_SAMPLED_OBJECT_ALLOC = 86
    JVMTI_MAX_EVENT_TYPE_VAL = 86
)



// jvmtiEnv corresponds to jvmtiEnv*
type jvmtiEnv uintptr

func (jvmti jvmtiEnv) raw() *C.jvmtiEnv {
	return (*C.jvmtiEnv)(unsafe.Pointer(jvmti))
}

func (jvmti jvmtiEnv) asPointer() unsafe.Pointer {
	return unsafe.Pointer(jvmti)
}


func (jvmti jvmtiEnv) setEventNotificationMode(mode C.jvmtiEventMode, event_type C.jvmtiEvent, event_thread C.jthread) C.jvmtiError {
  return C.jvmtiSetEventNotificationMode(jvmti.raw(), mode, event_type, event_thread)
}

func (jvmti jvmtiEnv) getAllModules(module_count_ptr *C.jint, modules_ptr **C.jobject) C.jvmtiError {
  return C.jvmtiGetAllModules(jvmti.raw(), module_count_ptr, modules_ptr)
}

func (jvmti jvmtiEnv) getAllThreads(threads_count_ptr *C.jint, threads_ptr **C.jthread) C.jvmtiError {
  return C.jvmtiGetAllThreads(jvmti.raw(), threads_count_ptr, threads_ptr)
}

func (jvmti jvmtiEnv) suspendThread(thread C.jthread) C.jvmtiError {
  return C.jvmtiSuspendThread(jvmti.raw(), thread)
}

func (jvmti jvmtiEnv) resumeThread(thread C.jthread) C.jvmtiError {
  return C.jvmtiResumeThread(jvmti.raw(), thread)
}

func (jvmti jvmtiEnv) stopThread(thread C.jthread, exception C.jobject) C.jvmtiError {
  return C.jvmtiStopThread(jvmti.raw(), thread, exception)
}

func (jvmti jvmtiEnv) interruptThread(thread C.jthread) C.jvmtiError {
  return C.jvmtiInterruptThread(jvmti.raw(), thread)
}

func (jvmti jvmtiEnv) getThreadInfo(thread C.jthread, info_ptr *C.jvmtiThreadInfo) C.jvmtiError {
  return C.jvmtiGetThreadInfo(jvmti.raw(), thread, info_ptr)
}

func (jvmti jvmtiEnv) getOwnedMonitorInfo(thread C.jthread, owned_monitor_count_ptr *C.jint, owned_monitors_ptr **C.jobject) C.jvmtiError {
  return C.jvmtiGetOwnedMonitorInfo(jvmti.raw(), thread, owned_monitor_count_ptr, owned_monitors_ptr)
}

func (jvmti jvmtiEnv) getCurrentContendedMonitor(thread C.jthread, monitor_ptr *C.jobject) C.jvmtiError {
  return C.jvmtiGetCurrentContendedMonitor(jvmti.raw(), thread, monitor_ptr)
}

func (jvmti jvmtiEnv) runAgentThread(thread C.jthread, proc C.jvmtiStartFunction, arg unsafe.Pointer, priority C.jint) C.jvmtiError {
  return C.jvmtiRunAgentThread(jvmti.raw(), thread, proc, arg, priority)
}

func (jvmti jvmtiEnv) getTopThreadGroups(group_count_ptr *C.jint, groups_ptr **C.jthreadGroup) C.jvmtiError {
  return C.jvmtiGetTopThreadGroups(jvmti.raw(), group_count_ptr, groups_ptr)
}

func (jvmti jvmtiEnv) getThreadGroupInfo(group C.jthreadGroup, info_ptr *C.jvmtiThreadGroupInfo) C.jvmtiError {
  return C.jvmtiGetThreadGroupInfo(jvmti.raw(), group, info_ptr)
}

func (jvmti jvmtiEnv) getThreadGroupChildren(group C.jthreadGroup, thread_count_ptr *C.jint, threads_ptr **C.jthread, group_count_ptr *C.jint, groups_ptr **C.jthreadGroup) C.jvmtiError {
  return C.jvmtiGetThreadGroupChildren(jvmti.raw(), group, thread_count_ptr, threads_ptr, group_count_ptr, groups_ptr)
}

func (jvmti jvmtiEnv) getFrameCount(thread C.jthread, count_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetFrameCount(jvmti.raw(), thread, count_ptr)
}

func (jvmti jvmtiEnv) getThreadState(thread C.jthread, thread_state_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetThreadState(jvmti.raw(), thread, thread_state_ptr)
}

func (jvmti jvmtiEnv) getCurrentThread(thread_ptr *C.jthread) C.jvmtiError {
  return C.jvmtiGetCurrentThread(jvmti.raw(), thread_ptr)
}

func (jvmti jvmtiEnv) getFrameLocation(thread C.jthread, depth C.jint, method_ptr *C.jmethodID, location_ptr *C.jlocation) C.jvmtiError {
  return C.jvmtiGetFrameLocation(jvmti.raw(), thread, depth, method_ptr, location_ptr)
}

func (jvmti jvmtiEnv) notifyFramePop(thread C.jthread, depth C.jint) C.jvmtiError {
  return C.jvmtiNotifyFramePop(jvmti.raw(), thread, depth)
}

func (jvmti jvmtiEnv) getLocalObject(thread C.jthread, depth C.jint, slot C.jint, value_ptr *C.jobject) C.jvmtiError {
  return C.jvmtiGetLocalObject(jvmti.raw(), thread, depth, slot, value_ptr)
}

func (jvmti jvmtiEnv) getLocalInt(thread C.jthread, depth C.jint, slot C.jint, value_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetLocalInt(jvmti.raw(), thread, depth, slot, value_ptr)
}

func (jvmti jvmtiEnv) getLocalLong(thread C.jthread, depth C.jint, slot C.jint, value_ptr *C.jlong) C.jvmtiError {
  return C.jvmtiGetLocalLong(jvmti.raw(), thread, depth, slot, value_ptr)
}

func (jvmti jvmtiEnv) getLocalFloat(thread C.jthread, depth C.jint, slot C.jint, value_ptr *C.jfloat) C.jvmtiError {
  return C.jvmtiGetLocalFloat(jvmti.raw(), thread, depth, slot, value_ptr)
}

func (jvmti jvmtiEnv) getLocalDouble(thread C.jthread, depth C.jint, slot C.jint, value_ptr *C.jdouble) C.jvmtiError {
  return C.jvmtiGetLocalDouble(jvmti.raw(), thread, depth, slot, value_ptr)
}

func (jvmti jvmtiEnv) setLocalObject(thread C.jthread, depth C.jint, slot C.jint, value C.jobject) C.jvmtiError {
  return C.jvmtiSetLocalObject(jvmti.raw(), thread, depth, slot, value)
}

func (jvmti jvmtiEnv) setLocalInt(thread C.jthread, depth C.jint, slot C.jint, value C.jint) C.jvmtiError {
  return C.jvmtiSetLocalInt(jvmti.raw(), thread, depth, slot, value)
}

func (jvmti jvmtiEnv) setLocalLong(thread C.jthread, depth C.jint, slot C.jint, value C.jlong) C.jvmtiError {
  return C.jvmtiSetLocalLong(jvmti.raw(), thread, depth, slot, value)
}

func (jvmti jvmtiEnv) setLocalFloat(thread C.jthread, depth C.jint, slot C.jint, value C.jfloat) C.jvmtiError {
  return C.jvmtiSetLocalFloat(jvmti.raw(), thread, depth, slot, value)
}

func (jvmti jvmtiEnv) setLocalDouble(thread C.jthread, depth C.jint, slot C.jint, value C.jdouble) C.jvmtiError {
  return C.jvmtiSetLocalDouble(jvmti.raw(), thread, depth, slot, value)
}

func (jvmti jvmtiEnv) createRawMonitor(name *C.char, monitor_ptr *C.jrawMonitorID) C.jvmtiError {
  return C.jvmtiCreateRawMonitor(jvmti.raw(), name, monitor_ptr)
}

func (jvmti jvmtiEnv) destroyRawMonitor(monitor C.jrawMonitorID) C.jvmtiError {
  return C.jvmtiDestroyRawMonitor(jvmti.raw(), monitor)
}

func (jvmti jvmtiEnv) rawMonitorEnter(monitor C.jrawMonitorID) C.jvmtiError {
  return C.jvmtiRawMonitorEnter(jvmti.raw(), monitor)
}

func (jvmti jvmtiEnv) rawMonitorExit(monitor C.jrawMonitorID) C.jvmtiError {
  return C.jvmtiRawMonitorExit(jvmti.raw(), monitor)
}

func (jvmti jvmtiEnv) rawMonitorWait(monitor C.jrawMonitorID, millis C.jlong) C.jvmtiError {
  return C.jvmtiRawMonitorWait(jvmti.raw(), monitor, millis)
}

func (jvmti jvmtiEnv) rawMonitorNotify(monitor C.jrawMonitorID) C.jvmtiError {
  return C.jvmtiRawMonitorNotify(jvmti.raw(), monitor)
}

func (jvmti jvmtiEnv) rawMonitorNotifyAll(monitor C.jrawMonitorID) C.jvmtiError {
  return C.jvmtiRawMonitorNotifyAll(jvmti.raw(), monitor)
}

func (jvmti jvmtiEnv) setBreakpoint(method C.jmethodID, location C.jlocation) C.jvmtiError {
  return C.jvmtiSetBreakpoint(jvmti.raw(), method, location)
}

func (jvmti jvmtiEnv) clearBreakpoint(method C.jmethodID, location C.jlocation) C.jvmtiError {
  return C.jvmtiClearBreakpoint(jvmti.raw(), method, location)
}

func (jvmti jvmtiEnv) getNamedModule(class_loader C.jobject, package_name *C.char, module_ptr *C.jobject) C.jvmtiError {
  return C.jvmtiGetNamedModule(jvmti.raw(), class_loader, package_name, module_ptr)
}

func (jvmti jvmtiEnv) setFieldAccessWatch(klass C.jclass, field C.jfieldID) C.jvmtiError {
  return C.jvmtiSetFieldAccessWatch(jvmti.raw(), klass, field)
}

func (jvmti jvmtiEnv) clearFieldAccessWatch(klass C.jclass, field C.jfieldID) C.jvmtiError {
  return C.jvmtiClearFieldAccessWatch(jvmti.raw(), klass, field)
}

func (jvmti jvmtiEnv) setFieldModificationWatch(klass C.jclass, field C.jfieldID) C.jvmtiError {
  return C.jvmtiSetFieldModificationWatch(jvmti.raw(), klass, field)
}

func (jvmti jvmtiEnv) clearFieldModificationWatch(klass C.jclass, field C.jfieldID) C.jvmtiError {
  return C.jvmtiClearFieldModificationWatch(jvmti.raw(), klass, field)
}

func (jvmti jvmtiEnv) isModifiableClass(klass C.jclass, is_modifiable_class_ptr *C.jboolean) C.jvmtiError {
  return C.jvmtiIsModifiableClass(jvmti.raw(), klass, is_modifiable_class_ptr)
}

func (jvmti jvmtiEnv) allocate(size C.jlong, mem_ptr **C.uchar) C.jvmtiError {
  return C.jvmtiAllocate(jvmti.raw(), size, mem_ptr)
}

func (jvmti jvmtiEnv) deallocate(mem *C.uchar) C.jvmtiError {
  return C.jvmtiDeallocate(jvmti.raw(), mem)
}

func (jvmti jvmtiEnv) getClassSignature(klass C.jclass, signature_ptr **C.char, generic_ptr **C.char) C.jvmtiError {
  return C.jvmtiGetClassSignature(jvmti.raw(), klass, signature_ptr, generic_ptr)
}

func (jvmti jvmtiEnv) getClassStatus(klass C.jclass, status_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetClassStatus(jvmti.raw(), klass, status_ptr)
}

func (jvmti jvmtiEnv) getSourceFileName(klass C.jclass, source_name_ptr **C.char) C.jvmtiError {
  return C.jvmtiGetSourceFileName(jvmti.raw(), klass, source_name_ptr)
}

func (jvmti jvmtiEnv) getClassModifiers(klass C.jclass, modifiers_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetClassModifiers(jvmti.raw(), klass, modifiers_ptr)
}

func (jvmti jvmtiEnv) getClassMethods(klass C.jclass, method_count_ptr *C.jint, methods_ptr **C.jmethodID) C.jvmtiError {
  return C.jvmtiGetClassMethods(jvmti.raw(), klass, method_count_ptr, methods_ptr)
}

func (jvmti jvmtiEnv) getClassFields(klass C.jclass, field_count_ptr *C.jint, fields_ptr **C.jfieldID) C.jvmtiError {
  return C.jvmtiGetClassFields(jvmti.raw(), klass, field_count_ptr, fields_ptr)
}

func (jvmti jvmtiEnv) getImplementedInterfaces(klass C.jclass, interface_count_ptr *C.jint, interfaces_ptr **C.jclass) C.jvmtiError {
  return C.jvmtiGetImplementedInterfaces(jvmti.raw(), klass, interface_count_ptr, interfaces_ptr)
}

func (jvmti jvmtiEnv) isInterface(klass C.jclass, is_interface_ptr *C.jboolean) C.jvmtiError {
  return C.jvmtiIsInterface(jvmti.raw(), klass, is_interface_ptr)
}

func (jvmti jvmtiEnv) isArrayClass(klass C.jclass, is_array_class_ptr *C.jboolean) C.jvmtiError {
  return C.jvmtiIsArrayClass(jvmti.raw(), klass, is_array_class_ptr)
}

func (jvmti jvmtiEnv) getClassLoader(klass C.jclass, classloader_ptr *C.jobject) C.jvmtiError {
  return C.jvmtiGetClassLoader(jvmti.raw(), klass, classloader_ptr)
}

func (jvmti jvmtiEnv) getObjectHashCode(object C.jobject, hash_code_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetObjectHashCode(jvmti.raw(), object, hash_code_ptr)
}

func (jvmti jvmtiEnv) getObjectMonitorUsage(object C.jobject, info_ptr *C.jvmtiMonitorUsage) C.jvmtiError {
  return C.jvmtiGetObjectMonitorUsage(jvmti.raw(), object, info_ptr)
}

func (jvmti jvmtiEnv) getFieldName(klass C.jclass, field C.jfieldID, name_ptr **C.char, signature_ptr **C.char, generic_ptr **C.char) C.jvmtiError {
  return C.jvmtiGetFieldName(jvmti.raw(), klass, field, name_ptr, signature_ptr, generic_ptr)
}

func (jvmti jvmtiEnv) getFieldDeclaringClass(klass C.jclass, field C.jfieldID, declaring_class_ptr *C.jclass) C.jvmtiError {
  return C.jvmtiGetFieldDeclaringClass(jvmti.raw(), klass, field, declaring_class_ptr)
}

func (jvmti jvmtiEnv) getFieldModifiers(klass C.jclass, field C.jfieldID, modifiers_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetFieldModifiers(jvmti.raw(), klass, field, modifiers_ptr)
}

func (jvmti jvmtiEnv) isFieldSynthetic(klass C.jclass, field C.jfieldID, is_synthetic_ptr *C.jboolean) C.jvmtiError {
  return C.jvmtiIsFieldSynthetic(jvmti.raw(), klass, field, is_synthetic_ptr)
}

func (jvmti jvmtiEnv) getMethodName(method C.jmethodID, name_ptr **C.char, signature_ptr **C.char, generic_ptr **C.char) C.jvmtiError {
  return C.jvmtiGetMethodName(jvmti.raw(), method, name_ptr, signature_ptr, generic_ptr)
}

func (jvmti jvmtiEnv) getMethodDeclaringClass(method C.jmethodID, declaring_class_ptr *C.jclass) C.jvmtiError {
  return C.jvmtiGetMethodDeclaringClass(jvmti.raw(), method, declaring_class_ptr)
}

func (jvmti jvmtiEnv) getMethodModifiers(method C.jmethodID, modifiers_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetMethodModifiers(jvmti.raw(), method, modifiers_ptr)
}

func (jvmti jvmtiEnv) getMaxLocals(method C.jmethodID, max_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetMaxLocals(jvmti.raw(), method, max_ptr)
}

func (jvmti jvmtiEnv) getArgumentsSize(method C.jmethodID, size_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetArgumentsSize(jvmti.raw(), method, size_ptr)
}

func (jvmti jvmtiEnv) getLineNumberTable(method C.jmethodID, entry_count_ptr *C.jint, table_ptr **C.jvmtiLineNumberEntry) C.jvmtiError {
  return C.jvmtiGetLineNumberTable(jvmti.raw(), method, entry_count_ptr, table_ptr)
}

func (jvmti jvmtiEnv) getMethodLocation(method C.jmethodID, start_location_ptr *C.jlocation, end_location_ptr *C.jlocation) C.jvmtiError {
  return C.jvmtiGetMethodLocation(jvmti.raw(), method, start_location_ptr, end_location_ptr)
}

func (jvmti jvmtiEnv) getLocalVariableTable(method C.jmethodID, entry_count_ptr *C.jint, table_ptr **C.jvmtiLocalVariableEntry) C.jvmtiError {
  return C.jvmtiGetLocalVariableTable(jvmti.raw(), method, entry_count_ptr, table_ptr)
}

func (jvmti jvmtiEnv) setNativeMethodPrefix(prefix *C.char) C.jvmtiError {
  return C.jvmtiSetNativeMethodPrefix(jvmti.raw(), prefix)
}

func (jvmti jvmtiEnv) setNativeMethodPrefixes(prefix_count C.jint, prefixes **C.char) C.jvmtiError {
  return C.jvmtiSetNativeMethodPrefixes(jvmti.raw(), prefix_count, prefixes)
}

func (jvmti jvmtiEnv) getBytecodes(method C.jmethodID, bytecode_count_ptr *C.jint, bytecodes_ptr **C.uchar) C.jvmtiError {
  return C.jvmtiGetBytecodes(jvmti.raw(), method, bytecode_count_ptr, bytecodes_ptr)
}

func (jvmti jvmtiEnv) isMethodNative(method C.jmethodID, is_native_ptr *C.jboolean) C.jvmtiError {
  return C.jvmtiIsMethodNative(jvmti.raw(), method, is_native_ptr)
}

func (jvmti jvmtiEnv) isMethodSynthetic(method C.jmethodID, is_synthetic_ptr *C.jboolean) C.jvmtiError {
  return C.jvmtiIsMethodSynthetic(jvmti.raw(), method, is_synthetic_ptr)
}

func (jvmti jvmtiEnv) getLoadedClasses(class_count_ptr *C.jint, classes_ptr **C.jclass) C.jvmtiError {
  return C.jvmtiGetLoadedClasses(jvmti.raw(), class_count_ptr, classes_ptr)
}

func (jvmti jvmtiEnv) getClassLoaderClasses(initiating_loader C.jobject, class_count_ptr *C.jint, classes_ptr **C.jclass) C.jvmtiError {
  return C.jvmtiGetClassLoaderClasses(jvmti.raw(), initiating_loader, class_count_ptr, classes_ptr)
}

func (jvmti jvmtiEnv) popFrame(thread C.jthread) C.jvmtiError {
  return C.jvmtiPopFrame(jvmti.raw(), thread)
}

func (jvmti jvmtiEnv) forceEarlyReturnObject(thread C.jthread, value C.jobject) C.jvmtiError {
  return C.jvmtiForceEarlyReturnObject(jvmti.raw(), thread, value)
}

func (jvmti jvmtiEnv) forceEarlyReturnInt(thread C.jthread, value C.jint) C.jvmtiError {
  return C.jvmtiForceEarlyReturnInt(jvmti.raw(), thread, value)
}

func (jvmti jvmtiEnv) forceEarlyReturnLong(thread C.jthread, value C.jlong) C.jvmtiError {
  return C.jvmtiForceEarlyReturnLong(jvmti.raw(), thread, value)
}

func (jvmti jvmtiEnv) forceEarlyReturnFloat(thread C.jthread, value C.jfloat) C.jvmtiError {
  return C.jvmtiForceEarlyReturnFloat(jvmti.raw(), thread, value)
}

func (jvmti jvmtiEnv) forceEarlyReturnDouble(thread C.jthread, value C.jdouble) C.jvmtiError {
  return C.jvmtiForceEarlyReturnDouble(jvmti.raw(), thread, value)
}

func (jvmti jvmtiEnv) forceEarlyReturnVoid(thread C.jthread) C.jvmtiError {
  return C.jvmtiForceEarlyReturnVoid(jvmti.raw(), thread)
}

func (jvmti jvmtiEnv) redefineClasses(class_count C.jint, class_definitions *C.jvmtiClassDefinition) C.jvmtiError {
  return C.jvmtiRedefineClasses(jvmti.raw(), class_count, class_definitions)
}

func (jvmti jvmtiEnv) getVersionNumber(version_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetVersionNumber(jvmti.raw(), version_ptr)
}

func (jvmti jvmtiEnv) getCapabilities(capabilities_ptr *C.jvmtiCapabilities) C.jvmtiError {
  return C.jvmtiGetCapabilities(jvmti.raw(), capabilities_ptr)
}

func (jvmti jvmtiEnv) getSourceDebugExtension(klass C.jclass, source_debug_extension_ptr **C.char) C.jvmtiError {
  return C.jvmtiGetSourceDebugExtension(jvmti.raw(), klass, source_debug_extension_ptr)
}

func (jvmti jvmtiEnv) isMethodObsolete(method C.jmethodID, is_obsolete_ptr *C.jboolean) C.jvmtiError {
  return C.jvmtiIsMethodObsolete(jvmti.raw(), method, is_obsolete_ptr)
}

func (jvmti jvmtiEnv) suspendThreadList(request_count C.jint, request_list *C.jthread, results *C.jvmtiError) C.jvmtiError {
  return C.jvmtiSuspendThreadList(jvmti.raw(), request_count, request_list, results)
}

func (jvmti jvmtiEnv) resumeThreadList(request_count C.jint, request_list *C.jthread, results *C.jvmtiError) C.jvmtiError {
  return C.jvmtiResumeThreadList(jvmti.raw(), request_count, request_list, results)
}

func (jvmti jvmtiEnv) addModuleReads(module C.jobject, to_module C.jobject) C.jvmtiError {
  return C.jvmtiAddModuleReads(jvmti.raw(), module, to_module)
}

func (jvmti jvmtiEnv) addModuleExports(module C.jobject, pkg_name *C.char, to_module C.jobject) C.jvmtiError {
  return C.jvmtiAddModuleExports(jvmti.raw(), module, pkg_name, to_module)
}

func (jvmti jvmtiEnv) addModuleOpens(module C.jobject, pkg_name *C.char, to_module C.jobject) C.jvmtiError {
  return C.jvmtiAddModuleOpens(jvmti.raw(), module, pkg_name, to_module)
}

func (jvmti jvmtiEnv) addModuleUses(module C.jobject, service C.jclass) C.jvmtiError {
  return C.jvmtiAddModuleUses(jvmti.raw(), module, service)
}

func (jvmti jvmtiEnv) addModuleProvides(module C.jobject, service C.jclass, impl_class C.jclass) C.jvmtiError {
  return C.jvmtiAddModuleProvides(jvmti.raw(), module, service, impl_class)
}

func (jvmti jvmtiEnv) isModifiableModule(module C.jobject, is_modifiable_module_ptr *C.jboolean) C.jvmtiError {
  return C.jvmtiIsModifiableModule(jvmti.raw(), module, is_modifiable_module_ptr)
}

func (jvmti jvmtiEnv) getAllStackTraces(max_frame_count C.jint, stack_info_ptr **C.jvmtiStackInfo, thread_count_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetAllStackTraces(jvmti.raw(), max_frame_count, stack_info_ptr, thread_count_ptr)
}

func (jvmti jvmtiEnv) getThreadListStackTraces(thread_count C.jint, thread_list *C.jthread, max_frame_count C.jint, stack_info_ptr **C.jvmtiStackInfo) C.jvmtiError {
  return C.jvmtiGetThreadListStackTraces(jvmti.raw(), thread_count, thread_list, max_frame_count, stack_info_ptr)
}

func (jvmti jvmtiEnv) getThreadLocalStorage(thread C.jthread, data_ptr *unsafe.Pointer) C.jvmtiError {
  return C.jvmtiGetThreadLocalStorage(jvmti.raw(), thread, data_ptr)
}

func (jvmti jvmtiEnv) setThreadLocalStorage(thread C.jthread, data unsafe.Pointer) C.jvmtiError {
  return C.jvmtiSetThreadLocalStorage(jvmti.raw(), thread, data)
}

func (jvmti jvmtiEnv) getStackTrace(thread C.jthread, start_depth C.jint, max_frame_count C.jint, frame_buffer *C.jvmtiFrameInfo, count_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetStackTrace(jvmti.raw(), thread, start_depth, max_frame_count, frame_buffer, count_ptr)
}

func (jvmti jvmtiEnv) getTag(object C.jobject, tag_ptr *C.jlong) C.jvmtiError {
  return C.jvmtiGetTag(jvmti.raw(), object, tag_ptr)
}

func (jvmti jvmtiEnv) setTag(object C.jobject, tag C.jlong) C.jvmtiError {
  return C.jvmtiSetTag(jvmti.raw(), object, tag)
}

func (jvmti jvmtiEnv) iterateOverObjectsReachableFromObject(object C.jobject, object_reference_callback C.jvmtiObjectReferenceCallback, user_data unsafe.Pointer) C.jvmtiError {
  return C.jvmtiIterateOverObjectsReachableFromObject(jvmti.raw(), object, object_reference_callback, user_data)
}

func (jvmti jvmtiEnv) iterateOverReachableObjects(heap_root_callback C.jvmtiHeapRootCallback, stack_ref_callback C.jvmtiStackReferenceCallback, object_ref_callback C.jvmtiObjectReferenceCallback, user_data unsafe.Pointer) C.jvmtiError {
  return C.jvmtiIterateOverReachableObjects(jvmti.raw(), heap_root_callback, stack_ref_callback, object_ref_callback, user_data)
}

func (jvmti jvmtiEnv) iterateOverHeap(object_filter C.jvmtiHeapObjectFilter, heap_object_callback C.jvmtiHeapObjectCallback, user_data unsafe.Pointer) C.jvmtiError {
  return C.jvmtiIterateOverHeap(jvmti.raw(), object_filter, heap_object_callback, user_data)
}

func (jvmti jvmtiEnv) iterateOverInstancesOfClass(klass C.jclass, object_filter C.jvmtiHeapObjectFilter, heap_object_callback C.jvmtiHeapObjectCallback, user_data unsafe.Pointer) C.jvmtiError {
  return C.jvmtiIterateOverInstancesOfClass(jvmti.raw(), klass, object_filter, heap_object_callback, user_data)
}

func (jvmti jvmtiEnv) getObjectsWithTags(tag_count C.jint, tags *C.jlong, count_ptr *C.jint, object_result_ptr **C.jobject, tag_result_ptr **C.jlong) C.jvmtiError {
  return C.jvmtiGetObjectsWithTags(jvmti.raw(), tag_count, tags, count_ptr, object_result_ptr, tag_result_ptr)
}

func (jvmti jvmtiEnv) followReferences(heap_filter C.jint, klass C.jclass, initial_object C.jobject, callbacks *C.jvmtiHeapCallbacks, user_data unsafe.Pointer) C.jvmtiError {
  return C.jvmtiFollowReferences(jvmti.raw(), heap_filter, klass, initial_object, callbacks, user_data)
}

func (jvmti jvmtiEnv) iterateThroughHeap(heap_filter C.jint, klass C.jclass, callbacks *C.jvmtiHeapCallbacks, user_data unsafe.Pointer) C.jvmtiError {
  return C.jvmtiIterateThroughHeap(jvmti.raw(), heap_filter, klass, callbacks, user_data)
}

func (jvmti jvmtiEnv) setJNIFunctionTable(function_table *C.jniNativeInterface) C.jvmtiError {
  return C.jvmtiSetJNIFunctionTable(jvmti.raw(), function_table)
}

func (jvmti jvmtiEnv) getJNIFunctionTable(function_table **C.jniNativeInterface) C.jvmtiError {
  return C.jvmtiGetJNIFunctionTable(jvmti.raw(), function_table)
}

func (jvmti jvmtiEnv) setEventCallbacks(callbacks *C.jvmtiEventCallbacks, size_of_callbacks C.jint) C.jvmtiError {
  return C.jvmtiSetEventCallbacks(jvmti.raw(), callbacks, size_of_callbacks)
}

func (jvmti jvmtiEnv) generateEvents(event_type C.jvmtiEvent) C.jvmtiError {
  return C.jvmtiGenerateEvents(jvmti.raw(), event_type)
}

func (jvmti jvmtiEnv) getExtensionFunctions(extension_count_ptr *C.jint, extensions **C.jvmtiExtensionFunctionInfo) C.jvmtiError {
  return C.jvmtiGetExtensionFunctions(jvmti.raw(), extension_count_ptr, extensions)
}

func (jvmti jvmtiEnv) getExtensionEvents(extension_count_ptr *C.jint, extensions **C.jvmtiExtensionEventInfo) C.jvmtiError {
  return C.jvmtiGetExtensionEvents(jvmti.raw(), extension_count_ptr, extensions)
}

func (jvmti jvmtiEnv) setExtensionEventCallback(extension_event_index C.jint, callback C.jvmtiExtensionEvent) C.jvmtiError {
  return C.jvmtiSetExtensionEventCallback(jvmti.raw(), extension_event_index, callback)
}

func (jvmti jvmtiEnv) getErrorName(error C.jvmtiError, name_ptr **C.char) C.jvmtiError {
  return C.jvmtiGetErrorName(jvmti.raw(), error, name_ptr)
}

func (jvmti jvmtiEnv) getJLocationFormat(format_ptr *C.jvmtiJlocationFormat) C.jvmtiError {
  return C.jvmtiGetJLocationFormat(jvmti.raw(), format_ptr)
}

func (jvmti jvmtiEnv) getSystemProperties(count_ptr *C.jint, property_ptr ***C.char) C.jvmtiError {
  return C.jvmtiGetSystemProperties(jvmti.raw(), count_ptr, property_ptr)
}

func (jvmti jvmtiEnv) getSystemProperty(property *C.char, value_ptr **C.char) C.jvmtiError {
  return C.jvmtiGetSystemProperty(jvmti.raw(), property, value_ptr)
}

func (jvmti jvmtiEnv) setSystemProperty(property *C.char, value_ptr *C.char) C.jvmtiError {
  return C.jvmtiSetSystemProperty(jvmti.raw(), property, value_ptr)
}

func (jvmti jvmtiEnv) getPhase(phase_ptr *C.jvmtiPhase) C.jvmtiError {
  return C.jvmtiGetPhase(jvmti.raw(), phase_ptr)
}

func (jvmti jvmtiEnv) getCurrentThreadCpuTimerInfo(info_ptr *C.jvmtiTimerInfo) C.jvmtiError {
  return C.jvmtiGetCurrentThreadCpuTimerInfo(jvmti.raw(), info_ptr)
}

func (jvmti jvmtiEnv) getCurrentThreadCpuTime(nanos_ptr *C.jlong) C.jvmtiError {
  return C.jvmtiGetCurrentThreadCpuTime(jvmti.raw(), nanos_ptr)
}

func (jvmti jvmtiEnv) getThreadCpuTimerInfo(info_ptr *C.jvmtiTimerInfo) C.jvmtiError {
  return C.jvmtiGetThreadCpuTimerInfo(jvmti.raw(), info_ptr)
}

func (jvmti jvmtiEnv) getThreadCpuTime(thread C.jthread, nanos_ptr *C.jlong) C.jvmtiError {
  return C.jvmtiGetThreadCpuTime(jvmti.raw(), thread, nanos_ptr)
}

func (jvmti jvmtiEnv) getTimerInfo(info_ptr *C.jvmtiTimerInfo) C.jvmtiError {
  return C.jvmtiGetTimerInfo(jvmti.raw(), info_ptr)
}

func (jvmti jvmtiEnv) getTime(nanos_ptr *C.jlong) C.jvmtiError {
  return C.jvmtiGetTime(jvmti.raw(), nanos_ptr)
}

func (jvmti jvmtiEnv) getPotentialCapabilities(capabilities_ptr *C.jvmtiCapabilities) C.jvmtiError {
  return C.jvmtiGetPotentialCapabilities(jvmti.raw(), capabilities_ptr)
}

func (jvmti jvmtiEnv) addCapabilities(capabilities_ptr *C.jvmtiCapabilities) C.jvmtiError {
  return C.jvmtiAddCapabilities(jvmti.raw(), capabilities_ptr)
}

func (jvmti jvmtiEnv) relinquishCapabilities(capabilities_ptr *C.jvmtiCapabilities) C.jvmtiError {
  return C.jvmtiRelinquishCapabilities(jvmti.raw(), capabilities_ptr)
}

func (jvmti jvmtiEnv) getAvailableProcessors(processor_count_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetAvailableProcessors(jvmti.raw(), processor_count_ptr)
}

func (jvmti jvmtiEnv) getClassVersionNumbers(klass C.jclass, minor_version_ptr *C.jint, major_version_ptr *C.jint) C.jvmtiError {
  return C.jvmtiGetClassVersionNumbers(jvmti.raw(), klass, minor_version_ptr, major_version_ptr)
}

func (jvmti jvmtiEnv) getConstantPool(klass C.jclass, constant_pool_count_ptr *C.jint, constant_pool_byte_count_ptr *C.jint, constant_pool_bytes_ptr **C.uchar) C.jvmtiError {
  return C.jvmtiGetConstantPool(jvmti.raw(), klass, constant_pool_count_ptr, constant_pool_byte_count_ptr, constant_pool_bytes_ptr)
}

func (jvmti jvmtiEnv) getEnvironmentLocalStorage(data_ptr *unsafe.Pointer) C.jvmtiError {
  return C.jvmtiGetEnvironmentLocalStorage(jvmti.raw(), data_ptr)
}

func (jvmti jvmtiEnv) setEnvironmentLocalStorage(data unsafe.Pointer) C.jvmtiError {
  return C.jvmtiSetEnvironmentLocalStorage(jvmti.raw(), data)
}

func (jvmti jvmtiEnv) addToBootstrapClassLoaderSearch(segment *C.char) C.jvmtiError {
  return C.jvmtiAddToBootstrapClassLoaderSearch(jvmti.raw(), segment)
}

func (jvmti jvmtiEnv) setVerboseFlag(flag C.jvmtiVerboseFlag, value C.jboolean) C.jvmtiError {
  return C.jvmtiSetVerboseFlag(jvmti.raw(), flag, value)
}

func (jvmti jvmtiEnv) addToSystemClassLoaderSearch(segment *C.char) C.jvmtiError {
  return C.jvmtiAddToSystemClassLoaderSearch(jvmti.raw(), segment)
}

func (jvmti jvmtiEnv) retransformClasses(class_count C.jint, classes *C.jclass) C.jvmtiError {
  return C.jvmtiRetransformClasses(jvmti.raw(), class_count, classes)
}

func (jvmti jvmtiEnv) getOwnedMonitorStackDepthInfo(thread C.jthread, monitor_info_count_ptr *C.jint, monitor_info_ptr **C.jvmtiMonitorStackDepthInfo) C.jvmtiError {
  return C.jvmtiGetOwnedMonitorStackDepthInfo(jvmti.raw(), thread, monitor_info_count_ptr, monitor_info_ptr)
}

func (jvmti jvmtiEnv) getObjectSize(object C.jobject, size_ptr *C.jlong) C.jvmtiError {
  return C.jvmtiGetObjectSize(jvmti.raw(), object, size_ptr)
}

func (jvmti jvmtiEnv) getLocalInstance(thread C.jthread, depth C.jint, value_ptr *C.jobject) C.jvmtiError {
  return C.jvmtiGetLocalInstance(jvmti.raw(), thread, depth, value_ptr)
}

func (jvmti jvmtiEnv) setHeapSamplingInterval(sampling_interval C.jint) C.jvmtiError {
  return C.jvmtiSetHeapSamplingInterval(jvmti.raw(), sampling_interval)
}

func describeJvmtiError(err int) string {
  switch (err) {

	case JVMTI_ERROR_NONE:
		return "JVMTI_ERROR_NONE"

	case JVMTI_ERROR_INVALID_THREAD:
		return "JVMTI_ERROR_INVALID_THREAD"

	case JVMTI_ERROR_INVALID_THREAD_GROUP:
		return "JVMTI_ERROR_INVALID_THREAD_GROUP"

	case JVMTI_ERROR_INVALID_PRIORITY:
		return "JVMTI_ERROR_INVALID_PRIORITY"

	case JVMTI_ERROR_THREAD_NOT_SUSPENDED:
		return "JVMTI_ERROR_THREAD_NOT_SUSPENDED"

	case JVMTI_ERROR_THREAD_SUSPENDED:
		return "JVMTI_ERROR_THREAD_SUSPENDED"

	case JVMTI_ERROR_THREAD_NOT_ALIVE:
		return "JVMTI_ERROR_THREAD_NOT_ALIVE"

	case JVMTI_ERROR_INVALID_OBJECT:
		return "JVMTI_ERROR_INVALID_OBJECT"

	case JVMTI_ERROR_INVALID_CLASS:
		return "JVMTI_ERROR_INVALID_CLASS"

	case JVMTI_ERROR_CLASS_NOT_PREPARED:
		return "JVMTI_ERROR_CLASS_NOT_PREPARED"

	case JVMTI_ERROR_INVALID_METHODID:
		return "JVMTI_ERROR_INVALID_METHODID"

	case JVMTI_ERROR_INVALID_LOCATION:
		return "JVMTI_ERROR_INVALID_LOCATION"

	case JVMTI_ERROR_INVALID_FIELDID:
		return "JVMTI_ERROR_INVALID_FIELDID"

	case JVMTI_ERROR_INVALID_MODULE:
		return "JVMTI_ERROR_INVALID_MODULE"

	case JVMTI_ERROR_NO_MORE_FRAMES:
		return "JVMTI_ERROR_NO_MORE_FRAMES"

	case JVMTI_ERROR_OPAQUE_FRAME:
		return "JVMTI_ERROR_OPAQUE_FRAME"

	case JVMTI_ERROR_TYPE_MISMATCH:
		return "JVMTI_ERROR_TYPE_MISMATCH"

	case JVMTI_ERROR_INVALID_SLOT:
		return "JVMTI_ERROR_INVALID_SLOT"

	case JVMTI_ERROR_DUPLICATE:
		return "JVMTI_ERROR_DUPLICATE"

	case JVMTI_ERROR_NOT_FOUND:
		return "JVMTI_ERROR_NOT_FOUND"

	case JVMTI_ERROR_INVALID_MONITOR:
		return "JVMTI_ERROR_INVALID_MONITOR"

	case JVMTI_ERROR_NOT_MONITOR_OWNER:
		return "JVMTI_ERROR_NOT_MONITOR_OWNER"

	case JVMTI_ERROR_INTERRUPT:
		return "JVMTI_ERROR_INTERRUPT"

	case JVMTI_ERROR_INVALID_CLASS_FORMAT:
		return "JVMTI_ERROR_INVALID_CLASS_FORMAT"

	case JVMTI_ERROR_CIRCULAR_CLASS_DEFINITION:
		return "JVMTI_ERROR_CIRCULAR_CLASS_DEFINITION"

	case JVMTI_ERROR_FAILS_VERIFICATION:
		return "JVMTI_ERROR_FAILS_VERIFICATION"

	case JVMTI_ERROR_UNSUPPORTED_REDEFINITION_METHOD_ADDED:
		return "JVMTI_ERROR_UNSUPPORTED_REDEFINITION_METHOD_ADDED"

	case JVMTI_ERROR_UNSUPPORTED_REDEFINITION_SCHEMA_CHANGED:
		return "JVMTI_ERROR_UNSUPPORTED_REDEFINITION_SCHEMA_CHANGED"

	case JVMTI_ERROR_INVALID_TYPESTATE:
		return "JVMTI_ERROR_INVALID_TYPESTATE"

	case JVMTI_ERROR_UNSUPPORTED_REDEFINITION_HIERARCHY_CHANGED:
		return "JVMTI_ERROR_UNSUPPORTED_REDEFINITION_HIERARCHY_CHANGED"

	case JVMTI_ERROR_UNSUPPORTED_REDEFINITION_METHOD_DELETED:
		return "JVMTI_ERROR_UNSUPPORTED_REDEFINITION_METHOD_DELETED"

	case JVMTI_ERROR_UNSUPPORTED_VERSION:
		return "JVMTI_ERROR_UNSUPPORTED_VERSION"

	case JVMTI_ERROR_NAMES_DONT_MATCH:
		return "JVMTI_ERROR_NAMES_DONT_MATCH"

	case JVMTI_ERROR_UNSUPPORTED_REDEFINITION_CLASS_MODIFIERS_CHANGED:
		return "JVMTI_ERROR_UNSUPPORTED_REDEFINITION_CLASS_MODIFIERS_CHANGED"

	case JVMTI_ERROR_UNSUPPORTED_REDEFINITION_METHOD_MODIFIERS_CHANGED:
		return "JVMTI_ERROR_UNSUPPORTED_REDEFINITION_METHOD_MODIFIERS_CHANGED"

	case JVMTI_ERROR_UNSUPPORTED_REDEFINITION_CLASS_ATTRIBUTE_CHANGED:
		return "JVMTI_ERROR_UNSUPPORTED_REDEFINITION_CLASS_ATTRIBUTE_CHANGED"

	case JVMTI_ERROR_UNMODIFIABLE_CLASS:
		return "JVMTI_ERROR_UNMODIFIABLE_CLASS"

	case JVMTI_ERROR_UNMODIFIABLE_MODULE:
		return "JVMTI_ERROR_UNMODIFIABLE_MODULE"

	case JVMTI_ERROR_NOT_AVAILABLE:
		return "JVMTI_ERROR_NOT_AVAILABLE"

	case JVMTI_ERROR_MUST_POSSESS_CAPABILITY:
		return "JVMTI_ERROR_MUST_POSSESS_CAPABILITY"

	case JVMTI_ERROR_NULL_POINTER:
		return "JVMTI_ERROR_NULL_POINTER"

	case JVMTI_ERROR_ABSENT_INFORMATION:
		return "JVMTI_ERROR_ABSENT_INFORMATION"

	case JVMTI_ERROR_INVALID_EVENT_TYPE:
		return "JVMTI_ERROR_INVALID_EVENT_TYPE"

	case JVMTI_ERROR_ILLEGAL_ARGUMENT:
		return "JVMTI_ERROR_ILLEGAL_ARGUMENT"

	case JVMTI_ERROR_NATIVE_METHOD:
		return "JVMTI_ERROR_NATIVE_METHOD"

	case JVMTI_ERROR_CLASS_LOADER_UNSUPPORTED:
		return "JVMTI_ERROR_CLASS_LOADER_UNSUPPORTED"

	case JVMTI_ERROR_OUT_OF_MEMORY:
		return "JVMTI_ERROR_OUT_OF_MEMORY"

	case JVMTI_ERROR_ACCESS_DENIED:
		return "JVMTI_ERROR_ACCESS_DENIED"

	case JVMTI_ERROR_WRONG_PHASE:
		return "JVMTI_ERROR_WRONG_PHASE"

	case JVMTI_ERROR_INTERNAL:
		return "JVMTI_ERROR_INTERNAL"

	case JVMTI_ERROR_UNATTACHED_THREAD:
		return "JVMTI_ERROR_UNATTACHED_THREAD"

	case JVMTI_ERROR_INVALID_ENVIRONMENT:
		return "JVMTI_ERROR_INVALID_ENVIRONMENT"
default:
		panic(fmt.Sprintf("Unknown JVMTI error code: %d", err))
	}
	return ""
}