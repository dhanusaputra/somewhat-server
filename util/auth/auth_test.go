package auth

import (
	"testing"

	pb "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
)

func TestSignJWT(t *testing.T) {
	user := &pb.User{
		Id:       "1",
		Username: "username",
		Password: "password",
	}
	tokenString, err := SignJWT(user)
	if err != nil {
		t.Errorf("error during signing JWT")
	}
	if tokenString == "" {
		t.Errorf("invalid JWT")
	}
}
