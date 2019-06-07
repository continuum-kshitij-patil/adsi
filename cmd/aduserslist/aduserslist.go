package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/continuum-nilesh-akhade/adsi"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	adsiObject, err := adsi.Open("LDAP://" + flag.Arg(0))
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
			user, err := child.ToUser()
			if err != nil {
				log.Fatal(err)
			} else {
				//list user attributes
				fullName, _ := user.FullName()
				fmt.Printf("Username: %v\n", fullName)
				passReq, _ := user.PasswordRequired()
				fmt.Printf("PasswordRequired: %v\n", passReq)
				accDisabled, _ := user.AccountDisabled()
				fmt.Printf("AccountDisabled: %v\n", accDisabled)
				accLocked, _ := user.IsAccountLocked()
				fmt.Printf("IsAccountLocked: %v\n", accLocked)
				passMinLen, _ := user.PasswordMinimumLength()
				fmt.Printf("PasswordMinimumLength: %v\n", passMinLen)
				reqUniqPass, _ := user.RequireUniquePassword()
				fmt.Printf("RequireUniquePassword: %v\n", reqUniqPass)
			}
		}
		fmt.Printf("\n%s %s \n", name, class)
		i++
	}
}
