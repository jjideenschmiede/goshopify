// Copyright 2023 J&J Ideenschmiede GmbH. All rights reserved.

package goshopify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// OrderReturn is to decode the order return form shopify
type OrderReturn struct {
	Orders []struct {
		Id                    int         `json:"id"`
		AdminGraphqlApiId     string      `json:"admin_graphql_api_id"`
		AppId                 int         `json:"app_id"`
		BrowserIp             string      `json:"browser_ip"`
		BuyerAcceptsMarketing bool        `json:"buyer_accepts_marketing"`
		CancelReason          interface{} `json:"cancel_reason"`
		CancelledAt           interface{} `json:"cancelled_at"`
		CartToken             string      `json:"cart_token"`
		CheckoutId            int         `json:"checkout_id"`
		CheckoutToken         string      `json:"checkout_token"`
		ClientDetails         struct {
			AcceptLanguage string      `json:"accept_language"`
			BrowserHeight  int         `json:"browser_height"`
			BrowserIp      string      `json:"browser_ip"`
			BrowserWidth   int         `json:"browser_width"`
			SessionHash    interface{} `json:"session_hash"`
			UserAgent      string      `json:"user_agent"`
		} `json:"client_details"`
		ClosedAt                interface{} `json:"closed_at"`
		ConfirmationNumber      string      `json:"confirmation_number"`
		Confirmed               bool        `json:"confirmed"`
		ContactEmail            string      `json:"contact_email"`
		CreatedAt               time.Time   `json:"created_at"`
		Currency                string      `json:"currency"`
		CurrentSubtotalPrice    string      `json:"current_subtotal_price"`
		CurrentSubtotalPriceSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"current_subtotal_price_set"`
		CurrentTotalDiscounts    string `json:"current_total_discounts"`
		CurrentTotalDiscountsSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"current_total_discounts_set"`
		CurrentTotalDutiesSet interface{} `json:"current_total_duties_set"`
		CurrentTotalPrice     string      `json:"current_total_price"`
		CurrentTotalPriceSet  struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"current_total_price_set"`
		CurrentTotalTax    string `json:"current_total_tax"`
		CurrentTotalTaxSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"current_total_tax_set"`
		CustomerLocale    string        `json:"customer_locale"`
		DeviceId          interface{}   `json:"device_id"`
		DiscountCodes     []interface{} `json:"discount_codes"`
		Email             string        `json:"email"`
		EstimatedTaxes    bool          `json:"estimated_taxes"`
		FinancialStatus   string        `json:"financial_status"`
		FulfillmentStatus interface{}   `json:"fulfillment_status"`
		LandingSite       string        `json:"landing_site"`
		LandingSiteRef    interface{}   `json:"landing_site_ref"`
		LocationId        interface{}   `json:"location_id"`
		Name              string        `json:"name"`
		Note              string        `json:"note"`
		NoteAttributes    []struct {
			Name  string      `json:"name"`
			Value interface{} `json:"value"`
		} `json:"note_attributes"`
		Number                 int         `json:"number"`
		OrderNumber            int         `json:"order_number"`
		OrderStatusUrl         string      `json:"order_status_url"`
		OriginalTotalDutiesSet interface{} `json:"original_total_duties_set"`
		PaymentGatewayNames    []string    `json:"payment_gateway_names"`
		Phone                  interface{} `json:"phone"`
		PresentmentCurrency    string      `json:"presentment_currency"`
		ProcessedAt            time.Time   `json:"processed_at"`
		ProcessingMethod       string      `json:"processing_method"`
		Reference              interface{} `json:"reference"`
		ReferringSite          string      `json:"referring_site"`
		SourceIdentifier       interface{} `json:"source_identifier"`
		SourceName             string      `json:"source_name"`
		SourceUrl              interface{} `json:"source_url"`
		SubtotalPrice          string      `json:"subtotal_price"`
		SubtotalPriceSet       struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"subtotal_price_set"`
		Tags              string        `json:"tags"`
		TaxLines          []interface{} `json:"tax_lines"`
		TaxesIncluded     bool          `json:"taxes_included"`
		Test              bool          `json:"test"`
		Token             string        `json:"token"`
		TotalDiscounts    string        `json:"total_discounts"`
		TotalDiscountsSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_discounts_set"`
		TotalLineItemsPrice    string `json:"total_line_items_price"`
		TotalLineItemsPriceSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_line_items_price_set"`
		TotalOutstanding string `json:"total_outstanding"`
		TotalPrice       string `json:"total_price"`
		TotalPriceSet    struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_price_set"`
		TotalPriceUsd         string `json:"total_price_usd"`
		TotalShippingPriceSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_shipping_price_set"`
		TotalTax    string `json:"total_tax"`
		TotalTaxSet struct {
			ShopMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"shop_money"`
			PresentmentMoney struct {
				Amount       string `json:"amount"`
				CurrencyCode string `json:"currency_code"`
			} `json:"presentment_money"`
		} `json:"total_tax_set"`
		TotalTipReceived string      `json:"total_tip_received"`
		TotalWeight      int         `json:"total_weight"`
		UpdatedAt        time.Time   `json:"updated_at"`
		UserId           interface{} `json:"user_id"`
		BillingAddress   struct {
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
		} `json:"billing_address"`
		Customer struct {
			Id                        int64         `json:"id"`
			Email                     string        `json:"email"`
			AcceptsMarketing          bool          `json:"accepts_marketing"`
			CreatedAt                 time.Time     `json:"created_at"`
			UpdatedAt                 time.Time     `json:"updated_at"`
			FirstName                 string        `json:"first_name"`
			LastName                  string        `json:"last_name"`
			OrdersCount               int           `json:"orders_count"`
			State                     string        `json:"state"`
			TotalSpent                string        `json:"total_spent"`
			LastOrderId               int64         `json:"last_order_id"`
			Note                      string        `json:"note"`
			VerifiedEmail             bool          `json:"verified_email"`
			MultipassIdentifier       interface{}   `json:"multipass_identifier"`
			TaxExempt                 bool          `json:"tax_exempt"`
			Phone                     interface{}   `json:"phone"`
			Tags                      string        `json:"tags"`
			LastOrderName             string        `json:"last_order_name"`
			Currency                  string        `json:"currency"`
			AcceptsMarketingUpdatedAt time.Time     `json:"accepts_marketing_updated_at"`
			MarketingOptInLevel       interface{}   `json:"marketing_opt_in_level"`
			TaxExemptions             []interface{} `json:"tax_exemptions"`
			AdminGraphqlApiId         string        `json:"admin_graphql_api_id"`
			DefaultAddress            struct {
				Id           int64  `json:"id"`
				CustomerId   int64  `json:"customer_id"`
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
			} `json:"default_address"`
		} `json:"customer"`
		DiscountApplications []interface{} `json:"discount_applications"`
		Fulfillments         []interface{} `json:"fulfillments"`
		LineItems            []struct {
			Id                  int64       `json:"id"`
			AdminGraphqlApiId   string      `json:"admin_graphql_api_id"`
			FulfillableQuantity int         `json:"fulfillable_quantity"`
			FulfillmentService  string      `json:"fulfillment_service"`
			FulfillmentStatus   interface{} `json:"fulfillment_status"`
			GiftCard            bool        `json:"gift_card"`
			Grams               int         `json:"grams"`
			Name                string      `json:"name"`
			OriginLocation      struct {
				Id           int64  `json:"id"`
				CountryCode  string `json:"country_code"`
				ProvinceCode string `json:"province_code"`
				Name         string `json:"name"`
				Address1     string `json:"address1"`
				Address2     string `json:"address2"`
				City         string `json:"city"`
				Zip          string `json:"zip"`
			} `json:"origin_location"`
			Price    string `json:"price"`
			PriceSet struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"price_set"`
			ProductExists bool  `json:"product_exists"`
			ProductId     int64 `json:"product_id"`
			Properties    []struct {
				Name  string      `json:"name"`
				Value interface{} `json:"value"`
			} `json:"properties"`
			Quantity         int    `json:"quantity"`
			RequiresShipping bool   `json:"requires_shipping"`
			Sku              string `json:"sku"`
			Taxable          bool   `json:"taxable"`
			Title            string `json:"title"`
			TotalDiscount    string `json:"total_discount"`
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
			VariantId                  int64  `json:"variant_id"`
			VariantInventoryManagement string `json:"variant_inventory_management"`
			VariantTitle               string `json:"variant_title"`
			Vendor                     string `json:"vendor"`
			TaxLines                   []struct {
				ChannelLiable bool   `json:"channel_liable"`
				Price         string `json:"price"`
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
				Rate  float64 `json:"rate"`
				Title string  `json:"title"`
			} `json:"tax_lines"`
			Duties              []interface{} `json:"duties"`
			DiscountAllocations []struct {
				Amount    string `json:"amount"`
				AmountSet struct {
					ShopMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"shop_money"`
					PresentmentMoney struct {
						Amount       string `json:"amount"`
						CurrencyCode string `json:"currency_code"`
					} `json:"presentment_money"`
				} `json:"amount_set"`
				DiscountApplicationIndex int `json:"discount_application_index"`
			} `json:"discount_allocations"`
		} `json:"line_items"`
		Refunds         []interface{} `json:"refunds"`
		ShippingAddress struct {
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
		} `json:"shipping_address"`
		ShippingLines []struct {
			Id                 int64       `json:"id"`
			CarrierIdentifier  interface{} `json:"carrier_identifier"`
			Code               string      `json:"code"`
			DeliveryCategory   interface{} `json:"delivery_category"`
			DiscountedPrice    string      `json:"discounted_price"`
			DiscountedPriceSet struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"discounted_price_set"`
			Phone    interface{} `json:"phone"`
			Price    string      `json:"price"`
			PriceSet struct {
				ShopMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"shop_money"`
				PresentmentMoney struct {
					Amount       string `json:"amount"`
					CurrencyCode string `json:"currency_code"`
				} `json:"presentment_money"`
			} `json:"price_set"`
			RequestedFulfillmentServiceId interface{}   `json:"requested_fulfillment_service_id"`
			Source                        string        `json:"source"`
			Title                         string        `json:"title"`
			TaxLines                      []interface{} `json:"tax_lines"`
			DiscountAllocations           []interface{} `json:"discount_allocations"`
		} `json:"shipping_lines"`
		PaymentDetails struct {
			CreditCardBin     string      `json:"credit_card_bin"`
			AvsResultCode     interface{} `json:"avs_result_code"`
			CvvResultCode     string      `json:"cvv_result_code"`
			CreditCardNumber  string      `json:"credit_card_number"`
			CreditCardCompany string      `json:"credit_card_company"`
		} `json:"payment_details,omitempty"`
	} `json:"orders"`
}

// OrderTransactionsReturn is to decode the order transactions return form shopify
type OrderTransactionsReturn struct {
	Transactions []struct {
		Id             int         `json:"id"`
		OrderId        int         `json:"order_id"`
		Kind           string      `json:"kind"`
		Gateway        string      `json:"gateway"`
		Status         string      `json:"status"`
		Message        string      `json:"message"`
		CreatedAt      time.Time   `json:"created_at"`
		Test           bool        `json:"test"`
		Authorization  string      `json:"authorization"`
		LocationId     interface{} `json:"location_id"`
		UserId         interface{} `json:"user_id"`
		ParentId       interface{} `json:"parent_id"`
		ProcessedAt    time.Time   `json:"processed_at"`
		DeviceId       interface{} `json:"device_id"`
		ErrorCode      interface{} `json:"error_code"`
		SourceName     string      `json:"source_name"`
		PaymentDetails struct {
			CreditCardBin             string      `json:"credit_card_bin"`
			AvsResultCode             interface{} `json:"avs_result_code"`
			CvvResultCode             interface{} `json:"cvv_result_code"`
			CreditCardNumber          string      `json:"credit_card_number"`
			CreditCardCompany         string      `json:"credit_card_company"`
			BuyerActionInfo           interface{} `json:"buyer_action_info"`
			CreditCardName            string      `json:"credit_card_name"`
			CreditCardWallet          interface{} `json:"credit_card_wallet"`
			CreditCardExpirationMonth int         `json:"credit_card_expiration_month"`
			CreditCardExpirationYear  int         `json:"credit_card_expiration_year"`
		} `json:"payment_details"`
		Receipt struct {
			PaidAmount string `json:"paid_amount"`
		} `json:"receipt"`
		Amount            string `json:"amount"`
		Currency          string `json:"currency"`
		PaymentId         string `json:"payment_id"`
		TotalUnsettledSet struct {
			PresentmentMoney struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"presentment_money"`
			ShopMoney struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"shop_money"`
		} `json:"total_unsettled_set"`
		AdminGraphqlApiId string `json:"admin_graphql_api_id"`
	} `json:"transactions"`
}

// Orders is to get a list of all orders since the id
func Orders(id int, r Request) (OrderReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/orders.json?limit=200&since_id=%d", id), http.MethodGet, nil}

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

// Order is to get the last order
func Order(r Request) (OrderReturn, error) {

	// Set config for new request
	c := Config{"/orders.json?limit=1", http.MethodGet, nil}

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

// OrderTransactions is to get all transactions from an order
func OrderTransactions(id int, r Request) (OrderTransactionsReturn, error) {

	// Set config for new request
	c := Config{fmt.Sprintf("/orders/%d/transactions.json", id), http.MethodGet, nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return OrderTransactionsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode OrderTransactionsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return OrderTransactionsReturn{}, err
	}

	// Return data
	return decode, err

}
