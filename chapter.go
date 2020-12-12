package mangadex

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
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

// ChapterOptions contains options that can be passed to the endpoint.
type ChapterOptions struct {
	Server   string `url:"server,omitempty"`    // Override location-based server assignment. Possible values: na, na2.
	Saver    bool   `url:"saver,omitempty"`     // Use low quality images (data saver).
	MarkRead bool   `url:"mark_read,omitempty"` // Mark the chapter as read.
}

// Chapter returns a chapter.
func (c *Client) Chapter(ctx context.Context, id int, opts *ChapterOptions) (Chapter, error) {
	values, err := query.Values(opts)
	if err != nil {
		return Chapter{}, errors.Wrap(err, "could not encode options")
	}
	raw, err := c.get(ctx, fmt.Sprintf("/chapter/%d", id), values)
	if err != nil {
		return Chapter{}, errors.Wrapf(err, "could not get chapter %d", id)
	}
	var res Chapter
	if err := json.Unmarshal(raw, &res); err != nil {
		return Chapter{}, errors.Wrapf(err, "could not unmarshal chapter %d", id)
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
