package mangadex

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/pkg/errors"
)

// Chapter contains a chapter.
type Chapter struct {
	PreviewChapter
	Status         string   `json:"status"`
	Pages          []string `json:"pages"`
	Server         string   `json:"server"`
	ServerFallback string   `json:"serverFallback"`
}

// Images returns a slice of URLs to the chapters pages.
func (c Chapter) Images() []string {
	is := make([]string, len(c.Pages))
	for i, p := range c.Pages {
		is[i] = c.Server + c.Hash + "/" + p
	}
	return is
}

// Chapter gets a chapter by ID or hash.
func (c *Client) Chapter(ctx context.Context, id string, query url.Values) (Chapter, error) {
	raw, err := c.get(ctx, "/chapter/"+id, query)
	if err != nil {
		return Chapter{}, errors.Wrapf(err, "could not get chapter %s", id)
	}
	var res Chapter
	if err := json.Unmarshal(raw, &res); err != nil {
		return Chapter{}, errors.Wrapf(err, "could not unmarshal chapter %s", id)
	}
	return res, nil
}

func (c *Client) previewChapters(ctx context.Context, endpoint string, query url.Values) ([]PreviewChapter, error) {
	raw, err := c.get(ctx, endpoint, query)
	if err != nil {
		return nil, errors.Wrap(err, "could not get chapters")
	}
	var res struct {
		Chapters []PreviewChapter `json:"chapters"`
		Groups   []PreviewGroup   `json:"groups"`
	}
	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, errors.Wrap(err, "could not unmarshal chapters")
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
