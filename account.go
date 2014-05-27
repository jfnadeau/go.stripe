package stripe

// Customer encapsulates details about a Customer registered in Stripe.
//
// see https://stripe.com/docs/api#account
type Account struct {
	Id                  string `json:"id"`
	Email               String `json:"email,omitempty"`
	StatementDescriptor String `json:"statement_descriptor,omitempty"`
	DisplayName         String `json:"display_name,omitempty"`
	TimeZone            string `json:"timezone"`
	DetailsSubmitted    bool   `json:"details_submitted"`
	ChargeEnabled       bool   `json:"charge_enabled"`
	TransferEnabled     bool   `json:"transfer_enabled"`
}

// AccountClient encapsulates querying account using the Stripe REST API.
type AccountClient struct {
	BaseClient
}

// Retrieves a Customer linked to the api key
//
// see https://stripe.com/docs/api#rretrieve_account
func (self *AccountClient) Retrieve() (*Account, error) {
	account := Account{}
	err := self.query("GET", "/v1/account", nil, &account)
	return &account, err
}
