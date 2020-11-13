package mangadex

import (
	"context"
	"net/url"

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
	Timestamp  int            `json:"timestamp"`
	Comments   int            `json:"comments"`
	Views      int            `json:"views"`
}

func (c PreviewChapter) String() string { return c.Title }

// MangaChapters returns chapters of a manga.
func (c *Client) MangaChapters(ctx context.Context, id string, query url.Values) ([]PreviewChapter, error) {
	cs, err := c.previewChapters(ctx, "/manga/"+id+"/chapters", query)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get chapters of manga %s", id)
	}
	return cs, nil
}
