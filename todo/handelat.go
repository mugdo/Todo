package todo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/auth"
)

type handlerStruct struct {
	todoService *Service
	authService auth.Service
}

func MakeHTTPHandlers(router *gin.RouterGroup, todoService *Service, service auth.Service) {
	h := &handlerStruct{
		todoService: todoService,
		authService: service,
	}
	router.GET("todo", h.getuser)
	router.POST("todo", h.addTodo)
	router.DELETE("todo", h.deleteTodo)
	router.PUT("todo", h.updateTodo)
}

type TodoRequest struct {
	Todo TodoDecode
}

func (h *handlerStruct) addTodo(c *gin.Context) {
	returnValue, Username := h.authService.IsLogin(c)
	if !returnValue {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	decodeTodo := TodoDecode{}

	err := c.ShouldBindJSON(&decodeTodo)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	VTodo := ITodo{}
	VTodo.Name = Username
	VTodo.Mssage = append(VTodo.Mssage, decodeTodo.Mssage)
	err = h.todoService.insertTodo(VTodo)
	if err != nil {
		fmt.Println("bad1")
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	response := &TodoDecode{
		Mssage: "Inserted",
	}
	c.JSON(http.StatusOK, &response)
}
func (h *handlerStruct) getuser(c *gin.Context) {
	returnValue, Username := h.authService.IsLogin(c)
	var users []ITodo
	if !returnValue {

		users = h.todoService.users()
		c.JSON(http.StatusOK, &users)
	} else {
		user := h.todoService.user(Username)
		c.JSON(http.StatusOK, user)
	}
}
func (h *handlerStruct) deleteTodo(c *gin.Context) {
	fmt.Println("deletetodo...")
	returnValue, Username := h.authService.IsLogin(c)
	if !returnValue {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	decodeTodo := TodoDecode{}

	err := c.ShouldBindJSON(&decodeTodo)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	VTodo := ITodo{}
	VTodo.Name = Username
	VTodo.Mssage = append(VTodo.Mssage, decodeTodo.Mssage)
	err = h.todoService.dlete(VTodo)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	response := &TodoDecode{
		Mssage: "Deketed",
	}
	c.JSON(http.StatusOK, &response)

}
func (h *handlerStruct) updateTodo(c *gin.Context) {
	fmt.Println("deletetodo...")
	returnValue, Username := h.authService.IsLogin(c)
	if !returnValue {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	decodeTodo := TodoDecode{}

	err := c.ShouldBindJSON(&decodeTodo)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	VTodo := ITodo{}
	VTodo.Name = Username
	VTodo.Mssage = append(VTodo.Mssage, decodeTodo.Mssage)
	err = h.todoService.dlete(VTodo)
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	response := &TodoDecode{
		Mssage: "Deketed",
	}
	c.JSON(http.StatusOK, &response)

}
