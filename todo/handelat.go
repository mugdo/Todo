package todo

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerStruct struct {
	todoService *Service
}

func MakeHTTPHandlers(router *gin.RouterGroup, todoService *Service) {
	h := &handlerStruct{
		todoService: todoService,
	}
	router.POST("todo", h.addTodo)
}

type TodoRequest struct {
	Todo TodoDecode
}

func (h *handlerStruct) addTodo(c *gin.Context) {
	returnValue, Username := h.todoService.isLogin(c)
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
		fmt.Println("bad3")
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	response := &TodoDecode{
		Mssage: "Inserted",
	}
	c.JSON(http.StatusOK, &response)

}
