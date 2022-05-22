package controller

import (
	"api/src/config"
	"api/src/models"
	"api/src/presenters"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users"))
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find User"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		presenters.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)

	if err != nil {
		presenters.Error(w, http.StatusBadRequest, err)
		return
	}

	err = user.Validate()

	if err != nil {
		presenters.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := config.ConnectDatabase()

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repositories.NewUserRepo(db)

	id, err := repo.CreateUser(&user)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	user.ID = id
	presenters.JSON(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update User"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
