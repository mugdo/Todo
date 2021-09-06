package todo

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"main.go/data"
)

type RepoInterface interface {
	InsertByName(req ITodo) error
	FindUserByName() []ITodo
	singleUser(name string) ITodo
	deleteTodoMessage(req ITodo) error
	updateTodoMessage(req ITodo) error
}

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
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	selector := bson.M{"name": req.Name}
	change := bson.M{"$pull": bson.M{"mssage": req.Mssage[0]}} //pull data from database
	fmt.Println(change)
	err := coll.Update(selector, change) //update data without change
	// coll.UpdateAll(selector, user)
	// err = coll.Update(user,user)
	fmt.Println(err)
	return err
}
func (r *repoStruct) updateTodoMessage(req ITodo) error {
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	selector := bson.M{"name": req.Name}
	change := bson.M{"$pull": bson.M{"mssage": req.Mssage[0]}}
	err := coll.Update(selector, change)
	if err != nil {
		return err
	}
	change = bson.M{"$push": bson.M{"mssage": req.Mssage[1]}}
	err = coll.Update(selector, change)
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
