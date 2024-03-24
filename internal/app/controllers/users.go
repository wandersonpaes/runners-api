package controllers

import "net/http"

func CreatingUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating a User!"))
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching for all Users!"))
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching a User!"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating a User!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting a User!"))
}
