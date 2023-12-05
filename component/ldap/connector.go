package ldap

import (
	"daemon_backend.bin/component/extractor"
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

// connect establishes an LDAP connection
func connect() (*ldap.Conn, error) {
	FQDN := extractor.ExtractStrFromFile("ldap", "FQDN")
	protocol := extractor.ExtractStrFromFile("ldap", "protocol")
	port := extractor.ExtractStrFromFile("ldap", "port")
	URL := fmt.Sprintf("%s://%s:%s", protocol, FQDN, port)

	l, err := ldap.DialURL(URL)
	if err != nil {
		return nil, fmt.Errorf("LDAP dial error: %s", err)
	}

	return l, nil
}
