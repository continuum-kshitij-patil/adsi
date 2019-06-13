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

	dirSearch, err := adsi.OpenDirectorySearch(flag.Arg(0)) // use LDAP: for root LDAP path
	//TODO:Add filter &((objectCategory=person)(objectClass=user))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer dirSearch.Close()
	searchHandle, err := dirSearch.ExecuteSearch("(&(objectClass=user)(objectCategory=person))", []string{"name", "distinguishedName"})
	if err != nil {
		log.Fatal(err)
		return
	}
	defer dirSearch.CloseSearchHandle(searchHandle)

	err = dirSearch.GetFirstRow(searchHandle)
	if err != nil {
		log.Fatal(err)
		return
	}
	columnName, err := dirSearch.GetNextColumnName(searchHandle)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(columnName)

}
