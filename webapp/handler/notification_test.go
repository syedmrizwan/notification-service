package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"notification_service_webapp/model"
	"notification_service_webapp/model/mocks"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetNotifications(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		var now time.Time
		mockNotificationResp := []model.Notification{{ID: 1, Priority: "High", UserId: 1,
			NotificationText: &model.NotificationText{ID: 1, Message: "Hello"}, NotificationTextID: 1,
			NotificationHandler: &model.NotificationHandler{ID: 1, Name: "SMS", RatePerMinute: 1}, NotificationHandlerID: 1, CreatedAt: now}}

		mockNotificationService := new(mocks.MockNotificationService)
		mockNotificationService.On("GetAll", mock.AnythingOfType("*context.timerCtx")).Return(mockNotificationResp, nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:                   router,
			NotificationService: mockNotificationService,
			BaseURL:             "",
		})

		request, err := http.NewRequest(http.MethodGet, "/notifications", nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"notifications": mockNotificationResp,
		})
		assert.NoError(t, err)

		fmt.Println(rr.Body.Bytes())
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockNotificationService.AssertExpectations(t) // assert that UserService.Get was called
	})
}

func TestCreateNotifications(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		notificationPostBody := &model.NotificationPostBody{
			NotificationText: "Temp",
			UserId:           1,
			Priority:         "Low",
			NotificationMode: "SMS",
		}
		var now time.Time
		mockNotification := &model.Notification{
			ID:                    1,
			Priority:              "High",
			UserId:                1,
			NotificationText:      &model.NotificationText{ID: 1, Message: "Hello"},
			NotificationTextID:    1,
			NotificationHandler:   &model.NotificationHandler{ID: 1, Name: "SMS", RatePerMinute: 1},
			NotificationHandlerID: 1,
			CreatedAt:             now,
		}

		mockNotificationService := new(mocks.MockNotificationService)
		mockNotificationService.On("Create", mock.AnythingOfType("*context.timerCtx"), notificationPostBody).
			Return(mockNotification, nil)

		message, _ := json.Marshal(mockNotification)
		mockMessagingService := new(mocks.MockMessagingService)
		mockMessagingService.
			On("Write", mockNotification.Priority, message).
			Return(nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:                   router,
			NotificationService: mockNotificationService,
			MessagingService:    mockMessagingService,
			BaseURL:             "",
		})

		reqBody, err := json.Marshal(notificationPostBody)
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/notifications", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		assert.NoError(t, err)

		assert.Equal(t, 201, rr.Code)

		mockNotificationService.AssertExpectations(t) // assert that UserService.Get was called
	})
}

func TestCreateBulkNotifications(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		bulkNotificationBody := &model.BulkNotificationBody{
			NotificationText: "Temp",
			UserIds:          []int{1},
			NotificationMode: "SMS",
		}
		var now time.Time
		mockNotification := []*model.Notification{&model.Notification{
			ID:                    1,
			Priority:              "Low",
			UserId:                1,
			NotificationText:      &model.NotificationText{ID: 1, Message: "Hello"},
			NotificationTextID:    1,
			NotificationHandler:   &model.NotificationHandler{ID: 1, Name: "SMS", RatePerMinute: 1},
			NotificationHandlerID: 1,
			CreatedAt:             now}}

		mockNotificationService := new(mocks.MockNotificationService)
		mockNotificationService.On("BulkCreate", mock.AnythingOfType("*context.timerCtx"), bulkNotificationBody).
			Return(mockNotification, nil)

		mockMessagingService := new(mocks.MockMessagingService)
		mockMessagingService.
			On("Write", "Low", mock.Anything).
			Return(nil)

		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		router := gin.Default()

		NewHandler(&Config{
			R:                   router,
			NotificationService: mockNotificationService,
			MessagingService:    mockMessagingService,
			BaseURL:             "",
		})

		reqBody, err := json.Marshal(bulkNotificationBody)
		assert.NoError(t, err)

		request, err := http.NewRequest(http.MethodPost, "/bulk-notifications", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		assert.NoError(t, err)

		assert.Equal(t, 201, rr.Code)

		mockNotificationService.AssertExpectations(t) // assert that UserService.Get was called
	})
}
