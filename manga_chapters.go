package mangadex

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

// PreviewChapter contains a chapter without images.
type PreviewChapter struct {
	ID         int     `json:"id"`
	Hash       string  `json:"hash"`
	MangaID    int     `json:"mangaId"`
	MangaTitle string  `json:"mangaTitle"`
	Volume     string  `json:"volume"`
	Chapter    string  `json:"chapter"`
	Title      string  `json:"title"`
	Language   string  `json:"language"`
	Groups     []Group `json:"groups"`
	Uploader   int     `json:"uploader"`
	Timestamp  int     `json:"timestamp"`
	Comments   int     `json:"comments"`
	Views      int     `json:"views"`
}

// MangaChapters returns chapters of a manga.
func (c *Client) MangaChapters(ctx context.Context, id string, query url.Values) ([]PreviewChapter, error) {
	raw, err := c.get(ctx, "/manga/"+id+"/chapters", query)
	if err != nil {
		return nil, errors.Wrapf(err, "could not get chapters of manga %s", id)
	}
	var res struct {
		Chapters []PreviewChapter `json:"chapters"`
		Groups   []Group          `json:"groups"`
	}
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, errors.Wrapf(err, "could not unmarshal chapters of manga %s", id)
	}
	// Merge groups into chapters.
	gs := map[int]string{}
	for _, g := range res.Groups {
		gs[g.ID] = g.Name
	}
	for i, c := range res.Chapters {
		for j, g := range c.Groups {
			res.Chapters[i].Groups[j].Name = gs[g.ID]
		}
	}
	return res.Chapters, nil
}

func (c PreviewChapter) String() string { return c.Title }
