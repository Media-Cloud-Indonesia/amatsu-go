package test

import (
	"testing"
	"time"

	amatsugo "github.com/Media-Cloud-Indonesia/amatsu-go"
)

func TestMailPayload(t *testing.T) {
	nProducer := amatsugo.NewNotificationProducer()

	t.Run("Invoice Paid", func(t *testing.T) {
		result, err := nProducer.Mail.PrepareResellerInvoicePaid(amatsugo.RawPayload{}, []amatsugo.MailNotification[amatsugo.ResellerInvoicePaidData]{
			{
				Subject: "Your Invoice Has Been Paid",
				To: []string{"user@mediacloud.id"},
				Payload: amatsugo.ResellerInvoicePaidData{
					Receiver: "User",
					InvoiceID: "INV000001",
					Status: "PAID",
					CreatedAt: "19 March 2025",
					PaidAt: "19 March 2025",
					TotalAmount: "Rp20.000,00",
					InvoiceUrl: "https://mediacloud.id/client-area/invoice/INV000001",
				},
			},
		})

		if err != nil {
			t.Fatalf("failed to get invoice paid json string. reason: %s", err.Error())
		}

		mustBeResult := `{"event":"notify","type":"mail","appId":"RESELLER_DASHBOARD","data":[{"subject":"Your Invoice Has Been Paid","to":["user@mediacloud.id"],"payload":{"receiver":"User","invoiceId":"INV000001","status":"PAID","createdAt":"19 March 2025","paidAt":"19 March 2025","totalAmount":"Rp20.000,00","invoiceUrl":"https://mediacloud.id/client-area/invoice/INV000001"}}]}`

		if result != mustBeResult {
			t.Fatalf("json string for invoice paid not matched. result: %s", result)
		}

		nextThreeHours := time.Now().Add(3 * time.Hour)
		
		nProducer.Push.PrepareNotifyPush(amatsugo.RegisteredAppMainApp, amatsugo.RawPayload{}, []amatsugo.PushNotification{
			{
				ScheduleAt: &nextThreeHours,
				UserID: "ID000001",
				Title: "Your Domain Will Be Expired",
				Body: "One of your domain will be expired soon. Please check as soon as posible.",
				Url: "https://mediacloud.id/client-area/domains/DM000001",
			},
		})
	})
}