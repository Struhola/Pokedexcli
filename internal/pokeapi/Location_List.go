package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) List_Locations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	//Cache
	if data, ok := c.cache.Get(url); ok {
		locationsResp := Locations{}
		err := json.Unmarshal(data, &locationsResp)
		if err != nil {
			return Locations{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}

	//cache add
	c.cache.Add(url, data)

	locationsResp := Locations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return Locations{}, err
	}

	return locationsResp, nil
}
