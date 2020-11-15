package mangadex

import (
	"context"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// PreviewChapter contains a chapter without images.
type PreviewChapter struct {
	ID         int            `json:"id"`
	Hash       string         `json:"hash"`
	MangaID    int            `json:"mangaId"`
	MangaTitle string         `json:"mangaTitle"`
	Volume     string         `json:"volume"`
	Chapter    string         `json:"chapter"`
	Title      string         `json:"title"`
	Language   string         `json:"language"`
	Groups     []PreviewGroup `json:"groups"`
	Uploader   int            `json:"uploader"`
	Timestamp  Timestamp      `json:"timestamp"`
	Comments   int            `json:"comments"`
	Views      int            `json:"views"`
}

func (c PreviewChapter) String() string { return c.Title }

// MangaChaptersOptions contains options that can be passed to the endpoint.
type MangaChaptersOptions struct {
	Page  int `url:"p,omitempty"`
	Limit int `url:"limit,omitempty"`
}

// MangaChapters returns partial information about the chapters belonging to a
// manga.
func (c *Client) MangaChapters(ctx context.Context, id string, opts *MangaChaptersOptions) ([]PreviewChapter, error) {
	values, err := query.Values(opts)
	if err != nil {
		return nil, errors.Wrap(err, "could not encode options")
	}
	cs, err := c.previewChapters(ctx, "/manga/"+id+"/chapters", values)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get chapters of manga %s", id)
	}
	return cs, nil
}
