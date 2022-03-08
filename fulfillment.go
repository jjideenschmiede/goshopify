//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2022 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
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

// FulfillmentBody is to structure the body data
type FulfillmentBody struct {
	Fulfillment FulfillmentBodyFulfillment `json:"fulfillment"`
}

type FulfillmentBodyFulfillment struct {
	LocationId      int    `json:"location_id"`
	TrackingNumber  string `json:"tracking_number"`
	TrackingCompany string `json:"tracking_company"`
	TrackingUrl     string `json:"tracking_url"`
}

// FulfillmentReturn is to decode the json data
type FulfillmentReturn struct {
	Fulfillment struct {
		Id              int64       `json:"id"`
		OrderId         int64       `json:"order_id"`
		Status          string      `json:"status"`
		CreatedAt       time.Time   `json:"created_at"`
		Service         string      `json:"service"`
		UpdatedAt       time.Time   `json:"updated_at"`
		TrackingCompany string      `json:"tracking_company"`
		ShipmentStatus  interface{} `json:"shipment_status"`
		LocationId      int64       `json:"location_id"`
		OriginAddress   interface{} `json:"origin_address"`
		LineItems       []struct {
			Id                         int64         `json:"id"`
			VariantId                  int64         `json:"variant_id"`
			Title                      string        `json:"title"`
			Quantity                   int           `json:"quantity"`
			Sku                        string        `json:"sku"`
			VariantTitle               string        `json:"variant_title"`
			Vendor                     string        `json:"vendor"`
			FulfillmentService         string        `json:"fulfillment_service"`
			ProductId                  int64         `json:"product_id"`
			RequiresShipping           bool          `json:"requires_shipping"`
			Taxable                    bool          `json:"taxable"`
			GiftCard                   bool          `json:"gift_card"`
			Name                       string        `json:"name"`
			VariantInventoryManagement string        `json:"variant_inventory_management"`
			Properties                 []interface{} `json:"properties"`
			ProductExists              bool          `json:"product_exists"`
			FulfillableQuantity        int           `json:"fulfillable_quantity"`
			Grams                      int           `json:"grams"`
			Price                      string        `json:"price"`
			TotalDiscount              string        `json:"total_discount"`
			FulfillmentStatus          string        `json:"fulfillment_status"`
			PriceSet                   struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"price_set"`
			TotalDiscountSet struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"total_discount_set"`
			DiscountAllocations []interface{} `json:"discount_allocations"`
			Duties              []interface{} `json:"duties"`
			AdminGraphqlApiId   string        `json:"admin_graphql_api_id"`
			TaxLines            []struct {
				Title         string  `json:"title"`
				Price         string  `json:"price"`
				Rate          float64 `json:"rate"`
				ChannelLiable bool    `json:"channel_liable"`
				PriceSet      struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"price_set"`
			} `json:"tax_lines"`
			OriginLocation struct {
				Id           int64  `json:"id"`
				CountryCode  string `json:"country_code"`
				ProvinceCode string `json:"province_code"`
				Name         string `json:"name"`
				Address1     string `json:"address1"`
				Address2     string `json:"address2"`
				City         string `json:"city"`
				Zip          string `json:"zip"`
			} `json:"origin_location"`
		} `json:"line_items"`
		TrackingNumber  string   `json:"tracking_number"`
		TrackingNumbers []string `json:"tracking_numbers"`
		TrackingUrl     string   `json:"tracking_url"`
		TrackingUrls    []string `json:"tracking_urls"`
		Receipt         struct {
		} `json:"receipt"`
		Name              string `json:"name"`
		AdminGraphqlApiId string `json:"admin_graphql_api_id"`
	} `json:"fulfillment"`
}

// AddFulfillment is to add a new fulfillment
func AddFulfillment(id int, body FulfillmentBody, r Request) (FulfillmentReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return FulfillmentReturn{}, err
	}

	// Set config for new request
	c := Config{fmt.Sprintf("/orders/%d/fulfillments.json", id), "POST", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return FulfillmentReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode FulfillmentReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return FulfillmentReturn{}, err
	}

	// Return data
	return decode, err

}
