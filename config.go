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
	pwsAccessTokenUrl = "https://api.idealo.com/mer/businessaccount/api/v1/oauth/token"
	pwsBaseUrl        = "https://import.idealo.com"
	moaBaseUrl        = "https://orders.idealo.com"
	moaSandboxBaseUrl = "https://orders-sandbox.idealo.com"
)

// Config is to define config data
type Config struct {
	PwsAccessToken, MoaAccessToken, MoaSandboxAccessToken, Pws, Moa, MoaSandbox bool
	Path, Method                                                                string
	Body                                                                        []byte
}

// Request is to define the request data
type Request struct {
	ClientId, ClientPassword, AccessToken string
	Sandbox                               bool
	Header                                map[string]string
}

// Send is to send a new request
func (c Config) Send(r Request) (*http.Response, error) {

	// Set url
	var url string

	switch {
	case c.PwsAccessToken:
		url = pwsAccessTokenUrl
	case c.MoaAccessToken:
		url = moaBaseUrl + c.Path
	case c.MoaSandboxAccessToken:
		url = moaSandboxBaseUrl + c.Path
	case c.Pws:
		url = pwsBaseUrl + c.Path
	case c.Moa:
		url = moaBaseUrl + c.Path
	case c.MoaSandbox:
		url = moaSandboxBaseUrl + c.Path
	}

	// Define client
	client := &http.Client{}

	// Request
	request, err := http.NewRequest(c.Method, url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	// Set header
	for index, value := range r.Header {
		request.Header.Set(index, value)
	}

	// Check api type & add an api header
	switch {
	case c.Pws || c.Moa || c.MoaSandbox:
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Authorization", "Bearer "+r.AccessToken)
	default:
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
