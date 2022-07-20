package repositories

import (
	"api/src/models"
	"database/sql"

	"github.com/google/uuid"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) PostRepository {
	return PostRepository{db}
}

func (repo *PostRepository) CreatePost(post models.Post) (string, error) {
	stmt, err := repo.db.Prepare(`
		INSERT INTO posts (id, title, content, author_id) VALUES (?, ?, ?, ?)
	`)

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	post.ID = uuid.New().String()

	_, err = stmt.Exec(post.ID, post.Title, post.Content, post.Author.ID)

	if err != nil {
		return "", err
	}

	return post.ID, nil
}
