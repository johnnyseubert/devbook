package security

import (
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
