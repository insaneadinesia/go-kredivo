package kredivo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type CoreGateway struct {
	Client Client
}

func (gateway *CoreGateway) PaymentType(req *PaymentTypeRequest) (PaymentTypeResponse, error) {
	req.ServerKey = gateway.Client.ServerKey

	path := gateway.Client.APIEnvType.String() + "/payments"
	resp := PaymentTypeResponse{}
	jsonReq, _ := json.Marshal(req)

	err := gateway.Client.Call("POST", path, bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		fmt.Println("Error Payment Type : ", err)
		return resp, err
	}

	if resp.Status != "OK" {
		return resp, errors.New("Payment Type Response Error")
	}

	return resp, nil
}

func (gateway *CoreGateway) Checkout(req *CheckoutRequest) (CheckoutResponse, error) {
	req.ServerKey = gateway.Client.ServerKey
	req.PushUri = gateway.Client.PushUri
	req.BackToStoreUri = gateway.Client.BackToStoreUri

	path := gateway.Client.APIEnvType.String() + "/checkout_url"
	resp := CheckoutResponse{}
	jsonReq, _ := json.Marshal(req)

	err := gateway.Client.Call("POST", path, bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		fmt.Println("Error Checkout : ", err)
		return resp, err
	}

	if resp.Status != "OK" {
		return resp, errors.New(resp.Error.Message)
	}

	return resp, nil
}

func (gateway *CoreGateway) ConfirmPurchase(req *ConfirmPurchaseRequest) (ConfirmPurchaseResponse, error) {
	path := gateway.Client.APIEnvType.String() + "/update?transaction_id=" + req.TransactionID + "&signature_key=" + req.SignatureKey
	resp := ConfirmPurchaseResponse{}

	err := gateway.Client.Call("GET", path, nil, &resp)
	if err != nil {
		fmt.Println("Error Confirm Purchase : ", err)
		return resp, err
	}

	if resp.Status != "OK" {
		return resp, errors.New("Confirm Purchase Response Error")
	}

	return resp, nil
}

func (gateway *CoreGateway) CancelPurchase(req *CancelPurchaseRequest) (CancelPurchaseResponse, error) {
	req.ServerKey = gateway.Client.ServerKey

	path := gateway.Client.APIEnvType.String() + "/cancel_transaction"
	resp := CancelPurchaseResponse{}
	jsonReq, _ := json.Marshal(req)

	err := gateway.Client.Call("POST", path, bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		fmt.Println("Error Cancel Purchase : ", err)
		return resp, err
	}

	if resp.Status != "OK" {
		return resp, errors.New(resp.Error.Message)
	}

	return resp, nil
}

func (gateway *CoreGateway) CheckTransactionStatus(req *TransactionStatusRequest) (TransactionStatusResponse, error) {
	req.ServerKey = gateway.Client.ServerKey

	path := gateway.Client.APIEnvType.String() + "/transaction/status"
	resp := TransactionStatusResponse{}
	jsonReq, _ := json.Marshal(req)

	err := gateway.Client.Call("POST", path, bytes.NewBuffer(jsonReq), &resp)
	if err != nil {
		fmt.Println("Error Check Transaction Status : ", err)
		return resp, err
	}

	if resp.Status != "OK" {
		return resp, errors.New("Check Transaction Status Response Error")
	}

	return resp, nil
}
