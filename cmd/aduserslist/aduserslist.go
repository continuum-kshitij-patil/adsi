package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/continuum-nilesh-akhade/adsi"
)

const (
	msNPAllowDialin    = "msNPAllowDialin"
	objectSid          = "objectSid"
	sAMAccountName     = "sAMAccountName"
	userAccountControl = "userAccountControl"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	adsiObject, err := adsi.Open(flag.Arg(0)) // use LDAP: for root LDAP path
	//TODO:Add filter &((objectCategory=person)(objectClass=user))
	if err != nil {
		log.Fatal(err)
	}
	defer adsiObject.Close()

	printUsers(adsiObject)
}

func printUsers(parent *adsi.Object) {
	c, err := parent.ToContainer()
	if err != nil {
		return
	}
	defer c.Close()

	iter, err := c.Children()
	if err != nil {
		log.Fatal(err)
	}
	defer iter.Close()

	i := 0
	for child, err := iter.Next(); err == nil; child, err = iter.Next() {
		defer child.Close()
		name, err := child.Name()
		if err != nil {
			log.Fatal(err)
		}
		class, err := child.Class()
		if err != nil {
			log.Fatal(err)
		}
		if class == "user" {
			fmt.Printf("\n%v-> %s %s \n", i, name, class)
			user, err := child.ToUser()
			if err != nil {
				log.Fatal(err)
			} else {
				//list user attributes
				fullName, err := user.FullName()
				if err != nil {
					fmt.Printf("IADsUser.FullName() err: %v\n", err)
				} else {
					fmt.Printf("IADsUser.FullName(): %v\n", fullName)
				}
				passReq, err := user.PasswordRequired()
				if err != nil {
					fmt.Printf("IADsUser.PasswordRequired() err: %v\n", err)
				} else {
					fmt.Printf("IADsUser.PasswordRequired(): %v\n", passReq)
				}
				accDisabled, err := user.AccountDisabled()
				if err != nil {
					fmt.Printf("IADsUser.AccountDisabled() err: %v\n", err)
				} else {
					fmt.Printf("IADsUser.AccountDisabled(): %v\n", accDisabled)
				}
				accLocked, err := user.IsAccountLocked()
				if err != nil {
					fmt.Printf("IADsUser.IsAccountLocked() err: %v\n", err)
				} else {
					fmt.Printf("IADsUser.IsAccountLocked(): %v\n", accLocked)
				}
				passMinLen, err := user.PasswordMinimumLength()
				if err != nil {
					fmt.Printf("IADsUser.PasswordMinimumLength() err: %v\n", err)
				} else {
					fmt.Printf("IADsUser.PasswordMinimumLength(): %v\n", passMinLen)
				}
				reqUniqPass, err := user.RequireUniquePassword()
				if err != nil {
					fmt.Printf("IADsUser.RequireUniquePassword() err: %v\n", err)
				} else {
					fmt.Printf("IADsUser.RequireUniquePassword(): %v\n", reqUniqPass)
				}
				// lastLogin, err := user.LastLogin()
				// if err != nil {
				// 	fmt.Printf("IADsUser.LastLogin() err: %v\n", err)
				// } else {
				// 	fmt.Printf("IADsUser.LastLogin(): %v\n", lastLogin)
				// }
				passExpDate, err := user.PasswordExpirationDate()
				if err != nil {
					fmt.Printf("IADsUser.PasswordExpirationDate() err: %v\n", err)
				} else {
					fmt.Printf("IADsUser.PasswordExpirationDate(): %v\n", passExpDate)
				}

				propValue, err := user.Get(msNPAllowDialin)
				if err != nil {
					fmt.Printf("IADs.Get(\"%v\") err: %v\n", msNPAllowDialin, err)
				} else {
					fmt.Printf("IADs.Get(\"%v\"): %v\n", msNPAllowDialin, propValue.Value())
				}
				propValue, err = user.Get(objectSid)
				if err != nil {
					fmt.Printf("IADs.Get(\"%v\") err: %v\n", objectSid, err)
				} else {
					fmt.Printf("IADs.Get(\"%v\"): %v\n", objectSid, propValue.Value())
				}
				propValue, err = user.Get(sAMAccountName)
				if err != nil {
					fmt.Printf("IADs.Get(\"%v\") err: %v\n", sAMAccountName, err)
				} else {
					fmt.Printf("IADs.Get(\"%v\"): %v\n", sAMAccountName, propValue.Value())
				}
				propValue, err = user.Get(userAccountControl)
				if err != nil {
					fmt.Printf("IADs.Get(\"%v\") err: %v\n", userAccountControl, err)
				} else {
					fmt.Printf("IADs.Get(\"%v\"): %v\n", userAccountControl, propValue.Value())
				}
			}
		}
		i++
		if i > 10 {
			fmt.Println("\n--------------10 users are enough----------------------")
			break
		}
	}
}
