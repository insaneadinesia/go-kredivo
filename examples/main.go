package main

import (
	"fmt"

	kredivo "github.com/insaneadinesia/go-kredivo"
)

var kredivoClient kredivo.Client
var coreGateway kredivo.CoreGateway

func setup() {
	kredivoClient = kredivo.NewClient()
	kredivoClient.ServerKey = "8tLHIx8V0N6KtnSpS9Nbd6zROFFJH7"
	kredivoClient.PushUri = "https://misteraladin.com/push/notification"
	kredivoClient.BackToStoreUri = "https://misteraladin.com"
	kredivoClient.APIEnvType = kredivo.Sandbox

	coreGateway = kredivo.CoreGateway{
		Client: kredivoClient,
	}
}

func main() {
	itemDetails := []kredivo.Item{
		kredivo.Item{
			ID:       "1111111",
			Name:     "Universal Studios Singapore - Dewasa",
			Price:    700000,
			URL:      "https://tours.misteraladin.com/product/universal-studios-singapore/",
			Quantity: 2,
		},
		kredivo.Item{
			ID:       "discount",
			Name:     "Discount Promo Code",
			Price:    100000,
			Quantity: 1,
		},
	}

	customerDetails := &kredivo.CustomerDetails{
		FirstName: "Rachmat Adi",
		LastName:  "Prakoso",
		Email:     "rachmat.adi.p@gmail.com",
		Phone:     "089999999999",
	}

	billingAddress := &kredivo.Address{
		FirstName:   "Rachmat Adi",
		LastName:    "Prakoso",
		Address:     "Jl. Pelangi Indah",
		City:        "Jakarta",
		PostalCode:  "123456",
		Phone:       "089999999999",
		CountryCode: "IDN",
	}

	shippingAddress := kredivo.Address{
		FirstName:   "Rachmat Adi",
		LastName:    "Prakoso",
		Address:     "Jl. Pelangi Indah",
		City:        "Jakarta",
		PostalCode:  "123456",
		Phone:       "089999999999",
		CountryCode: "IDN",
	}

	transactionDetails := kredivo.TransactionDetails{
		Amount:  600000,
		OrderID: "QWERTY",
		Item:    itemDetails,
	}

	checkoutRequest := &kredivo.CheckoutRequest{
		PaymentType:        "30_days",
		TransactionDetails: transactionDetails,
		CustomerDetails:    customerDetails,
		BillingAddress:     billingAddress,
		ShippingAddress:    shippingAddress,
	}

	response, err := coreGateway.Checkout(checkoutRequest)
	if err != nil {
		fmt.Println("ERROR CHECKOUT : ", err)
		return
	}

	fmt.Println("RESPONSE CHECKOUT : ", response)
	return
}
