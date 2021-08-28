package model

import "github.com/dgrijalva/jwt-go"

type Userinfo struct {
	Name     string `json:"name"`
	Passward string ` json:"passward"`
}
type Claim struct {
	Name               string `json:"username"`
	jwt.StandardClaims        //embading another struct jwt standeer libary
}
type MessageDecode struct {
	Mssage string `json:"mssage"`
}

type StoreInfo struct {
	Name   string   `json:"name"`
	Mssage []string `json:"mssage"`
}

// For other message..
type TokenMsg struct {
	Token string `json:"token"`
}
