package auth

import (
	"fmt"
)

// GetUserDN generates the user's Distinguished Name (DN) based on the username.
func GetUserDN(username, baseDN string) string {
	return fmt.Sprintf("cn=%s,%s", username, baseDN)
}
