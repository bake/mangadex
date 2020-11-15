package mangadex

import (
	"context"
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// Cover contains a URL to a volumes cover image.
type Cover struct {
	Volume string `json:"volume"`
	URL    string `json:"url"`
}

// MangaCoversOptions contains options that can be passed to the endpoint.
type MangaCoversOptions struct{}

// MangaCovers returns a list of covers belonging to a manga.
func (c *Client) MangaCovers(ctx context.Context, id string, opts *MangaCoversOptions) ([]Cover, error) {
	values, err := query.Values(opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not encode options")
	}
	raw, err := c.get(ctx, "/manga/"+id+"/covers", values)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get manga %s", id)
	}
	var res []Cover
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, errors.Wrapf(err, "could not unmarshal manga %s", id)
	}
	return res, nil
}
