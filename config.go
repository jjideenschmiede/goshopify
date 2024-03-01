// Copyright 2023 J&J Ideenschmiede GmbH. All rights reserved.

package goshopify

import (
	"bytes"
	"net/http"
)

const (
	transferProtocol = "https://"
	baseUrl          = ".myshopify.com/admin/api/"
	apiVersion       = "2024-04"
)

// Config is to define config data
type Config struct {
	Path, Method string
	Body         []byte
}

// Request is to define the request data
type Request struct {
	StoreName, AccessToken string
}

// Send is to send a new request
func (c Config) Send(r Request) (*http.Response, error) {

	// Set url
	url := transferProtocol + r.StoreName + baseUrl + apiVersion + c.Path

	// Define client
	client := &http.Client{}

	// Request
	request, err := http.NewRequest(c.Method, url, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	// Define header
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Shopify-Access-Token", r.AccessToken)

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}
