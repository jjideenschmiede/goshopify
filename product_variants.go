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

// ProductVariantReturn is to decode the json data
type ProductVariantReturn struct {
	Variants []struct {
		Id                   int64       `json:"id"`
		ProductId            int64       `json:"product_id"`
		Title                string      `json:"title"`
		Price                string      `json:"price"`
		Sku                  string      `json:"sku"`
		Position             int         `json:"position"`
		InventoryPolicy      string      `json:"inventory_policy"`
		CompareAtPrice       string      `json:"compare_at_price"`
		FulfillmentService   string      `json:"fulfillment_service"`
		InventoryManagement  string      `json:"inventory_management"`
		Option1              string      `json:"option1"`
		Option2              interface{} `json:"option2"`
		Option3              interface{} `json:"option3"`
		CreatedAt            time.Time   `json:"created_at"`
		UpdatedAt            time.Time   `json:"updated_at"`
		Taxable              bool        `json:"taxable"`
		Barcode              string      `json:"barcode"`
		Grams                int         `json:"grams"`
		ImageId              int64       `json:"image_id"`
		Weight               float64     `json:"weight"`
		WeightUnit           string      `json:"weight_unit"`
		InventoryItemId      int64       `json:"inventory_item_id"`
		InventoryQuantity    int         `json:"inventory_quantity"`
		OldInventoryQuantity int         `json:"old_inventory_quantity"`
		RequiresShipping     bool        `json:"requires_shipping"`
		AdminGraphqlApiId    string      `json:"admin_graphql_api_id"`
	} `json:"variants"`
}

// ProductVariants is to get a list of all product variants
func ProductVariants(id int, r Request) (ProductVariantReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/products/%d/variants.json", id), "GET", nil}

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
