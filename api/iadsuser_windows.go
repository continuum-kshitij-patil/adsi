// +build windows

package api

import (
	"syscall"
	"unsafe"

	ole "github.com/go-ole/go-ole"
)

// PasswordRequired field
func (v *IADsUser) PasswordRequired() (passReq bool, err error) {
	var bstr *bool
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().PasswordRequired),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if hr == 0 {
		passReq = *bstr
	} else {
		return false, convertHresultToError(hr)
	}
	return
}

// Description retrieves the description of the user
func (v *IADsUser) Description() (description string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Description),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr == 0 {
		description = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	} else {
		return "", convertHresultToError(hr)
	}
	return
}
