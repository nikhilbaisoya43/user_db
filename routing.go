package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/user/{eid}", GetUserById).Methods("GET")
	r.HandleFunc("/user", CreateUser).Methods("POST")
	r.HandleFunc("/user/{eid}", UpdateUser).Methods("put")
	r.HandleFunc("/user/{eid}", DeleteUserById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
