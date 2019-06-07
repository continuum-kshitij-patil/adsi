package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	adsi "gopkg.in/adsi.v0"
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
			//list user attributes
			user, err := child.ToUser()
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(user.Description())
			}
		}
		fmt.Printf("\n%s %s", name, class)
		i++
	}
}
