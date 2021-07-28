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

import "encoding/json"

// AddProductBody is to structure the product data
type AddProductBody struct {
	Product AddProductBodyProduct `json:"product"`
}

type AddProductBodyProduct struct {
	Title       string                   `json:"title"`
	BodyHtml    string                   `json:"body_html"`
	Vendor      string                   `json:"vendor"`
	ProductType string                   `json:"product_type"`
	Status      string                   `json:"status"`
	Tags        string                   `json:"tags"`
	Images      []AddProductBodyImages   `json:"images"`
	Variants    []AddProductBodyVariants `json:"variants"`
	Options     []AddProductBodyOptions  `json:"options"`
}

type AddProductBodyImages struct {
	Src string `json:"src"`
}

type AddProductBodyVariants struct {
	Title               string  `json:"title"`
	Price               string  `json:"price"`
	Sku                 string  `json:"sku"`
	FulfillmentService  string  `json:"fulfillment_service"`
	InventoryManagement string  `json:"inventory_management"`
	Option1             string  `json:"option1"`
	Option2             string  `json:"option2"`
	Option3             string  `json:"option3"`
	Taxable             bool    `json:"taxable"`
	Barcode             string  `json:"barcode"`
	Grams               int     `json:"grams"`
	Weight              float64 `json:"weight"`
	WeightUnit          string  `json:"weight_unit"`
	InventoryQuantity   int     `json:"inventory_quantity"`
	RequiresShipping    bool    `json:"requires_shipping"`
}

type AddProductBodyOptions struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

// AddProductReturn is to decode the json return
type AddProductReturn struct {
	Product AddProductReturnProduct `json:"product"`
}

type AddProductReturnProduct struct {
	Id                int                        `json:"id"`
	Title             string                     `json:"title"`
	BodyHtml          string                     `json:"body_html"`
	Vendor            string                     `json:"vendor"`
	ProductType       string                     `json:"product_type"`
	CreatedAt         string                     `json:"created_at"`
	Handle            string                     `json:"handle"`
	UpdatedAt         string                     `json:"updated_at"`
	PublishedAt       string                     `json:"published_at"`
	TemplateSuffix    string                     `json:"template_suffix"`
	Status            string                     `json:"status"`
	PublishedScope    string                     `json:"published_scope"`
	Tags              string                     `json:"tags"`
	AdminGraphqlApiId string                     `json:"admin_graphql_api_id"`
	Variants          []AddProductReturnVariants `json:"variants"`
	Options           []AddProductReturnOptions  `json:"options"`
	Images            []AddProductReturnImages   `json:"images"`
	Image             AddProductReturnImage      `json:"image"`
}

type AddProductReturnVariants struct {
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

type AddProductReturnOptions struct {
	Id        int      `json:"id"`
	ProductId int      `json:"product_id"`
	Name      string   `json:"name"`
	Positions int      `json:"positions"`
	Values    []string `json:"values"`
}

type AddProductReturnImages struct {
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

type AddProductReturnImage struct {
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
func AddProduct(body *AddProductBody, r *Request) (*AddProductReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Set config for new request
	c := Config{"/products.json", "POST", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return nil, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode AddProductReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return nil, err
	}

	// Return data
	return &decode, err

}
