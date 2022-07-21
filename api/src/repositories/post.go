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

func (repo *PostRepository) GetPosts(userID string) ([]models.Post, error) {
	rows, err := repo.db.Query(`
		SELECT
			p.id, p.title, p.content, p.likes, p.created_at, p.updated_at,
			u.id, u.first_name, u.last_name, u.username
		FROM posts p
		INNER JOIN users u ON p.author_id = u.id
		WHERE p.author_id = ?
		ORDER BY p.created_at DESC
	`, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Author.ID,
			&post.Author.FirstName,
			&post.Author.LastName,
			&post.Author.Username,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (repo *PostRepository) FindPost(postID string) (models.Post, error) {
	rows, err := repo.db.Query(`
		SELECT
			p.id, p.title, p.content, p.likes, p.created_at, p.updated_at,
			u.id, u.first_name, u.last_name, u.username
		FROM posts p
		INNER JOIN users u ON p.author_id = u.id
		WHERE p.id = ?
		LIMIT 1
	`, postID)

	if err != nil {
		return models.Post{}, err
	}

	defer rows.Close()

	var post models.Post

	for rows.Next() {
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.UpdatedAt,
			&post.Author.ID,
			&post.Author.FirstName,
			&post.Author.LastName,
			&post.Author.Username,
		)

		if err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
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
