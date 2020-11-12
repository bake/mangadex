package mangadex

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Chapter contains a chapter.
type Chapter struct {
	PreviewChapter
	Status         string   `json:"status"`
	Pages          []string `json:"pages"`
	Server         string   `json:"server"`
	ServerFallback string   `json:"serverFallback"`
}

// Chapter gets a chapter by ID or hash.
func (c *Client) Chapter(id string) (Chapter, error) {
	raw, err := c.get("/chapter/"+id, nil)
	if err != nil {
		return Chapter{}, errors.Wrapf(err, "could not get chapter %s", id)
	}
	var res Chapter
	if err := json.Unmarshal(raw, &res); err != nil {
		return Chapter{}, errors.Wrapf(err, "could not unmarshal chapter %s", id)
	}
	return res, nil
}

// Images returns a slice of URLs to the chapters pages.
func (c Chapter) Images() []string {
	is := make([]string, len(c.Pages))
	for i, p := range c.Pages {
		is[i] = c.Server + c.Hash + "/" + p
	}
	return is
}
