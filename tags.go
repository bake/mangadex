package mangadex

import (
	"context"
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// Tag contains information about a tag.
type Tag struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Group       string `json:"group"`
	Description string `json:"description"`
}

// TagsOptions contains options that can be passed to the endpoint.
type TagsOptions struct{}

// Tags all tags.
func (c *Client) Tags(ctx context.Context, opts *TagsOptions) ([]Tag, error) {
	values, err := query.Values(opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not encode options")
	}
	raw, err := c.get(ctx, "/tag", values)
	if err != nil {
		return nil, errors.Wrap(err, "could not get tags")
	}
	var res []Tag
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal tags")
	}
	return res, nil
}

// TagOptions contains options that can be passed to the endpoint.
type TagOptions struct{}

// Tag returns a tag.
func (c *Client) Tag(ctx context.Context, id string, opts *TagOptions) (Tag, error) {
	values, err := query.Values(opts)
	if err != nil {
		return Tag{}, errors.Wrap(err, "could not encode options")
	}
	raw, err := c.get(ctx, "/tag/"+id, values)
	if err != nil {
		return Tag{}, errors.Wrapf(err, "could not get tag %s", id)
	}
	var res Tag
	if err := json.Unmarshal(raw, &res); err != nil {
		return Tag{}, errors.Wrapf(err, "could not unmarshal tag %s", id)
	}
	return res, nil
}

func (t Tag) String() string { return t.Name }
