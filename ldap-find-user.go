package main

import (
	"log"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func main() {
	if err := findUser(); err != nil {
		log.Fatal(err)
	}
}

func findUser() error {
	ldapURL := "ldap://0.0.0.0:389"
	adminusername := "cn=admin,dc=example,dc=org"
	adminpassword := "adminpassword"
	baseDN := "dc=example,dc=org"

	l, err := ldap.DialURL(ldapURL)
	if err != nil {
		return err
	}
	l.Start()

	err = l.Bind(adminusername, adminpassword)
	if err != nil {
		return err
	}

	searchAll := &ldap.SearchRequest{
		BaseDN: baseDN,
		Scope:  ldap.ScopeWholeSubtree,
		Filter: "(objectClass=*)",
	}

	fmt.Println("Search: (objectClass=*)")
	sr, err := l.Search(searchAll)
	if err != nil {
		return err
	}

	printResult(sr.Entries)

	fmt.Println("")
	fmt.Println("(&(objectClass=inetOrgPerson)(uid=user01))")
	searchSpecificUser := &ldap.SearchRequest{
		BaseDN: baseDN,
		Scope:  ldap.ScopeWholeSubtree,
		Filter: "(&(objectClass=inetOrgPerson)(uid=user01))",
	}

	sr, err = l.Search(searchSpecificUser)
	if err != nil {
		return err
	}

	if len(sr.Entries) != 1 {
		fmt.Printf("%+v\n", sr)
		log.Fatal("User does not exist or too many entries returned")
	} else {
		fmt.Printf("****** FOUND! ******")
	}
	return nil
}

func printResult(entries []*ldap.Entry) {
	for _, entry := range entries {
		fmt.Println("DN:", entry.DN)
		for _, attr := range entry.Attributes {
			for i := 0; i < len(attr.Values); i++ {
				fmt.Printf("%s: %s\n", attr.Name, attr.Values[i])
			}
		}
		fmt.Println()
	}
}
