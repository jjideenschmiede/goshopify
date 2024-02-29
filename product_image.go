// Copyright 2024 J&J Ideenschmiede GmbH. All rights reserved.

package goshopify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ProductImageBody is to structure the data
type ProductImageBody struct {
	Image ProductImageBodyImage `json:"image"`
}

type ProductImageBodyImage struct {
	Id         int    `json:"id,omitempty"`
	Position   int    `json:"position,omitempty"`
	Src        string `json:"src"`
	Alt        string `json:"alt,omitempty"`
	VariantIds []int  `json:"variant_ids"`
}

// ProductImageReturn is to decode the json data
type ProductImageReturn struct {
	Image struct {
		Width             int       `json:"width"`
		Height            int       `json:"height"`
		Position          int       `json:"position"`
		Alt               string    `json:"alt"`
		Id                int       `json:"id"`
		ProductId         int       `json:"product_id"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		Src               string    `json:"src"`
		VariantIds        []int     `json:"variant_ids"`
		AdminGraphqlApiId string    `json:"admin_graphql_api_id"`
	} `json:"image"`
}

// AddProductImage is to add a new image to a product
func AddProductImage(id int, body ProductImageBody, r Request) (ProductImageReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductImageReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/products/%d/images.json", id), http.MethodPost, convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductImageReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductImageReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductImageReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteProductImage is to delete an image of a product
func DeleteProductImage(productId, imageId int, r Request) error {

	// Set config for new request
	c := Config{fmt.Sprintf("/products/%d/images/%d.json", productId, imageId), http.MethodDelete, nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return err
	}

	// Close request
	defer response.Body.Close()

	// Return nothing
	return nil
}
