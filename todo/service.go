package todo

import (
	"github.com/gin-gonic/gin"
)

type Service struct {
	repoService *repoStruct
}

func NewAuthService(repo *repoStruct) *Service {
	return &Service{
		repoService: repo,
	}
}
func (todoService *Service) isLogin(c *gin.Context) (bool, string) {
	value, usename := todoService.repoService.tokenValid(c)
	if !value {
		return false, ""
	}
	return value, usename
}
func (todoService *Service) insertTodo(S ITodo) error {
	err := todoService.repoService.InsertByName(S)
	return err

}
