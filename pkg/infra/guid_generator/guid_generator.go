package guidGenerator

import (
	"github.com/google/uuid"

	"jwemanager/pkg/app/interfaces"
)

type guidGenerator struct{}

func (guidGenerator) V4() string {
	return uuid.New().String()
}

func NewGuidGenerator() interfaces.IGuidGenerator {
	return guidGenerator{}
}
