package guidGenerator

import (
	"github.com/google/uuid"

	"jwemanager/pkg/app/interfaces"
)

type guidGenerator struct{}

var NewUUID = uuid.New

func (guidGenerator) V4() string {
	return NewUUID().String()
}

func NewGuidGenerator() interfaces.IGuidGenerator {
	return guidGenerator{}
}
