package ldap

import (
	"daemon_backend.bin/component/extractor"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

// Bind connects to LDAP and binds with admin credentials
func Bind() (*ldap.Conn, error) {
	BindUsername := extractor.ExtractStrFromFile("ldap", "bindUsername")
	BindPassword := extractor.ExtractStrFromFile("ldap", "bindPassword")

	// Connect to LDAP
	l, err := connect()
	if err != nil {
		return nil, err
	}

	// Bind with admin credentials
	err = l.Bind(BindUsername, BindPassword)
	if err != nil {
		return nil, fmt.Errorf("LDAP bind error: %s", err)
	}

	return l, nil
}
