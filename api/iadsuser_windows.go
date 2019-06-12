// +build windows

package api

import (
	"syscall"
	"unsafe"

	ole "github.com/go-ole/go-ole"
)

// FullName retrieves the FullName of the user
func (v *IADsUser) FullName() (fullName string, err error) {
	var bstr *int16
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().FullName),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&bstr)),
		0)
	if bstr != nil {
		defer ole.SysFreeString(bstr)
	}
	if hr == 0 {
		fullName = ole.BstrToString((*uint16)(unsafe.Pointer(bstr)))
	} else {
		return "", convertHresultToError(hr)
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

// PasswordRequired field
func (v *IADsUser) PasswordRequired() (passReq bool, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().PasswordRequired),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&passReq)),
		0)
	if hr == 0 {
		return
	} else {
		return false, convertHresultToError(hr)
	}
	return
}

// AccountDisabled field
func (v *IADsUser) AccountDisabled() (accDisabled bool, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().AccountDisabled),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&accDisabled)),
		0)
	if hr == 0 {
		return
	} else {
		return false, convertHresultToError(hr)
	}
}

// IsAccountLocked field
func (v *IADsUser) IsAccountLocked() (accLocked bool, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().IsAccountLocked),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&accLocked)),
		0)
	if hr == 0 {
		return
	} else {
		return false, convertHresultToError(hr)
	}
}

// RequireUniquePassword field
func (v *IADsUser) RequireUniquePassword() (reqUniqPass bool, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().RequireUniquePassword),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&reqUniqPass)),
		0)
	if hr == 0 {
		return
	} else {
		return false, convertHresultToError(hr)
	}
}

// PasswordMinimumLength field
func (v *IADsUser) PasswordMinimumLength() (minPassLen int64, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().PasswordMinimumLength),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&minPassLen)),
		0)
	if hr == 0 {
		return
	} else {
		return 0, convertHresultToError(hr)
	}
}

// LastLogin field
func (v *IADsUser) LastLogin() (lastLoginDate *IADsLargeInteger, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().LastLogin),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&lastLoginDate)),
		0)
	if hr == 0 {
		return
	} else {
		return nil, convertHresultToError(hr)
	}
}

// PasswordExpirationDate field
func (v *IADsUser) PasswordExpirationDate() (passExpDate int64, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().PasswordExpirationDate),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&passExpDate)),
		0)
	if hr == 0 {
		return
	} else {
		return 0, convertHresultToError(hr)
	}
}
