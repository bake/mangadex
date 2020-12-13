package mangadex

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// Cover contains a URL to a volumes cover image.
type Cover struct {
	Volume string `json:"volume"`
	URL    string `json:"url"`
}

func (c Cover) String() string { return c.URL }

// MangaCoversOptions contains options that can be passed to the endpoint.
type MangaCoversOptions struct{}

// MangaCovers returns a list of covers belonging to a manga.
func (c *Client) MangaCovers(ctx context.Context, id int, opts *MangaCoversOptions) ([]Cover, error) {
	values, err := query.Values(opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not encode options")
	}
	raw, err := c.get(ctx, fmt.Sprintf("/manga/%d/covers", id), values)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get manga %d", id)
	}
	var res []Cover
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, errors.Wrapf(err, "could not unmarshal manga %d", id)
	}
	return res, nil
}
