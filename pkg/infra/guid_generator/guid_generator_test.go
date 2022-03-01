package guidGenerator

import (
	"testing"

	"jwemanager/pkg/app/interfaces"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("should return uuid correctly", func(t *testing.T) {
		sut := makeGuidGenSut()

		NewUUID = func() uuid.UUID {
			return [16]byte{}
		}

		result := sut.guiGen.V4()

		assert.NotEmpty(t, result)
	})
}

type guidGenSutRtn struct {
	guiGen interfaces.IGuidGenerator
}

func makeGuidGenSut() guidGenSutRtn {
	guidGen := NewGuidGenerator()

	return guidGenSutRtn{guidGen}
}
