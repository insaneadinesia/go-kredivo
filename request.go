package kredivo

type PaymentTypeRequest struct {
	ServerKey string  `json:"server_key"`
	Amount    float64 `json:"amount"`
	Item      []Item  `json:"items"`
}

type CheckoutRequest struct {
	ServerKey          string             `json:"server_key"`
	PaymentType        string             `json:"payment_type"`
	TransactionDetails TransactionDetails `json:"transaction_details"`
	Seller             *[]Seller          `json:"sellers,omitempty"`
	CustomerDetails    *CustomerDetails   `json:"customer_details,omitempty"`
	BillingAddress     *Address           `json:"billing_address,omitempty"`
	ShippingAddress    Address            `json:"shipping_address"`
	PushUri            string             `json:"push_uri"`
	UserCancelUri      string             `json:"user_cancel_uri,omitempty"`
	BackToStoreUri     string             `json:"back_to_store_uri"`
	ExpirationTime     string             `json:"expiration_time,omitempty"`
}

type NotificationRequest struct {
	Status            string               `json:"status"`
	Amount            string               `json:"amount"`
	PaymentType       string               `json:"payment_type"`
	TransactionStatus string               `json:"transaction_status"`
	OrderID           string               `json:"order_id"`
	Message           string               `json:"message"`
	TransactionTime   int64                `json:"transaction_time"`
	TransactionID     string               `json:"transaction_id"`
	SignatureKey      string               `json:"signature_key"`
	ShippingAddress   *NotificationAddress `json:"shipping_address,omitempty"`
}

type ConfirmPurchaseRequest struct {
	TransactionID string `json:"transaction_id"`
	SignatureKey  string `json:"signature_key"`
}

type TransactionStatusRequest struct {
	ServerKey string `json:"server_key"`
	OrderID   string `json:"order_id"`
}

type TransactionDetails struct {
	Amount  float64 `json:"amount"`
	OrderID string  `json:"order_id"`
	Item    []Item  `json:"items"`
}

type CustomerDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type Item struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	URL        string  `json:"url"`
	ImageURL   string  `json:"image_url,omitempty"`
	Type       string  `json:"type"`
	Quantity   int32   `json:"quantity"`
	ParentType string  `json:"parent_type,omitempty"`
	ParentID   string  `json:"parent_id,omitempty"`
}

type Seller struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	LegalID string  `json:"legal_id,omitempty"`
	Email   string  `json:"email"`
	URL     string  `json:"url,omitempty"`
	Address Address `json:"address"`
}

type Address struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	Address     string `json:"address"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
}

type NotificationAddress struct {
	City            string `json:"city,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
	Countrycode     string `json:"countrycode,omitempty"`
	CreationDate    string `json:"creation_date,omitempty"`
	Phone           string `json:"phone,omitempty"`
	State           string `json:"state,omitempty"`
	Transaction     int64  `json:"transaction,omitempty"`
	Postcode        string `json:"postcode,omitempty"`
	LocationDetails string `json:"location_details,omitempty"`
}
