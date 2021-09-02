package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"main.go/auth"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	r.Use(cors.New(config))

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
	auth.Init(router, dbSession)
}
