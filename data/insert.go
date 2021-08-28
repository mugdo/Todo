package data

import (
	"main.go/model"
)

var Key = []byte("Secrect-key")
var LoginInfo = map[string]string{
	"name1": "pass1",
	"nmae2": "pass2",
	"user1": "pass1",
	"new1":  "p1",
}
var Imsg = []model.StoreInfo{
	{
		Name: "me1",
		Mssage: []string{
			"Go jogging",
			"Buy groceries",
			"Feed the dog",
		},
	},
	{
		Name: "John",
		Mssage: []string{
			"Water",
			"Food",
			"Shelter",
		},
	},
	{
		Name: "user1",
		Mssage: []string{
			"Build an app",
			"Eat Icecream",
			"Learn how to code",
			"Go to gym",
		},
	},
	{
		Name: "new1",
		Mssage: []string{
			"T1",
			"T2",
			"T3",
			"T4",
		},
	},
}
