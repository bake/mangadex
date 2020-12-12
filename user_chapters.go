package mangadex

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// UserChaptersOptions contains options that can be passed to the endpoint.
type UserChaptersOptions struct {
	Page  int `url:"p,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// UserChapters partial information about the chapters uploaded by the user.
func (c *Client) UserChapters(ctx context.Context, id int, opts *UserChaptersOptions) ([]PreviewChapter, error) {
	values, err := query.Values(opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not encode options")
	}
	cs, err := c.previewChapters(ctx, fmt.Sprintf("/user/%d/chapters", id), values)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get chapters of user %d", id)
	}
	return cs, nil
}
