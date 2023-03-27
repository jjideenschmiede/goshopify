// Copyright 2023 J&J Ideenschmiede GmbH. All rights reserved.

package goshopify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// FulfillmentOrdersReturn is to decode the json data
type FulfillmentOrdersReturn struct {
	FulfillmentOrders []struct {
		Id                 int64    `json:"id"`
		ShopId             int64    `json:"shop_id"`
		OrderId            int64    `json:"order_id"`
		AssignedLocationId int64    `json:"assigned_location_id"`
		RequestStatus      string   `json:"request_status"`
		Status             string   `json:"status"`
		SupportedActions   []string `json:"supported_actions"`
		Destination        struct {
			Id        int64       `json:"id"`
			Address1  string      `json:"address1"`
			Address2  interface{} `json:"address2"`
			City      string      `json:"city"`
			Company   string      `json:"company"`
			Country   string      `json:"country"`
			Email     string      `json:"email"`
			FirstName string      `json:"first_name"`
			LastName  string      `json:"last_name"`
			Phone     interface{} `json:"phone"`
			Province  interface{} `json:"province"`
			Zip       string      `json:"zip"`
		} `json:"destination"`
		LineItems []struct {
			Id                  int64 `json:"id"`
			ShopId              int64 `json:"shop_id"`
			FulfillmentOrderId  int64 `json:"fulfillment_order_id"`
			Quantity            int   `json:"quantity"`
			LineItemId          int64 `json:"line_item_id"`
			InventoryItemId     int64 `json:"inventory_item_id"`
			FulfillableQuantity int   `json:"fulfillable_quantity"`
			VariantId           int64 `json:"variant_id"`
		} `json:"line_items"`
		FulfillAt           time.Time     `json:"fulfill_at"`
		FulfillBy           interface{}   `json:"fulfill_by"`
		InternationalDuties interface{}   `json:"international_duties"`
		FulfillmentHolds    []interface{} `json:"fulfillment_holds"`
		DeliveryMethod      struct {
			Id                  int64       `json:"id"`
			MethodType          string      `json:"method_type"`
			MinDeliveryDateTime interface{} `json:"min_delivery_date_time"`
			MaxDeliveryDateTime interface{} `json:"max_delivery_date_time"`
		} `json:"delivery_method"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		AssignedLocation struct {
			Address1    string      `json:"address1"`
			Address2    interface{} `json:"address2"`
			City        string      `json:"city"`
			CountryCode string      `json:"country_code"`
			LocationId  int64       `json:"location_id"`
			Name        string      `json:"name"`
			Phone       string      `json:"phone"`
			Province    interface{} `json:"province"`
			Zip         string      `json:"zip"`
		} `json:"assigned_location"`
		MerchantRequests []interface{} `json:"merchant_requests"`
	} `json:"fulfillment_orders"`
}

// FulfillmentOrders is to get the fulfillment orders for an order
func FulfillmentOrders(id int, r Request) (FulfillmentOrdersReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/orders/%d/fulfillment_orders.json", id), http.MethodGet, nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return FulfillmentOrdersReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode FulfillmentOrdersReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return FulfillmentOrdersReturn{}, err
	}

	// Return data
	return decode, err

}
