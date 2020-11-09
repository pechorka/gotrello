package trello

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const trelloAPIBaseURL = `https://api.trello.com`

type Client struct {
	key   string
	token string
}

func NewClient(key, token string) *Client {
	return &Client{
		key:   key,
		token: token,
	}
}

// GetCards returns cards in list
func (c *Client) GetCards(listID string) ([]Card, error) {
	u, _ := url.Parse(trelloAPIBaseURL)
	u.Path = "/1/lists/" + listID + "/cards"
	values := u.Query()
	values.Add("key", c.key)
	values.Add("token", c.token)
	u.RawQuery = values.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error code %d", resp.StatusCode)
	}
	var cards []Card
	json.NewDecoder(resp.Body).Decode(&cards)
	return cards, err
}
