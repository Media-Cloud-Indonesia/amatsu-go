package amatsugo

import (
	"encoding/json"
	"time"
)

type BasePayload struct {
	Event    Event                `json:"event"`
	Type     NotificationType     `json:"type"`
	AppID    RegisteredApp        `json:"appId"`
	Category NotificationCategory `json:"category,omitempty"`
	Group    string               `json:"group,omitempty"`
	Source   SourceType           `json:"source,omitempty"`
	Campaign string               `json:"campaign,omitempty"`
}

type RawPayload struct {
	Category  NotificationCategory
	Source    SourceType
	BatchName string
	Campaign  string
	Group     string
}

type MailNotificationPayload[T TemplatesData] struct {
	BasePayload
	Data []MailNotification[T] `json:"data"`
}

type PushNotificationPayload struct {
	BasePayload
	Data []PushNotification `json:"data"`
}

type MailNotification[T TemplatesData] struct {
	ScheduleAt  *time.Time   `json:"scheduleAt,omitempty"`
	Subject     string       `json:"subject,omitempty"`
	To          []string     `json:"to"`
	UserID      string       `json:"userId,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	Payload     T            `json:"payload"`
}

type PushNotification struct {
	ScheduleAt *time.Time `json:"scheduleAt,omitempty"`
	UserID     string     `json:"userId,omitempty"`
	Title      string     `json:"title"`
	Body       string     `json:"body"`
	Url        string     `json:"url"`
	Icon       string     `json:"icon,omitempty"`
	ImageUrl   string     `json:"imageUrl,omitempty"`
}

func ToStringJSON(payload any) (string, error) {
	byteJson, err := json.Marshal(payload)

	if err != nil {
		return "", nil
	}

	return string(byteJson), nil
}
