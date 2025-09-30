package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User created"))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All users"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User by ID"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User updated"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User deleted"))
}
