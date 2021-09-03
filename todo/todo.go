package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func Init(r *gin.RouterGroup, dbSession *mgo.Session) {
	repoService := NewRepository(dbSession)
	todoService := NewAuthService(repoService)
	MakeHTTPHandlers(r, todoService)

}
