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
		IdealoOrderId string    `json:"idealoOrderId"`
		Created       time.Time `json:"created"`
		Updated       time.Time `json:"updated"`
		Status        string    `json:"status"`
		Currency      string    `json:"currency"`
		OffersPrice   string    `json:"offersPrice"`
		GrossPrice    string    `json:"grossPrice"`
		ShippingCosts string    `json:"shippingCosts"`
		LineItems     []struct {
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
	} `json:"content"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
}

// OrderReturn is to decode the json data
type OrderReturn struct {
	IdealoOrderId string    `json:"idealoOrderId"`
	Created       time.Time `json:"created"`
	Updated       time.Time `json:"updated"`
	Status        string    `json:"status"`
	Currency      string    `json:"currency"`
	OffersPrice   string    `json:"offersPrice"`
	GrossPrice    string    `json:"grossPrice"`
	ShippingCosts string    `json:"shippingCosts"`
	LineItems     []struct {
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
	err = pwsStatusCodes(response.Status)
	if err != nil {
		return OrdersReturn{}, err
	}

	// Decode data
	var decode OrdersReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OrdersReturn{}, err
	}

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
	err = pwsStatusCodes(response.Status)
	if err != nil {
		return OrderReturn{}, err
	}

	// Decode data
	var decode OrderReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OrderReturn{}, err
	}

	// Return data
	return decode, nil

}
