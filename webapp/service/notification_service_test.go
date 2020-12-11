package service

import (
	"context"
	"fmt"
	"notification_service_webapp/model"
	"notification_service_webapp/model/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		var now time.Time
		mockNotificationResp := []model.Notification{{ID: 1, Priority: "High", UserId: 1,
			NotificationText: &model.NotificationText{ID: 1, Message: "Hello"}, NotificationTextID: 1,
			NotificationHandler: &model.NotificationHandler{ID: 1, Name: "SMS", RatePerMinute: 1}, NotificationHandlerID: 1, CreatedAt: now}}

		mockNotificationRepository := new(mocks.MockNotificationRepository)
		ns := NewNotificationService(&NSConfig{
			NotificationRepository: mockNotificationRepository,
		})
		mockNotificationRepository.On("GetAll", mock.Anything).Return(mockNotificationResp, nil)

		ctx := context.TODO()
		u, err := ns.GetAll(ctx)

		assert.NoError(t, err)
		assert.Equal(t, u, mockNotificationResp)
		mockNotificationRepository.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		mockNotificationRepository := new(mocks.MockNotificationRepository)
		ns := NewNotificationService(&NSConfig{
			NotificationRepository: mockNotificationRepository,
		})

		mockNotificationRepository.On("GetAll", mock.Anything).Return(nil, fmt.Errorf("Some error down the call chain"))

		ctx := context.TODO()
		u, err := ns.GetAll(ctx)

		assert.Nil(t, u)
		assert.Error(t, err)
		mockNotificationRepository.AssertExpectations(t)
	})

}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		notificationPostBody := &model.NotificationPostBody{NotificationText: "Temp", UserId: 1, Priority: "Low", NotificationMode: "SMS"}
		var now time.Time
		mockNotification := &model.Notification{ID: 1, Priority: "High", UserId: 1,
			NotificationText: &model.NotificationText{ID: 1, Message: "Hello"}, NotificationTextID: 1,
			NotificationHandler: &model.NotificationHandler{ID: 1, Name: "SMS", RatePerMinute: 1}, NotificationHandlerID: 1, CreatedAt: now}

		mockNotificationRepository := new(mocks.MockNotificationRepository)
		ns := NewNotificationService(&NSConfig{
			NotificationRepository: mockNotificationRepository,
		})

		// We can use Run method to modify the user when the Create method is called.
		//  We can then chain on a Return method to return no error
		mockNotificationRepository.
			On("Insert", mock.Anything, notificationPostBody).
			Return(mockNotification, nil)

		ctx := context.TODO()
		u, err := ns.Create(ctx, notificationPostBody)

		assert.NoError(t, err)

		assert.Equal(t, u, mockNotification)

		mockNotificationRepository.AssertExpectations(t)
	})

}

func TestBulkCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		bulkNotificationBody := &model.BulkNotificationBody{NotificationText: "Temp", UserIds: []int{1}, NotificationMode: "SMS"}
		var now time.Time
		mockNotificationResp := []*model.Notification{&model.Notification{
			ID:                    1,
			Priority:              "High",
			UserId:                1,
			NotificationText:      &model.NotificationText{ID: 1, Message: "Hello"},
			NotificationTextID:    1,
			NotificationHandler:   &model.NotificationHandler{ID: 1, Name: "SMS", RatePerMinute: 1},
			NotificationHandlerID: 1,
			CreatedAt:             now}}

		mockNotificationRepository := new(mocks.MockNotificationRepository)
		ns := NewNotificationService(&NSConfig{
			NotificationRepository: mockNotificationRepository,
		})

		// We can use Run method to modify the user when the Create method is called.
		//  We can then chain on a Return method to return no error
		mockNotificationRepository.
			On("BulkInsert", mock.Anything, bulkNotificationBody).
			Return(mockNotificationResp, nil)

		ctx := context.TODO()
		u, err := ns.BulkCreate(ctx, bulkNotificationBody)

		assert.NoError(t, err)

		assert.Equal(t, u, mockNotificationResp)

		mockNotificationRepository.AssertExpectations(t)
	})

}
