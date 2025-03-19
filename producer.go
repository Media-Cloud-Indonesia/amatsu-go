package amatsugo

type NotificationProducer struct {
	Mail *Mail
	Push *Push
}

func NewNotificationProducer() *NotificationProducer {
	npc := &NotificationProducer{}

	npc.Mail = NewMail()
	npc.Push = NewPush()

	return npc
}