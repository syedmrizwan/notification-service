package service

import (
	"context"
	"notification_service_webapp/model"
)

// userService acts as a struct for injecting an implementation of UserRepository
// for use in service methods
type notificationService struct {
	NotificationRepository model.NotificationRepository
}

// USConfig will hold repositories that will eventually be injected into this
// this service layer
type NSConfig struct {
	NotificationRepository model.NotificationRepository
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewNotificationService(c *NSConfig) model.NotificationService {
	return &notificationService{
		NotificationRepository: c.NotificationRepository,
	}
}

// Get retrieves a user based on their uuid
func (s *notificationService) GetAll(ctx context.Context) ([]model.Notification, error) {
	return s.NotificationRepository.GetAll(ctx)
}

// Get retrieves a user based on their uuid
func (s *notificationService) Create(ctx context.Context, n *model.NotificationPostBody) (*model.Notification, error) {
	return s.NotificationRepository.Insert(ctx, n)
}

// Get retrieves a user based on their uuid
func (s *notificationService) BulkCreate(ctx context.Context, n *model.BulkNotificationBody) ([]*model.Notification, error) {
	return s.NotificationRepository.BulkInsert(ctx, n)

}
