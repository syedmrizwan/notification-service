package repository

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"notification_service_webapp/env"
	"notification_service_webapp/model"
)

// natsMessagingRepository is messaging/repository implementation
// of service layer TokenRepository
type natsMessagingRepository struct {
	Conn stan.Conn
}

func NewMessagingRepository(conn stan.Conn) model.MessagingRepository {
	return &natsMessagingRepository{
		Conn: conn,
	}
}

func (r *natsMessagingRepository) Publish(messagePriority string, m []byte) error {
	err := r.Conn.Publish(fmt.Sprintf("%s.%s", env.Env.RateLimiterChannel, messagePriority), m)
	if err != nil {
		return err
	}
	return nil
}
