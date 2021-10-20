//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of goidealo.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package goidealo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// OrdersReturn is to decode the json return
type OrdersReturn struct {
	Content []struct {
		IdealoOrderId       string    `json:"idealoOrderId"`
		MerchantOrderNumber string    `json:"merchantOrderNumber"`
		Created             time.Time `json:"created"`
		Updated             time.Time `json:"updated"`
		Status              string    `json:"status"`
		Currency            string    `json:"currency"`
		OffersPrice         string    `json:"offersPrice"`
		GrossPrice          string    `json:"grossPrice"`
		ShippingCosts       string    `json:"shippingCosts"`
		LineItems           []struct {
			Title                string `json:"title"`
			Price                string `json:"price"`
			Quantity             int    `json:"quantity"`
			Sku                  string `json:"sku"`
			MerchantDeliveryText string `json:"merchantDeliveryText"`
		} `json:"lineItems"`
		Customer struct {
			Email string `json:"email"`
		} `json:"customer"`
		Payment struct {
			PaymentMethod string `json:"paymentMethod"`
			TransactionId string `json:"transactionId"`
		} `json:"payment"`
		BillingAddress struct {
			Salutation   string `json:"salutation"`
			FirstName    string `json:"firstName"`
			LastName     string `json:"lastName"`
			AddressLine1 string `json:"addressLine1"`
			PostalCode   string `json:"postalCode"`
			City         string `json:"city"`
			CountryCode  string `json:"countryCode"`
			AddressLine2 string `json:"addressLine2,omitempty"`
		} `json:"billingAddress"`
		ShippingAddress struct {
			Salutation   string `json:"salutation"`
			FirstName    string `json:"firstName"`
			LastName     string `json:"lastName"`
			AddressLine1 string `json:"addressLine1"`
			PostalCode   string `json:"postalCode"`
			City         string `json:"city"`
			CountryCode  string `json:"countryCode"`
		} `json:"shippingAddress"`
		Fulfillment struct {
			Method   string `json:"method"`
			Tracking []struct {
				Code    string `json:"code"`
				Carrier string `json:"carrier"`
			} `json:"tracking"`
			Options []interface{} `json:"options"`
		} `json:"fulfillment"`
		Refunds []interface{} `json:"refunds"`
	} `json:"content"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
}

// OrderReturn is to decode the json data
type OrderReturn struct {
	IdealoOrderId       string    `json:"idealoOrderId"`
	MerchantOrderNumber string    `json:"merchantOrderNumber"`
	Created             time.Time `json:"created"`
	Updated             time.Time `json:"updated"`
	Status              string    `json:"status"`
	Currency            string    `json:"currency"`
	OffersPrice         string    `json:"offersPrice"`
	GrossPrice          string    `json:"grossPrice"`
	ShippingCosts       string    `json:"shippingCosts"`
	LineItems           []struct {
		Title                string `json:"title"`
		Price                string `json:"price"`
		Quantity             int    `json:"quantity"`
		Sku                  string `json:"sku"`
		MerchantDeliveryText string `json:"merchantDeliveryText"`
	} `json:"lineItems"`
	Customer struct {
		Email string `json:"email"`
	} `json:"customer"`
	Payment struct {
		PaymentMethod string `json:"paymentMethod"`
		TransactionId string `json:"transactionId"`
	} `json:"payment"`
	BillingAddress struct {
		Salutation   string `json:"salutation"`
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
		AddressLine1 string `json:"addressLine1"`
		AddressLine2 string `json:"addressLine2"`
		PostalCode   string `json:"postalCode"`
		City         string `json:"city"`
		CountryCode  string `json:"countryCode"`
	} `json:"billingAddress"`
	ShippingAddress struct {
		Salutation   string `json:"salutation"`
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
		AddressLine1 string `json:"addressLine1"`
		PostalCode   string `json:"postalCode"`
		City         string `json:"city"`
		CountryCode  string `json:"countryCode"`
	} `json:"shippingAddress"`
	Fulfillment struct {
		Method   string        `json:"method"`
		Tracking []interface{} `json:"tracking"`
		Options  []interface{} `json:"options"`
	} `json:"fulfillment"`
	Refunds []interface{} `json:"refunds"`
}

// MerchantOrderNumberBody is to structure the data
type MerchantOrderNumberBody struct {
	MerchantOrderNumber string `json:"merchantOrderNumber"`
}

// MerchantOrderNumberReturn is to decode the json return
type MerchantOrderNumberReturn struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Instance    string `json:"instance"`
	FieldErrors []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	} `json:"fieldErrors"`
}

// FulfillmentInformationBody is to structure the data
type FulfillmentInformationBody struct {
	Carrier      string   `json:"carrier"`
	TrackingCode []string `json:"trackingCode"`
}

// FulfillmentInformationReturn is to decode the json return
type FulfillmentInformationReturn struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Instance    string `json:"instance"`
	FieldErrors []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	} `json:"fieldErrors"`
}

// RevokeOrderBody is to structure the data
type RevokeOrderBody struct {
	Sku               string `json:"sku"`
	RemainingQuantity int    `json:"remainingQuantity"`
	Reason            string `json:"reason"`
	Comment           string `json:"comment"`
}

// RevokeOrderReturn is to decode the json return
type RevokeOrderReturn struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Instance    string `json:"instance"`
	FieldErrors []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	} `json:"fieldErrors"`
}

// RefundOrderBody is to structure the data
type RefundOrderBody struct {
	RefundAmount float64 `json:"refundAmount"`
	Currency     string  `json:"currency"`
}

// RefundOrderReturn is to decode the json return
type RefundOrderReturn struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Instance string `json:"instance"`
	Reason   string `json:"reason"`
}

// RefundsReturn is to decode the json return
type RefundsReturn struct {
	RefundId            string    `json:"refundId"`
	RefundTransactionId string    `json:"refundTransactionId"`
	Status              string    `json:"status"`
	Currency            string    `json:"currency"`
	RefundAmount        float64   `json:"refundAmount"`
	Created             time.Time `json:"created"`
	Updated             time.Time `json:"updated"`
}

// Orders is to get a list of orders from the moa api
func Orders(shopId int, parameter map[string]string, r Request) (OrdersReturn, error) {

	// Config new request
	c := Config{
		Moa:    true,
		Path:   fmt.Sprintf("/api/v2/shops/%d/orders?pageSize=250", shopId),
		Method: "GET",
	}

	// Check sandbox
	if r.Sandbox {
		c.Moa = false
		c.MoaSandbox = true
	}

	// Parse url & add attributes
	parse, err := url.Parse(c.Path)
	if err != nil {
		return OrdersReturn{}, err
	}

	newUrl, err := url.ParseQuery(parse.RawQuery)
	if err != nil {
		return OrdersReturn{}, err
	}

	for index, value := range parameter {
		newUrl.Add(index, value)
	}

	// Set new url
	parse.RawQuery = newUrl.Encode()
	c.Path = fmt.Sprintf("%s", parse)

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return OrdersReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = statusCodes(response.Status)
	if err != nil {
		return OrdersReturn{}, err
	}

	// Decode data
	var decode OrdersReturn
	json.NewDecoder(response.Body).Decode(&decode)

	// Return data
	return decode, nil

}

// Order is to get a specific order
func Order(shopId int, id string, r Request) (OrderReturn, error) {

	// Config new request
	c := Config{
		Moa:    true,
		Path:   fmt.Sprintf("/api/v2/shops/%d/orders/%s", shopId, id),
		Method: "GET",
	}

	// Check sandbox
	if r.Sandbox {
		c.Moa = false
		c.MoaSandbox = true
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return OrderReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = statusCodes(response.Status)
	if err != nil {
		return OrderReturn{}, err
	}

	// Decode data
	var decode OrderReturn
	json.NewDecoder(response.Body).Decode(&decode)

	// Return data
	return decode, nil

}

// MerchantOrderNumber is to set the merchant order number
func MerchantOrderNumber(shopId int, id string, body MerchantOrderNumberBody, r Request) (MerchantOrderNumberReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return MerchantOrderNumberReturn{}, err
	}

	// Config new request
	c := Config{
		Moa:    true,
		Path:   fmt.Sprintf("/api/v2/shops/%d/orders/%s/merchant-order-number", shopId, id),
		Method: "POST",
		Body:   convert,
	}

	// Check sandbox
	if r.Sandbox {
		c.Moa = false
		c.MoaSandbox = true
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return MerchantOrderNumberReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = statusCodes(response.Status)
	if err != nil {
		return MerchantOrderNumberReturn{}, err
	}

	// Decode data
	var decode MerchantOrderNumberReturn
	json.NewDecoder(response.Body).Decode(&decode)

	// Return data
	return decode, nil

}

// FulfillmentInformation is to set the fulfillment information
func FulfillmentInformation(shopId int, id string, body FulfillmentInformationBody, r Request) (FulfillmentInformationReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return FulfillmentInformationReturn{}, err
	}

	// Config new request
	c := Config{
		Moa:    true,
		Path:   fmt.Sprintf("/api/v2/shops/%d/orders/%s/fulfillment", shopId, id),
		Method: "POST",
		Body:   convert,
	}

	// Check sandbox
	if r.Sandbox {
		c.Moa = false
		c.MoaSandbox = true
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return FulfillmentInformationReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = statusCodes(response.Status)
	if err != nil {
		return FulfillmentInformationReturn{}, err
	}

	// Decode data
	var decode FulfillmentInformationReturn
	json.NewDecoder(response.Body).Decode(&decode)

	// Return data
	return decode, nil

}

// RevokeOrder is to set a revoke for an order
func RevokeOrder(shopId int, id string, body RevokeOrderBody, r Request) (RevokeOrderReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return RevokeOrderReturn{}, err
	}

	// Config new request
	c := Config{
		Moa:    true,
		Path:   fmt.Sprintf("/api/v2/shops/%d/orders/%s/revocations", shopId, id),
		Method: "POST",
		Body:   convert,
	}

	// Check sandbox
	if r.Sandbox {
		c.Moa = false
		c.MoaSandbox = true
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return RevokeOrderReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = statusCodes(response.Status)
	if err != nil {
		return RevokeOrderReturn{}, err
	}

	// Decode data
	var decode RevokeOrderReturn
	json.NewDecoder(response.Body).Decode(&decode)

	// Return data
	return decode, nil

}

// RefundOrder is to refund an order
func RefundOrder(shopId int, id string, body RefundOrderBody, r Request) (RefundOrderReturn, error) {

	// Convert body data
	convert, err := json.Marshal(body)
	if err != nil {
		return RefundOrderReturn{}, err
	}

	// Config new request
	c := Config{
		Moa:    true,
		Path:   fmt.Sprintf("/api/v2/shops/%d/orders/%s/refunds", shopId, id),
		Method: "POST",
		Body:   convert,
	}

	// Check sandbox
	if r.Sandbox {
		c.Moa = false
		c.MoaSandbox = true
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return RefundOrderReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = statusCodes(response.Status)
	if err != nil {
		return RefundOrderReturn{}, err
	}

	// Decode data
	var decode RefundOrderReturn
	json.NewDecoder(response.Body).Decode(&decode)

	// Return data
	return decode, nil

}

// Refunds is to check the refunds for an order
func Refunds(shopId int, id string, r Request) ([]RefundsReturn, error) {

	// Config new request
	c := Config{
		Moa:    true,
		Path:   fmt.Sprintf("/api/v2/shops/%d/orders/%s/refunds", shopId, id),
		Method: "GET",
	}

	// Check sandbox
	if r.Sandbox {
		c.Moa = false
		c.MoaSandbox = true
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return nil, err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = statusCodes(response.Status)
	if err != nil {
		return nil, err
	}

	// Decode data
	var decode []RefundsReturn
	json.NewDecoder(response.Body).Decode(&decode)

	// Return data
	return decode, nil

}
