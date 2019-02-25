package mangadex

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type mangaResponse struct {
	Manga    Manga                   `json:"manga"`
	Chapters map[json.Number]Chapter `json:"chapter"`
	Status   string                  `json:"status"`
}

// Manga contains information about a given manga.
type Manga struct {
	ID          json.Number        `json:"id"`
	CoverURL    string             `json:"cover_url"`
	Description string             `json:"description"`
	Title       string             `json:"title"`
	Artist      string             `json:"artist"`
	Author      string             `json:"author"`
	Status      int                `json:"status"`
	Genres      []int              `json:"genres"`
	LastChapter string             `json:"last_chapter"`
	LangName    string             `json:"lang_name"`
	LangFlag    string             `json:"lang_flag"`
	Hentai      int                `json:"hentai"`
	Links       map[string]string  `json:"links"`
	Chapters    map[string]Chapter `json:"chapters"`
}

// Manga fetches a manga. The returned chapter slice is a second representation
// of the mangas Chapters map.
func (c *Client) Manga(id string) (Manga, []Chapter, error) {
	raw, err := c.get(id, "manga")
	if err != nil {
		return Manga{}, nil, errors.Wrapf(err, "could not get manga %s", id)
	}
	var res mangaResponse
	if err := json.Unmarshal(raw, &res); err != nil {
		return Manga{}, nil, errors.Wrapf(err, "could not unmarshal manga %s", id)
	}
	if res.Status != "OK" {
		return Manga{}, nil, errors.Errorf("could not get manga %s: got unexpected status: %s", id, res.Status)
	}
	res.Manga.ID = json.Number(id)
	var chapters []Chapter
	for id, chapter := range res.Chapters {
		chapter.ID = id
		chapters = append(chapters, chapter)
	}
	return res.Manga, chapters, nil
}

func (m Manga) String() string { return m.Title }
