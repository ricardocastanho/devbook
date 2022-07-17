package controller

import (
	"api/src/config"
	"api/src/models"
	"api/src/presenters"
	"api/src/repositories"
	"api/src/support"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := config.ConnectDatabase()

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repositories.NewUserRepo(db)

	userDB, err := repo.FindUserByUsername(user.Username)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	err = support.CompareHashAndPassword(userDB.Password, user.Password)

	if err != nil {
		presenters.Error(w, http.StatusUnauthorized, errors.New("invalid username or password"))
		return
	}

	token, err := support.GenerateToken(userDB.ID)

	if err != nil {
		presenters.Error(w, http.StatusUnauthorized, err)
		return
	}

	w.Write([]byte(token))
}
