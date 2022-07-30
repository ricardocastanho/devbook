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
		SELECT DISTINCT
			p.id, p.title, p.content, p.likes, p.created_at, p.updated_at,
			u.id, u.first_name, u.last_name, u.username
		FROM posts p
		INNER JOIN users u ON u.id = p.author_id
		LEFT JOIN followers f ON f.user_id = p.author_id
		WHERE (p.author_id = ? OR f.user_id = p.author_id) AND p.deleted_at IS NULL
		ORDER BY p.created_at DESC
		`,
		userID,
	)

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
		WHERE p.id = ? AND p.deleted_at IS NULL
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

func (repo *PostRepository) UpdatePost(post models.Post) error {
	stmt, err := repo.db.Prepare(`
		UPDATE posts SET title = ?, content = ? WHERE id = ?
	`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Content, post.ID)

	if err != nil {
		return err
	}

	return nil
}

func (repo *PostRepository) DeletePost(postID string) error {
	stmt, err := repo.db.Prepare(
		"UPDATE posts SET deleted_at = now() WHERE id = ?",
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(postID)

	if err != nil {
		return err
	}

	return nil
}

func (repo *PostRepository) GetPostsByUser(userID string) ([]models.Post, error) {
	rows, err := repo.db.Query(`
		SELECT DISTINCT
			p.id, p.title, p.content, p.likes, p.created_at, p.updated_at,
			u.id, u.first_name, u.last_name, u.username
		FROM posts p
		INNER JOIN users u ON p.author_id = u.id
		WHERE p.author_id = ? AND p.deleted_at IS NULL
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

func (repo *PostRepository) LikePost(postID string) error {
	stmt, err := repo.db.Prepare(`
		UPDATE posts SET likes = likes + 1 WHERE id = ?
	`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(postID)

	if err != nil {
		return err
	}

	return nil
}

func (repo *PostRepository) UnlikePost(postID string) error {
	stmt, err := repo.db.Prepare(`
		UPDATE posts SET likes =
			CASE WHEN likes > 0 THEN likes - 1 ELSE 0 END 
		WHERE id = ? AND deleted_at IS NULL
	`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(postID)

	if err != nil {
		return err
	}

	return nil
}
