package todo

import "github.com/dgrijalva/jwt-go"

type TodoDecode struct {
	Mssage string `json:"mssage"`
}

type claim struct {
	Name               string `json:"username"`
	jwt.StandardClaims        //embading another struct jwt standeer libary
}
type ITodo struct {
	Name   string   `json:"name"`
	Mssage []string `json:"mssage"`
}

