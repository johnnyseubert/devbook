package security

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/johnnyseubert/devbook/src/config"
)

func GenerateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix() // 12381237189471
	permissions["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.JWT_SECRET))
}

func ValidateToken(request *http.Request) error {
	tokenString := excractToken(request)

	token, err := jwt.Parse(tokenString, retrieveTokenKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func excractToken(request *http.Request) string {
	token := request.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func retrieveTokenKey(token *jwt.Token) (any, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		return nil, http.ErrAbortHandler
	}

	return []byte(config.JWT_SECRET), nil
}
