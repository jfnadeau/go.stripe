package stripe

type Event struct {
	Id              string    `json:"id"`
	Livemode        bool      `json:"livemode"`
	Created         int64     `json:"created"`
	Data            EventData `json:"data"`
	PendingWebhooks int       `json:"pending_webhooks"`
	Type            string    `json:"type"`
	Request         string    `json:"request"`
	UserId          string    `json:"user_id,omitempty"` // Stripe connect
}

type EventData struct {
	Object             map[string]interface{} `json:"object"`
	PreviousAttributes map[string]interface{} `json:"previous_attributes"`
}

const (
	EventType_AccountUpdated                 = "account.updated"
	EventType_InvoiceCreated                 = "invoice.created"
	EventType_SubscriptionDeleted            = "customer.subscription.deleted"
	EventType_CustomerDiscountCreated        = "customer.discount.created"
	EventType_CustomerDiscountUpdated        = "customer.discount.updated"
	EventType_CustomerDiscountDeleted        = "customer.discount.deleted"
	EventType_CustomerDeleted                = "customer.deleted"
	EventType_CustomerUpdated                = "customer.updated"
	EventType_SubscriptionUpdated            = "customer.subscription.updated"
	EventType_InvoiceItemCreated             = "invoiceitem.created"
	EventType_AccountApplicationDeauthorized = "account.application.deauthorized"
)
