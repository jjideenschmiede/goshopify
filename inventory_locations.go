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
	"time"
)

// InventoryLocationsReturn is to decode the json data
type InventoryLocationsReturn struct {
	Locations []struct {
		Id                    int         `json:"id"`
		Name                  string      `json:"name"`
		Address1              string      `json:"address1"`
		Address2              interface{} `json:"address2"`
		City                  string      `json:"city"`
		Zip                   string      `json:"zip"`
		Province              interface{} `json:"province"`
		Country               string      `json:"country"`
		Phone                 string      `json:"phone"`
		CreatedAt             time.Time   `json:"created_at"`
		UpdatedAt             time.Time   `json:"updated_at"`
		CountryCode           string      `json:"country_code"`
		CountryName           string      `json:"country_name"`
		ProvinceCode          interface{} `json:"province_code"`
		Legacy                bool        `json:"legacy"`
		Active                bool        `json:"active"`
		AdminGraphqlApiId     string      `json:"admin_graphql_api_id"`
		LocalizedCountryName  string      `json:"localized_country_name"`
		LocalizedProvinceName interface{} `json:"localized_province_name"`
	} `json:"locations"`
}

// InventoryLocations is to get a list of all locations
func InventoryLocations(r Request) (InventoryLocationsReturn, error) {

	// Set config for new request
	c := Config{"/locations.json", "GET", nil}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return InventoryLocationsReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode InventoryLocationsReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return InventoryLocationsReturn{}, err
	}

	// Return data
	return decode, err

}
