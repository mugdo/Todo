package api

import (
	"net/http"
	"strings"

	"main.go/data"

	"github.com/dgrijalva/jwt-go"
	"main.go/model"
)

func IsLogin(W http.ResponseWriter, r *http.Request) (bool, string) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, " ")
	if len(splitToken) != 2 {
		W.WriteHeader(http.StatusUnauthorized)
		return false,""
	}
	claims := &model.Claim{}
	tkn, err := jwt.ParseWithClaims(splitToken[1], claims,
		func(t *jwt.Token) (interface{}, error) {
			return data.Key, nil

		})
	if err != nil {
		W.WriteHeader(http.StatusBadRequest)
	}

	if !tkn.Valid {
		return false, ""
	}

	return true, claims.Name

}
