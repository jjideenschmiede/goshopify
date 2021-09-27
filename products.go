//**********************************************************
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of goshopify.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor
//
//**********************************************************

package goshopify

import (
	"encoding/json"
	"fmt"
)

// ProductBody is to structure the product data
type ProductBody struct {
	Product ProductBodyProduct `json:"product"`
}

type ProductBodyProduct struct {
	Id          int                   `json:"id,omitempty"`
	Title       string                `json:"title,omitempty"`
	BodyHtml    string                `json:"body_html,omitempty"`
	Vendor      string                `json:"vendor,omitempty"`
	ProductType string                `json:"product_type,omitempty"`
	Status      string                `json:"status,omitempty"`
	Tags        string                `json:"tags,omitempty"`
	Images      []ProductBodyImages   `json:"images,omitempty"`
	Variants    []ProductBodyVariants `json:"variants,omitempty"`
	Options     []ProductBodyOptions  `json:"options,omitempty"`
}

type ProductBodyImages struct {
	Src string `json:"src,omitempty"`
}

type ProductBodyVariants struct {
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

type ProductBodyOptions struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}

// ProductReturn is to decode the json return
type ProductReturn struct {
	Product ProductReturnProduct `json:"product"`
}

type ProductReturnProduct struct {
	Id                int                     `json:"id"`
	Title             string                  `json:"title"`
	BodyHtml          string                  `json:"body_html"`
	Vendor            string                  `json:"vendor"`
	ProductType       string                  `json:"product_type"`
	CreatedAt         string                  `json:"created_at"`
	Handle            string                  `json:"handle"`
	UpdatedAt         string                  `json:"updated_at"`
	PublishedAt       string                  `json:"published_at"`
	TemplateSuffix    string                  `json:"template_suffix"`
	Status            string                  `json:"status"`
	PublishedScope    string                  `json:"published_scope"`
	Tags              string                  `json:"tags"`
	AdminGraphqlApiId string                  `json:"admin_graphql_api_id"`
	Variants          []ProductReturnVariants `json:"variants"`
	Options           []ProductReturnOptions  `json:"options"`
	Images            []ProductReturnImages   `json:"images"`
	Image             ProductReturnImage      `json:"image"`
}

type ProductReturnVariants struct {
	Id                   int     `json:"id"`
	ProductId            int     `json:"product_id"`
	Title                string  `json:"title"`
	Price                string  `json:"price"`
	Sku                  string  `json:"sku"`
	Position             int     `json:"position"`
	InventoryPolicy      string  `json:"inventory_policy"`
	CompareAtPrice       string  `json:"compare_at_price"`
	FulfillmentService   string  `json:"fulfillment_service"`
	InventoryManagement  string  `json:"inventory_management"`
	Option1              string  `json:"option1"`
	Option2              string  `json:"option2"`
	Option3              string  `json:"option3"`
	CreatedAt            string  `json:"created_at"`
	UpdatedAt            string  `json:"updated_at"`
	Taxable              bool    `json:"taxable"`
	Barcode              string  `json:"barcode"`
	Grams                int     `json:"grams"`
	ImageId              string  `json:"image_id"`
	Weight               float64 `json:"weight"`
	WeightUnit           string  `json:"weight_unit"`
	InventoryItemId      int     `json:"inventory_item_id"`
	OldInventoryQuantity int     `json:"old_inventory_quantity"`
	RequiresShipping     bool    `json:"requires_shipping"`
	AdminGraphqlApiId    string  `json:"admin_graphql_api_id"`
}

type ProductReturnOptions struct {
	Id        int      `json:"id"`
	ProductId int      `json:"product_id"`
	Name      string   `json:"name"`
	Positions int      `json:"positions"`
	Values    []string `json:"values"`
}

type ProductReturnImages struct {
	Id                int           `json:"id"`
	ProductId         int           `json:"product_id"`
	Position          int           `json:"position"`
	CreatedAt         string        `json:"created_at"`
	UpdatedAt         string        `json:"updated_at"`
	Alt               string        `json:"alt"`
	Width             int           `json:"width"`
	Height            int           `json:"height"`
	Src               string        `json:"src"`
	VariantsIds       []interface{} `json:"variants_ids"`
	AdminGraphqlApiId string        `json:"admin_graphql_api_id"`
}

type ProductReturnImage struct {
	Id                int           `json:"id"`
	ProductId         int           `json:"product_id"`
	Position          int           `json:"position"`
	CreatedAt         string        `json:"created_at"`
	UpdatedAt         string        `json:"updated_at"`
	Alt               string        `json:"alt"`
	Width             int           `json:"width"`
	Height            int           `json:"height"`
	Src               string        `json:"src"`
	VariantsIds       []interface{} `json:"variants_ids"`
	AdminGraphqlApiId string        `json:"admin_graphql_api_id"`
}

// AddProduct is to create a new single or variant product
func AddProduct(body ProductBody, r Request) (ProductReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductReturn{}, err
	}

	// Set config for new request
	c := Config{"/products.json", "POST", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductReturn{}, err
	}

	// Return data
	return decode, err

}

// UpdateProduct is to create a new single or variant product
func UpdateProduct(body ProductBody, r Request) (ProductReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return ProductReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/products/%d.json", body.Product.Id), "PUT", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return ProductReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode ProductReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductReturn{}, err
	}

	// Return data
	return decode, err

}

// DeleteProduct is to remove a product form shop
func DeleteProduct(id int, r Request) error {

	// Set config for new request
	c := Config{fmt.Sprintf("/products/%d.json", id), "DELETE", nil}

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
