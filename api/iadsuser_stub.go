// +build !windows

package api

import ole "github.com/go-ole/go-ole"

// PasswordRequired retrieves the password required attribute
func (v *IADsUser) PasswordRequired() (passReq bool, err error) {
	return false, ole.NewError(ole.E_NOTIMPL)
}

// Description retrieves the description
func (v *IADsUser) Description() (description string, err error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}
