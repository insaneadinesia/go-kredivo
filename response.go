package kredivo

type PaymentTypeResponse struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Payment []Payment `json:"payments"`
}

type CheckoutResponse struct {
	Status      string                 `json:"status"`
	Message     string                 `json:"message"`
	RedirectURL string                 `json:"redirect_url"`
	Error       *CheckoutErrorResponse `json:"error,omitempty"`
}

type NotificationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ConfirmPurchaseResponse struct {
	Status            string `json:"status"`
	LegalName         string `json:"legal_name"`
	FraudStatus       string `json:"fraud_status"`
	OrderID           string `json:"order_id"`
	TransactionTime   int64  `json:"transaction_time"`
	Amount            string `json:"amount"`
	PaymentType       string `json:"payment_type"`
	TransactionStatus string `json:"transaction_status"`
	Message           string `json:"message"`
	TransactionID     string `json:"transaction_id"`
}

type TransactionStatusResponse struct {
	Status            string `json:"status"`
	LegalName         string `json:"legal_name"`
	FraudStatus       string `json:"fraud_status"`
	OrderID           string `json:"order_id"`
	TransactionTime   int64  `json:"transaction_time"`
	ExternalUserid    string `json:"external_userid"`
	Amount            string `json:"amount"`
	PaymentType       string `json:"payment_type"`
	TransactionStatus string `json:"transaction_status"`
	Message           string `json:"message"`
	TransactionID     string `json:"transaction_id"`
}

type CheckoutErrorResponse struct {
	Message string `json:"message"`
	Kind    string `json:"kind"`
}

type Payment struct {
	ID                           string  `json:"id"`
	Name                         string  `json:"name"`
	RawMonthlyInstallment        float64 `json:"raw_monthly_installment"`
	InterestRateTransitionTerm   float64 `json:"interest_rate_transition_term"`
	Amount                       float64 `json:"amount"`
	InstallmentAmount            float64 `json:"installment_amount"`
	RawAmount                    float64 `json:"raw_amount"`
	Rate                         float64 `json:"rate"`
	DownPayment                  float64 `json:"down_payment"`
	MonthlyInstallment           float64 `json:"monthly_installment"`
	DiscountedMonthlyInstallment float64 `json:"discounted_monthly_installment"`
	Tenure                       int16   `json:"tenure"`
}
