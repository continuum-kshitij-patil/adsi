package adsi

import (
	"github.com/continuum-nilesh-akhade/adsi/api"
	"github.com/scjalliance/comshim"
)

// User provides access to Active Directory users.
type User struct {
	object
	iface *api.IADsUser
}

// NewUser returns a user that manages the given COM interface.
func NewUser(iface *api.IADsUser) *User {
	comshim.Add(1)
	return &User{iface: iface, object: object{iface: &iface.IADs}}
}

func (u *User) closed() bool {
	return (u.iface == nil)
}

// Close will release resources consumed by the user. It should be
// called when the user is no longer needed.
func (u *User) Close() {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return
	}
	defer comshim.Done()
	u.iface.Release()
	u.object.iface = nil
	u.iface = nil
}

// PasswordRequired retrieves the attribute of the user
func (u *User) PasswordRequired() (passReq bool, err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return false, ErrClosed
	}
	passReq, err = u.iface.PasswordRequired()
	return
}

// Description retrieves the description of the user.
func (u *User) Description() (description string, err error) {
	u.m.Lock()
	defer u.m.Unlock()
	if u.closed() {
		return "", ErrClosed
	}
	description, err = u.iface.Description()
	return
}
