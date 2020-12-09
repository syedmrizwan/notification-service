package model

import "time"

type NotificationText struct {
	ID      int    `pg:",pk" json:"id"`
	Message string `json:"message"`
}

type NotificationHandler struct {
	ID            int    `pg:",pk" json:"id"`
	Name          string `json:"name"`
	RatePerMinute int    `json:"rate_per_minute"`
}

type Notification struct {
	ID                    int64                `pg:",pk" json:"id"`
	Priority              string               `json:"Priority"`
	UserId                int                  `json:"user_id"`
	NotificationTextID    int                  `pg:",fk" json:"notification_text_id"`
	NotificationText      *NotificationText    `json:"notification_text"`
	NotificationHandlerID int                  `pg:",fk" json:"notification_handler_id"`
	NotificationHandler   *NotificationHandler `json:"notification_handler"`
	CreatedAt             time.Time            `pg:"default:now()" json:"created_at"`
}
