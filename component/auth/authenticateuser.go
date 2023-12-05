package auth

import (
	"daemon_backend.bin/component/extractor"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

// AuthenticateUser performs a bind to authenticate the user
func AuthenticateUser(l *ldap.Conn, username, password string) error {
	BaseDN := extractor.ExtractStrFromFile("ldap", "baseDN")
	userDN := fmt.Sprintf("cn=%s,%s", username, BaseDN)

	err := l.Bind(userDN, password)
	if err != nil {
		return fmt.Errorf("authentication failed: %s", err)
	}

	return nil
}
