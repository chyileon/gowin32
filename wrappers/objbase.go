/*
 * Copyright (c) 2014 MongoDB, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the license is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package wrappers

import (
	"syscall"
	"unsafe"
)

const (
	COINIT_APARTMENTTHREADED = 0x00000002
	COINIT_MULTITHREADED     = 0x00000000
	COINIT_DISABLE_OLE1DDE   = 0x00000004
	COINIT_SPEED_OVER_MEMORY = 0x00000008
)

var (
	modole32 = syscall.NewLazyDLL("ole32.dll")

	procCoCreateInstance = modole32.NewProc("CoCreateInstance")
	procCoInitializeEx   = modole32.NewProc("CoInitializeEx")
	procCoUninitialize   = modole32.NewProc("CoUninitialize")
)

func CoCreateInstance(clsid *syscall.GUID, outer *IUnknown, clsContext uint32, iid *syscall.GUID, object *uintptr) error {
	r1, _, _ := procCoCreateInstance.Call(
		uintptr(unsafe.Pointer(clsid)),
		uintptr(unsafe.Pointer(outer)),
		uintptr(clsContext),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(object)))
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func CoInitializeEx(flags uint32) error {
	r1, _, _ := procCoInitializeEx.Call(0, uintptr(flags))
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func CoUninitialize() {
	procCoUninitialize.Call()
}