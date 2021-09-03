package todo

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"main.go/data"
)

type repoStruct struct {
	DBSession *mgo.Session
	DBName    string
	DBTable   string
}

func (r *repoStruct) InsertByName(req ITodo) error {
	var user ITodo
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{"name": req.Name}).One(&user)
	fmt.Println("user: ", user.Name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Mssage = append(req.Mssage, user.Mssage...)
	err = coll.Update(user, req)
	return err
}
func (r *repoStruct) FindUserByName() []ITodo {
	var user []ITodo
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{}).All(&user)
	if err != nil {
		fmt.Println(err)
		return user
	}
	return user
}
func (r *repoStruct) singleUser(name string) ITodo {
	var user ITodo
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{"name": name}).One(&user)
	fmt.Println("user: ", user.Name)
	if err != nil {
		fmt.Println(err)
		return ITodo{}
	}
	return user
}
func (r *repoStruct) deleteTodoMessage(req ITodo) error {
	var user ITodo
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{"name": req.Name}).One(&user)
	fmt.Println("user1: ", req.Name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("len: ", len(user.Mssage))
	for ind, value := range user.Mssage {
		if value == req.Mssage[0] {
			fmt.Println("message : ", value)
			user.Mssage = append(user.Mssage[:ind], user.Mssage[ind+1:]...)
			fmt.Println(user.Mssage)
			break
		}
	}
	selector := bson.M{"name": req.Name}
	// info, err := coll.Upsert(selector, bson.M{"$set":{"mssage":user.Mssage}})
	change := bson.M{"$pull": bson.M{"mssage": req.Mssage[0]}}
	fmt.Println(change)
	err = coll.Update(selector, change)
	// coll.UpdateAll(selector, user)

	// err = coll.Update(user,user)
	fmt.Println(err)
	return err
}
func NewRepository(dbSession *mgo.Session) *repoStruct {
	return &repoStruct{
		DBSession: dbSession,
		DBName:    data.DBName,
		DBTable:   data.TodoTable,
	}
}
