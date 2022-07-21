package controller

import (
	"api/src/config"
	"api/src/models"
	"api/src/presenters"
	"api/src/repositories"
	"api/src/support"
	"encoding/json"
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
