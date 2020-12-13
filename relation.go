package mangadex

import (
	"context"
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// Relation contains information about a relation.
type Relation struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	PairID int    `json:"pairId"`
}

func (r Relation) String() string { return r.Name }

// RelationsOptions contains options that can be passed to the endpoint.
type RelationsOptions struct{}

// Relations returns all manga relation types.
func (c *Client) Relations(ctx context.Context, opts *RelationsOptions) ([]Relation, error) {
	values, err := query.Values(opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not encode options")
	}
	raw, err := c.get(ctx, "/relations", values)
	if err != nil {
		return nil, errors.Wrap(err, "could not get relations")
	}
	res := map[int]Relation{}
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal relations")
	}
	var rs []Relation
	for _, r := range res {
		rs = append(rs, r)
	}
	return rs, nil
}
