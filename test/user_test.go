package test

import (
	"net/http"
	"testing"
)

// TestInvalidCredentialsLoginUsingLDAP tests the scenario where login using LDAP fails due to invalid credentials.
func TestInvalidCredentialsLoginUsingLDAP(t *testing.T) {
	method := "POST"
	endpoint := "http://127.0.0.1:8080/api/v2/user/login"
	payload := map[string]string{
		"username": "test",
		"password": "someRandomPassword",
	}
	expectedStatusCode := http.StatusUnauthorized
	expectedResponse := `{"error":"INVALID_CREDENTIALS"}`

	performHTTPRequest(t, method, endpoint, payload, expectedStatusCode, expectedResponse)
}

// TestNoInputLoginUsingLDAP tests the scenario where no input fields are provided for login using LDAP.
func TestNoInputLoginUsingLDAP(t *testing.T) {
	method := "POST"
	endpoint := "http://127.0.0.1:8080/api/v2/user/login"
	payload := map[string]string{}
	expectedStatusCode := http.StatusBadRequest
	expectedResponse := `{"error":"missing fields: username, password"}`

	performHTTPRequest(t, method, endpoint, payload, expectedStatusCode, expectedResponse)
}

// TestNoUsernameLoginUsingLDAP tests the scenario where no username is provided for login using LDAP.
func TestNoUsernameLoginUsingLDAP(t *testing.T) {
	method := "POST"
	endpoint := "http://127.0.0.1:8080/api/v2/user/login"
	payload := map[string]string{
		"password": "someRandomPassword",
	}
	expectedStatusCode := http.StatusBadRequest
	expectedResponse := `{"error":"missing fields: username"}`

	performHTTPRequest(t, method, endpoint, payload, expectedStatusCode, expectedResponse)
}

// TestNoPasswordLoginUsingLDAP tests the scenario where no password is provided for login using LDAP.
func TestNoPasswordLoginUsingLDAP(t *testing.T) {
	method := "POST"
	endpoint := "http://127.0.0.1:8080/api/v2/user/login"
	payload := map[string]string{
		"username": "someRandomUsername",
	}
	expectedStatusCode := http.StatusBadRequest
	expectedResponse := `{"error":"missing fields: password"}`

	performHTTPRequest(t, method, endpoint, payload, expectedStatusCode, expectedResponse)
}

// TestExpiredUserLoginUsingLDAP tests the scenario where account is expired for login using LDAP.
func TestExpiredUserLoginUsingLDAP(t *testing.T) {
	method := "POST"
	endpoint := "http://127.0.0.1:8080/api/v2/user/login"
	payload := map[string]string{
		"username": "testUserExpired",
		"password": "Zxc:123456",
	}
	expectedStatusCode := http.StatusUnauthorized
	expectedResponse := `{"error":"ACCOUNT_EXPIRED"}`

	performHTTPRequest(t, method, endpoint, payload, expectedStatusCode, expectedResponse)
}

// TestExpiredPasswordLoginUsingLDAP tests the scenario where password is expired for login using LDAP.
func TestExpiredPasswordLoginUsingLDAP(t *testing.T) {
	method := "POST"
	endpoint := "http://127.0.0.1:8080/api/v2/user/login"
	payload := map[string]string{
		"username": "testPassExpired",
		"password": "Zxc:123456",
	}
	expectedStatusCode := http.StatusUnauthorized
	expectedResponse := `{"error":"PASSWORD_EXPIRED"}`

	performHTTPRequest(t, method, endpoint, payload, expectedStatusCode, expectedResponse)
}

// TestUserMustChangePasswordLoginUsingLDAP tests the scenario where user must change password for login using LDAP.
func TestUserMustChangePasswordLoginUsingLDAP(t *testing.T) {
	method := "POST"
	endpoint := "http://127.0.0.1:8080/api/v2/user/login"
	payload := map[string]string{
		"username": "testUserMustChangePass",
		"password": "Zxc:123456",
	}
	expectedStatusCode := http.StatusUnauthorized
	expectedResponse := `{"error":"USER_MUST_RESET_PASSWORD"}`

	performHTTPRequest(t, method, endpoint, payload, expectedStatusCode, expectedResponse)
}

// TestAccountDisabledLoginUsingLDAP tests the scenario where account is disabled for login using LDAP.
func TestAccountDisabledLoginUsingLDAP(t *testing.T) {
	method := "POST"
	endpoint := "http://127.0.0.1:8080/api/v2/user/login"
	payload := map[string]string{
		"username": "testAccountDisabled",
		"password": "Zxc:123456",
	}
	expectedStatusCode := http.StatusUnauthorized
	expectedResponse := `{"error":"ACCOUNT_DISABLED"}`

	performHTTPRequest(t, method, endpoint, payload, expectedStatusCode, expectedResponse)
}

// TestUserNotPermittedLogonNowLoginUsingLDAP tests the scenario where account is not permitted to log in at a certain time for login using LDAP.
func TestUserNotPermittedLogonNowLoginUsingLDAP(t *testing.T) {
	method := "POST"
	endpoint := "http://127.0.0.1:8080/api/v2/user/login"
	payload := map[string]string{
		"username": "testNotPermittedLogonNow",
		"password": "Zxc:123456",
	}
	expectedStatusCode := http.StatusUnauthorized
	expectedResponse := `{"error":"NOT_PERMITTED_TO_LOGON_AT_THIS_TIME"}`

	performHTTPRequest(t, method, endpoint, payload, expectedStatusCode, expectedResponse)
}
