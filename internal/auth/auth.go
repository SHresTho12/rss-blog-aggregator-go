package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extracts the api key from a HTTP request header and returns
// Authorization : ApiKey{api_key}
func GetApi(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("no api key provided")
	}
	val := strings.Split(apiKey, " ")

	if val[0] != "ApiKey" {
		return "", errors.New("invalid api key provided")
	}

	if len(val) != 2 {
		return "", errors.New("invalid api key provided")
	}

	return val[1], nil

}
