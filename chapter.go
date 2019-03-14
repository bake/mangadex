package mangadex

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Chapter contains information about a chapter. Some fields may be left empty
// when returned by the Manga function.
type Chapter struct {
	ID         json.Number `json:"id"`
	Volume     json.Number `json:"volume"`
	Chapter    json.Number `json:"chapter"`
	Title      string      `json:"title"`
	LangCode   string      `json:"lang_code"`
	GroupID    int         `json:"group_id"`
	GroupName  string      `json:"group_name"`
	GroupID2   int         `json:"group_id_2"`
	GroupName2 string      `json:"group_name_2"`
	GroupID3   int         `json:"group_id_3"`
	GroupName3 string      `json:"group_name_3"`
	Timestamp  int64       `json:"timestamp"`

	// The following fields are only filled when the chapter is requested
	// through the Client.Chapter function.
	LangName  string   `json:"lang_name"`
	Hash      string   `json:"hash"`
	Server    string   `json:"server"`
	Pages     []string `json:"page_array"`
	LongStrip int      `json:"long_strip"`
	Status    string   `json:"status"`
}

// Chapter fetches a mangas chapter.
func (c *Client) Chapter(id string) (Chapter, error) {
	raw, err := c.get(id, "chapter")
	if err != nil {
		return Chapter{}, errors.Wrapf(err, "could not get chapter %s", id)
	}
	var chapter Chapter
	if err := json.Unmarshal(raw, &chapter); err != nil {
		return Chapter{}, errors.Wrapf(err, "could not unmarshal chapter %s", id)
	}
	if chapter.Status != "OK" {
		return Chapter{}, errors.Errorf("could not get chapter %s: got unexpected status: %s", id, chapter.Status)
	}
	if chapter.Server == "" || chapter.Server[0] == '/' {
		chapter.Server = c.base + chapter.Server
	}
	return chapter, nil
}

func (ch Chapter) String() string { return ch.Title }

// Images returns a slice of URLs to the chapters pages.
func (ch Chapter) Images() []string {
	images := make([]string, len(ch.Pages))
	base := ch.Server + ch.Hash + "/"
	for i, page := range ch.Pages {
		images[i] = base + page
	}
	return images
}
