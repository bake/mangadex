package mangadex

import (
	"context"
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// Follow contains information about a follow.
type Follow struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// FollowsOptions contains options that can be passed to the endpoint.
type FollowsOptions struct{}

// Follows returns all follow types.
func (c *Client) Follows(ctx context.Context, opts *FollowsOptions) ([]Follow, error) {
	values, err := query.Values(opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not encode options")
	}
	raw, err := c.get(ctx, "/follows", values)
	if err != nil {
		return nil, errors.Wrap(err, "could not get follows")
	}
	var res []Follow
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal follows")
	}
	return res, nil
}
