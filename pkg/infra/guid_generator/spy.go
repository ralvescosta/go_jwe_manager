package guidGenerator

import (
	"github.com/stretchr/testify/mock"
)

type GuidGeneratorSpy struct {
	mock.Mock
}

func (pst GuidGeneratorSpy) V4() string {
	args := pst.Called()

	return args.String(0)
}

func NewGuidGeneratorSpy() *GuidGeneratorSpy {
	return new(GuidGeneratorSpy)
}
