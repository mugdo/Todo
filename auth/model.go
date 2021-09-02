package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type Login struct {
	Name     string `json:"name"`
	Passward string ` json:"passward"`
}
type User struct {
	ID       []uint8 `json:"_id" bson:"_id"`
	Name     string  `json:"name" bson:"name"`
	Password string  `json:"passward" bson:"passward"`
}
type claim struct {
	Name               string `json:"username"`
	jwt.StandardClaims        //embading another struct jwt standeer libary
}
type loginResponse struct {
	Token string `json:"token"`
}
