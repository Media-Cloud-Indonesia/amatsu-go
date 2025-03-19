package amatsugo

type Mail struct {
}

func NewMail() *Mail {
	return &Mail{}
}

func (m *Mail) PrepareResellerInvoicePaid(r RawPayload, data []MailNotification[ResellerInvoicePaidData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerInvoiceReminder(r RawPayload, data []MailNotification[ResellerInvoiceReminderData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerPasswordInformation(r RawPayload, data []MailNotification[ResellerPasswordInformationData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerDomainExpiration(r RawPayload, data []MailNotification[ResellerDomainExpirationData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerDomainRegistration(r RawPayload, data []MailNotification[ResellerDomainRegistrationData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerDomainRenewal(r RawPayload, data []MailNotification[ResellerDomainRenewalData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerDomainTransferRequest(r RawPayload, data []MailNotification[ResellerDomainTransferRequestData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerDomainTransferSuccess(r RawPayload, data []MailNotification[ResellerDomainTransferSuccessData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerMonthlyReportSummary(r RawPayload, data []MailNotification[ResellerMonthlyReportSummaryData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerContactCreate(r RawPayload, data []MailNotification[ResellerContactCreateData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerGeneralReport(r RawPayload, data []MailNotification[ResellerGeneralReportData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerMultiDomains(r RawPayload, data []MailNotification[ResellerMultiDomainsData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerForgotPassword(r RawPayload, data []MailNotification[ResellerForgotPasswordData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerEPPPassword(r RawPayload, data []MailNotification[ResellerEPPPasswordData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareResellerRegistrationUpdateStatus(r RawPayload, data []MailNotification[ResellerRegistrationUpdateStatusData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppResellerDashboard, r, data)
}

func (m *Mail) PrepareMainTransferOutgoingRequest(r RawPayload, data []MailNotification[MainTransferOutgoingRequestData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppMainApp, r, data)
}

func (m *Mail) PrepareMainTransferOutgoingCancel(r RawPayload, data []MailNotification[MainTransferOutgoingCancelData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppMainApp, r, data)
}

func (m *Mail) PrepareMainTransferOutgoingReject(r RawPayload, data []MailNotification[MainTransferOutgoingRejectData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppMainApp, r, data)
}

func (m *Mail) PrepareMainTransferOutgoingSuccess(r RawPayload, data []MailNotification[MainTransferOutgoingSuccessData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppMainApp, r, data)
}

func (m *Mail) PrepareMainCPanelInformation(r RawPayload, data []MailNotification[MainCPanelInformationData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppMainApp, r, data)
}

func (m *Mail) PrepareMainBillingCycle(r RawPayload, data []MailNotification[MainBillingCycleData]) (string, error) {
	return PrepareNotifyMail(RegisteredAppMainApp, r, data)
}

func PrepareNotifyMail[T TemplatesData](appID RegisteredApp, r RawPayload, data []MailNotification[T]) (string, error) {
	payload := MailNotificationPayload[T]{
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