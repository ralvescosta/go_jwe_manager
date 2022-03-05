package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NotFoundError(t *testing.T) {
	t.Run("should create a notFoundError correctly", func(t *testing.T) {
		err := NewNotFoundError("some error")

		assert.Error(t, err)
		assert.IsType(t, NotFoundError{}, err)
	})
}
