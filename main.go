package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"main.go/auth"
	"main.go/todo"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("DB not connect")
		return
	}
	initializeAllServices(v1, session)
	r.Run(":3002")
}
func initializeAllServices(router *gin.RouterGroup, dbSession *mgo.Session) {
	Service := auth.Init(router, dbSession)
	todo.Init(router, dbSession, *Service)
}
