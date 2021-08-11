package authutil

import (
	"testing"

	pb "github.com/dhanusaputra/somewhat-server/pkg/api/v1"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestSignJWT(t *testing.T) {
	user := &pb.User{
		Id:       "1",
		Username: "username",
		Password: "password",
	}
	tokenString, err := SignJWT(user)
	assert.Nil(t, err)
	assert.NotNil(t, tokenString)
}

func TestValidateJWT(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkX2F0IjpudWxsLCJpZCI6IjEiLCJpc3MiOiJzb21ldGhpbmciLCJ1c2VybmFtZSI6InVzZXJuYW1lIn0.laPjiS5zWxCaihlGzYTI9jJ1lGuTWsTd4IJdEMgZwuc"
	want := jwt.MapClaims{
		"created_at": nil,
		"id":         "1",
		"iss":        "something",
		"username":   "username",
	}
	_, got, err := ValidateJWT(tokenString)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
