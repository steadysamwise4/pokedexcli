package pokeapi

import(
	"fmt"
	"io"
	"encoding/json"
	// "net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
    // 1. Determine which URL to use
    url := baseURL + "/location-area"
    if pageURL != nil {
        url = *pageURL
    }

    // 2. Create the request
    // 3. Do the request using c.httpClient.Get(url)
		res, err := c.httpClient.Get(url)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer res.Body.Close()
    // 4. Check the status code (ensure it's 200)
		if res.StatusCode > 299 {
			return RespShallowLocations{}, fmt.Errorf("bad status code: %v", res.StatusCode)
		}
    // 5. Read the body and Unmarshal it into your struct
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}
		locationsResp := RespShallowLocations{}
		err = json.Unmarshal(data, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
}