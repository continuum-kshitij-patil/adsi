package adsi

import (
	"github.com/continuum-nilesh-akhade/adsi/api"
	"github.com/scjalliance/comshim"
)

// ADSI Objects of LDAP:  https://msdn.microsoft.com/library/aa772208
// ADSI Objects of WinNT: https://msdn.microsoft.com/library/aa772211

// Object provides access to Active Directory objects.
type DirectorySearch struct {
	object
	iface *api.IDirectorySearch
}

// NewDirectorySearch returns an object that manages the given COM interface.
func NewDirectorySearch(iface *api.IDirectorySearch) *DirectorySearch {
	comshim.Add(1)
	return &DirectorySearch{iface: iface}
}

func (o *DirectorySearch) closed() bool {
	return (o.iface == nil)
}

// Close will release resources consumed by the object. It should be
// called when the object is no longer needed.
func (o *DirectorySearch) Close() {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return
	}
	defer comshim.Done()
	o.iface.Release() // FIXME: What happens if release returns an error?
	o.iface = nil
}

// ExecuteSearch retrieves the name of the object.
func (o *DirectorySearch) ExecuteSearch(searchFilter string, attrNames []string) (phSearchResult uintptr, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return 0, ErrClosed
	}
	phSearchResult, err = o.iface.ExecuteSearch(searchFilter, attrNames)
	return
}

func (o *DirectorySearch) AbandonSearch(hSearchResult uintptr) (err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return ErrClosed
	}
	err = o.iface.AbandonSearch(hSearchResult)
	return
}

func (o *DirectorySearch) CloseSearchHandle(hSearchResult uintptr) (err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return ErrClosed
	}
	err = o.iface.CloseSearchHandle(hSearchResult)
	return
}

func (o *DirectorySearch) FreeColumn(pSearchColumn uintptr) (err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return ErrClosed
	}
	err = o.iface.FreeColumn(pSearchColumn)
	return
}

func (o *DirectorySearch) GetColumn(hSearchResult uintptr, columnName string, pSearchColumn uintptr) (err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return ErrClosed
	}
	err = o.iface.GetColumn(hSearchResult, columnName, pSearchColumn)
	return
}

func (o *DirectorySearch) GetFirstRow(hSearchResult uintptr) (err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return ErrClosed
	}
	err = o.iface.GetFirstRow(hSearchResult)
	return
}

func (o *DirectorySearch) GetNextColumnName(hSearchResult uintptr) (columnName string, err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return "", ErrClosed
	}
	columnName, err = o.iface.GetNextColumnName(hSearchResult)
	return
}

func (o *DirectorySearch) GetNextRow(hSearchResult uintptr) (err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return ErrClosed
	}
	err = o.iface.GetNextRow(hSearchResult)
	return
}

func (o *DirectorySearch) GetPreviousRow(hSearchResult uintptr) (err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return ErrClosed
	}
	err = o.iface.GetPreviousRow(hSearchResult)
	return
}

func (o *DirectorySearch) SetSearchPreference(searchPrefs api.PADsSearchPrefInfo, dwNumPrefs int64) (err error) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.closed() {
		return ErrClosed
	}
	err = o.iface.SetSearchPreference(searchPrefs, dwNumPrefs)
	return
}
