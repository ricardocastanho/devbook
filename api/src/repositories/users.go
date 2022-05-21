package repositories

import (
	"api/src/models"
	"database/sql"

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

	uuid := uuid.New()

	_, err = stmt.Exec(
		uuid.String(),
		user.FirstName,
		user.LastName,
		user.Username,
		user.Password,
	)

	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}