package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return UserRepository{db}
}

func (repo *UserRepository) CreateUser(user *models.User) (string, error) {
	stmt, err := repo.db.Prepare(
		"INSERT INTO users (id, first_name, last_name, username, password) VALUES (?, ?, ?, ?, ?)",
	)

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	user.ID = uuid.New().String()

	_, err = stmt.Exec(
		user.ID,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Password,
	)

	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (repo *UserRepository) GetUsers(username string) ([]models.User, error) {
	username = fmt.Sprintf("%%%s%%", username)

	rows, err := repo.db.Query(
		"SELECT id, first_name, last_name, username, created_at, updated_at FROM users WHERE username LIKE ?",
		username,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepository) FindUser(id string) (models.User, error) {
	rows, err := repo.db.Query(
		"SELECT id, first_name, last_name, username, created_at, updated_at FROM users WHERE id = ?",
		id,
	)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User

	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repo *UserRepository) UpdateUser(userID string, user *models.User) error {
	stmt, err := repo.db.Prepare(
		"UPDATE users SET first_name = ?, last_name = ?, username = ? WHERE id = ?",
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Username,
		userID,
	)

	if err != nil {
		return err
	}

	return nil
}
