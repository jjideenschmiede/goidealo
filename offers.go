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
	Deposit                  string `json:"deposit"`
	Size                     string `json:"size"`
	Colour                   string `json:"colour"`
	Gender                   string `json:"gender"`
	Material                 string `json:"material"`
	Replica                  bool   `json:"replica"`
	Used                     bool   `json:"used"`
	Download                 bool   `json:"download"`
	DynamicProductAttributes struct {
		Field1 []string `json:"22337"`
		Field2 []string `json:"19326"`
	} `json:"dynamicProductAttributes"`
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

	// Decode data
	var decode OfferReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OfferReturn{}, err
	}

	// Return data
	return decode, nil

}
