package pokeapi

import(
	"io"
	"encoding/json"
	"net/http"
)

func (c *Client) GetLocation(name string) (Location, error) {
	url := baseURL + "/location-area/" + name

	cached, ok := c.cache.Get(url)
	if ok {
		location := Location{}
		err := json.Unmarshal(cached, &location)
		if err != nil {
			return Location{}, err
		}
		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	location := Location{}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)
	return location, nil
}