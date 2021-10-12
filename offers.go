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
)

// OfferBody is to structure the body data
type OfferBody struct {
	Sku                      string                  `json:"sku,omitempty"`
	Title                    string                  `json:"title,omitempty"`
	Price                    string                  `json:"price,omitempty"`
	Url                      string                  `json:"url,omitempty"`
	PaymentCosts             OfferBodyPaymentCosts   `json:"paymentCosts,omitempty"`
	DeliveryCosts            OfferBodyDeliveryCosts  `json:"deliveryCosts,omitempty"`
	BasePrice                string                  `json:"basePrice,omitempty"`
	PackagingUnit            int                     `json:"packagingUnit,omitempty"`
	VoucherCode              string                  `json:"voucherCode,omitempty"`
	BranchId                 string                  `json:"branchId,omitempty"`
	Brand                    string                  `json:"brand,omitempty"`
	Oens                     []string                `json:"oens,omitempty"`
	CategoryPath             []string                `json:"categoryPath,omitempty"`
	Description              string                  `json:"description,omitempty"`
	ImageUrls                []string                `json:"imageUrls,omitempty"`
	Eans                     []string                `json:"eans,omitempty"`
	Hans                     []string                `json:"hans,omitempty"`
	Pzns                     []string                `json:"pzns,omitempty"`
	Kbas                     []string                `json:"kbas,omitempty"`
	MerchantName             string                  `json:"merchantName,omitempty"`
	MerchantId               string                  `json:"merchantId,omitempty"`
	DeliveryComment          string                  `json:"deliveryComment,omitempty"`
	Delivery                 string                  `json:"delivery,omitempty"`
	MaxOrderProcessingTime   int                     `json:"maxOrderProcessingTime,omitempty"`
	FreeReturnDays           int                     `json:"freeReturnDays,omitempty"`
	Checkout                 bool                    `json:"checkout,omitempty"`
	CheckoutLimitPerPeriod   int                     `json:"checkoutLimitPerPeriod,omitempty"`
	QuantityPerOrder         int                     `json:"quantityPerOrder,omitempty"`
	MinimumPrice             string                  `json:"minimumPrice,omitempty"`
	FulfillmentType          string                  `json:"fulfillmentType,omitempty"`
	TwoManHandlingFee        string                  `json:"twoManHandlingFee,omitempty"`
	DisposalFee              string                  `json:"disposalFee,omitempty"`
	Eec                      string                  `json:"eec,omitempty"`
	EnergyLabels             []OfferBodyEnergyLabels `json:"energyLabels,omitempty"`
	Deposit                  string                  `json:"deposit,omitempty"`
	Size                     string                  `json:"size,omitempty"`
	Colour                   string                  `json:"colour,omitempty"`
	Gender                   string                  `json:"gender,omitempty"`
	Material                 string                  `json:"material,omitempty"`
	Replica                  bool                    `json:"replica,omitempty"`
	Used                     bool                    `json:"used,omitempty"`
	Download                 bool                    `json:"download,omitempty"`
	DynamicProductAttributes interface{}             `json:"dynamicProductAttributes,omitempty"`
}

type OfferBodyPaymentCosts struct {
	PAYPAL string `json:"PAYPAL,omitempty"`
}

type OfferBodyDeliveryCosts struct {
	DHL string `json:"DHL,omitempty"`
}

type OfferBodyEnergyLabels struct {
	EfficiencyClass           string `json:"efficiencyClass,omitempty"`
	Spectrum                  string `json:"spectrum,omitempty"`
	LabelUrl                  string `json:"labelUrl,omitempty"`
	DataSheetUrl              string `json:"dataSheetUrl,omitempty"`
	FuelEfficiencyClass       string `json:"fuelEfficiencyClass,omitempty"`
	WetGripClass              string `json:"wetGripClass,omitempty"`
	ExternalRollingNoise      int    `json:"externalRollingNoise,omitempty"`
	ExternalRollingNoiseClass string `json:"externalRollingNoiseClass,omitempty"`
	SnowGrip                  bool   `json:"snowGrip,omitempty"`
	IceGrip                   bool   `json:"iceGrip,omitempty"`
	Version                   int    `json:"version,omitempty"`
}

// OfferReturn is to decode the json data
type OfferReturn struct {
	Sku           string   `json:"sku"`
	Title         string   `json:"title"`
	Price         string   `json:"price"`
	Url           string   `json:"url"`
	BasePrice     string   `json:"basePrice"`
	PackagingUnit int      `json:"packagingUnit"`
	VoucherCode   string   `json:"voucherCode"`
	BranchId      string   `json:"branchId"`
	Brand         string   `json:"brand"`
	Oens          []string `json:"oens"`
	CategoryPath  []string `json:"categoryPath"`
	Description   string   `json:"description"`
	ImageUrls     []string `json:"imageUrls"`
	Eans          []string `json:"eans"`
	Hans          []string `json:"hans"`
	Pzns          []string `json:"pzns"`
	Kbas          []string `json:"kbas"`
	MerchantName  string   `json:"merchantName"`
	MerchantId    string   `json:"merchantId"`
	PaymentCosts  struct {
		PAYPAL        string `json:"PAYPAL"`
		CASHINADVANCE string `json:"CASH_IN_ADVANCE"`
		INVOICE       string `json:"INVOICE"`
		CREDITCARD    string `json:"CREDIT_CARD"`
	} `json:"paymentCosts"`
	DeliveryCosts struct {
		DHL string `json:"DHL"`
		DPD string `json:"DPD"`
	} `json:"deliveryCosts"`
	DeliveryComment        string `json:"deliveryComment"`
	Delivery               string `json:"delivery"`
	MaxOrderProcessingTime int    `json:"maxOrderProcessingTime"`
	FreeReturnDays         int    `json:"freeReturnDays"`
	Checkout               bool   `json:"checkout"`
	CheckoutLimitPerPeriod int    `json:"checkoutLimitPerPeriod"`
	QuantityPerOrder       int    `json:"quantityPerOrder"`
	MinimumPrice           string `json:"minimumPrice"`
	FulfillmentType        string `json:"fulfillmentType"`
	TwoManHandlingFee      string `json:"twoManHandlingFee"`
	DisposalFee            string `json:"disposalFee"`
	Eec                    string `json:"eec"`
	EnergyLabels           []struct {
		EfficiencyClass string `json:"efficiencyClass"`
		Spectrum        string `json:"spectrum"`
		LabelUrl        string `json:"labelUrl"`
		DataSheetUrl    string `json:"dataSheetUrl"`
		Version         int    `json:"version"`
	} `json:"energyLabels"`
	Deposit                  string      `json:"deposit"`
	Size                     string      `json:"size"`
	Colour                   string      `json:"colour"`
	Gender                   string      `json:"gender"`
	Material                 string      `json:"material"`
	Replica                  bool        `json:"replica"`
	Used                     bool        `json:"used"`
	Download                 bool        `json:"download"`
	DynamicProductAttributes interface{} `json:"dynamicProductAttributes"`
	FieldErrors              []struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	} `json:"fieldErrors"`
	GeneralErrors []string `json:"generalErrors,omitempty"`
}

// Offer is to get an offer by id
func Offer(shopId int, sku string, r Request) (OfferReturn, error) {

	// Config new request
	c := Config{
		Pws:    true,
		Path:   fmt.Sprintf("/shop/%d/offer/%s", shopId, sku),
		Method: "GET",
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return OfferReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = pwsStatusCodes(response.Status)
	if err != nil {
		return OfferReturn{}, err
	}

	// Decode data
	var decode OfferReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OfferReturn{}, err
	}

	// Return data
	return decode, nil

}

// CreateOffer is to create an new offer
func CreateOffer(shopId int, body OfferBody, r Request) (OfferReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return OfferReturn{}, err
	}

	// Config new request
	c := Config{
		Pws:    true,
		Path:   fmt.Sprintf("/shop/%d/offer/%s", shopId, body.Sku),
		Method: "PUT",
		Body:   convert,
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return OfferReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = pwsStatusCodes(response.Status)
	if err != nil {
		return OfferReturn{}, err
	}

	// Decode data
	var decode OfferReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OfferReturn{}, err
	}

	// Return data
	return decode, nil

}

// DeleteOffer is to delete an offer
func DeleteOffer(shopId int, sku string, r Request) error {

	// Config new request
	c := Config{
		Pws:    true,
		Path:   fmt.Sprintf("/shop/%d/offer/%s", shopId, sku),
		Method: "DELETE",
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = pwsStatusCodes(response.Status)
	if err != nil {
		return err
	}

	// Return nothing
	return nil

}

// DeleteAllOffers is to delete all existing offers
func DeleteAllOffers(shopId int, r Request) error {

	// Config new request
	c := Config{
		Pws:    true,
		Path:   fmt.Sprintf("/shop/%d/offer", shopId),
		Method: "DELETE",
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = pwsStatusCodes(response.Status)
	if err != nil {
		return err
	}

	// Return nothing
	return nil

}

// UpdateOffersTimestamp is to update the timestamp of all offers
func UpdateOffersTimestamp(shopId int, r Request) error {

	// Config new request
	c := Config{
		Pws:    true,
		Path:   fmt.Sprintf("/shop/%d/offer", shopId),
		Method: "PUT",
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = pwsStatusCodes(response.Status)
	if err != nil {
		return err
	}

	// Return nothing
	return nil

}
