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
	Object map[string]interface{} `json:"object"`
}

const (
	EventType_AccountUpdated = "account.updated"
	EventType_InvoiceCreated = "invoice.created"
)
