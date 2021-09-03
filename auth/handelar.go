package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	login(c *gin.Context)
}
type handlerStruct struct {
	authService *Service
}
type loginRequest struct {
	Login Login
}

func MakeHTTPHandlers(router *gin.RouterGroup, authService *Service) {
	h := &handlerStruct{
		authService: authService,
	}

	router.POST("auth/login", h.login)
	router.POST("auth", h.validToken)
}
func (h *handlerStruct) login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req.Login); err != nil {
		c.JSON(http.StatusInternalServerError, loginResponse{
			Token: "",
		})
		return
	}
	user, err := h.authService.loginUser(req.Login)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Name != req.Login.Name || user.Password != req.Login.Passward {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	expararTime := time.Now().Add(time.Minute * 10)
	clam := &claim{
		Name: req.Login.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expararTime.Unix(),
		},
	}
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, clam)
	tokenString, err := Token.SignedString([]byte("key"))
	if err != nil {
		fmt.Println(err)
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	response := &loginResponse{
		Token: tokenString,
	}

	c.JSON(http.StatusOK, &response)

}
func (h *handlerStruct) validToken(c *gin.Context) {
	returnValue, Username := h.authService.IsLogin(c)
	if !returnValue {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	response := &tokenValid{
		Name:   Username,
		Mssage: "Token valid",
	}
	c.JSON(http.StatusOK, &response)

}
