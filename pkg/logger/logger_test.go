package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	err := Init(-1, "2006-01-02T15:04:05.999999999Z07:00")
	assert.Nil(t, err)
}
