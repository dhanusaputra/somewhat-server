package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrigins(t *testing.T) {
	var testFunc = func() bool {
		return true
	}
	func() {
		defer Origins([]interface{}{&testFunc})
		testFunc = func() bool {
			return false
		}
	}()
	assert.Equal(t, true, testFunc())
}
