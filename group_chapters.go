package mangadex

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// GroupChaptersOptions contains options that can be passed to the endpoint.
type GroupChaptersOptions struct {
	Page  int `url:"p,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// GroupChapters returns partial information about the chapters belonging to the
// group.
func (c *Client) GroupChapters(ctx context.Context, id int, opts *GroupChaptersOptions) ([]PreviewChapter, error) {
	values, err := query.Values(opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not encode options")
	}
	cs, err := c.previewChapters(ctx, fmt.Sprintf("/group/%d/chapters", id), values)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get chapters of group %d", id)
	}
	return cs, nil
}
