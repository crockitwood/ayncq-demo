package task

import (
	"encoding/json"
	"github.com/hibiken/asynq"
)

const (
	TypeEmailDelivery = "email:deliver"
)

type EmailDeliveryPayload struct {
	UserID     int
	TemplateID string
}

func NewEmailDeliveryTask(userID int, tmplID string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailDeliveryPayload{UserID: userID, TemplateID: tmplID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeEmailDelivery, payload), nil
}
