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
	"bytes"
	"encoding/base64"
	"net/http"
)

const (
	accessTokenUrl = "https://api.idealo.com/mer/businessaccount/api/v1/oauth/token"
	pwsBaseUrl     = "https://import.idealo.com"
	moaBaseUrl     = "https://orders.idealo.com"
)

// Config is to define config data
type Config struct {
	accessToken, pws, moa bool
	Path, Method          string
	Body                  []byte
}

// Request is to define the request data
type Request struct {
	ClientId, ClientPassword, AccessToken string
}

// Send is to send a new request
func (c Config) Send(r Request) (*http.Response, error) {

	// Set url
	var url string

	switch {
	case c.accessToken:
		url = accessTokenUrl
	case c.pws:
		url = pwsBaseUrl + c.Path
	case c.moa:
		url = moaBaseUrl + c.Path
	}

	// Define client
	client := &http.Client{}

	// Request
	request, err := http.NewRequest(c.Method, url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	// Check api type & add an api header
	if c.pws || c.moa {
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+r.AccessToken)
	} else {
		request.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(r.ClientId+":"+r.ClientPassword)))
	}

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}
