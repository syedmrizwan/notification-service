package mocks

import "github.com/stretchr/testify/mock"

// MockMessagingRepository is a mock type for model.MessagingService
type MockMessagingRepository struct {
	mock.Mock
}

// Publish is mock of MockMessagingRepository Publish
func (m *MockMessagingRepository) Publish(messagePriority string, message []byte) error {
	ret := m.Called(messagePriority, message)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}
