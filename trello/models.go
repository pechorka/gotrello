package trello

type Card struct {
	ID     string `json:"id"`
	Closed bool   `json:"closed"`
	Desc   string `json:"desc"`
	Name   string `json:"name"`
	URL    string `json:"url"`
}
