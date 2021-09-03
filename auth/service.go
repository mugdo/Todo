package auth

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
func (authService *Service) loginUser(Login Login) (User, error) {
	userf, err := authService.repoService.FindByName(Login.Name)
	if err != nil {
		return User{}, err
	}
	return userf, nil

}
func (authService *Service) IsLogin(c *gin.Context) (bool, string) {
	value, usename := authService.repoService.tokenValid(c)
	if !value {
		return false, ""
	}
	return value, usename
}
