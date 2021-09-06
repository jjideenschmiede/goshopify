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
)

// OrderReturn is to decode the order return form shopify
type OrderReturn struct {
	Orders []OrderReturnOrders `json:"orders"`
}

type OrderReturnOrders struct {
	Id                       int                        `json:"id"`
	AdminGraphqlApiId        string                     `json:"admin_graphql_api_id"`
	AppId                    int                        `json:"app_id"`
	BrowserIp                string                     `json:"browser_ip"`
	BuyerAcceptsMarketing    bool                       `json:"buyer_accepts_marketing"`
	CancelReason             interface{}                `json:"cancel_reason"`
	CancelledAt              interface{}                `json:"cancelled_at"`
	CartToken                string                     `json:"cart_token"`
	CheckoutId               int                        `json:"checkout_id"`
	CheckoutToken            string                     `json:"checkout_token"`
	ClientDetails            OrderReturnClientDetails   `json:"client_details"`
	ClosedAt                 interface{}                `json:"closed_at"`
	Confirmed                bool                       `json:"confirmed"`
	ContactEmail             string                     `json:"contact_email"`
	CreateAt                 string                     `json:"create_at"`
	Currency                 string                     `json:"currency"`
	CurrentSubtotalPrice     string                     `json:"current_subtotal_price"`
	CurrentSubtotalPriceSet  OrderReturnPriceSet        `json:"current_subtotal_price_set"`
	CurrentTotalDiscounts    string                     `json:"current_total_discounts"`
	CurrentTotalDiscountsSet OrderReturnPriceSet        `json:"current_total_discounts_set"`
	CurrentTotalDutiesSet    interface{}                `json:"current_total_duties_set"`
	CurrentTotalPrice        string                     `json:"current_total_price"`
	CurrentTotalPriceSet     OrderReturnPriceSet        `json:"current_total_price_set"`
	CurrentTotalTax          string                     `json:"current_total_tax"`
	CurrentTotalTaxSet       OrderReturnPriceSet        `json:"current_total_tax_set"`
	TotalTipReceived         string                     `json:"total_tip_received"`
	TotalWeight              int                        `json:"total_weight"`
	UpdatedAt                string                     `json:"updated_at"`
	UserId                   interface{}                `json:"user_id"`
	BillingAddress           OrderReturnBillingAddress  `json:"billing_address"`
	Customer                 OrderReturnCustomer        `json:"customer"`
	DiscountApplications     []interface{}              `json:"discount_applications"`
	Fulfillments             []interface{}              `json:"fulfillments"`
	LineItems                []OrderReturnLineItems     `json:"line_items"`
	PaymentDetails           OrderReturnPaymentDetails  `json:"pament_details"`
	Refunds                  []interface{}              `json:"refunds"`
	ShippingLines            []OrderReturnShippingLines `json:"shipping_lines"`
}

type OrderReturnClientDetails struct {
	AcceptLanguage string      `json:"accept_language"`
	BrowserHeight  int         `json:"browser_height"`
	BrowserIp      string      `json:"browser_ip"`
	BrowserWidth   int         `json:"browser_width"`
	SessionHash    interface{} `json:"session_hash"`
	UserAgent      string      `json:"user_agent"`
}

type OrderReturnPriceSet struct {
	ShopMoney        OrderReturnMoney `json:"shop_money"`
	PresentmentMoney OrderReturnMoney `json:"presentment_money"`
}

type OrderReturnMoney struct {
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}

type OrderReturnBillingAddress struct {
	FirstName    string  `json:"first_name"`
	Address1     string  `json:"address1"`
	Phone        string  `json:"phone"`
	City         string  `json:"city"`
	Zip          string  `json:"zip"`
	Province     string  `json:"province"`
	Country      string  `json:"country"`
	LastName     string  `json:"last_name"`
	Address2     string  `json:"address2"`
	Company      string  `json:"company"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Name         string  `json:"name"`
	CountryCode  string  `json:"country_code"`
	ProvinceCode string  `json:"province_code"`
}

type OrderReturnCustomer struct {
	Id                        int                       `json:"id"`
	Email                     string                    `json:"email"`
	AcceptsMarketing          bool                      `json:"accepts_marketing"`
	CreatedAt                 string                    `json:"created_at"`
	UpdatedAt                 string                    `json:"updated_at"`
	FirstName                 string                    `json:"first_name"`
	LastName                  string                    `json:"last_name"`
	OrdersCount               int                       `json:"orders_count"`
	State                     string                    `json:"state"`
	TotalSpent                string                    `json:"total_spent"`
	LastOrderId               interface{}               `json:"last_order_id"`
	Note                      interface{}               `json:"note"`
	VerifiedEmail             bool                      `json:"verified_email"`
	MultipassIdentifier       interface{}               `json:"multipass_identifier"`
	TaxExempt                 bool                      `json:"tax_exempt"`
	Phone                     interface{}               `json:"phone"`
	Tags                      string                    `json:"tags"`
	LastOrderName             interface{}               `json:"last_order_name"`
	Currency                  string                    `json:"currency"`
	AcceptsMarketingUpdatedAt string                    `json:"accepts_marketing_updated_at"`
	MarketingOptInLevel       interface{}               `json:"marketing_opt_in_level"`
	TaxExemptions             []interface{}             `json:"tax_exemptions"`
	AdminGraphqpApiId         string                    `json:"admin_graphqp_api_id"`
	DefaultAddress            OrderReturnDefaultAddress `json:"default_address"`
}

type OrderReturnDefaultAddress struct {
	Id           int    `json:"id"`
	CustomerId   int    `json:"customer_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Company      string `json:"company"`
	Address1     string `json:"address1"`
	Address2     string `json:"address2"`
	City         string `json:"city"`
	Province     string `json:"province"`
	Country      string `json:"country"`
	Zip          string `json:"zip"`
	Phone        string `json:"phone"`
	Name         string `json:"name"`
	ProvinceCode string `json:"province_code"`
	CountryCode  string `json:"country_code"`
	CountryName  string `json:"country_name"`
	Default      bool   `json:"default"`
}

type OrderReturnLineItems struct {
	Id                         int                       `json:"id"`
	AdminGraphqlApiId          string                    `json:"admin_graphql_api_id"`
	FulfillableQuantity        int                       `json:"fulfillable_quantity"`
	FulfillmentService         string                    `json:"fulfillment_service"`
	FulfillmentStatus          interface{}               `json:"fulfillment_status"`
	GiftCard                   bool                      `json:"gift_card"`
	Grams                      int                       `json:"grams"`
	Name                       string                    `json:"name"`
	OriginLocation             OrderReturnOriginLocation `json:"origin_location"`
	Price                      string                    `json:"price"`
	PriceSet                   OrderReturnPriceSet       `json:"price_set"`
	ProductExists              bool                      `json:"product_exists"`
	ProductId                  interface{}               `json:"product_id"`
	Properties                 []interface{}             `json:"properties"`
	Quantity                   int                       `json:"quantity"`
	RequiresShipping           bool                      `json:"requires_shipping"`
	Sku                        string                    `json:"sku"`
	Taxable                    bool                      `json:"taxable"`
	Title                      string                    `json:"title"`
	TotalDiscount              string                    `json:"total_discount"`
	TotalDiscountSet           OrderReturnPriceSet       `json:"total_discount_set"`
	VariantId                  interface{}               `json:"variant_id"`
	VariantInventoryManagement interface{}               `json:"variant_inventory_management"`
	VariantTitle               string                    `json:"variant_title"`
	Vendor                     string                    `json:"vendor"`
	TaxLines                   []OrderReturnTaxLines     `json:"tax_lines"`
	Duties                     []interface{}             `json:"duties"`
	DiscountAllocations        []interface{}             `json:"discount_allocations"`
}

type OrderReturnOriginLocation struct {
	Id           int    `json:"id"`
	CountryCode  string `json:"country_code"`
	ProvinceCode string `json:"province_code"`
	Name         string `json:"name"`
	Address1     string `json:"address1"`
	Address2     string `json:"address2"`
	City         string `json:"city"`
	Zip          string `json:"zip"`
}

type OrderReturnTaxLines struct {
	ChannelLiable bool                `json:"channel_liable"`
	Price         string              `json:"price"`
	PriceSet      OrderReturnPriceSet `json:"price_set"`
	Rate          float64             `json:"rate"`
	Title         string              `json:"title"`
}

type OrderReturnPaymentDetails struct {
	CreditCardBin     string      `json:"credit_card_bin"`
	AvsResultCode     interface{} `json:"avs_result_code"`
	CvvResultCode     interface{} `json:"cvv_result_code"`
	CreditCardNumber  string      `json:"credit_card_number"`
	CreditCardCompany string      `json:"credit_card_company"`
}

type OrderReturnShippingLines struct {
	Id                            int                   `json:"id"`
	CarrierIdentifiert            interface{}           `json:"carrier_identifiert"`
	Code                          string                `json:"code"`
	DeliveryCategory              interface{}           `json:"delivery_category"`
	DiscountedPrice               string                `json:"discounted_price"`
	DiscountedPriceSet            OrderReturnPriceSet   `json:"discounted_price_set"`
	Phone                         interface{}           `json:"phone"`
	Price                         string                `json:"price"`
	PriceSet                      OrderReturnPriceSet   `json:"price_set"`
	RequestedFulfillmentServiceId interface{}           `json:"requested_fulfillment_service_id"`
	Source                        string                `json:"source"`
	Title                         string                `json:"title"`
	TaxLines                      []OrderReturnTaxLines `json:"tax_lines"`
}

// Orders is to get a list of all orders since the id
func Orders(id int, r Request) (OrderReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/orders.json?limit=200&since_id=%d", id), "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return OrderReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode OrderReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OrderReturn{}, err
	}

	// Return data
	return decode, err

}
