package repositories

import (
	"api/src/models"
	"database/sql"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) AuthRepository {
	return AuthRepository{db}
}

func (repo *AuthRepository) Login(username string, password string) (string, error) {
	var token string

	rows, err := repo.db.Query("SELECT id, password FROM users WHERE username = ?", username)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.Username, &user.Password)

		if err != nil {
			return "", err
		}
	}

	return token, nil
}
