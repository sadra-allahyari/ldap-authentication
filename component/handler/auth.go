package handler

import (
	"daemon_backend.bin/component/respond"
	"fmt"
	"net/http"
)

// handleLdapError responds to the client with the appropriate LDAP error message based on the LDAP error code.
// It uses a map to look up the error message corresponding to the given LDAP error code.
// If the error code is not found in the map, it responds with an "Unknown LDAP Error" message.
func handleLdapError(w http.ResponseWriter, ldapErrCode, addsErrCode string) {
	// Create a map to store LDAP error messages
	ldapErrorMessages := map[string]string{
		"0": "LDAP_SUCCESS",
		"1": "LDAP_OPERATIONS_ERROR",
		"2": "LDAP_PROTOCOL_ERROR",
		"3": "LDAP_TIMELIMIT_EXCEEDED",
		"4": "LDAP_SIZELIMIT_EXCEEDED",
		"5": "LDAP_COMPARE_FALSE",
		"6": "LDAP_COMPARE_TRUE",
		"7": "LDAP_AUTH_METHOD_NOT_SUPPORTED",
		// Standard LDAP: "7":   "LDAP_STRONG_AUTH_NOT_SUPPORTED"
		"8": "LDAP_STRONG_AUTH_REQUIRED",
		"9": "LDAP_PARTIAL_RESULTS",
		//"9":   "LDAP_REFERRAL_V2"
		"10": "LDAP_REFERRAL",
		"11": "LDAP_ADMIN_LIMIT_EXCEEDED",
		"12": "LDAP_UNAVAILABLE_CRIT_EXTENSION",
		// Standard LDAP: "12":  "LDAP_UNAVAILABLE_CRITICAL_EXTENSION",
		"13": "LDAP_CONFIDENTIALITY_REQUIRED",
		"14": "LDAP_SASL_BIND_IN_PROGRESS",
		// Standard LDAP: "14":  "LDAP_SASLBIND_IN_PROGRESS",
		"16": "LDAP_NO_SUCH_ATTRIBUTE",
		"17": "LDAP_UNDEFINED_TYPE",
		"18": "LDAP_INAPPROPRIATE_MATCHING",
		"19": "LDAP_CONSTRAINT_VIOLATION",
		"20": "LDAP_ATTRIBUTE_OR_VALUE_EXISTS",
		// Standard LDAP: "20":  "LDAP_TYPE_OR_VALUE_EXISTS",
		"21": "LDAP_INVALID_SYNTAX",
		"32": "LDAP_NO_SUCH_OBJECT",
		"33": "LDAP_ALIAS_PROBLEM",
		"34": "LDAP_INVALID_DN_SYNTAX",
		"35": "LDAP_IS_LEAF",
		"36": "LDAP_ALIAS_DEREF_PROBLEM",
		"48": "LDAP_INAPPROPRIATE_AUTH",
		"49": handleADDSError(addsErrCode),
		"50": "LDAP_INSUFFICIENT_RIGHTS",
		// Standard LDAP: "50":  "LDAP_INSUFFICIENT_ACCESS",
		"51": "LDAP_BUSY",
		"52": "LDAP_UNAVAILABLE",
		"53": "LDAP_UNWILLING_TO_PERFORM",
		"54": "LDAP_LOOP_DETECT",
		"60": "LDAP_SORT_CONTROL_MISSING",
		"61": "LDAP_OFFSET_RANGE_ERROR",
		"64": "LDAP_NAMING_VIOLATION",
		"65": "LDAP_OBJECT_CLASS_VIOLATION",
		"66": "LDAP_NOT_ALLOWED_ON_NONLEAF",
		"67": "LDAP_NOT_ALLOWED_ON_RDN",
		"68": "LDAP_ALREADY_EXISTS",
		"69": "LDAP_NO_OBJECT_CLASS_MODS",
		"70": "LDAP_RESULTS_TOO_LARGE",
		"71": "LDAP_AFFECTS_MULTIPLE_DSAS",
		"77": "LDAP_VIRTUAL_LIST_VIEW_ERROR",
		"80": "LDAP_OTHER",
		"81": "LDAP_SERVER_DOWN",
		"82": "LDAP_LOCAL_ERROR",
		"83": "LDAP_ENCODING_ERROR",
		"84": "LDAP_DECODING_ERROR",
		"85": "LDAP_TIMEOUT",
		"86": "LDAP_AUTH_UNKNOWN",
		"87": "LDAP_FILTER_ERROR",
		"88": "LDAP_USER_CANCELLED",
		"89": "LDAP_PARAM_ERROR",
		"90": "LDAP_NO_MEMORY",
		"91": "LDAP_CONNECT_ERROR",
		"92": "LDAP_NOT_SUPPORTED",
		"93": "LDAP_NO_RESULTS_RETURNED",
		// Standard LDAP: "93":  "LDAP_CONTROL_NOT_FOUND",
		"94": "LDAP_CONTROL_NOT_FOUND",
		// Standard LDAP: "94":  "LDAP_NO_RESULTS_RETURNED",
		"95": "LDAP_MORE_RESULTS_TO_RETURN",
		"96": "LDAP_CLIENT_LOOP",
		// Standard LDAP: "96":  "LDAP_URL_ERR_NOTLDAP",
		"97": "LDAP_REFERRAL_LIMIT_EXCEEDED",
		// Standard LDAP: "97":  "LDAP_URL_ERR_NODN",
		"98":  "LDAP_URL_ERR_BADSCOPE",
		"99":  "LDAP_URL_ERR_MEM",
		"100": "LDAP_CLIENT_LOOP",
		"101": "LDAP_REFERRAL_LIMIT_EXCEEDED",
		"112": "LDAP_SSL_ALREADY_INITIALIZED",
		"113": "LDAP_SSL_INITIALIZE_FAILED",
		"114": "LDAP_SSL_CLIENT_INIT_NOT_CALLED",
		"115": "LDAP_SSL_PARAM_ERROR",
		"116": "LDAP_SSL_HANDSHAKE_FAILED",
		"117": "LDAP_SSL_GET_CIPHER_FAILED",
		"118": "LDAP_SSL_NOT_AVAILABLE",
		"128": "LDAP_NO_EXPLICIT_OWNER",
		"129": "LDAP_NO_LOCK",
		"133": "LDAP_DNS_NO_SERVERS",
		"134": "LDAP_DNS_TRUNCATED",
		"135": "LDAP_DNS_INVALID_DATA",
		"136": "LDAP_DNS_RESOLVE_ERROR",
		"137": "LDAP_DNS_CONF_FILE_ERROR",
		"160": "LDAP_XLATE_E2BIG",
		"161": "LDAP_XLATE_EINVAL",
		"162": "LDAP_XLATE_EILSEQ",
		"163": "LDAP_XLATE_NO_ENTRY",
		"176": "LDAP_REG_FILE_NOT_FOUND",
		"177": "LDAP_REG_CANNOT_OPEN",
		"178": "LDAP_REG_ENTRY_NOT_FOUND",
		"192": "LDAP_CONF_FILE_NOT_OPENED",
		"193": "LDAP_PLUGIN_NOT_LOADED",
		"194": "LDAP_PLUGIN_FUNCTION_NOT_RESOLVED",
		"195": "LDAP_PLUGIN_NOT_INITIALIZED",
		"196": "LDAP_PLUGIN_COULD_NOT_BIND",
		"208": "LDAP_SASL_GSS_NO_SEC_CONTEXT",
	}

	// Look up the LDAP error message based on the LDAP error code
	errorMessage, ok := ldapErrorMessages[ldapErrCode]
	if !ok {
		// If the LDAP error code is not found in the map, set the error message to "Unknown LDAP Error" with the LDAP error code
		errorMessage = fmt.Sprintf("Unknown LDAP Error (Code: %s)", ldapErrCode)
	}

	// Respond with the appropriate LDAP error message
	respond.ErrRespond(w, http.StatusUnauthorized, fmt.Sprintf(errorMessage))
}

// handleADDSError handles ADDS (Active Directory Domain Services) errors.
// It takes a http.ResponseWriter and an LDAP error message as input, extracts the ADDS error code,
// looks up the corresponding error message, and responds with the appropriate ADDS error message.
func handleADDSError(addsErrCode string) string {
	// Create a map to store ADDS error messages
	addsErrorMessages := map[string]string{
		"52e": "INVALID_CREDENTIALS",
		"525": "INVALID_CREDENTIALS",
		// the 525 error code originally means USER_NOT_FOUND but due to security measurements changed to INVALID_CREDENTIALS,
		"530": "NOT_PERMITTED_TO_LOGON_AT_THIS_TIME",
		"531": "RESTRICTED_TO_SPECIFIC_MACHINES",
		"532": "PASSWORD_EXPIRED",
		"533": "ACCOUNT_DISABLED",
		"568": "ERROR_TOO_MANY_CONTEXT_IDS",
		"701": "ACCOUNT_EXPIRED",
		"773": "USER_MUST_RESET_PASSWORD",
	}

	// Look up the ADDS error message based on the ADDS error code
	errorText, ok := addsErrorMessages[addsErrCode]
	if !ok {
		// If the ADDS error code is not found in the map, set the error message to "Unknown ADDS Error" with the ADDS error code
		errorText = fmt.Sprintf("Unknown ADDS Error (Code: %s)", addsErrCode)
	}

	return errorText
}
