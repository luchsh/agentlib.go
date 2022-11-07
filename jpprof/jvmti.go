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

package jpprof

// #include <jvmti.h>
//
// static jvmtiError SetEventNotificationMode(jvmtiEnv* env, jvmtiEventMode mode, jvmtiEvent event_type, jthread event_thread, ... ) {
//   return (*env)->SetEventNotificationMode(env, mode, event_type, event_thread);
// }
// static jvmtiError GetAllModules(jvmtiEnv* env, jint* module_count_ptr, jobject** modules_ptr) {
//   return (*env)->GetAllModules(env, module_count_ptr, modules_ptr);
// }
// static jvmtiError GetAllThreads(jvmtiEnv* env, jint* threads_count_ptr, jthread** threads_ptr) {
//   return (*env)->GetAllThreads(env, threads_count_ptr, threads_ptr);
// }
// static jvmtiError SuspendThread(jvmtiEnv* env, jthread thread) {
//   return (*env)->SuspendThread(env, thread);
// }
// static jvmtiError ResumeThread(jvmtiEnv* env, jthread thread) {
//   return (*env)->ResumeThread(env, thread);
// }
// static jvmtiError StopThread(jvmtiEnv* env, jthread thread, jobject exception) {
//   return (*env)->StopThread(env, thread, exception);
// }
// static jvmtiError InterruptThread(jvmtiEnv* env, jthread thread) {
//   return (*env)->InterruptThread(env, thread);
// }
// static jvmtiError GetThreadInfo(jvmtiEnv* env, jthread thread, jvmtiThreadInfo* info_ptr) {
//   return (*env)->GetThreadInfo(env, thread, info_ptr);
// }
// static jvmtiError GetOwnedMonitorInfo(jvmtiEnv* env, jthread thread, jint* owned_monitor_count_ptr, jobject** owned_monitors_ptr) {
//   return (*env)->GetOwnedMonitorInfo(env, thread, owned_monitor_count_ptr, owned_monitors_ptr);
// }
// static jvmtiError GetCurrentContendedMonitor(jvmtiEnv* env, jthread thread, jobject* monitor_ptr) {
//   return (*env)->GetCurrentContendedMonitor(env, thread, monitor_ptr);
// }
// static jvmtiError RunAgentThread(jvmtiEnv* env, jthread thread, jvmtiStartFunction proc, const void* arg, jint priority) {
//   return (*env)->RunAgentThread(env, thread, proc, arg, priority);
// }
// static jvmtiError GetTopThreadGroups(jvmtiEnv* env, jint* group_count_ptr, jthreadGroup** groups_ptr) {
//   return (*env)->GetTopThreadGroups(env, group_count_ptr, groups_ptr);
// }
// static jvmtiError GetThreadGroupInfo(jvmtiEnv* env, jthreadGroup group, jvmtiThreadGroupInfo* info_ptr) {
//   return (*env)->GetThreadGroupInfo(env, group, info_ptr);
// }
// static jvmtiError GetThreadGroupChildren(jvmtiEnv* env, jthreadGroup group, jint* thread_count_ptr, jthread** threads_ptr, jint* group_count_ptr, jthreadGroup** groups_ptr) {
//   return (*env)->GetThreadGroupChildren(env, group, thread_count_ptr, threads_ptr, group_count_ptr, groups_ptr);
// }
// static jvmtiError GetFrameCount(jvmtiEnv* env, jthread thread, jint* count_ptr) {
//   return (*env)->GetFrameCount(env, thread, count_ptr);
// }
// static jvmtiError GetThreadState(jvmtiEnv* env, jthread thread, jint* thread_state_ptr) {
//   return (*env)->GetThreadState(env, thread, thread_state_ptr);
// }
// static jvmtiError GetCurrentThread(jvmtiEnv* env, jthread* thread_ptr) {
//   return (*env)->GetCurrentThread(env, thread_ptr);
// }
// static jvmtiError GetFrameLocation(jvmtiEnv* env, jthread thread, jint depth, jmethodID* method_ptr, jlocation* location_ptr) {
//   return (*env)->GetFrameLocation(env, thread, depth, method_ptr, location_ptr);
// }
// static jvmtiError NotifyFramePop(jvmtiEnv* env, jthread thread, jint depth) {
//   return (*env)->NotifyFramePop(env, thread, depth);
// }
// static jvmtiError GetLocalObject(jvmtiEnv* env, jthread thread, jint depth, jint slot, jobject* value_ptr) {
//   return (*env)->GetLocalObject(env, thread, depth, slot, value_ptr);
// }
// static jvmtiError GetLocalInt(jvmtiEnv* env, jthread thread, jint depth, jint slot, jint* value_ptr) {
//   return (*env)->GetLocalInt(env, thread, depth, slot, value_ptr);
// }
// static jvmtiError GetLocalLong(jvmtiEnv* env, jthread thread, jint depth, jint slot, jlong* value_ptr) {
//   return (*env)->GetLocalLong(env, thread, depth, slot, value_ptr);
// }
// static jvmtiError GetLocalFloat(jvmtiEnv* env, jthread thread, jint depth, jint slot, jfloat* value_ptr) {
//   return (*env)->GetLocalFloat(env, thread, depth, slot, value_ptr);
// }
// static jvmtiError GetLocalDouble(jvmtiEnv* env, jthread thread, jint depth, jint slot, jdouble* value_ptr) {
//   return (*env)->GetLocalDouble(env, thread, depth, slot, value_ptr);
// }
// static jvmtiError SetLocalObject(jvmtiEnv* env, jthread thread, jint depth, jint slot, jobject value) {
//   return (*env)->SetLocalObject(env, thread, depth, slot, value);
// }
// static jvmtiError SetLocalInt(jvmtiEnv* env, jthread thread, jint depth, jint slot, jint value) {
//   return (*env)->SetLocalInt(env, thread, depth, slot, value);
// }
// static jvmtiError SetLocalLong(jvmtiEnv* env, jthread thread, jint depth, jint slot, jlong value) {
//   return (*env)->SetLocalLong(env, thread, depth, slot, value);
// }
// static jvmtiError SetLocalFloat(jvmtiEnv* env, jthread thread, jint depth, jint slot, jfloat value) {
//   return (*env)->SetLocalFloat(env, thread, depth, slot, value);
// }
// static jvmtiError SetLocalDouble(jvmtiEnv* env, jthread thread, jint depth, jint slot, jdouble value) {
//   return (*env)->SetLocalDouble(env, thread, depth, slot, value);
// }
// static jvmtiError CreateRawMonitor(jvmtiEnv* env, const char* name, jrawMonitorID* monitor_ptr) {
//   return (*env)->CreateRawMonitor(env, name, monitor_ptr);
// }
// static jvmtiError DestroyRawMonitor(jvmtiEnv* env, jrawMonitorID monitor) {
//   return (*env)->DestroyRawMonitor(env, monitor);
// }
// static jvmtiError RawMonitorEnter(jvmtiEnv* env, jrawMonitorID monitor) {
//   return (*env)->RawMonitorEnter(env, monitor);
// }
// static jvmtiError RawMonitorExit(jvmtiEnv* env, jrawMonitorID monitor) {
//   return (*env)->RawMonitorExit(env, monitor);
// }
// static jvmtiError RawMonitorWait(jvmtiEnv* env, jrawMonitorID monitor, jlong millis) {
//   return (*env)->RawMonitorWait(env, monitor, millis);
// }
// static jvmtiError RawMonitorNotify(jvmtiEnv* env, jrawMonitorID monitor) {
//   return (*env)->RawMonitorNotify(env, monitor);
// }
// static jvmtiError RawMonitorNotifyAll(jvmtiEnv* env, jrawMonitorID monitor) {
//   return (*env)->RawMonitorNotifyAll(env, monitor);
// }
// static jvmtiError SetBreakpoint(jvmtiEnv* env, jmethodID method, jlocation location) {
//   return (*env)->SetBreakpoint(env, method, location);
// }
// static jvmtiError ClearBreakpoint(jvmtiEnv* env, jmethodID method, jlocation location) {
//   return (*env)->ClearBreakpoint(env, method, location);
// }
// static jvmtiError GetNamedModule(jvmtiEnv* env, jobject class_loader, const char* package_name, jobject* module_ptr) {
//   return (*env)->GetNamedModule(env, class_loader, package_name, module_ptr);
// }
// static jvmtiError SetFieldAccessWatch(jvmtiEnv* env, jclass klass, jfieldID field) {
//   return (*env)->SetFieldAccessWatch(env, klass, field);
// }
// static jvmtiError ClearFieldAccessWatch(jvmtiEnv* env, jclass klass, jfieldID field) {
//   return (*env)->ClearFieldAccessWatch(env, klass, field);
// }
// static jvmtiError SetFieldModificationWatch(jvmtiEnv* env, jclass klass, jfieldID field) {
//   return (*env)->SetFieldModificationWatch(env, klass, field);
// }
// static jvmtiError ClearFieldModificationWatch(jvmtiEnv* env, jclass klass, jfieldID field) {
//   return (*env)->ClearFieldModificationWatch(env, klass, field);
// }
// static jvmtiError IsModifiableClass(jvmtiEnv* env, jclass klass, jboolean* is_modifiable_class_ptr) {
//   return (*env)->IsModifiableClass(env, klass, is_modifiable_class_ptr);
// }
// static jvmtiError Allocate(jvmtiEnv* env, jlong size, unsigned char** mem_ptr) {
//   return (*env)->Allocate(env, size, mem_ptr);
// }
// static jvmtiError Deallocate(jvmtiEnv* env, unsigned char* mem) {
//   return (*env)->Deallocate(env, mem);
// }
// static jvmtiError GetClassSignature(jvmtiEnv* env, jclass klass, char** signature_ptr, char** generic_ptr) {
//   return (*env)->GetClassSignature(env, klass, signature_ptr, generic_ptr);
// }
// static jvmtiError GetClassStatus(jvmtiEnv* env, jclass klass, jint* status_ptr) {
//   return (*env)->GetClassStatus(env, klass, status_ptr);
// }
// static jvmtiError GetSourceFileName(jvmtiEnv* env, jclass klass, char** source_name_ptr) {
//   return (*env)->GetSourceFileName(env, klass, source_name_ptr);
// }
// static jvmtiError GetClassModifiers(jvmtiEnv* env, jclass klass, jint* modifiers_ptr) {
//   return (*env)->GetClassModifiers(env, klass, modifiers_ptr);
// }
// static jvmtiError GetClassMethods(jvmtiEnv* env, jclass klass, jint* method_count_ptr, jmethodID** methods_ptr) {
//   return (*env)->GetClassMethods(env, klass, method_count_ptr, methods_ptr);
// }
// static jvmtiError GetClassFields(jvmtiEnv* env, jclass klass, jint* field_count_ptr, jfieldID** fields_ptr) {
//   return (*env)->GetClassFields(env, klass, field_count_ptr, fields_ptr);
// }
// static jvmtiError GetImplementedInterfaces(jvmtiEnv* env, jclass klass, jint* interface_count_ptr, jclass** interfaces_ptr) {
//   return (*env)->GetImplementedInterfaces(env, klass, interface_count_ptr, interfaces_ptr);
// }
// static jvmtiError IsInterface(jvmtiEnv* env, jclass klass, jboolean* is_interface_ptr) {
//   return (*env)->IsInterface(env, klass, is_interface_ptr);
// }
// static jvmtiError IsArrayClass(jvmtiEnv* env, jclass klass, jboolean* is_array_class_ptr) {
//   return (*env)->IsArrayClass(env, klass, is_array_class_ptr);
// }
// static jvmtiError GetClassLoader(jvmtiEnv* env, jclass klass, jobject* classloader_ptr) {
//   return (*env)->GetClassLoader(env, klass, classloader_ptr);
// }
// static jvmtiError GetObjectHashCode(jvmtiEnv* env, jobject object, jint* hash_code_ptr) {
//   return (*env)->GetObjectHashCode(env, object, hash_code_ptr);
// }
// static jvmtiError GetObjectMonitorUsage(jvmtiEnv* env, jobject object, jvmtiMonitorUsage* info_ptr) {
//   return (*env)->GetObjectMonitorUsage(env, object, info_ptr);
// }
// static jvmtiError GetFieldName(jvmtiEnv* env, jclass klass, jfieldID field, char** name_ptr, char** signature_ptr, char** generic_ptr) {
//   return (*env)->GetFieldName(env, klass, field, name_ptr, signature_ptr, generic_ptr);
// }
// static jvmtiError GetFieldDeclaringClass(jvmtiEnv* env, jclass klass, jfieldID field, jclass* declaring_class_ptr) {
//   return (*env)->GetFieldDeclaringClass(env, klass, field, declaring_class_ptr);
// }
// static jvmtiError GetFieldModifiers(jvmtiEnv* env, jclass klass, jfieldID field, jint* modifiers_ptr) {
//   return (*env)->GetFieldModifiers(env, klass, field, modifiers_ptr);
// }
// static jvmtiError IsFieldSynthetic(jvmtiEnv* env, jclass klass, jfieldID field, jboolean* is_synthetic_ptr) {
//   return (*env)->IsFieldSynthetic(env, klass, field, is_synthetic_ptr);
// }
// static jvmtiError GetMethodName(jvmtiEnv* env, jmethodID method, char** name_ptr, char** signature_ptr, char** generic_ptr) {
//   return (*env)->GetMethodName(env, method, name_ptr, signature_ptr, generic_ptr);
// }
// static jvmtiError GetMethodDeclaringClass(jvmtiEnv* env, jmethodID method, jclass* declaring_class_ptr) {
//   return (*env)->GetMethodDeclaringClass(env, method, declaring_class_ptr);
// }
// static jvmtiError GetMethodModifiers(jvmtiEnv* env, jmethodID method, jint* modifiers_ptr) {
//   return (*env)->GetMethodModifiers(env, method, modifiers_ptr);
// }
// static jvmtiError GetMaxLocals(jvmtiEnv* env, jmethodID method, jint* max_ptr) {
//   return (*env)->GetMaxLocals(env, method, max_ptr);
// }
// static jvmtiError GetArgumentsSize(jvmtiEnv* env, jmethodID method, jint* size_ptr) {
//   return (*env)->GetArgumentsSize(env, method, size_ptr);
// }
// static jvmtiError GetLineNumberTable(jvmtiEnv* env, jmethodID method, jint* entry_count_ptr, jvmtiLineNumberEntry** table_ptr) {
//   return (*env)->GetLineNumberTable(env, method, entry_count_ptr, table_ptr);
// }
// static jvmtiError GetMethodLocation(jvmtiEnv* env, jmethodID method, jlocation* start_location_ptr, jlocation* end_location_ptr) {
//   return (*env)->GetMethodLocation(env, method, start_location_ptr, end_location_ptr);
// }
// static jvmtiError GetLocalVariableTable(jvmtiEnv* env, jmethodID method, jint* entry_count_ptr, jvmtiLocalVariableEntry** table_ptr) {
//   return (*env)->GetLocalVariableTable(env, method, entry_count_ptr, table_ptr);
// }
// static jvmtiError SetNativeMethodPrefix(jvmtiEnv* env, const char* prefix) {
//   return (*env)->SetNativeMethodPrefix(env, prefix);
// }
// static jvmtiError SetNativeMethodPrefixes(jvmtiEnv* env, jint prefix_count, char** prefixes) {
//   return (*env)->SetNativeMethodPrefixes(env, prefix_count, prefixes);
// }
// static jvmtiError GetBytecodes(jvmtiEnv* env, jmethodID method, jint* bytecode_count_ptr, unsigned char** bytecodes_ptr) {
//   return (*env)->GetBytecodes(env, method, bytecode_count_ptr, bytecodes_ptr);
// }
// static jvmtiError IsMethodNative(jvmtiEnv* env, jmethodID method, jboolean* is_native_ptr) {
//   return (*env)->IsMethodNative(env, method, is_native_ptr);
// }
// static jvmtiError IsMethodSynthetic(jvmtiEnv* env, jmethodID method, jboolean* is_synthetic_ptr) {
//   return (*env)->IsMethodSynthetic(env, method, is_synthetic_ptr);
// }
// static jvmtiError GetLoadedClasses(jvmtiEnv* env, jint* class_count_ptr, jclass** classes_ptr) {
//   return (*env)->GetLoadedClasses(env, class_count_ptr, classes_ptr);
// }
// static jvmtiError GetClassLoaderClasses(jvmtiEnv* env, jobject initiating_loader, jint* class_count_ptr, jclass** classes_ptr) {
//   return (*env)->GetClassLoaderClasses(env, initiating_loader, class_count_ptr, classes_ptr);
// }
// static jvmtiError PopFrame(jvmtiEnv* env, jthread thread) {
//   return (*env)->PopFrame(env, thread);
// }
// static jvmtiError ForceEarlyReturnObject(jvmtiEnv* env, jthread thread, jobject value) {
//   return (*env)->ForceEarlyReturnObject(env, thread, value);
// }
// static jvmtiError ForceEarlyReturnInt(jvmtiEnv* env, jthread thread, jint value) {
//   return (*env)->ForceEarlyReturnInt(env, thread, value);
// }
// static jvmtiError ForceEarlyReturnLong(jvmtiEnv* env, jthread thread, jlong value) {
//   return (*env)->ForceEarlyReturnLong(env, thread, value);
// }
// static jvmtiError ForceEarlyReturnFloat(jvmtiEnv* env, jthread thread, jfloat value) {
//   return (*env)->ForceEarlyReturnFloat(env, thread, value);
// }
// static jvmtiError ForceEarlyReturnDouble(jvmtiEnv* env, jthread thread, jdouble value) {
//   return (*env)->ForceEarlyReturnDouble(env, thread, value);
// }
// static jvmtiError ForceEarlyReturnVoid(jvmtiEnv* env, jthread thread) {
//   return (*env)->ForceEarlyReturnVoid(env, thread);
// }
// static jvmtiError RedefineClasses(jvmtiEnv* env, jint class_count, const jvmtiClassDefinition* class_definitions) {
//   return (*env)->RedefineClasses(env, class_count, class_definitions);
// }
// static jvmtiError GetVersionNumber(jvmtiEnv* env, jint* version_ptr) {
//   return (*env)->GetVersionNumber(env, version_ptr);
// }
// static jvmtiError GetCapabilities(jvmtiEnv* env, jvmtiCapabilities* capabilities_ptr) {
//   return (*env)->GetCapabilities(env, capabilities_ptr);
// }
// static jvmtiError GetSourceDebugExtension(jvmtiEnv* env, jclass klass, char** source_debug_extension_ptr) {
//   return (*env)->GetSourceDebugExtension(env, klass, source_debug_extension_ptr);
// }
// static jvmtiError IsMethodObsolete(jvmtiEnv* env, jmethodID method, jboolean* is_obsolete_ptr) {
//   return (*env)->IsMethodObsolete(env, method, is_obsolete_ptr);
// }
// static jvmtiError SuspendThreadList(jvmtiEnv* env, jint request_count, const jthread* request_list, jvmtiError* results) {
//   return (*env)->SuspendThreadList(env, request_count, request_list, results);
// }
// static jvmtiError ResumeThreadList(jvmtiEnv* env, jint request_count, const jthread* request_list, jvmtiError* results) {
//   return (*env)->ResumeThreadList(env, request_count, request_list, results);
// }
// static jvmtiError AddModuleReads(jvmtiEnv* env, jobject module, jobject to_module) {
//   return (*env)->AddModuleReads(env, module, to_module);
// }
// static jvmtiError AddModuleExports(jvmtiEnv* env, jobject module, const char* pkg_name, jobject to_module) {
//   return (*env)->AddModuleExports(env, module, pkg_name, to_module);
// }
// static jvmtiError AddModuleOpens(jvmtiEnv* env, jobject module, const char* pkg_name, jobject to_module) {
//   return (*env)->AddModuleOpens(env, module, pkg_name, to_module);
// }
// static jvmtiError AddModuleUses(jvmtiEnv* env, jobject module, jclass service) {
//   return (*env)->AddModuleUses(env, module, service);
// }
// static jvmtiError AddModuleProvides(jvmtiEnv* env, jobject module, jclass service, jclass impl_class) {
//   return (*env)->AddModuleProvides(env, module, service, impl_class);
// }
// static jvmtiError IsModifiableModule(jvmtiEnv* env, jobject module, jboolean* is_modifiable_module_ptr) {
//   return (*env)->IsModifiableModule(env, module, is_modifiable_module_ptr);
// }
// static jvmtiError GetAllStackTraces(jvmtiEnv* env, jint max_frame_count, jvmtiStackInfo** stack_info_ptr, jint* thread_count_ptr) {
//   return (*env)->GetAllStackTraces(env, max_frame_count, stack_info_ptr, thread_count_ptr);
// }
// static jvmtiError GetThreadListStackTraces(jvmtiEnv* env, jint thread_count, const jthread* thread_list, jint max_frame_count, jvmtiStackInfo** stack_info_ptr) {
//   return (*env)->GetThreadListStackTraces(env, thread_count, thread_list, max_frame_count, stack_info_ptr);
// }
// static jvmtiError GetThreadLocalStorage(jvmtiEnv* env, jthread thread, void** data_ptr) {
//   return (*env)->GetThreadLocalStorage(env, thread, data_ptr);
// }
// static jvmtiError SetThreadLocalStorage(jvmtiEnv* env, jthread thread, const void* data) {
//   return (*env)->SetThreadLocalStorage(env, thread, data);
// }
// static jvmtiError GetStackTrace(jvmtiEnv* env, jthread thread, jint start_depth, jint max_frame_count, jvmtiFrameInfo* frame_buffer, jint* count_ptr) {
//   return (*env)->GetStackTrace(env, thread, start_depth, max_frame_count, frame_buffer, count_ptr);
// }
// static jvmtiError GetTag(jvmtiEnv* env, jobject object, jlong* tag_ptr) {
//   return (*env)->GetTag(env, object, tag_ptr);
// }
// static jvmtiError SetTag(jvmtiEnv* env, jobject object, jlong tag) {
//   return (*env)->SetTag(env, object, tag);
// }
// static jvmtiError IterateOverObjectsReachableFromObject(jvmtiEnv* env, jobject object, jvmtiObjectReferenceCallback object_reference_callback, const void* user_data) {
//   return (*env)->IterateOverObjectsReachableFromObject(env, object, object_reference_callback, user_data);
// }
// static jvmtiError IterateOverReachableObjects(jvmtiEnv* env, jvmtiHeapRootCallback heap_root_callback, jvmtiStackReferenceCallback stack_ref_callback, jvmtiObjectReferenceCallback object_ref_callback, const void* user_data) {
//   return (*env)->IterateOverReachableObjects(env, heap_root_callback, stack_ref_callback, object_ref_callback, user_data);
// }
// static jvmtiError IterateOverHeap(jvmtiEnv* env, jvmtiHeapObjectFilter object_filter, jvmtiHeapObjectCallback heap_object_callback, const void* user_data) {
//   return (*env)->IterateOverHeap(env, object_filter, heap_object_callback, user_data);
// }
// static jvmtiError IterateOverInstancesOfClass(jvmtiEnv* env, jclass klass, jvmtiHeapObjectFilter object_filter, jvmtiHeapObjectCallback heap_object_callback, const void* user_data) {
//   return (*env)->IterateOverInstancesOfClass(env, klass, object_filter, heap_object_callback, user_data);
// }
// static jvmtiError GetObjectsWithTags(jvmtiEnv* env, jint tag_count, const jlong* tags, jint* count_ptr, jobject** object_result_ptr, jlong** tag_result_ptr) {
//   return (*env)->GetObjectsWithTags(env, tag_count, tags, count_ptr, object_result_ptr, tag_result_ptr);
// }
// static jvmtiError FollowReferences(jvmtiEnv* env, jint heap_filter, jclass klass, jobject initial_object, const jvmtiHeapCallbacks* callbacks, const void* user_data) {
//   return (*env)->FollowReferences(env, heap_filter, klass, initial_object, callbacks, user_data);
// }
// static jvmtiError IterateThroughHeap(jvmtiEnv* env, jint heap_filter, jclass klass, const jvmtiHeapCallbacks* callbacks, const void* user_data) {
//   return (*env)->IterateThroughHeap(env, heap_filter, klass, callbacks, user_data);
// }
// static jvmtiError SetJNIFunctionTable(jvmtiEnv* env, const jniNativeInterface* function_table) {
//   return (*env)->SetJNIFunctionTable(env, function_table);
// }
// static jvmtiError GetJNIFunctionTable(jvmtiEnv* env, jniNativeInterface** function_table) {
//   return (*env)->GetJNIFunctionTable(env, function_table);
// }
// static jvmtiError SetEventCallbacks(jvmtiEnv* env, const jvmtiEventCallbacks* callbacks, jint size_of_callbacks) {
//   return (*env)->SetEventCallbacks(env, callbacks, size_of_callbacks);
// }
// static jvmtiError GenerateEvents(jvmtiEnv* env, jvmtiEvent event_type) {
//   return (*env)->GenerateEvents(env, event_type);
// }
// static jvmtiError GetExtensionFunctions(jvmtiEnv* env, jint* extension_count_ptr, jvmtiExtensionFunctionInfo** extensions) {
//   return (*env)->GetExtensionFunctions(env, extension_count_ptr, extensions);
// }
// static jvmtiError GetExtensionEvents(jvmtiEnv* env, jint* extension_count_ptr, jvmtiExtensionEventInfo** extensions) {
//   return (*env)->GetExtensionEvents(env, extension_count_ptr, extensions);
// }
// static jvmtiError SetExtensionEventCallback(jvmtiEnv* env, jint extension_event_index, jvmtiExtensionEvent callback) {
//   return (*env)->SetExtensionEventCallback(env, extension_event_index, callback);
// }
// static jvmtiError GetErrorName(jvmtiEnv* env, jvmtiError error, char** name_ptr) {
//   return (*env)->GetErrorName(env, error, name_ptr);
// }
// static jvmtiError GetJLocationFormat(jvmtiEnv* env, jvmtiJlocationFormat* format_ptr) {
//   return (*env)->GetJLocationFormat(env, format_ptr);
// }
// static jvmtiError GetSystemProperties(jvmtiEnv* env, jint* count_ptr, char*** property_ptr) {
//   return (*env)->GetSystemProperties(env, count_ptr, property_ptr);
// }
// static jvmtiError GetSystemProperty(jvmtiEnv* env, const char* property, char** value_ptr) {
//   return (*env)->GetSystemProperty(env, property, value_ptr);
// }
// static jvmtiError SetSystemProperty(jvmtiEnv* env, const char* property, const char* value_ptr) {
//   return (*env)->SetSystemProperty(env, property, value_ptr);
// }
// static jvmtiError GetPhase(jvmtiEnv* env, jvmtiPhase* phase_ptr) {
//   return (*env)->GetPhase(env, phase_ptr);
// }
// static jvmtiError GetCurrentThreadCpuTimerInfo(jvmtiEnv* env, jvmtiTimerInfo* info_ptr) {
//   return (*env)->GetCurrentThreadCpuTimerInfo(env, info_ptr);
// }
// static jvmtiError GetCurrentThreadCpuTime(jvmtiEnv* env, jlong* nanos_ptr) {
//   return (*env)->GetCurrentThreadCpuTime(env, nanos_ptr);
// }
// static jvmtiError GetThreadCpuTimerInfo(jvmtiEnv* env, jvmtiTimerInfo* info_ptr) {
//   return (*env)->GetThreadCpuTimerInfo(env, info_ptr);
// }
// static jvmtiError GetThreadCpuTime(jvmtiEnv* env, jthread thread, jlong* nanos_ptr) {
//   return (*env)->GetThreadCpuTime(env, thread, nanos_ptr);
// }
// static jvmtiError GetTimerInfo(jvmtiEnv* env, jvmtiTimerInfo* info_ptr) {
//   return (*env)->GetTimerInfo(env, info_ptr);
// }
// static jvmtiError GetTime(jvmtiEnv* env, jlong* nanos_ptr) {
//   return (*env)->GetTime(env, nanos_ptr);
// }
// static jvmtiError GetPotentialCapabilities(jvmtiEnv* env, jvmtiCapabilities* capabilities_ptr) {
//   return (*env)->GetPotentialCapabilities(env, capabilities_ptr);
// }
// static jvmtiError AddCapabilities(jvmtiEnv* env, const jvmtiCapabilities* capabilities_ptr) {
//   return (*env)->AddCapabilities(env, capabilities_ptr);
// }
// static jvmtiError RelinquishCapabilities(jvmtiEnv* env, const jvmtiCapabilities* capabilities_ptr) {
//   return (*env)->RelinquishCapabilities(env, capabilities_ptr);
// }
// static jvmtiError GetAvailableProcessors(jvmtiEnv* env, jint* processor_count_ptr) {
//   return (*env)->GetAvailableProcessors(env, processor_count_ptr);
// }
// static jvmtiError GetClassVersionNumbers(jvmtiEnv* env, jclass klass, jint* minor_version_ptr, jint* major_version_ptr) {
//   return (*env)->GetClassVersionNumbers(env, klass, minor_version_ptr, major_version_ptr);
// }
// static jvmtiError GetConstantPool(jvmtiEnv* env, jclass klass, jint* constant_pool_count_ptr, jint* constant_pool_byte_count_ptr, unsigned char** constant_pool_bytes_ptr) {
//   return (*env)->GetConstantPool(env, klass, constant_pool_count_ptr, constant_pool_byte_count_ptr, constant_pool_bytes_ptr);
// }
// static jvmtiError GetEnvironmentLocalStorage(jvmtiEnv* env, void** data_ptr) {
//   return (*env)->GetEnvironmentLocalStorage(env, data_ptr);
// }
// static jvmtiError SetEnvironmentLocalStorage(jvmtiEnv* env, const void* data) {
//   return (*env)->SetEnvironmentLocalStorage(env, data);
// }
// static jvmtiError AddToBootstrapClassLoaderSearch(jvmtiEnv* env, const char* segment) {
//   return (*env)->AddToBootstrapClassLoaderSearch(env, segment);
// }
// static jvmtiError SetVerboseFlag(jvmtiEnv* env, jvmtiVerboseFlag flag, jboolean value) {
//   return (*env)->SetVerboseFlag(env, flag, value);
// }
// static jvmtiError AddToSystemClassLoaderSearch(jvmtiEnv* env, const char* segment) {
//   return (*env)->AddToSystemClassLoaderSearch(env, segment);
// }
// static jvmtiError RetransformClasses(jvmtiEnv* env, jint class_count, const jclass* classes) {
//   return (*env)->RetransformClasses(env, class_count, classes);
// }
// static jvmtiError GetOwnedMonitorStackDepthInfo(jvmtiEnv* env, jthread thread, jint* monitor_info_count_ptr, jvmtiMonitorStackDepthInfo** monitor_info_ptr) {
//   return (*env)->GetOwnedMonitorStackDepthInfo(env, thread, monitor_info_count_ptr, monitor_info_ptr);
// }
// static jvmtiError GetObjectSize(jvmtiEnv* env, jobject object, jlong* size_ptr) {
//   return (*env)->GetObjectSize(env, object, size_ptr);
// }
// static jvmtiError GetLocalInstance(jvmtiEnv* env, jthread thread, jint depth, jobject* value_ptr) {
//   return (*env)->GetLocalInstance(env, thread, depth, value_ptr);
// }
// static jvmtiError SetHeapSamplingInterval(jvmtiEnv* env, jint sampling_interval) {
//   return (*env)->SetHeapSamplingInterval(env, sampling_interval);
// }
import "C"

import "unsafe"

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
    JVMTI_CLASS_STATUS_VERIFIED = 1
    JVMTI_CLASS_STATUS_PREPARED = 2
    JVMTI_CLASS_STATUS_INITIALIZED = 4
    JVMTI_CLASS_STATUS_ERROR = 8
    JVMTI_CLASS_STATUS_ARRAY = 16
    JVMTI_CLASS_STATUS_PRIMITIVE = 32
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
    JVMTI_RESOURCE_EXHAUSTED_OOM_ERROR = 0x0001
    JVMTI_RESOURCE_EXHAUSTED_JAVA_HEAP = 0x0002
    JVMTI_RESOURCE_EXHAUSTED_THREADS = 0x0004
)



// JvmtiEnv corresponds to jvmtiEnv*
type JvmtiEnv uintptr

func (jvmti JvmtiEnv) raw() *C.jvmtiEnv {
	return (*C.jvmtiEnv)(unsafe.Pointer(jvmti))
}

func (jvmti JvmtiEnv) asPointer() unsafe.Pointer {
	return unsafe.Pointer(jvmti)
}


func (jvmti JvmtiEnv) SetEventNotificationMode(mode C.jvmtiEventMode, event_type C.jvmtiEvent, event_thread C.jthread, args []interface{}) C.jvmtiError {
  return C.SetEventNotificationMode(jvmti.raw(), mode, event_type, event_thread )
}

func (jvmti JvmtiEnv) GetAllModules(module_count_ptr *C.jint, modules_ptr **C.jobject) C.jvmtiError {
  return C.GetAllModules(jvmti.raw(), module_count_ptr, modules_ptr)
}

func (jvmti JvmtiEnv) GetAllThreads(threads_count_ptr *C.jint, threads_ptr **C.jthread) C.jvmtiError {
  return C.GetAllThreads(jvmti.raw(), threads_count_ptr, threads_ptr)
}

func (jvmti JvmtiEnv) SuspendThread(thread C.jthread) C.jvmtiError {
  return C.SuspendThread(jvmti.raw(), thread)
}

func (jvmti JvmtiEnv) ResumeThread(thread C.jthread) C.jvmtiError {
  return C.ResumeThread(jvmti.raw(), thread)
}

func (jvmti JvmtiEnv) StopThread(thread C.jthread, exception C.jobject) C.jvmtiError {
  return C.StopThread(jvmti.raw(), thread, exception)
}

func (jvmti JvmtiEnv) InterruptThread(thread C.jthread) C.jvmtiError {
  return C.InterruptThread(jvmti.raw(), thread)
}

func (jvmti JvmtiEnv) GetThreadInfo(thread C.jthread, info_ptr *C.jvmtiThreadInfo) C.jvmtiError {
  return C.GetThreadInfo(jvmti.raw(), thread, info_ptr)
}

func (jvmti JvmtiEnv) GetOwnedMonitorInfo(thread C.jthread, owned_monitor_count_ptr *C.jint, owned_monitors_ptr **C.jobject) C.jvmtiError {
  return C.GetOwnedMonitorInfo(jvmti.raw(), thread, owned_monitor_count_ptr, owned_monitors_ptr)
}

func (jvmti JvmtiEnv) GetCurrentContendedMonitor(thread C.jthread, monitor_ptr *C.jobject) C.jvmtiError {
  return C.GetCurrentContendedMonitor(jvmti.raw(), thread, monitor_ptr)
}

func (jvmti JvmtiEnv) RunAgentThread(thread C.jthread, proc C.jvmtiStartFunction, arg unsafe.Pointer, priority C.jint) C.jvmtiError {
  return C.RunAgentThread(jvmti.raw(), thread, proc, arg, priority)
}

func (jvmti JvmtiEnv) GetTopThreadGroups(group_count_ptr *C.jint, groups_ptr **C.jthreadGroup) C.jvmtiError {
  return C.GetTopThreadGroups(jvmti.raw(), group_count_ptr, groups_ptr)
}

func (jvmti JvmtiEnv) GetThreadGroupInfo(group C.jthreadGroup, info_ptr *C.jvmtiThreadGroupInfo) C.jvmtiError {
  return C.GetThreadGroupInfo(jvmti.raw(), group, info_ptr)
}

func (jvmti JvmtiEnv) GetThreadGroupChildren(group C.jthreadGroup, thread_count_ptr *C.jint, threads_ptr **C.jthread, group_count_ptr *C.jint, groups_ptr **C.jthreadGroup) C.jvmtiError {
  return C.GetThreadGroupChildren(jvmti.raw(), group, thread_count_ptr, threads_ptr, group_count_ptr, groups_ptr)
}

func (jvmti JvmtiEnv) GetFrameCount(thread C.jthread, count_ptr *C.jint) C.jvmtiError {
  return C.GetFrameCount(jvmti.raw(), thread, count_ptr)
}

func (jvmti JvmtiEnv) GetThreadState(thread C.jthread, thread_state_ptr *C.jint) C.jvmtiError {
  return C.GetThreadState(jvmti.raw(), thread, thread_state_ptr)
}

func (jvmti JvmtiEnv) GetCurrentThread(thread_ptr *C.jthread) C.jvmtiError {
  return C.GetCurrentThread(jvmti.raw(), thread_ptr)
}

func (jvmti JvmtiEnv) GetFrameLocation(thread C.jthread, depth C.jint, method_ptr *C.jmethodID, location_ptr *C.jlocation) C.jvmtiError {
  return C.GetFrameLocation(jvmti.raw(), thread, depth, method_ptr, location_ptr)
}

func (jvmti JvmtiEnv) NotifyFramePop(thread C.jthread, depth C.jint) C.jvmtiError {
  return C.NotifyFramePop(jvmti.raw(), thread, depth)
}

func (jvmti JvmtiEnv) GetLocalObject(thread C.jthread, depth C.jint, slot C.jint, value_ptr *C.jobject) C.jvmtiError {
  return C.GetLocalObject(jvmti.raw(), thread, depth, slot, value_ptr)
}

func (jvmti JvmtiEnv) GetLocalInt(thread C.jthread, depth C.jint, slot C.jint, value_ptr *C.jint) C.jvmtiError {
  return C.GetLocalInt(jvmti.raw(), thread, depth, slot, value_ptr)
}

func (jvmti JvmtiEnv) GetLocalLong(thread C.jthread, depth C.jint, slot C.jint, value_ptr *C.jlong) C.jvmtiError {
  return C.GetLocalLong(jvmti.raw(), thread, depth, slot, value_ptr)
}

func (jvmti JvmtiEnv) GetLocalFloat(thread C.jthread, depth C.jint, slot C.jint, value_ptr *C.jfloat) C.jvmtiError {
  return C.GetLocalFloat(jvmti.raw(), thread, depth, slot, value_ptr)
}

func (jvmti JvmtiEnv) GetLocalDouble(thread C.jthread, depth C.jint, slot C.jint, value_ptr *C.jdouble) C.jvmtiError {
  return C.GetLocalDouble(jvmti.raw(), thread, depth, slot, value_ptr)
}

func (jvmti JvmtiEnv) SetLocalObject(thread C.jthread, depth C.jint, slot C.jint, value C.jobject) C.jvmtiError {
  return C.SetLocalObject(jvmti.raw(), thread, depth, slot, value)
}

func (jvmti JvmtiEnv) SetLocalInt(thread C.jthread, depth C.jint, slot C.jint, value C.jint) C.jvmtiError {
  return C.SetLocalInt(jvmti.raw(), thread, depth, slot, value)
}

func (jvmti JvmtiEnv) SetLocalLong(thread C.jthread, depth C.jint, slot C.jint, value C.jlong) C.jvmtiError {
  return C.SetLocalLong(jvmti.raw(), thread, depth, slot, value)
}

func (jvmti JvmtiEnv) SetLocalFloat(thread C.jthread, depth C.jint, slot C.jint, value C.jfloat) C.jvmtiError {
  return C.SetLocalFloat(jvmti.raw(), thread, depth, slot, value)
}

func (jvmti JvmtiEnv) SetLocalDouble(thread C.jthread, depth C.jint, slot C.jint, value C.jdouble) C.jvmtiError {
  return C.SetLocalDouble(jvmti.raw(), thread, depth, slot, value)
}

func (jvmti JvmtiEnv) CreateRawMonitor(name *C.char, monitor_ptr *C.jrawMonitorID) C.jvmtiError {
  return C.CreateRawMonitor(jvmti.raw(), name, monitor_ptr)
}

func (jvmti JvmtiEnv) DestroyRawMonitor(monitor C.jrawMonitorID) C.jvmtiError {
  return C.DestroyRawMonitor(jvmti.raw(), monitor)
}

func (jvmti JvmtiEnv) RawMonitorEnter(monitor C.jrawMonitorID) C.jvmtiError {
  return C.RawMonitorEnter(jvmti.raw(), monitor)
}

func (jvmti JvmtiEnv) RawMonitorExit(monitor C.jrawMonitorID) C.jvmtiError {
  return C.RawMonitorExit(jvmti.raw(), monitor)
}

func (jvmti JvmtiEnv) RawMonitorWait(monitor C.jrawMonitorID, millis C.jlong) C.jvmtiError {
  return C.RawMonitorWait(jvmti.raw(), monitor, millis)
}

func (jvmti JvmtiEnv) RawMonitorNotify(monitor C.jrawMonitorID) C.jvmtiError {
  return C.RawMonitorNotify(jvmti.raw(), monitor)
}

func (jvmti JvmtiEnv) RawMonitorNotifyAll(monitor C.jrawMonitorID) C.jvmtiError {
  return C.RawMonitorNotifyAll(jvmti.raw(), monitor)
}

func (jvmti JvmtiEnv) SetBreakpoint(method C.jmethodID, location C.jlocation) C.jvmtiError {
  return C.SetBreakpoint(jvmti.raw(), method, location)
}

func (jvmti JvmtiEnv) ClearBreakpoint(method C.jmethodID, location C.jlocation) C.jvmtiError {
  return C.ClearBreakpoint(jvmti.raw(), method, location)
}

func (jvmti JvmtiEnv) GetNamedModule(class_loader C.jobject, package_name *C.char, module_ptr *C.jobject) C.jvmtiError {
  return C.GetNamedModule(jvmti.raw(), class_loader, package_name, module_ptr)
}

func (jvmti JvmtiEnv) SetFieldAccessWatch(klass C.jclass, field C.jfieldID) C.jvmtiError {
  return C.SetFieldAccessWatch(jvmti.raw(), klass, field)
}

func (jvmti JvmtiEnv) ClearFieldAccessWatch(klass C.jclass, field C.jfieldID) C.jvmtiError {
  return C.ClearFieldAccessWatch(jvmti.raw(), klass, field)
}

func (jvmti JvmtiEnv) SetFieldModificationWatch(klass C.jclass, field C.jfieldID) C.jvmtiError {
  return C.SetFieldModificationWatch(jvmti.raw(), klass, field)
}

func (jvmti JvmtiEnv) ClearFieldModificationWatch(klass C.jclass, field C.jfieldID) C.jvmtiError {
  return C.ClearFieldModificationWatch(jvmti.raw(), klass, field)
}

func (jvmti JvmtiEnv) IsModifiableClass(klass C.jclass, is_modifiable_class_ptr *C.jboolean) C.jvmtiError {
  return C.IsModifiableClass(jvmti.raw(), klass, is_modifiable_class_ptr)
}

func (jvmti JvmtiEnv) Allocate(size C.jlong, mem_ptr **C.uchar) C.jvmtiError {
  return C.Allocate(jvmti.raw(), size, mem_ptr)
}

func (jvmti JvmtiEnv) Deallocate(mem *C.uchar) C.jvmtiError {
  return C.Deallocate(jvmti.raw(), mem)
}

func (jvmti JvmtiEnv) GetClassSignature(klass C.jclass, signature_ptr **C.char, generic_ptr **C.char) C.jvmtiError {
  return C.GetClassSignature(jvmti.raw(), klass, signature_ptr, generic_ptr)
}

func (jvmti JvmtiEnv) GetClassStatus(klass C.jclass, status_ptr *C.jint) C.jvmtiError {
  return C.GetClassStatus(jvmti.raw(), klass, status_ptr)
}

func (jvmti JvmtiEnv) GetSourceFileName(klass C.jclass, source_name_ptr **C.char) C.jvmtiError {
  return C.GetSourceFileName(jvmti.raw(), klass, source_name_ptr)
}

func (jvmti JvmtiEnv) GetClassModifiers(klass C.jclass, modifiers_ptr *C.jint) C.jvmtiError {
  return C.GetClassModifiers(jvmti.raw(), klass, modifiers_ptr)
}

func (jvmti JvmtiEnv) GetClassMethods(klass C.jclass, method_count_ptr *C.jint, methods_ptr **C.jmethodID) C.jvmtiError {
  return C.GetClassMethods(jvmti.raw(), klass, method_count_ptr, methods_ptr)
}

func (jvmti JvmtiEnv) GetClassFields(klass C.jclass, field_count_ptr *C.jint, fields_ptr **C.jfieldID) C.jvmtiError {
  return C.GetClassFields(jvmti.raw(), klass, field_count_ptr, fields_ptr)
}

func (jvmti JvmtiEnv) GetImplementedInterfaces(klass C.jclass, interface_count_ptr *C.jint, interfaces_ptr **C.jclass) C.jvmtiError {
  return C.GetImplementedInterfaces(jvmti.raw(), klass, interface_count_ptr, interfaces_ptr)
}

func (jvmti JvmtiEnv) IsInterface(klass C.jclass, is_interface_ptr *C.jboolean) C.jvmtiError {
  return C.IsInterface(jvmti.raw(), klass, is_interface_ptr)
}

func (jvmti JvmtiEnv) IsArrayClass(klass C.jclass, is_array_class_ptr *C.jboolean) C.jvmtiError {
  return C.IsArrayClass(jvmti.raw(), klass, is_array_class_ptr)
}

func (jvmti JvmtiEnv) GetClassLoader(klass C.jclass, classloader_ptr *C.jobject) C.jvmtiError {
  return C.GetClassLoader(jvmti.raw(), klass, classloader_ptr)
}

func (jvmti JvmtiEnv) GetObjectHashCode(object C.jobject, hash_code_ptr *C.jint) C.jvmtiError {
  return C.GetObjectHashCode(jvmti.raw(), object, hash_code_ptr)
}

func (jvmti JvmtiEnv) GetObjectMonitorUsage(object C.jobject, info_ptr *C.jvmtiMonitorUsage) C.jvmtiError {
  return C.GetObjectMonitorUsage(jvmti.raw(), object, info_ptr)
}

func (jvmti JvmtiEnv) GetFieldName(klass C.jclass, field C.jfieldID, name_ptr **C.char, signature_ptr **C.char, generic_ptr **C.char) C.jvmtiError {
  return C.GetFieldName(jvmti.raw(), klass, field, name_ptr, signature_ptr, generic_ptr)
}

func (jvmti JvmtiEnv) GetFieldDeclaringClass(klass C.jclass, field C.jfieldID, declaring_class_ptr *C.jclass) C.jvmtiError {
  return C.GetFieldDeclaringClass(jvmti.raw(), klass, field, declaring_class_ptr)
}

func (jvmti JvmtiEnv) GetFieldModifiers(klass C.jclass, field C.jfieldID, modifiers_ptr *C.jint) C.jvmtiError {
  return C.GetFieldModifiers(jvmti.raw(), klass, field, modifiers_ptr)
}

func (jvmti JvmtiEnv) IsFieldSynthetic(klass C.jclass, field C.jfieldID, is_synthetic_ptr *C.jboolean) C.jvmtiError {
  return C.IsFieldSynthetic(jvmti.raw(), klass, field, is_synthetic_ptr)
}

func (jvmti JvmtiEnv) GetMethodName(method C.jmethodID, name_ptr **C.char, signature_ptr **C.char, generic_ptr **C.char) C.jvmtiError {
  return C.GetMethodName(jvmti.raw(), method, name_ptr, signature_ptr, generic_ptr)
}

func (jvmti JvmtiEnv) GetMethodDeclaringClass(method C.jmethodID, declaring_class_ptr *C.jclass) C.jvmtiError {
  return C.GetMethodDeclaringClass(jvmti.raw(), method, declaring_class_ptr)
}

func (jvmti JvmtiEnv) GetMethodModifiers(method C.jmethodID, modifiers_ptr *C.jint) C.jvmtiError {
  return C.GetMethodModifiers(jvmti.raw(), method, modifiers_ptr)
}

func (jvmti JvmtiEnv) GetMaxLocals(method C.jmethodID, max_ptr *C.jint) C.jvmtiError {
  return C.GetMaxLocals(jvmti.raw(), method, max_ptr)
}

func (jvmti JvmtiEnv) GetArgumentsSize(method C.jmethodID, size_ptr *C.jint) C.jvmtiError {
  return C.GetArgumentsSize(jvmti.raw(), method, size_ptr)
}

func (jvmti JvmtiEnv) GetLineNumberTable(method C.jmethodID, entry_count_ptr *C.jint, table_ptr **C.jvmtiLineNumberEntry) C.jvmtiError {
  return C.GetLineNumberTable(jvmti.raw(), method, entry_count_ptr, table_ptr)
}

func (jvmti JvmtiEnv) GetMethodLocation(method C.jmethodID, start_location_ptr *C.jlocation, end_location_ptr *C.jlocation) C.jvmtiError {
  return C.GetMethodLocation(jvmti.raw(), method, start_location_ptr, end_location_ptr)
}

func (jvmti JvmtiEnv) GetLocalVariableTable(method C.jmethodID, entry_count_ptr *C.jint, table_ptr **C.jvmtiLocalVariableEntry) C.jvmtiError {
  return C.GetLocalVariableTable(jvmti.raw(), method, entry_count_ptr, table_ptr)
}

func (jvmti JvmtiEnv) SetNativeMethodPrefix(prefix *C.char) C.jvmtiError {
  return C.SetNativeMethodPrefix(jvmti.raw(), prefix)
}

func (jvmti JvmtiEnv) SetNativeMethodPrefixes(prefix_count C.jint, prefixes **C.char) C.jvmtiError {
  return C.SetNativeMethodPrefixes(jvmti.raw(), prefix_count, prefixes)
}

func (jvmti JvmtiEnv) GetBytecodes(method C.jmethodID, bytecode_count_ptr *C.jint, bytecodes_ptr **C.uchar) C.jvmtiError {
  return C.GetBytecodes(jvmti.raw(), method, bytecode_count_ptr, bytecodes_ptr)
}

func (jvmti JvmtiEnv) IsMethodNative(method C.jmethodID, is_native_ptr *C.jboolean) C.jvmtiError {
  return C.IsMethodNative(jvmti.raw(), method, is_native_ptr)
}

func (jvmti JvmtiEnv) IsMethodSynthetic(method C.jmethodID, is_synthetic_ptr *C.jboolean) C.jvmtiError {
  return C.IsMethodSynthetic(jvmti.raw(), method, is_synthetic_ptr)
}

func (jvmti JvmtiEnv) GetLoadedClasses(class_count_ptr *C.jint, classes_ptr **C.jclass) C.jvmtiError {
  return C.GetLoadedClasses(jvmti.raw(), class_count_ptr, classes_ptr)
}

func (jvmti JvmtiEnv) GetClassLoaderClasses(initiating_loader C.jobject, class_count_ptr *C.jint, classes_ptr **C.jclass) C.jvmtiError {
  return C.GetClassLoaderClasses(jvmti.raw(), initiating_loader, class_count_ptr, classes_ptr)
}

func (jvmti JvmtiEnv) PopFrame(thread C.jthread) C.jvmtiError {
  return C.PopFrame(jvmti.raw(), thread)
}

func (jvmti JvmtiEnv) ForceEarlyReturnObject(thread C.jthread, value C.jobject) C.jvmtiError {
  return C.ForceEarlyReturnObject(jvmti.raw(), thread, value)
}

func (jvmti JvmtiEnv) ForceEarlyReturnInt(thread C.jthread, value C.jint) C.jvmtiError {
  return C.ForceEarlyReturnInt(jvmti.raw(), thread, value)
}

func (jvmti JvmtiEnv) ForceEarlyReturnLong(thread C.jthread, value C.jlong) C.jvmtiError {
  return C.ForceEarlyReturnLong(jvmti.raw(), thread, value)
}

func (jvmti JvmtiEnv) ForceEarlyReturnFloat(thread C.jthread, value C.jfloat) C.jvmtiError {
  return C.ForceEarlyReturnFloat(jvmti.raw(), thread, value)
}

func (jvmti JvmtiEnv) ForceEarlyReturnDouble(thread C.jthread, value C.jdouble) C.jvmtiError {
  return C.ForceEarlyReturnDouble(jvmti.raw(), thread, value)
}

func (jvmti JvmtiEnv) ForceEarlyReturnVoid(thread C.jthread) C.jvmtiError {
  return C.ForceEarlyReturnVoid(jvmti.raw(), thread)
}

func (jvmti JvmtiEnv) RedefineClasses(class_count C.jint, class_definitions *C.jvmtiClassDefinition) C.jvmtiError {
  return C.RedefineClasses(jvmti.raw(), class_count, class_definitions)
}

func (jvmti JvmtiEnv) GetVersionNumber(version_ptr *C.jint) C.jvmtiError {
  return C.GetVersionNumber(jvmti.raw(), version_ptr)
}

func (jvmti JvmtiEnv) GetCapabilities(capabilities_ptr *C.jvmtiCapabilities) C.jvmtiError {
  return C.GetCapabilities(jvmti.raw(), capabilities_ptr)
}

func (jvmti JvmtiEnv) GetSourceDebugExtension(klass C.jclass, source_debug_extension_ptr **C.char) C.jvmtiError {
  return C.GetSourceDebugExtension(jvmti.raw(), klass, source_debug_extension_ptr)
}

func (jvmti JvmtiEnv) IsMethodObsolete(method C.jmethodID, is_obsolete_ptr *C.jboolean) C.jvmtiError {
  return C.IsMethodObsolete(jvmti.raw(), method, is_obsolete_ptr)
}

func (jvmti JvmtiEnv) SuspendThreadList(request_count C.jint, request_list *C.jthread, results *C.jvmtiError) C.jvmtiError {
  return C.SuspendThreadList(jvmti.raw(), request_count, request_list, results)
}

func (jvmti JvmtiEnv) ResumeThreadList(request_count C.jint, request_list *C.jthread, results *C.jvmtiError) C.jvmtiError {
  return C.ResumeThreadList(jvmti.raw(), request_count, request_list, results)
}

func (jvmti JvmtiEnv) AddModuleReads(module C.jobject, to_module C.jobject) C.jvmtiError {
  return C.AddModuleReads(jvmti.raw(), module, to_module)
}

func (jvmti JvmtiEnv) AddModuleExports(module C.jobject, pkg_name *C.char, to_module C.jobject) C.jvmtiError {
  return C.AddModuleExports(jvmti.raw(), module, pkg_name, to_module)
}

func (jvmti JvmtiEnv) AddModuleOpens(module C.jobject, pkg_name *C.char, to_module C.jobject) C.jvmtiError {
  return C.AddModuleOpens(jvmti.raw(), module, pkg_name, to_module)
}

func (jvmti JvmtiEnv) AddModuleUses(module C.jobject, service C.jclass) C.jvmtiError {
  return C.AddModuleUses(jvmti.raw(), module, service)
}

func (jvmti JvmtiEnv) AddModuleProvides(module C.jobject, service C.jclass, impl_class C.jclass) C.jvmtiError {
  return C.AddModuleProvides(jvmti.raw(), module, service, impl_class)
}

func (jvmti JvmtiEnv) IsModifiableModule(module C.jobject, is_modifiable_module_ptr *C.jboolean) C.jvmtiError {
  return C.IsModifiableModule(jvmti.raw(), module, is_modifiable_module_ptr)
}

func (jvmti JvmtiEnv) GetAllStackTraces(max_frame_count C.jint, stack_info_ptr **C.jvmtiStackInfo, thread_count_ptr *C.jint) C.jvmtiError {
  return C.GetAllStackTraces(jvmti.raw(), max_frame_count, stack_info_ptr, thread_count_ptr)
}

func (jvmti JvmtiEnv) GetThreadListStackTraces(thread_count C.jint, thread_list *C.jthread, max_frame_count C.jint, stack_info_ptr **C.jvmtiStackInfo) C.jvmtiError {
  return C.GetThreadListStackTraces(jvmti.raw(), thread_count, thread_list, max_frame_count, stack_info_ptr)
}

func (jvmti JvmtiEnv) GetThreadLocalStorage(thread C.jthread, data_ptr *unsafe.Pointer) C.jvmtiError {
  return C.GetThreadLocalStorage(jvmti.raw(), thread, data_ptr)
}

func (jvmti JvmtiEnv) SetThreadLocalStorage(thread C.jthread, data unsafe.Pointer) C.jvmtiError {
  return C.SetThreadLocalStorage(jvmti.raw(), thread, data)
}

func (jvmti JvmtiEnv) GetStackTrace(thread C.jthread, start_depth C.jint, max_frame_count C.jint, frame_buffer *C.jvmtiFrameInfo, count_ptr *C.jint) C.jvmtiError {
  return C.GetStackTrace(jvmti.raw(), thread, start_depth, max_frame_count, frame_buffer, count_ptr)
}

func (jvmti JvmtiEnv) GetTag(object C.jobject, tag_ptr *C.jlong) C.jvmtiError {
  return C.GetTag(jvmti.raw(), object, tag_ptr)
}

func (jvmti JvmtiEnv) SetTag(object C.jobject, tag C.jlong) C.jvmtiError {
  return C.SetTag(jvmti.raw(), object, tag)
}

func (jvmti JvmtiEnv) IterateOverObjectsReachableFromObject(object C.jobject, object_reference_callback C.jvmtiObjectReferenceCallback, user_data unsafe.Pointer) C.jvmtiError {
  return C.IterateOverObjectsReachableFromObject(jvmti.raw(), object, object_reference_callback, user_data)
}

func (jvmti JvmtiEnv) IterateOverReachableObjects(heap_root_callback C.jvmtiHeapRootCallback, stack_ref_callback C.jvmtiStackReferenceCallback, object_ref_callback C.jvmtiObjectReferenceCallback, user_data unsafe.Pointer) C.jvmtiError {
  return C.IterateOverReachableObjects(jvmti.raw(), heap_root_callback, stack_ref_callback, object_ref_callback, user_data)
}

func (jvmti JvmtiEnv) IterateOverHeap(object_filter C.jvmtiHeapObjectFilter, heap_object_callback C.jvmtiHeapObjectCallback, user_data unsafe.Pointer) C.jvmtiError {
  return C.IterateOverHeap(jvmti.raw(), object_filter, heap_object_callback, user_data)
}

func (jvmti JvmtiEnv) IterateOverInstancesOfClass(klass C.jclass, object_filter C.jvmtiHeapObjectFilter, heap_object_callback C.jvmtiHeapObjectCallback, user_data unsafe.Pointer) C.jvmtiError {
  return C.IterateOverInstancesOfClass(jvmti.raw(), klass, object_filter, heap_object_callback, user_data)
}

func (jvmti JvmtiEnv) GetObjectsWithTags(tag_count C.jint, tags *C.jlong, count_ptr *C.jint, object_result_ptr **C.jobject, tag_result_ptr **C.jlong) C.jvmtiError {
  return C.GetObjectsWithTags(jvmti.raw(), tag_count, tags, count_ptr, object_result_ptr, tag_result_ptr)
}

func (jvmti JvmtiEnv) FollowReferences(heap_filter C.jint, klass C.jclass, initial_object C.jobject, callbacks *C.jvmtiHeapCallbacks, user_data unsafe.Pointer) C.jvmtiError {
  return C.FollowReferences(jvmti.raw(), heap_filter, klass, initial_object, callbacks, user_data)
}

func (jvmti JvmtiEnv) IterateThroughHeap(heap_filter C.jint, klass C.jclass, callbacks *C.jvmtiHeapCallbacks, user_data unsafe.Pointer) C.jvmtiError {
  return C.IterateThroughHeap(jvmti.raw(), heap_filter, klass, callbacks, user_data)
}

func (jvmti JvmtiEnv) SetJNIFunctionTable(function_table *C.jniNativeInterface) C.jvmtiError {
  return C.SetJNIFunctionTable(jvmti.raw(), function_table)
}

func (jvmti JvmtiEnv) GetJNIFunctionTable(function_table **C.jniNativeInterface) C.jvmtiError {
  return C.GetJNIFunctionTable(jvmti.raw(), function_table)
}

func (jvmti JvmtiEnv) SetEventCallbacks(callbacks *C.jvmtiEventCallbacks, size_of_callbacks C.jint) C.jvmtiError {
  return C.SetEventCallbacks(jvmti.raw(), callbacks, size_of_callbacks)
}

func (jvmti JvmtiEnv) GenerateEvents(event_type C.jvmtiEvent) C.jvmtiError {
  return C.GenerateEvents(jvmti.raw(), event_type)
}

func (jvmti JvmtiEnv) GetExtensionFunctions(extension_count_ptr *C.jint, extensions **C.jvmtiExtensionFunctionInfo) C.jvmtiError {
  return C.GetExtensionFunctions(jvmti.raw(), extension_count_ptr, extensions)
}

func (jvmti JvmtiEnv) GetExtensionEvents(extension_count_ptr *C.jint, extensions **C.jvmtiExtensionEventInfo) C.jvmtiError {
  return C.GetExtensionEvents(jvmti.raw(), extension_count_ptr, extensions)
}

func (jvmti JvmtiEnv) SetExtensionEventCallback(extension_event_index C.jint, callback C.jvmtiExtensionEvent) C.jvmtiError {
  return C.SetExtensionEventCallback(jvmti.raw(), extension_event_index, callback)
}

func (jvmti JvmtiEnv) GetErrorName(error C.jvmtiError, name_ptr **C.char) C.jvmtiError {
  return C.GetErrorName(jvmti.raw(), error, name_ptr)
}

func (jvmti JvmtiEnv) GetJLocationFormat(format_ptr *C.jvmtiJlocationFormat) C.jvmtiError {
  return C.GetJLocationFormat(jvmti.raw(), format_ptr)
}

func (jvmti JvmtiEnv) GetSystemProperties(count_ptr *C.jint, property_ptr ***C.char) C.jvmtiError {
  return C.GetSystemProperties(jvmti.raw(), count_ptr, property_ptr)
}

func (jvmti JvmtiEnv) GetSystemProperty(property *C.char, value_ptr **C.char) C.jvmtiError {
  return C.GetSystemProperty(jvmti.raw(), property, value_ptr)
}

func (jvmti JvmtiEnv) SetSystemProperty(property *C.char, value_ptr *C.char) C.jvmtiError {
  return C.SetSystemProperty(jvmti.raw(), property, value_ptr)
}

func (jvmti JvmtiEnv) GetPhase(phase_ptr *C.jvmtiPhase) C.jvmtiError {
  return C.GetPhase(jvmti.raw(), phase_ptr)
}

func (jvmti JvmtiEnv) GetCurrentThreadCpuTimerInfo(info_ptr *C.jvmtiTimerInfo) C.jvmtiError {
  return C.GetCurrentThreadCpuTimerInfo(jvmti.raw(), info_ptr)
}

func (jvmti JvmtiEnv) GetCurrentThreadCpuTime(nanos_ptr *C.jlong) C.jvmtiError {
  return C.GetCurrentThreadCpuTime(jvmti.raw(), nanos_ptr)
}

func (jvmti JvmtiEnv) GetThreadCpuTimerInfo(info_ptr *C.jvmtiTimerInfo) C.jvmtiError {
  return C.GetThreadCpuTimerInfo(jvmti.raw(), info_ptr)
}

func (jvmti JvmtiEnv) GetThreadCpuTime(thread C.jthread, nanos_ptr *C.jlong) C.jvmtiError {
  return C.GetThreadCpuTime(jvmti.raw(), thread, nanos_ptr)
}

func (jvmti JvmtiEnv) GetTimerInfo(info_ptr *C.jvmtiTimerInfo) C.jvmtiError {
  return C.GetTimerInfo(jvmti.raw(), info_ptr)
}

func (jvmti JvmtiEnv) GetTime(nanos_ptr *C.jlong) C.jvmtiError {
  return C.GetTime(jvmti.raw(), nanos_ptr)
}

func (jvmti JvmtiEnv) GetPotentialCapabilities(capabilities_ptr *C.jvmtiCapabilities) C.jvmtiError {
  return C.GetPotentialCapabilities(jvmti.raw(), capabilities_ptr)
}

func (jvmti JvmtiEnv) AddCapabilities(capabilities_ptr *C.jvmtiCapabilities) C.jvmtiError {
  return C.AddCapabilities(jvmti.raw(), capabilities_ptr)
}

func (jvmti JvmtiEnv) RelinquishCapabilities(capabilities_ptr *C.jvmtiCapabilities) C.jvmtiError {
  return C.RelinquishCapabilities(jvmti.raw(), capabilities_ptr)
}

func (jvmti JvmtiEnv) GetAvailableProcessors(processor_count_ptr *C.jint) C.jvmtiError {
  return C.GetAvailableProcessors(jvmti.raw(), processor_count_ptr)
}

func (jvmti JvmtiEnv) GetClassVersionNumbers(klass C.jclass, minor_version_ptr *C.jint, major_version_ptr *C.jint) C.jvmtiError {
  return C.GetClassVersionNumbers(jvmti.raw(), klass, minor_version_ptr, major_version_ptr)
}

func (jvmti JvmtiEnv) GetConstantPool(klass C.jclass, constant_pool_count_ptr *C.jint, constant_pool_byte_count_ptr *C.jint, constant_pool_bytes_ptr **C.uchar) C.jvmtiError {
  return C.GetConstantPool(jvmti.raw(), klass, constant_pool_count_ptr, constant_pool_byte_count_ptr, constant_pool_bytes_ptr)
}

func (jvmti JvmtiEnv) GetEnvironmentLocalStorage(data_ptr *unsafe.Pointer) C.jvmtiError {
  return C.GetEnvironmentLocalStorage(jvmti.raw(), data_ptr)
}

func (jvmti JvmtiEnv) SetEnvironmentLocalStorage(data unsafe.Pointer) C.jvmtiError {
  return C.SetEnvironmentLocalStorage(jvmti.raw(), data)
}

func (jvmti JvmtiEnv) AddToBootstrapClassLoaderSearch(segment *C.char) C.jvmtiError {
  return C.AddToBootstrapClassLoaderSearch(jvmti.raw(), segment)
}

func (jvmti JvmtiEnv) SetVerboseFlag(flag C.jvmtiVerboseFlag, value C.jboolean) C.jvmtiError {
  return C.SetVerboseFlag(jvmti.raw(), flag, value)
}

func (jvmti JvmtiEnv) AddToSystemClassLoaderSearch(segment *C.char) C.jvmtiError {
  return C.AddToSystemClassLoaderSearch(jvmti.raw(), segment)
}

func (jvmti JvmtiEnv) RetransformClasses(class_count C.jint, classes *C.jclass) C.jvmtiError {
  return C.RetransformClasses(jvmti.raw(), class_count, classes)
}

func (jvmti JvmtiEnv) GetOwnedMonitorStackDepthInfo(thread C.jthread, monitor_info_count_ptr *C.jint, monitor_info_ptr **C.jvmtiMonitorStackDepthInfo) C.jvmtiError {
  return C.GetOwnedMonitorStackDepthInfo(jvmti.raw(), thread, monitor_info_count_ptr, monitor_info_ptr)
}

func (jvmti JvmtiEnv) GetObjectSize(object C.jobject, size_ptr *C.jlong) C.jvmtiError {
  return C.GetObjectSize(jvmti.raw(), object, size_ptr)
}

func (jvmti JvmtiEnv) GetLocalInstance(thread C.jthread, depth C.jint, value_ptr *C.jobject) C.jvmtiError {
  return C.GetLocalInstance(jvmti.raw(), thread, depth, value_ptr)
}

func (jvmti JvmtiEnv) SetHeapSamplingInterval(sampling_interval C.jint) C.jvmtiError {
  return C.SetHeapSamplingInterval(jvmti.raw(), sampling_interval)
}

