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
)

// PwsAccessTokenReturn is to decode the json return
type PwsAccessTokenReturn struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	ShopId      int    `json:"shop_id"`
}

// PwsAccessToken is to generate a bearer token for the pws api
func PwsAccessToken(r Request) (PwsAccessTokenReturn, error) {

	// Config new request
	c := Config{
		PwsAccessToken: true,
		Method:         "POST",
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return PwsAccessTokenReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = pwsStatusCodes(response.Status)
	if err != nil {
		return PwsAccessTokenReturn{}, err
	}

	// Decode data
	var decode PwsAccessTokenReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return PwsAccessTokenReturn{}, err
	}

	// Return data
	return decode, nil

}
