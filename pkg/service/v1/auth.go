package v1

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	v1 "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
)

var (
	key = []byte(os.Getenv("KEY"))
)

// Auth ...
type Auth interface {
	// SignJWT ...
	SignJWT(user *v1.User) (string, error)
}

type auth struct {
	appName             string
	expiredTimeInMinute time.Duration
}

// NewAuth ...
func NewAuth(appName string, expiredTimeInMinute time.Duration) Auth {
	return &auth{
		appName:             appName,
		expiredTimeInMinute: expiredTimeInMinute,
	}
}

// SignJWT ...
func (a *auth) SignJWT(user *v1.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.Id,
		"created_at": user.CreatedAt,
		"exp":        time.Now().Add(time.Second * a.expiredTimeInMinute).Unix(),
		"iss":        a.appName,
	})

	return token.SignedString([]byte(key))
}
