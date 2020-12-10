package model

type NotificationPostBody struct {
	NotificationText string `json:"notification_text"`
	Priority         string `json:"Priority"`
	UserId           int    `json:"user_id"`
	NotificationMode string `json:"notification_mode"`
}

type BulkNotificationBody struct {
	NotificationText string `json:"notification_text"`
	NotificationMode string `json:"notification_mode"`
	UserIds          []int  `json:"user_ids"`
}

type ResponseBody struct {
	Message string `json:"message"`
}

