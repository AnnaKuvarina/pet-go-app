package utils

import (
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

func ExtractToken(req *http.Request) (string, error) {
	authHeader := req.Header.Get("Authorization")
	authHeaderContext := strings.Split(authHeader, " ")
	if len(authHeaderContext) < 2 {
		return "", errors.New("Invalid token")
	}

	return authHeaderContext[1], nil
}
