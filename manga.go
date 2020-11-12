package mangadex

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

// Manga contains information about a given manga.
type Manga struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	AltTitles   []string `json:"altTitles"`
	Description string   `json:"description"`
	Artist      []string `json:"artist"`
	Author      []string `json:"author"`
	Publication struct {
		Language    string `json:"language"`
		Status      int    `json:"status"`
		Demographic int    `json:"demographic"`
	} `json:"publication"`
	Tags        []int             `json:"tags"`
	LastChapter interface{}       `json:"lastChapter"`
	LastVolume  interface{}       `json:"lastVolume"`
	IsHentai    bool              `json:"isHentai"`
	Links       map[string]string `json:"links"`
	Relations   []struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Type     int    `json:"type"`
		IsHentai bool   `json:"isHentai"`
	} `json:"relations"`
	Rating struct {
		Bayesian float64 `json:"bayesian"`
		Mean     float64 `json:"mean"`
		Users    int     `json:"users"`
	} `json:"rating"`
	Views        int    `json:"views"`
	Follows      int    `json:"follows"`
	Comments     int    `json:"comments"`
	LastUploaded int    `json:"lastUploaded"`
	MainCover    string `json:"mainCover"`
}

// Manga fetches a manga.
func (c *Client) Manga(ctx context.Context, id string) (Manga, error) {
	raw, err := c.get(ctx, "/manga/"+id, nil)
	if err != nil {
		return Manga{}, errors.Wrapf(err, "could not get manga %s", id)
	}
	var res Manga
	if err := json.Unmarshal(raw, &res); err != nil {
		return Manga{}, errors.Wrapf(err, "could not unmarshal manga %s", id)
	}
	return res, nil
}

func (m Manga) String() string { return m.Title }
