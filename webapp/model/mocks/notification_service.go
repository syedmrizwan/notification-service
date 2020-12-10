package mocks

import (
	"context"
	"notification_service_webapp/model"

	"github.com/stretchr/testify/mock"
)

// MockNotificationService is a mock type for model.UserService
type MockNotificationService struct {
	mock.Mock
}

// Get is mock of NotificationService Get
func (m *MockNotificationService) GetAll(ctx context.Context) ([]model.Notification, error) {
	// args that will be passed to "Return" in the tests, when function
	// is called Hence the name "ret"
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

// Create is a mock of NotificationService.Create
func (m *MockNotificationService) Create(ctx context.Context, n *model.NotificationPostBody) (*model.Notification, error) {
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

// BulkCreate is a mock of NotificationService.BulkCreate
func (m *MockNotificationService) BulkCreate(ctx context.Context, n *model.BulkNotificationBody) ([]*model.Notification, error) {
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
