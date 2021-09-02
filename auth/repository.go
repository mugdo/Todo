package auth

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"main.go/data"
)

// type repo interface {
// 	FindByName(name string) (User, error)
// }
type repoStruct struct {
	DBSession *mgo.Session
	DBName    string
	DBTable   string
}

func NewRepository(dbSession *mgo.Session) *repoStruct {
	return &repoStruct{
		DBSession: dbSession,
		DBName:    data.DBName,
		DBTable:   data.AuthTable,
	}
}
func (r *repoStruct) FindByName(name string) (User, error) {
	var user User
	fmt.Println("r.DBName", r.DBName)
	fmt.Println("r.DBTable", r.DBTable)
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{"name": name}).One(&user)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	fmt.Println("user: ", user)
	return user, nil
}
