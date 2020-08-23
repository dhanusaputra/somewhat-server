package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	v1 "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
)

var (
	key                        = []byte(os.Getenv("KEY"))
	defaultExpiredTimeInMinute = 30 * time.Second
	defaultAppName             = "something"
)

// SignJWT ...
func SignJWT(user *v1.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.Id,
		"created_at": user.CreatedAt,
		"exp":        time.Now().Add(time.Second * defaultExpiredTimeInMinute).Unix(),
		"iss":        defaultAppName,
	})

	return token.SignedString([]byte(key))
}

// ValidateJWT ...
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	return token, err
}
