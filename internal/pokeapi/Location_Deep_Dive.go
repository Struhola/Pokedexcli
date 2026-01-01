package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Location_Deep_Dive(Location_ID any) (LocationAreaById, error) {
	if Location_ID == nil {
		return LocationAreaById{}, nil
	}
	// This works whether 'v' is an int, a string, or even a boolean
	ID := fmt.Sprintf("%v", Location_ID)
	url := baseURL + "/location-area/" + ID

	//Cache
	if data, ok := c.cache.Get(url); ok {
		Location_Deep_Dive_Resp := LocationAreaById{}
		err := json.Unmarshal(data, &Location_Deep_Dive_Resp)
		if err != nil {
			return LocationAreaById{}, err
		}
		return Location_Deep_Dive_Resp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaById{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaById{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaById{}, err
	}

	//cache add
	c.cache.Add(url, data)

	Location_Deep_Dive_Resp := LocationAreaById{}
	err = json.Unmarshal(data, &Location_Deep_Dive_Resp)
	if err != nil {
		return LocationAreaById{}, err
	}

	return Location_Deep_Dive_Resp, nil
}
