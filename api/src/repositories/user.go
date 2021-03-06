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
		"SELECT id, first_name, last_name, username, created_at, updated_at FROM users WHERE username LIKE ? AND deleted_at IS NULL",
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

func (repo *UserRepository) FindUserByUsername(username string) (models.User, error) {
	rows, err := repo.db.Query(
		"SELECT id, username, password FROM users WHERE username = ?",
		username,
	)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User

	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
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

func (repo *UserRepository) DeleteUser(userId string) error {
	stmt, err := repo.db.Prepare(
		"UPDATE users SET deleted_at = now() WHERE id = ?",
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		userId,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) FollowUser(userId string, followerId string) error {
	stmt, err := repo.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)",
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, followerId)

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) UnfollowUser(userId string, followerId string) error {
	stmt, err := repo.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? AND follower_id = ?",
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, followerId)

	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) GetFollowers(userID string) ([]models.User, error) {
	rows, err := repo.db.Query(`
		SELECT 
			u.id, u.first_name, u.last_name, u.username, u.created_at, u.updated_at
		FROM
			users u
		INNER JOIN
			followers f ON f.follower_id = u.id
		WHERE
			f.user_id = ?
			AND u.deleted_at IS NULL;
	`, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var followers []models.User

	for rows.Next() {
		var follower models.User

		err := rows.Scan(
			&follower.ID,
			&follower.FirstName,
			&follower.LastName,
			&follower.Username,
			&follower.CreatedAt,
			&follower.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

func (repo *UserRepository) GetFollowing(userID string) ([]models.User, error) {
	rows, err := repo.db.Query(`
		SELECT 
			u.id, u.first_name, u.last_name, u.username, u.created_at, u.updated_at
		FROM
			users u
		INNER JOIN
			followers f ON f.user_id = u.id
		WHERE
			f.follower_id = ?
			AND u.deleted_at IS NULL;
	`, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var following []models.User

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

		following = append(following, user)
	}

	return following, nil
}

func (repo *UserRepository) GetPassword(userID string) (string, error) {
	rows, err := repo.db.Query(
		"SELECT password FROM users WHERE id = ? AND deleted_at IS NULL",
		userID,
	)

	if err != nil {
		return "", err
	}

	defer rows.Close()

	var password string

	for rows.Next() {
		err := rows.Scan(
			&password,
		)

		if err != nil {
			return "", err
		}
	}

	return password, nil
}

func (repo *UserRepository) ChangePassword(userID, password string) error {
	stmt, err := repo.db.Prepare(
		"UPDATE users SET password = ? WHERE id = ?",
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(password, userID)

	if err != nil {
		return err
	}

	return nil
}
