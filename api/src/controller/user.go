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
	"strings"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	username := strings.ToLower(r.URL.Query().Get("username"))

	db, err := config.ConnectDatabase()

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repositories.NewUserRepo(db)

	users, err := repo.GetUsers(username)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	presenters.JSON(w, http.StatusOK, users)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	userID := params["id"]

	db, err := config.ConnectDatabase()

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repositories.NewUserRepo(db)

	user, err := repo.FindUser(userID)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	presenters.JSON(w, http.StatusOK, user)
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

	err = user.Validate("add")

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
	params := mux.Vars(r)

	userID := params["id"]

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		presenters.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	userIDFromToken, err := support.GetUserLoggedFromToken(r)

	if err != nil {
		presenters.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDFromToken {
		presenters.Error(w, http.StatusForbidden, errors.New("you can only update your own user"))
		return
	}

	var user models.User

	err = json.Unmarshal(body, &user)

	if err != nil {
		presenters.Error(w, http.StatusBadRequest, err)
		return
	}

	err = user.Validate("edit")

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

	err = repo.UpdateUser(userID, &user)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	presenters.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID := params["id"]

	userIDFromToken, err := support.GetUserLoggedFromToken(r)

	if err != nil {
		presenters.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDFromToken {
		presenters.Error(w, http.StatusForbidden, errors.New("you can only delete your own user"))
		return
	}

	db, err := config.ConnectDatabase()

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repositories.NewUserRepo(db)

	err = repo.DeleteUser(userID)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	presenters.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	userLoggedID, err := support.GetUserLoggedFromToken(r)

	if err != nil {
		presenters.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	followerID := params["id"]

	if userLoggedID == followerID {
		presenters.Error(w, http.StatusForbidden, errors.New("you can't follow yourself"))
		return
	}

	db, err := config.ConnectDatabase()

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repositories.NewUserRepo(db)

	err = repo.FollowUser(userLoggedID, followerID)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	presenters.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	userLoggedID, err := support.GetUserLoggedFromToken(r)

	if err != nil {
		presenters.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	followerID := params["id"]

	db, err := config.ConnectDatabase()

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repositories.NewUserRepo(db)

	err = repo.UnfollowUser(userLoggedID, followerID)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	presenters.JSON(w, http.StatusNoContent, nil)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	db, err := config.ConnectDatabase()

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repositories.NewUserRepo(db)

	followers, err := repo.GetFollowers(userID)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	presenters.JSON(w, http.StatusOK, followers)
}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	db, err := config.ConnectDatabase()

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repositories.NewUserRepo(db)

	following, err := repo.GetFollowing(userID)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	presenters.JSON(w, http.StatusOK, following)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	userLoggedID, err := support.GetUserLoggedFromToken(r)

	if err != nil {
		presenters.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userLoggedID {
		presenters.Error(w, http.StatusForbidden, errors.New("you can only change your own password"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		presenters.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var changePassword models.ChangePassword

	err = json.Unmarshal(body, &changePassword)

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

	actualPassword, err := repo.GetPassword(userID)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	err = support.CompareHashAndPassword(actualPassword, changePassword.Old)

	if err != nil {
		presenters.Error(w, http.StatusUnauthorized, errors.New("the passwords don't match"))
		return
	}

	hashedPassword, err := support.Hash(changePassword.New)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	err = repo.ChangePassword(userID, string(hashedPassword))

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	presenters.JSON(w, http.StatusNoContent, nil)
}
