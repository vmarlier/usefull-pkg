package main

import (
	"crypto/tls"
	"fmt"
	"log"

	"golang.org/x/text/encoding/unicode"
	"gopkg.in/ldap.v2"
)

func main() {
	dn := checkAdUser("testReset")
	fmt.Println(dn)
	updatePassword(dn)

}

const (
	admin     = "testLDAP"
	adminPass = "Welcome12"
	// OU where you can find users
	dc = "OU=Users,OU=Workspace Services,OU=D2SIGROUP,DC=d2sim7,DC=loc"
	// AD hostname or IP
	ad = "parpcads01.d2sim7.loc"
	// AD port
	adp = 389
)

func updatePassword(dn string) error {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ad, adp))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// Reconnect with TLS
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}

	err = l.Bind(admin, adminPass)
	if err != nil {
		log.Fatal(err)
	}

	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	// According to the MS docs in the links above
	// The password needs to be enclosed in quotes
	pwdEncoded, _ := utf16.NewEncoder().String("\"Welcome10\"")
	passReq := &ldap.ModifyRequest{
		DN: dn, // DN for the user we're resetting
		ReplaceAttributes: []ldap.PartialAttribute{
			{"unicodePwd", []string{pwdEncoded}},
			//{"userPassword", []string{pwdEncoded}},
		},
	}
	fmt.Println(l.Modify(passReq))

	return nil
}

func checkAdUser(userToCheck string) string {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ad, adp))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// Reconnect with TLS
	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}

	// First bind with a user
	err = l.Bind(admin, adminPass)
	if err != nil {
		log.Fatal(err)
	}

	// Search for a specific user
	searchRequest := ldap.NewSearchRequest(dc, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(sAMAccountName=%s))", userToCheck), []string{"dn"}, nil)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	if len(sr.Entries) == 0 {
		log.Fatal("User does not exist")
	} else if len(sr.Entries) > 1 {
		log.Fatal("Too many entries returned")
	}

	return sr.Entries[0].DN
}
