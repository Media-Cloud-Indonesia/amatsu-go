package amatsugo

import "encoding/json"

// Registered App in Amatsu System
type RegisteredApp int

const (
	RegisteredAppMainApp = RegisteredApp(iota)
	RegisteredAppResellerDashboard
)

func (ra RegisteredApp) String() string {
	mapping := map[RegisteredApp]string{
		RegisteredAppMainApp:           "MAIN_APP",
		RegisteredAppResellerDashboard: "RESELLER_DASHBOARD",
	}

	return mapping[ra]
}

func (ra RegisteredApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(ra.String())
}

// Available Notification Category in Amatsu System
type NotificationCategory int

const (
	NotificationCategoryInfo = NotificationCategory(iota)
	NotificationCategoryPromotion
	NotificationCategoryOther
)

func (nc NotificationCategory) String() string {
	mapping := map[NotificationCategory]string{
		NotificationCategoryInfo:      "info",
		NotificationCategoryPromotion: "promotion",
		NotificationCategoryOther:     "other",
	}

	return mapping[nc]
}

func (nc NotificationCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(nc.String())
}

// Available Events for trigger Amatsu System
type Event int

const (
	EventNotify = Event(iota)
	EventSubscribe
)

func (e Event) String() string {
	mapping := map[Event]string{
		EventNotify: "notify",
		EventSubscribe: "subscribe",
	}

	return mapping[e]
}

func (e Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// Source Type Notification
type SourceType int

const (
	SourceTypeFrontend = SourceType(iota)
	SourceTypeBackend
)

func (st SourceType) String() string {
	mapping := map[SourceType]string{
		SourceTypeBackend: "backend",
		SourceTypeFrontend: "frontend",
	}

	return mapping[st]
}

func (st SourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(st.String())
}

// Available Template in Amatsu System
type ResellerDashboardTemplate int

const (
	ResellerDashboardTemplateInvoicePaid = ResellerDashboardTemplate(iota)
	ResellerDashboardTemplateInvoiceReminder
	ResellerDashboardTemplatePasswordInformation
	ResellerDashboardTemplateDomainExpiration
	ResellerDashboardTemplateDomainRegistration
	ResellerDashboardTemplateDomainRenewal
	ResellerDashboardTemplateDomainTransferRequest
	ResellerDashboardTemplateDomainTransferSuccess
	ResellerDashboardTemplateMonthlyReportSummary
	ResellerDashboardTemplateContactCreate
	ResellerDashboardTemplateGeneralReport
	ResellerDashboardTemplateMultiDomains
	ResellerDashboardTemplateResetPassword
	ResellerDashboardTemplateEppPasswordInformation
	ResellerDashboardTemplateResellerUpdateRegistration
)

func (rdt ResellerDashboardTemplate) String() string {
	mapping := map[ResellerDashboardTemplate]string{
		ResellerDashboardTemplateInvoicePaid: "invoice-paid",
		ResellerDashboardTemplateInvoiceReminder: "invoice-reminder",
		ResellerDashboardTemplatePasswordInformation: "password-information",
		ResellerDashboardTemplateDomainExpiration: "domain-expiration",
		ResellerDashboardTemplateDomainRegistration: "domain-registration",
		ResellerDashboardTemplateDomainRenewal: "domain-renewal",
		ResellerDashboardTemplateDomainTransferRequest: "domain-transfer-request",
		ResellerDashboardTemplateDomainTransferSuccess: "domain-transfer-success",
		ResellerDashboardTemplateMonthlyReportSummary: "monthly-report-summary",
		ResellerDashboardTemplateContactCreate: "contact-create",
		ResellerDashboardTemplateGeneralReport: "general-report",
		ResellerDashboardTemplateMultiDomains: "multi-domains",
		ResellerDashboardTemplateResetPassword: "reset-password",
		ResellerDashboardTemplateEppPasswordInformation: "epp-password-information",
		ResellerDashboardTemplateResellerUpdateRegistration: "reseller-update-registration",
	}

	return mapping[rdt]
}

func (rdt ResellerDashboardTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(rdt.String())
}

// 
type MainAppTemplate int

const (
	MainAppTemplateBillingCycle = MainAppTemplate(iota)
	MainAppTemplateCPanelInformation
	MainAppTemplateTransferOutgoingRequest
	MainAppTemplateTransferOutgoingSuccess
	MainAppTemplateTransferOutgoingCancel
	MainAppTemplateTransferOutgoingReject
)

func (mat MainAppTemplate) String() string {
	mapping := map[MainAppTemplate]string{
		MainAppTemplateBillingCycle: "billing-cycle",
		MainAppTemplateCPanelInformation: "cpanel-information",
		MainAppTemplateTransferOutgoingRequest: "transfer-outgoing-request",
		MainAppTemplateTransferOutgoingSuccess: "transfer-outgoing-success",
		MainAppTemplateTransferOutgoingCancel: "transfer-outgoing-cancel",
		MainAppTemplateTransferOutgoingReject: "transfer-outgoing-reject",
	}

	return mapping[mat]
}

func (mat MainAppTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(mat.String())
}

// Notification Type
type NotificationType int

const (
	NotificationTypeMail = NotificationType(iota)
	NotificationTypePush
)

func (nt NotificationType) String() string {
	mapping := map[NotificationType]string{
		NotificationTypeMail: "mail",
		NotificationTypePush: "push-notification",
	}

	return mapping[nt]
}

func (nt NotificationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(nt.String())
}