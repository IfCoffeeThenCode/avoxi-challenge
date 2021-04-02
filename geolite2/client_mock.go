package geolite2

import "github.com/stretchr/testify/mock"

type ClientMock struct {
	mock.Mock
}

func (cm *ClientMock) Get(ipAddress string) (*Response, error) {
	args := cm.Called(ipAddress)
	return args.Get(0).(*Response), args.Error(1)
}
