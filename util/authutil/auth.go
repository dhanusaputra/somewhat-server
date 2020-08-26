package authutil

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	v1 "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
	"github.com/dhanusaputra/somewhat-server/util/envutil"
)

var (
	key = []byte(os.Getenv("KEY"))
)

const (
	defaultExpiredTimeInMinute = 15 //mins
	defaultAppName             = "something"
)

// SignJWT ...
func SignJWT(user *v1.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.Id,
		"username":   user.Username,
		"created_at": user.CreatedAt,
		"exp":        time.Now().Add(time.Duration(envutil.GetEnvAsInt("JWT_EXPIRED_TIME_IN_MINUTE", defaultExpiredTimeInMinute)) * time.Minute).Unix(),
		"iss":        defaultAppName,
	})
	return token.SignedString(key)
}

// ValidateJWT ...
func ValidateJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	return token, claims, err
}
