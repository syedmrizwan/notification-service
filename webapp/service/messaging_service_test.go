package service

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"notification_service_webapp/model"
	"notification_service_webapp/model/mocks"
	"testing"
	"time"
)

func TestWrite(t *testing.T) {
	t.Run("Success", func(t *testing.T) {

		var now time.Time
		notificationBody := model.Notification{
			ID:                    1,
			Priority:              "High",
			UserId:                1,
			NotificationText:      &model.NotificationText{ID: 1, Message: "Hello"},
			NotificationTextID:    1,
			NotificationHandler:   &model.NotificationHandler{ID: 1, Name: "SMS", RatePerMinute: 1},
			NotificationHandlerID: 1,
			CreatedAt:             now}
		message, _ := json.Marshal(notificationBody)
		mockMessagingRepository := new(mocks.MockMessagingRepository)
		ms := NewMessagingService(&MSConfig{
			MessagingRepository: mockMessagingRepository,
		})

		// We can use Run method to modify the user when the Create method is called.
		//  We can then chain on a Return method to return no error
		mockMessagingRepository.
			On("Publish", notificationBody.Priority, message).
			Return(nil)

		err := ms.Write(notificationBody.Priority, message)

		assert.NoError(t, err)

		mockMessagingRepository.AssertExpectations(t)
	})

}
