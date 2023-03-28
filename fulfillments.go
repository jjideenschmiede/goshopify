// Copyright 2023 J&J Ideenschmiede GmbH. All rights reserved.

package goshopify

import (
	"encoding/json"
	"net/http"
	"time"
)

// FulfillmentsBody is to structure the body data
type FulfillmentsBody struct {
	Fulfillment FulfillmentsBodyFulfillment `json:"fulfillment"`
}

type FulfillmentsBodyFulfillment struct {
	Message                     string                                      `json:"message"`
	NotifyCustomer              bool                                        `json:"notify_customer"`
	TrackingInfo                FulfillmentsBodyTrackingInfo                `json:"tracking_info"`
	LineItemsByFulfillmentOrder FulfillmentsBodyLineItemsByFulfillmentOrder `json:"line_items_by_fulfillment_order"`
}

type FulfillmentsBodyTrackingInfo struct {
	Number  string `json:"number"`
	Url     string `json:"url"`
	Company string `json:"company"`
}

type FulfillmentsBodyLineItemsByFulfillmentOrder struct {
	FulfillmentOrderId        int           `json:"fulfillment_order_id"`
	FulfillmentOrderLineItems []interface{} `json:"fulfillment_order_line_items"`
}

// FulfillmentsReturn is to decode the json data
type FulfillmentsReturn struct {
	Fulfillment struct {
		Id              int         `json:"id"`
		OrderId         int         `json:"order_id"`
		Status          string      `json:"status"`
		CreatedAt       time.Time   `json:"created_at"`
		Service         string      `json:"service"`
		UpdatedAt       time.Time   `json:"updated_at"`
		TrackingCompany string      `json:"tracking_company"`
		ShipmentStatus  interface{} `json:"shipment_status"`
		LocationId      int         `json:"location_id"`
		OriginAddress   interface{} `json:"origin_address"`
		LineItems       []struct {
			Id                         int           `json:"id"`
			VariantId                  int           `json:"variant_id"`
			Title                      string        `json:"title"`
			Quantity                   int           `json:"quantity"`
			Sku                        string        `json:"sku"`
			VariantTitle               string        `json:"variant_title"`
			Vendor                     string        `json:"vendor"`
			FulfillmentService         string        `json:"fulfillment_service"`
			ProductId                  int           `json:"product_id"`
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

// Fulfillments is to create a fulfillment for an order or multiple
func Fulfillments(body FulfillmentsBody, r Request) (FulfillmentsReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return FulfillmentsReturn{}, err
	}

	// Set config for new request
	c := Config{"/fulfillments.json", http.MethodPost, convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return FulfillmentsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode FulfillmentsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return FulfillmentsReturn{}, err
	}

	// Return data
	return decode, err

}
