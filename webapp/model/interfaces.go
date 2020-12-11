package model

import "context"

// NotificationService defines methods the handler layer expects
// any service it interacts with to implement
type NotificationService interface {
	GetAll(ctx context.Context) ([]Notification, error)
	Create(ctx context.Context, n *NotificationPostBody) (*Notification, error)
	BulkCreate(ctx context.Context, n *BulkNotificationBody) ([]*Notification, error)
}

// NotificationRepository defines methods the service layer expects
// any repository it interacts with to implement
type NotificationRepository interface {
	GetAll(ctx context.Context) ([]Notification, error)
	Insert(ctx context.Context, n *NotificationPostBody) (*Notification, error)
	BulkInsert(ctx context.Context, n *BulkNotificationBody) ([]*Notification, error)
}

// MessagingService defines methods the handler layer expects
// any service it interacts with to implement
type MessagingService interface {
	Write(messagePriority string, m []byte) error
}

// Messaging defines methods the service layer expects
// any repository it interacts with to implement
type MessagingRepository interface {
	Publish(messagePriority string, m []byte) error
}
