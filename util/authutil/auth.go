package authutil

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	v1 "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
)

var (
	key = []byte(os.Getenv("KEY"))
)

const (
	defaultExpiredTimeInMinute = 1 * time.Minute
	defaultAppName             = "something"
)

// SignJWT ...
func SignJWT(user *v1.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.Id,
		"username":   user.Username,
		"created_at": user.CreatedAt,
		"exp":        time.Now().Add(defaultExpiredTimeInMinute).Unix(),
		"iss":        defaultAppName,
	})
	return token.SignedString(key)
}

// ValidateJWT ...
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	return token, err
}
