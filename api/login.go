package api

import (
	"encoding/json"
	"net/http"
	"time"

	"main.go/data"

	"github.com/dgrijalva/jwt-go"

	"main.go/model"
)

func Login(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json")
	var user model.Userinfo
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		W.WriteHeader(http.StatusBadRequest)
		return
	}
	expectPasswar, ok := data.LoginInfo[user.Name]
	if !ok || expectPasswar != user.Passward {
		W.WriteHeader(http.StatusBadRequest)
		return
	}
	expararTime := time.Now().Add(time.Minute * 10)
	clam := &model.Claim{
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expararTime.Unix(),
		},
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, clam)
	tokenString, err := Token.SignedString(data.Key)
	if err != nil {
		W.WriteHeader(http.StatusBadRequest)
		return
	}
	var returnJson = &model.TokenMsg{
		Token: tokenString,
	}
	json.NewEncoder(W).Encode(returnJson)

}
