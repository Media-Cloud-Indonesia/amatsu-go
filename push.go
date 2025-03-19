package amatsugo

type Push struct {
}

func NewPush() *Push {
	return &Push{}
}

func (p *Push) PrepareNotifyPush(appID RegisteredApp, r RawPayload, data []PushNotification) (string, error) {
	payload := PushNotificationPayload{
		Data: data,
		BasePayload: BasePayload{
			Event:    EventNotify,
			Type:     NotificationTypeMail,
			Category: r.Category,
			AppID:    appID,
			Group:    r.Group,
			Source:   r.Source,
			Campaign: r.Campaign,
		},
	}

	return ToStringJSON(payload)
}