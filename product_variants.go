// Copyright 2023 J&J Ideenschmiede GmbH. All rights reserved.

package goshopify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ProductVariantBody is to structure the data
type ProductVariantBody struct {
	Variant ProductVariantBodyVariant `json:"variant"`
}

type ProductVariantBodyVariant struct {
	Id                  int     `json:"id,omitempty"`
	Title               string  `json:"title,omitempty"`
	Price               string  `json:"price,omitempty"`
	Sku                 string  `json:"sku,omitempty"`
	CompareAtPrice      string  `json:"compare_at_price"`
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
	} `json:"variant"`
	Errors interface{} `json:"errors,omitempty"`
}

// ProductVariantPriceBody is to structure the data
type ProductVariantPriceBody struct {
	Variant ProductVariantPriceBodyVariant `json:"variant"`
}

type ProductVariantPriceBodyVariant struct {
	Id             int    `json:"id"`
	Price          string `json:"price,omitempty"`
	CompareAtPrice string `json:"compare_at_price"`
}

// ProductVariantMetafieldsReturn is to decode the json data
type ProductVariantMetafieldsReturn struct {
	Metafields []struct {
		Id                int         `json:"id"`
		Namespace         string      `json:"namespace"`
		Key               string      `json:"key"`
		Value             string      `json:"value"`
		Description       interface{} `json:"description"`
		OwnerId           int         `json:"owner_id"`
		CreatedAt         time.Time   `json:"created_at"`
		UpdatedAt         time.Time   `json:"updated_at"`
		OwnerResource     string      `json:"owner_resource"`
		Type              string      `json:"type"`
		AdminGraphqlApiId string      `json:"admin_graphql_api_id"`
	} `json:"metafields"`
	Errors interface{} `json:"errors,omitempty"`
}

// ProductVariants is to get a list of all product variants
func ProductVariants(id int, r Request) (ProductVariantsReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/products/%d/variants.json", id), http.MethodGet, nil}

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
func AddProductVariant(body ProductVariantBody, r Request) (ProductVariantReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductVariantReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/products/%d/variants.json", body.Variant.Id), http.MethodPost, convert}

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
func UpdateProductVariant(body ProductVariantBody, r Request) (ProductVariantReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductVariantReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/variants/%d.json", body.Variant.Id), http.MethodPut, convert}

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

// DeleteVariant is to remove a variant form a product
func DeleteVariant(productId, variantId int, r Request) error {

	// Set config for new request
	c := Config{fmt.Sprintf("/products/%d/variants/%d.json", productId, variantId), http.MethodDelete, nil}

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

// UpdateProductVariantPrice is to update the inventory quantity of a product variant
func UpdateProductVariantPrice(body ProductVariantPriceBody, r Request) (ProductVariantReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductVariantReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/variants/%d.json", body.Variant.Id), http.MethodPut, convert}

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

// ProductVariantMetafields is to get a list of all product variant metafields
func ProductVariantMetafields(productId, variantId int, r Request) (ProductVariantMetafieldsReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/products/%d/variants/%d/metafields.json", productId, variantId), http.MethodGet, nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductVariantMetafieldsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductVariantMetafieldsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductVariantMetafieldsReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteProductVariantMetafield is to remove a metafield from a product variant
func DeleteProductVariantMetafield(productId, variantId, metafieldId int, r Request) error {

	// Set config for new request
	c := Config{fmt.Sprintf("/products/%d/variants/%d/metafields/%d.json", productId, variantId, metafieldId), http.MethodDelete, nil}

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
