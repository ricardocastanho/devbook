package support

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Hash(s string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
}

func CompareHashAndPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(config.SecretKey)
}

func ValidateToken(r *http.Request) error {
	rawToken := getToken(r)

	if rawToken == "" {
		return errors.New("token not provided")
	}

	token, err := jwt.Parse(rawToken, getSignature)

	if err != nil {
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func GetUserLoggedFromToken(r *http.Request) (string, error) {
	rawToken := getToken(r)

	if rawToken == "" {
		return "", errors.New("token not provided")
	}

	token, err := jwt.Parse(rawToken, getSignature)

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userId := fmt.Sprintf("%s", claims["user_id"])

		return userId, nil
	}

	return "", errors.New("invalid token")
}

func getToken(r *http.Request) string {
	header := r.Header.Get("Authorization")

	token := strings.Split(header, " ")

	if len(token) == 2 {
		return token[1]
	}

	return ""
}

func getSignature(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		return nil, fmt.Errorf(
			"unexpected signing method: %v",
			token.Header["alg"],
		)
	}

	return config.SecretKey, nil
}
