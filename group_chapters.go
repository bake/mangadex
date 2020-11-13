package mangadex

import (
	"context"
	"net/url"

	"github.com/pkg/errors"
)

// GroupChapters returns chapters released by a group.
func (c *Client) GroupChapters(ctx context.Context, id string, query url.Values) ([]PreviewChapter, error) {
	cs, err := c.previewChapters(ctx, "/group/"+id+"/chapters", query)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get chapters of group %s", id)
	}
	return cs, nil
}
