// +build windows

package api

import (
	"syscall"
	"unsafe"
)

//AbandonSearch sdffdsf
func (v *IDirectorySearch) AbandonSearch(hSearchResult uintptr) (err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().AbandonSearch),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(hSearchResult),
		0)
	if hr == 0 {
		return
	}
	return convertHresultToError(hr)
}

//CloseSearchHandle sdffdsf
func (v *IDirectorySearch) CloseSearchHandle(hSearchResult uintptr) (err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().CloseSearchHandle),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(hSearchResult),
		0)
	if hr == 0 {
		return
	}
	return convertHresultToError(hr)
}

// ExecuteSearch func
func (v *IDirectorySearch) ExecuteSearch(searchFilter string, attrNames []string) (phSearchResult uintptr, err error) {
	pszSearchFilter, err := syscall.UTF16PtrFromString(searchFilter)
	if err != nil {
		return
	}
	// var numberOfAttrs int32
	// numberOfAttrs = int32(len(attrNames))
	// attrNamesArray := make([]*uint16, numberOfAttrs)
	// for i, attrName := range attrNames {
	// 	attrNamesArray[i], err = syscall.UTF16PtrFromString(attrName)
	// 	if err != nil {
	// 		return
	// 	}
	// }
	// hr, _, _ := syscall.Syscall6(
	// 	uintptr(v.VTable().ExecuteSearch),
	// 	5,
	// 	uintptr(unsafe.Pointer(v)),
	// 	uintptr(unsafe.Pointer(pszSearchFilter)),
	// 	uintptr(unsafe.Pointer(&attrNamesArray)),
	// 	uintptr(numberOfAttrs),
	// 	uintptr(unsafe.Pointer(&phSearchResult)),
	// 	0)
	hr, _, _ := syscall.Syscall6(
		uintptr(v.VTable().ExecuteSearch),
		5,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(pszSearchFilter)),
		uintptr(0),
		uintptr(0),
		uintptr(unsafe.Pointer(&phSearchResult)),
		0)
	if hr == 0 {
		return
	}
	return 0, convertHresultToError(hr)
}

//FreeColumn sdffdsf
func (v *IDirectorySearch) FreeColumn(pSearchColumn uintptr) (err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().FreeColumn),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(pSearchColumn),
		0)
	if hr == 0 {
		return
	}
	return convertHresultToError(hr)
}

//GetColumn sdffdsf
func (v *IDirectorySearch) GetColumn(hSearchResult uintptr, columnName string, pSearchColumn uintptr) (err error) {

	szColumnName, err := syscall.UTF16PtrFromString(columnName)
	if err != nil {
		return
	}

	hr, _, _ := syscall.Syscall6(
		uintptr(v.VTable().GetColumn),
		4,
		uintptr(unsafe.Pointer(v)),
		uintptr(hSearchResult),
		uintptr(unsafe.Pointer(szColumnName)),
		uintptr(pSearchColumn),
		0,
		0)
	if hr == 0 {
		return
	}
	return convertHresultToError(hr)
}

//GetFirstRow sdffdsf
func (v *IDirectorySearch) GetFirstRow(hSearchResult uintptr) (err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().GetFirstRow),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(hSearchResult),
		0)
	if hr == 0 {
		return
	}
	return convertHresultToError(hr)
}

//GetNextColumnName sdffdsf
func (v *IDirectorySearch) GetNextColumnName(hSearchResult uintptr) (columnName string, err error) {
	var szColumnName = make([]uint16, syscall.MAX_PATH+1)
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().GetNextColumnName),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&szColumnName[0])),
		0)
	if hr == 0 {
		columnName = syscall.UTF16ToString(szColumnName)
		return
	}
	return "", convertHresultToError(hr)
}

//GetNextRow sdffdsf
func (v *IDirectorySearch) GetNextRow(hSearchResult uintptr) (err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().GetNextRow),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(hSearchResult),
		0)
	if hr == 0 {
		return
	}
	return convertHresultToError(hr)
}

//GetPreviousRow sdffdsf
func (v *IDirectorySearch) GetPreviousRow(hSearchResult uintptr) (err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().GetPreviousRow),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(hSearchResult),
		0)
	if hr == 0 {
		return
	}
	return convertHresultToError(hr)
}

//SetSearchPreference sdffdsf
func (v *IDirectorySearch) SetSearchPreference(searchPrefs PADsSearchPrefInfo, dwNumPrefs int64) (err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().SetSearchPreference),
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&searchPrefs)),
		uintptr(dwNumPrefs))
	if hr == 0 {
		return
	}
	return convertHresultToError(hr)
}
