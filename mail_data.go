package amatsugo

type BaseData struct {
	LogoURL string `json:"logoUrl,omitempty"`
}

type Attachment struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
	Ext  string `json:"ext,omitempty"`
}

type ResellerInvoicePaidData struct {
	BaseData
	Receiver    string `json:"receiver"`
	InvoiceID   string `json:"invoiceId"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	PaidAt      string `json:"paidAt"`
	TotalAmount string `json:"totalAmount"`
	InvoiceUrl  string `json:"invoiceUrl"`
}

type ResellerInvoiceReminderData struct {
	BaseData
	Receiver    string `json:"receiver"`
	InvoiceID   string `json:"invoiceId"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	ExpiredAt   string `json:"expiredAt"`
	TotalAmount string `json:"totalAmount"`
	InvoiceUrl  string `json:"invoiceUrl"`
	PaymentLink string `json:"paymentLink"`
}

type ResellerPasswordInformationData struct {
	BaseData
	ProviderName string `json:"providerName"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ResellerID   string `json:"resellerId"`
	LoginUrl     string `json:"loginUrl"`
}

type ResellerDomainExpirationData struct {
	BaseData
	Receiver        string `json:"receiver"`
	DomainName      string `json:"domainName"`
	ExpirationDate  string `json:"expirationDate"`
	RemainingTime   string `json:"remainingTime"`
	DomainDetailUrl string `json:"domainDetailUrl"`
	RenewDomainUrl  string `json:"renewDomainUrl"`
}

type ResellerDomainRegistrationData struct {
	BaseData
	Receiver         string `json:"receiver"`
	DomainName       string `json:"domainName"`
	RegistrationDate string `json:"registrationDate"`
	ExpirationDate   string `json:"expirationDate"`
	DomainDetailsUrl string `json:"domainDetailsUrl"`
}

type ResellerDomainRenewalData struct {
	BaseData
	Receiver        string `json:"receiver"`
	DomainName      string `json:"domainName"`
	ExpirationDate  string `json:"expirationDate"`
	RenewalDate     string `json:"renewalDate"`
	DomainDetailUrl string `json:"domainDetailsUrl"`
}

type ResellerDomainTransferRequestData struct {
	BaseData
	Receiver               string `json:"receiver"`
	DomainName             string `json:"domainName"`
	TransferSubmissionDate string `json:"transferSubmissionDate"`
	DomainsUrl             string `json:"domainsUrl"`
	OldRegistrar           string `json:"oldRegistar"`
}

type ResellerDomainTransferSuccessData struct {
	BaseData
	Receiver               string `json:"receiver"`
	DomainName             string `json:"domainName"`
	TransferCompletionDate string `json:"transferCompletionDate"`
	DomainsUrl             string `json:"DomainsUrl"`
	OldRegistrar           string `json:"oldRegistrar"`
	NewRegistrar           string `json:"newRegistrar"`
}

type ResellerMonthlyReportSummaryData struct {
	BaseData
	Receiver         string `json:"receiver"`
	BalanceUsed      string `json:"balanceUsed"`
	RemainingBalance string `json:"remainingBalance"`
	DomainSold       string `json:"domainSold"`
	Period           string `json:"period"`
	DomainSalesUrl   string `json:"domainSalesUrl"`
	TopupUrl         string `json:"topupUrl"`
}

type ResellerContactCreateData struct {
	BaseData
	Receiver          string `json:"receiver"`
	ContactName       string `json:"contactName"`
	ContactEmail      string `json:"contactEmail"`
	CreatedAt         string `json:"createdAt"`
	ContactDetailsUrl string `json:"contactDetailsUrl"`
	DomainsUrl        string `json:"domainsUrl"`
}

type MainBillingCycleData struct {
	BaseData
	Receiver     string `json:"receiver"`
	Title        string `json:"title"`
	Descriptions string `json:"descriptions"`
	ProductName  string `json:"productName"`
	DomainName   string `json:"domainName"`
	ExpiredDate  string `json:"expiredDate"`
	Threats      string `json:"threats"`
	RenewLink    string `json:"renewLink"`
}

type ResellerGeneralReportData struct {
	BaseData
	Receiver    string `json:"receiver"`
	ExportKind  string `json:"exportKind"`
	RangePeriod string `json:"rangePeriod"`
}

type ResellerMultiDomainsData struct {
	BaseData
	Receiver string `json:"receiver"`
	Summary  string `json:"summary"`
}

type ResellerForgotPasswordData struct {
	BaseData
	Receiver           string `json:"receiver"`
	ExpireIn           string `json:"expireIn"`
	ChangePasswordLink string `json:"changePasswordLink"`
}

type ResellerEPPPasswordData struct {
	BaseData
	ProviderName string `json:"providerName"`
	ResellerID   string `json:"resellerId"`
	ClientID     string `json:"clientId"`
	EPPPassword  string `json:"eppPassword"`
}

type ResellerRegistrationUpdateStatusData struct {
	BaseData
	LoginUrl   string `json:"loginUrl,omitempty"`
	Status     string `json:"status"`
	Reason     string `json:"reason,omitempty"`
	ResellerID string `json:"reselerId"`
	Email      string `json:"email"`
}

type MainTransferOutgoingRequestData struct {
	ApproveLink string `json:"approveLink"`
	Receiver    string `json:"receiver"`
	DomainName  string `json:"domainName"`
}

type MainTransferOutgoingCancelData struct {
	DomainName string `json:"domainName"`
	Receiver   string `json:"receiver"`
}

type MainTransferOutgoingRejectData struct {
	DomainName string `json:"domainName"`
	Receiver   string `json:"receiver"`
}

type MainTransferOutgoingSuccessData struct {
	DomainName string `json:"domainName"`
	Receiver   string `json:"receiver"`
}

type MainCPanelInformationData struct {
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// List of Mail Templates
type TemplatesData interface {
	ResellerInvoicePaidData | ResellerInvoiceReminderData | ResellerPasswordInformationData |
		ResellerDomainExpirationData | ResellerDomainRegistrationData | ResellerDomainRenewalData |
		ResellerDomainTransferRequestData | ResellerDomainTransferSuccessData | ResellerMonthlyReportSummaryData |
		ResellerContactCreateData | MainBillingCycleData | ResellerGeneralReportData | ResellerMultiDomainsData |
		ResellerForgotPasswordData | ResellerEPPPasswordData | ResellerRegistrationUpdateStatusData | MainTransferOutgoingCancelData |
		MainTransferOutgoingRejectData | MainTransferOutgoingRequestData | MainTransferOutgoingSuccessData | MainCPanelInformationData
}
