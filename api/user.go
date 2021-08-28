package api

import (
	"encoding/json"
	"log"
	"net/http"

	"main.go/data"

	"main.go/model"
)

func AddInfo(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json")
	returnValue, Username := IsLogin(W, r)
	if !returnValue {
		W.WriteHeader(http.StatusUnauthorized)
		return
	}
	var userdata = model.MessageDecode{}
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		log.Fatal(err)
	}

	for ind, value := range data.Imsg {
		if value.Name == Username {
			value.Mssage = append(value.Mssage, userdata.Mssage)
			data.Imsg[ind].Mssage = value.Mssage
			break
		}
	}

	msg := model.MessageDecode{
		Mssage: "Data inserted Successfuly",
	}
	json.NewEncoder(W).Encode(msg)
}

func Vewinfo(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json")
	returnValue, Username := IsLogin(W, r)
	expectedvalue := true
	if returnValue != expectedvalue {
		W.WriteHeader(http.StatusUnauthorized)
		return
	}

	newUser := model.StoreInfo{}
	for _, value := range data.Imsg {
		if value.Name == Username {
			newUser = value
		}

	}
	json.NewEncoder(W).Encode(newUser)
}
func DeleteInfo(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json")
	returnValue, Username := IsLogin(W, r)
	if !returnValue {
		W.WriteHeader(http.StatusUnauthorized)
		return
	}

	var userdata = model.MessageDecode{}
	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		log.Fatal(err)
	}
	for ind, value := range data.Imsg {
		if value.Name == Username {
			for i, v := range value.Mssage {
				if v == userdata.Mssage {
					data.Imsg[ind].Mssage = append(data.Imsg[ind].Mssage[:i], data.Imsg[ind].Mssage[i+1:]...)
					break
				}
			}
		}
	}
	msg := model.MessageDecode{
		Mssage: "Delete message",
	}
	json.NewEncoder(W).Encode(msg)
}
func VewAll(W http.ResponseWriter, r *http.Request) {
	W.Header().Set("Content-Type", "application/json")
	returnValue, _ := IsLogin(W, r)
	if !returnValue {
		var User []string
		for _, value := range data.Imsg {
			User = append(User, value.Name)

		}
		json.NewEncoder(W).Encode(User)

	} else {
		json.NewEncoder(W).Encode(data.Imsg)
	}
}
