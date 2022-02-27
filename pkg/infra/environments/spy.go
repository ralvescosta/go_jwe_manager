package environments

import "github.com/stretchr/testify/mock"

type EnvironmentsSpy struct {
	mock.Mock
}

func (pst EnvironmentsSpy) Configure() error {
	args := pst.Called()

	return args.Error(0)
}

func NewEnvironmentsSpy() *EnvironmentsSpy {
	return new(EnvironmentsSpy)
}
