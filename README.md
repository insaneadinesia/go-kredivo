## [Unofficial] Kredivo Library For Go(lang)

### Please read kredivo api documentation ([Kredivo Docs](https://doc.kredivo.com/))

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/insaneadinesia/go-kredivo/blob/master/LICENSE)

## Installation
```shell
go get github.com/insaneadinesia/go-kredivo
```
or with dep (I love dep :heart:)
```shell
dep ensure -add github.com/insaneadinesia/go-kredivo
```
## Configuration
First, you need to setup some configuration for kredivo such as server key, push uri, back to store uri and kredivo environment (sandbox or production). By default, this library will set your environment to sandbox.

You can find how to setup the configuration in folder `example/main.go`.
```go
    kredivoClient = kredivo.NewClient()
    kredivoClient.ServerKey = "8tLHIx8V0N6KtnSpS9Nbd6zROFFJH7"
    kredivoClient.PushUri = "https://misteraladin.com/push/notification"
    kredivoClient.BackToStoreUri = "https://misteraladin.com"
    kredivoClient.APIEnvType = kredivo.Sandbox

    coreGateway = kredivo.CoreGateway{
        Client: kredivoClient,
    }
```
## Example
In example in folder `example/main.go`, I just create an example to process checkout with kredivo. For other process like get payment type, confirm purchase, etc is similar to the example I gave.
### Checkout
```go
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
```
## License
See [LICENSE](https://github.com/insaneadinesia/go-kredivo/blob/master/LICENSE).

