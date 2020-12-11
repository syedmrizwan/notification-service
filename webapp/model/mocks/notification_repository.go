package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"notification_service_webapp/model"
)

// MockUserRepository is a mock type for model.UserRepository
type MockNotificationRepository struct {
	mock.Mock
}

// GetAll is mock of UserRepository GetAll
func (m *MockNotificationRepository) GetAll(ctx context.Context) ([]model.Notification, error) {
	ret := m.Called(ctx)

	// first value passed to "Return"
	var r0 []model.Notification
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).([]model.Notification)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// Insert is a mock of NotificationService.Insert
func (m *MockNotificationRepository) Insert(ctx context.Context, n *model.NotificationPostBody) (*model.Notification, error) {
	// args that will be passed to "Return" in the tests, when function
	// is called Hence the name "ret"
	ret := m.Called(ctx, n)

	// first value passed to "Return"
	var r0 *model.Notification
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*model.Notification)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// BulkInsert is a mock of NotificationService.BulkInsert
func (m *MockNotificationRepository) BulkInsert(ctx context.Context, n *model.BulkNotificationBody) ([]*model.Notification, error) {
	// args that will be passed to "Return" in the tests, when function
	// is called Hence the name "ret"
	ret := m.Called(ctx, n)

	// first value passed to "Return"
	var r0 []*model.Notification
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).([]*model.Notification)
	}

	var r1 error
	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
