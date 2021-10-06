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

// InventoryLevelBody is to structure the data
type InventoryLevelBody struct {
	LocationId      int `json:"location_id"`
	InventoryItemId int `json:"inventory_item_id"`
	Available       int `json:"available"`
}

// InventoryLevelReturn is to decode the json data
type InventoryLevelReturn struct {
	InventoryLevel struct {
		InventoryItemId   int       `json:"inventory_item_id"`
		LocationId        int       `json:"location_id"`
		Available         int       `json:"available"`
		UpdatedAt         time.Time `json:"updated_at"`
		AdminGraphqlApiId string    `json:"admin_graphql_api_id"`
	} `json:"inventory_level"`
}

// InventoryLevels is to get a list of all locations
func InventoryLevels(body InventoryLevelBody, r Request) (InventoryLevelReturn, error) {

	// Convert body
	convert, err := json.Marshal(body)
	if err != nil {
		return InventoryLevelReturn{}, err
	}

	// Set config for new request
	c := Config{"/inventory_levels/set.json", "POST", convert}

	// Send request
	response, err := c.Send(r)
	if err != nil {
		return InventoryLevelReturn{}, err
	}

	// Close request
	defer response.Body.Close()

	// Decode data
	var decode InventoryLevelReturn

	err = json.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return InventoryLevelReturn{}, err
	}

	// Return data
	return decode, err

}
