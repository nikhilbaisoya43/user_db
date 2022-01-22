package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)
// CURD operations on user through api in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp User
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []User
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee User
	Database.Find(&employee, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(employee)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee User
	Database.Find(&employee, mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&employee)
	Database.Save(&employee)
	json.NewEncoder(w).Encode(employee)
}
func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp User
	Database.Delete(&emp, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode("emp is deleted")
}
