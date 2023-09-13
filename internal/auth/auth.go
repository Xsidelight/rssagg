package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts as API Key from the request
// Example header:
// Authorization: ApiKey {insert api key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authorization header is not found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("authorization header is not in the correct format")
	}


	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of Authorization header")
	}

	return vals[1], nil
}
