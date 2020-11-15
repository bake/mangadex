package mangadex

import (
	"context"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// UserChaptersOptions contains options that can be passed to the endpoint.
type UserChaptersOptions struct {
	Page  int `url:"p,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// UserChapters partial information about the chapters uploaded by the user.
func (c *Client) UserChapters(ctx context.Context, id string, opts *UserChaptersOptions) ([]PreviewChapter, error) {
	values, err := query.Values(opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not encode options")
	}
	cs, err := c.previewChapters(ctx, "/user/"+id+"/chapters", values)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get chapters of user %s", id)
	}
	return cs, nil
}
