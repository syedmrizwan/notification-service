package mocks

import "github.com/stretchr/testify/mock"

// MockMessagingService is a mock type for model.MessagingService
type MockMessagingService struct {
	mock.Mock
}

// Write is mock of MockMessagingService Write
func (m *MockMessagingService) Write(messagePriority string, message []byte) error {
	ret := m.Called(messagePriority, message)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
