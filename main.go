package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/api"
)

func main() {
	fmt.Println("Todo Start..")
	r := mux.NewRouter()

	r.HandleFunc("/api/login", api.Login).Methods("POST")
	r.HandleFunc("/api/task", api.AddInfo).Methods("POST")
	r.HandleFunc("/api/task", api.Vewinfo).Methods("GET")
	r.HandleFunc("/api/task", api.DeleteInfo).Methods("DELETE")
	r.HandleFunc("/api/users", api.VewAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))

}
