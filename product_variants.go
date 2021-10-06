//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of goshopify.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package goshopify

import (
	"encoding/json"
	"fmt"
	"time"
)

// ProductVariantBody is to structure the data
type ProductVariantBody struct {
	Variant ProductVariantBodyVariant `json:"variant"`
}

type ProductVariantBodyVariant struct {
	Id                  string  `json:"id,omitempty"`
	Title               string  `json:"title,omitempty"`
	Price               string  `json:"price,omitempty"`
	Sku                 string  `json:"sku,omitempty"`
	CompareAtPrice      string  `json:"compare_at_price,omitempty"`
	FulfillmentService  string  `json:"fulfillment_service,omitempty"`
	InventoryManagement string  `json:"inventory_management,omitempty"`
	Option1             string  `json:"option1,omitempty"`
	Option2             string  `json:"option2,omitempty"`
	Option3             string  `json:"option3,omitempty"`
	Taxable             bool    `json:"taxable,omitempty"`
	Barcode             string  `json:"barcode,omitempty"`
	Grams               int     `json:"grams,omitempty"`
	Weight              float64 `json:"weight,omitempty"`
	WeightUnit          string  `json:"weight_unit,omitempty"`
	InventoryQuantity   int     `json:"inventory_quantity,omitempty"`
	RequiresShipping    bool    `json:"requires_shipping,omitempty"`
}

// ProductVariantsReturn is to decode the json data
type ProductVariantsReturn struct {
	Variants []struct {
		Id                   int       `json:"id"`
		ProductId            int       `json:"product_id"`
		Title                string    `json:"title"`
		Price                string    `json:"price"`
		Sku                  string    `json:"sku"`
		Position             int       `json:"position"`
		InventoryPolicy      string    `json:"inventory_policy"`
		CompareAtPrice       string    `json:"compare_at_price"`
		FulfillmentService   string    `json:"fulfillment_service"`
		InventoryManagement  string    `json:"inventory_management"`
		Option1              string    `json:"option1"`
		Option2              string    `json:"option2"`
		Option3              string    `json:"option3"`
		CreatedAt            time.Time `json:"created_at"`
		UpdatedAt            time.Time `json:"updated_at"`
		Taxable              bool      `json:"taxable"`
		Barcode              string    `json:"barcode"`
		Grams                int       `json:"grams"`
		ImageId              int       `json:"image_id"`
		Weight               float64   `json:"weight"`
		WeightUnit           string    `json:"weight_unit"`
		InventoryItemId      int       `json:"inventory_item_id"`
		InventoryQuantity    int       `json:"inventory_quantity"`
		OldInventoryQuantity int       `json:"old_inventory_quantity"`
		RequiresShipping     bool      `json:"requires_shipping"`
		AdminGraphqlApiId    string    `json:"admin_graphql_api_id"`
	} `json:"variants"`
}

// ProductVariantReturn is to decode the json data
type ProductVariantReturn struct {
	Variant struct {
		Id                   int64     `json:"id"`
		ProductId            int64     `json:"product_id"`
		Title                string    `json:"title"`
		Price                string    `json:"price"`
		Sku                  string    `json:"sku"`
		Position             int       `json:"position"`
		InventoryPolicy      string    `json:"inventory_policy"`
		CompareAtPrice       string    `json:"compare_at_price"`
		FulfillmentService   string    `json:"fulfillment_service"`
		InventoryManagement  string    `json:"inventory_management"`
		Option1              string    `json:"option1"`
		Option2              string    `json:"option2"`
		Option3              string    `json:"option3"`
		CreatedAt            time.Time `json:"created_at"`
		UpdatedAt            time.Time `json:"updated_at"`
		Taxable              bool      `json:"taxable"`
		Barcode              string    `json:"barcode"`
		Grams                int       `json:"grams"`
		ImageId              string    `json:"image_id"`
		Weight               float64   `json:"weight"`
		WeightUnit           string    `json:"weight_unit"`
		InventoryItemId      int64     `json:"inventory_item_id"`
		InventoryQuantity    int       `json:"inventory_quantity"`
		OldInventoryQuantity int       `json:"old_inventory_quantity"`
		RequiresShipping     bool      `json:"requires_shipping"`
		AdminGraphqlApiId    string    `json:"admin_graphql_api_id"`
	} `json:"variant"`
}

// ProductVariants is to get a list of all product variants
func ProductVariants(id int, r Request) (ProductVariantsReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/products/%d/variants.json", id), "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductVariantsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductVariantsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductVariantsReturn{}, err
	}

	// Return data
	return decode, err

}

// AddProductVariant is to add a new variant to a product
func AddProductVariant(id int, body ProductVariantBody, r Request) (ProductVariantReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductVariantReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/products/%d/variants.json", id), "POST", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductVariantReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductVariantReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductVariantReturn{}, err
	}

	// Return data
	return decode, err

}

// UpdateProductVariant is to update an existing product variant
func UpdateProductVariant(id int, body ProductVariantBody, r Request) (ProductVariantReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductVariantReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/variants/%d.json", id), "PUT", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductVariantReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductVariantReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductVariantReturn{}, err
	}

	// Return data
	return decode, err

}
