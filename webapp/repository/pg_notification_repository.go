package repository

import (
	"context"
	"errors"
	"github.com/go-pg/pg/v9"
	"notification_service_webapp/model"
)

// PGUserRepository is data/repository implementation
// of service layer UserRepository
type pGNotificationRepository struct {
	DB *pg.DB
}

func NewNotificationRepository(db *pg.DB) model.NotificationRepository {
	return &pGNotificationRepository{
		DB: db,
	}
}

func (r *pGNotificationRepository) GetAll(ctx context.Context) ([]model.Notification, error) {
	notifications := make([]model.Notification, 0)
	if err := r.DB.ModelContext(ctx, &notifications).Select(); err != nil {
		return nil, err
	}
	return notifications, nil
}

func (r *pGNotificationRepository) Insert(ctx context.Context, notificationPostBody *model.NotificationPostBody) (*model.Notification, error) {
	notificationHandler := model.NotificationHandler{Name: notificationPostBody.NotificationMode}
	if err := r.DB.ModelContext(ctx, &notificationHandler).Where("name = ?name").Select(); err != nil {
		return nil, errors.New("notification Mode does not exists")
	}

	notificationText := model.NotificationText{Message: notificationPostBody.NotificationText}
	if _, err := r.DB.ModelContext(ctx, &notificationText).Where("message = ?message").SelectOrInsert(); err != nil {
		return nil, err
	}

	notification := model.Notification{
		Priority:              notificationPostBody.Priority,
		UserId:                notificationPostBody.UserId,
		NotificationHandlerID: notificationHandler.ID,
		NotificationHandler:   &notificationHandler,
		NotificationTextID:    notificationText.ID,
		NotificationText:      &notificationText,
	}

	if err := r.DB.Insert(&notification); err != nil {
		return nil, err
	}

	return &notification, nil
}

func (r *pGNotificationRepository) BulkInsert(ctx context.Context, bulkNotificationBody *model.BulkNotificationBody) ([]*model.Notification, error) {

	notificationHandler := model.NotificationHandler{Name: bulkNotificationBody.NotificationMode}
	if err := r.DB.ModelContext(ctx, &notificationHandler).Where("name = ?name").Select(); err != nil {
		return nil, errors.New("notification Mode does not exists")
	}

	notificationText := model.NotificationText{Message: bulkNotificationBody.NotificationText}
	if _, err := r.DB.ModelContext(ctx, &notificationText).Where("message = ?message").SelectOrInsert(); err != nil {
		return nil, err
	}

	var notifications []*model.Notification
	for _, userId := range bulkNotificationBody.UserIds {

		notification := &model.Notification{
			Priority:              "Low",
			UserId:                userId,
			NotificationHandlerID: notificationHandler.ID,
			NotificationHandler:   &notificationHandler,
			NotificationTextID:    notificationText.ID,
			NotificationText:      &notificationText,
		}
		notifications = append(notifications, notification)
	}

	//Bulk insertion
	if err := r.DB.Insert(&notifications); err != nil {
		return nil, err
	}
	return notifications, nil
}
