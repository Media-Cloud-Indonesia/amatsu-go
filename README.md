# Amatsu's Go Structs for Kafka Message Payloads

This package provides Go structs for generating JSON payloads used as messages in Kafka topics.

- Package version: 1.0.0

# Getting Started
## Installation
Install Amastu Go with:
```shell
go get github.com/Media-Cloud-Indonesia/amatsu-go
```

## Examples
Example with mail request
```go
package main

import (
  "fmt"
  "github.com/Media-Cloud-Indonesia/amatsu-go"
)

func main() {
  notificationProducer := amatsugo.NewNotificationProducer()

  // Create mail request
  result, err := notificationProducer.Mail.PrepareResellerInvoicePaid(amatsugo.RawPayload{},[]amatsugo.MailNotification[amatsugo.ResellerInvoicePaidData]{
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
    panic(err.Error())
  }

  fmt.Printf("JSON string result: %s", result)
  // result would be `{"event":"notify","type":"mail","appId":"RESELLER_DASHBOARD","data":[{"subject":"Your Invoice Has Been Paid","to":["user@mediacloud.id"],"payload":{"receiver":"User","invoiceId":"INV000001","status":"PAID","createdAt":"19 March 2025","paidAt":"19 March 2025","totalAmount":"Rp20.000,00","invoiceUrl":"https://mediacloud.id/client-area/invoice/INV000001"}}]}`
}
```
Example with push
```go
package main

import (
  "fmt"
  "time"
  "github.com/Media-Cloud-Indonesia/amatsu-go"
)

func main() {
  notificationProducer := amatsugo.NewNotificationProducer()

  nextThreeHours := time.Now().Add(3 * time.Hour) // You can set when notifications will be sent.
  
  // Create push request
  result, err := notificationProducer.Push.PrepareNotifyPush(amatsugo.RegisteredAppMainApp, amatsugo.RawPayload{}, []amatsugo.PushNotification{
    {
      ScheduleAt: &nextThreeHours,
      UserID: "ID000001",
      Title: "Your Domain Will Be Expired",
      Body: "One of your domains will expire soon. Please check as soon as possible.",
      Url: "https://mediacloud.id/client-area/domains/DM000001",
    },
  })
  
  if err != nil {
    panic(err.Error())
  }

  fmt.Printf("JSON string result: %s", result)
}
```