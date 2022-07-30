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

	"github.com/gorilla/mux"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := support.GetUserLoggedFromToken(r)

	if err != nil {
		presenters.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := config.ConnectDatabase()

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repositories.NewPostRepo(db)

	posts, err := repo.GetPosts(userID)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	presenters.JSON(w, http.StatusOK, posts)
}

func FindPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID := params["id"]

	db, err := config.ConnectDatabase()

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repositories.NewPostRepo(db)

	post, err := repo.FindPost(postID)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	presenters.JSON(w, http.StatusOK, post)
}

func CreatePosts(w http.ResponseWriter, r *http.Request) {
	userID, err := support.GetUserLoggedFromToken(r)

	if err != nil {
		presenters.Error(w, http.StatusUnauthorized, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		presenters.Error(w, http.StatusBadRequest, err)
		return
	}

	var post models.Post

	err = json.Unmarshal(body, &post)

	if err != nil {
		presenters.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.Author.ID = userID

	err = post.Validate()

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

	repo := repositories.NewPostRepo(db)

	postID, err := repo.CreatePost(post)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	post.ID = postID

	presenters.JSON(w, http.StatusCreated, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID := params["id"]

	userID, err := support.GetUserLoggedFromToken(r)

	if err != nil {
		presenters.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := config.ConnectDatabase()

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repo := repositories.NewPostRepo(db)

	post, err := repo.FindPost(postID)

	if err != nil {
		presenters.Error(w, http.StatusBadRequest, err)
		return
	}

	if post.Author.ID != userID {
		presenters.Error(w, http.StatusForbidden, errors.New("you are not the author of this post"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		presenters.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(body, &post)

	if err != nil {
		presenters.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = post.Validate()

	if err != nil {
		presenters.Error(w, http.StatusBadRequest, err)
		return
	}

	err = repo.UpdatePost(post)

	if err != nil {
		presenters.Error(w, http.StatusInternalServerError, err)
		return
	}

	presenters.JSON(w, http.StatusOK, post)
}
