package mangadex

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

// Cover contains a URL to a volumes cover image.
type Cover struct {
	Volume string `json:"volume"`
	URL    string `json:"url"`
}

// MangaCovers returns a slice of manga covers.
func (c *Client) MangaCovers(ctx context.Context, id string) ([]Cover, error) {
	raw, err := c.get(ctx, "/manga/"+id+"/covers", nil)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get manga %s", id)
	}
	var res []Cover
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, errors.Wrapf(err, "could not unmarshal manga %s", id)
	}
	return res, nil
}
