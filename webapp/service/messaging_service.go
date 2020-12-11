package service

import "notification_service_webapp/model"

// messagingService acts as a struct for injecting an implementation of MessagingRepository
// for use in service methods
type messagingService struct {
	MessagingRepository model.MessagingRepository
}

// MSConfig will hold repositories that will eventually be injected into this
// this service layer
type MSConfig struct {
	MessagingRepository model.MessagingRepository
}

// NewMessagingService is a factory function for
// initializing a MessagingService with its repository layer dependencies
func NewMessagingService(c *MSConfig) model.MessagingService {
	return &messagingService{
		MessagingRepository: c.MessagingRepository,
	}
}

func (s *messagingService) Write(messagePriority string, m []byte) error {
	return s.MessagingRepository.Publish(messagePriority, m)
}
