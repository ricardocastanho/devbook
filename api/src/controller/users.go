package controller

import "net/http"

func GetUsers(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Get Users"))
}

func FindUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Find User"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Create User"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Delete User"))
}
