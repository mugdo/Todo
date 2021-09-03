package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"main.go/auth"
)

func Init(r *gin.RouterGroup, dbSession *mgo.Session, authService auth.Service) {
	repoService := NewRepository(dbSession)
	todoService := NewTodohService(repoService)
	MakeHTTPHandlers(r, todoService, authService)

}
