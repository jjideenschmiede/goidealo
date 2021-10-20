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

import "encoding/json"

// MoaAccessTokenReturn is to decode the json return
type MoaAccessTokenReturn struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	ShopId      int    `json:"shop_id"`
}

// MoaAccessToken is to generate a bearer token for the moa api
func MoaAccessToken(r Request) (MoaAccessTokenReturn, error) {

	// Config new request
	c := Config{
		MoaAccessToken: true,
		Path:           "/api/v2/oauth/token",
		Method:         "POST",
	}

	// Check sandbox
	if r.Sandbox {
		c.MoaAccessToken = false
		c.MoaSandboxAccessToken = true
	}

	// Send new request
	response, err := c.Send(r)
	if err != nil {
		return MoaAccessTokenReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Check response status
	err = statusCodes(response.Status)
	if err != nil {
		return MoaAccessTokenReturn{}, err
	}

	// Decode data
	var decode MoaAccessTokenReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return MoaAccessTokenReturn{}, err
	}

	// Return data
	return decode, nil

}
