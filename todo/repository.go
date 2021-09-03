package todo

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"main.go/data"
)

type repoStruct struct {
	DBSession *mgo.Session
	DBName    string
	DBTable   string
}

func (s *repoStruct) tokenValid(c *gin.Context) (bool, string) {
	reqToken := c.GetHeader("Authorization")
	splitToken := strings.Split(reqToken, " ")
	if len(splitToken) != 2 {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return false, ""
	}
	claims := &claim{}
	tkn, err := jwt.ParseWithClaims(splitToken[1], claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte("key"), nil

		})
	if err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println("claim: ", claims.Name)

	if !tkn.Valid {
		return false, ""
	}

	return true, claims.Name
}
func (r *repoStruct) InsertByName(req ITodo) error {
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Insert(&req)
	return err
}
func NewRepository(dbSession *mgo.Session) *repoStruct {
	return &repoStruct{
		DBSession: dbSession,
		DBName:    data.DBName,
		DBTable:   data.TodoTable,
	}
}
