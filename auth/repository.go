package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"main.go/data"
)

type Repo interface {
	findByName(name string) (User, error)
	tokenValid(c *gin.Context) (bool, string)
}
type repoStruct struct {
	DBSession *mgo.Session
	DBName    string
	DBTable   string
}

func (r *repoStruct) FindByName(name string) (User, error) {
	var user User
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{"name": name}).One(&user)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	return user, nil
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

	if !tkn.Valid {
		return false, ""
	}

	return true, claims.Name
}
func NewRepository(dbSession *mgo.Session) *repoStruct {
	return &repoStruct{
		DBSession: dbSession,
		DBName:    data.DBName,
		DBTable:   data.AuthTable,
	}
}
