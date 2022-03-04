package guidGenerator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_V4(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewGuidGeneratorSpy()

		sut.On("V4").Return("some_guid")

		result := sut.V4()

		assert.Equal(t, "some_guid", result)
	})
}
