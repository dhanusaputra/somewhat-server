package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestore(t *testing.T) {
	testFunc := func() bool {
		return true
	}
	func() {
		defer NewPtrs([]interface{}{&testFunc}).Restore()
		testFunc = func() bool {
			return false
		}
	}()
	assert.Equal(t, true, testFunc())
}
