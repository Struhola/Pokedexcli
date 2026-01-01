package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Pokemon_Information(Pokemon_ID any) (Pokemon, error) {
	if Pokemon_ID == nil {
		return Pokemon{}, nil
	}
	// This works whether 'v' is an int, a string, or even a boolean
	ID := fmt.Sprintf("%v", Pokemon_ID)
	url := baseURL + "/pokemon/" + ID

	//Cache
	if data, ok := c.cache.Get(url); ok {
		Pokemon_Resp := Pokemon{}
		err := json.Unmarshal(data, &Pokemon_Resp)
		if err != nil {
			return Pokemon{}, err
		}
		return Pokemon_Resp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	//cache add
	c.cache.Add(url, data)

	Pokemon_Resp := Pokemon{}
	err = json.Unmarshal(data, &Pokemon_Resp)
	if err != nil {
		return Pokemon{}, err
	}

	return Pokemon_Resp, nil
}
