package extractor

import (
	"fmt"
	"regexp"
	"strings"
)

// ExtractErrorCode takes an error message as input and extracts the
// Active Directory Domain Services (ADDS) error code from it.
func ExtractErrorCode(errorMessage string) (string, string, error) {
	// Use a regular expression to find the LDAP result code in the error message.
	reLDAP := regexp.MustCompile(`LDAP Result Code (\d+)`)
	matchLDAP := reLDAP.FindStringSubmatch(errorMessage)

	// Check if the regular expression found a match.
	if len(matchLDAP) < 2 {
		return "", "", fmt.Errorf("LDAP error code not found in the error message")
	}

	// Check for "AcceptSecurityContext error" in the error message.
	if strings.Contains(errorMessage, "AcceptSecurityContext error") {
		// Use a regular expression to extract the error code from the error message.
		reADDS := regexp.MustCompile(`data (\w+),`)
		matchADDS := reADDS.FindStringSubmatch(errorMessage)

		// Check if the regular expression found a match.
		if len(matchADDS) < 2 {
			// If no match is found, return an error indicating that the ADDS error code
			// was not found in the error message.
			return "", "", fmt.Errorf("ADDS error code not found in the error message")
		}

		// Return the extracted LDAP result code and ADDS error code.
		return matchLDAP[1], matchADDS[1], nil
	}

	// If the error message does not contain "AcceptSecurityContext error"
	// or the expected pattern, return an error.
	return "", "", fmt.Errorf("ADDS error code not found in the error message")
}
