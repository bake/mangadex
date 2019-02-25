package mangadex

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Chapter contains information about a chapter. Some fields may be left empty
// when returned by the Manga function.
type Chapter struct {
	ID         int    `json:"id"`
	Volume     string `json:"volume"`
	Chapter    string `json:"chapter"`
	Title      string `json:"title"`
	LangCode   string `json:"lang_code"`
	GroupID    int    `json:"group_id"`
	GroupName  string `json:"group_name"`
	GroupID2   int    `json:"group_id_2"`
	GroupName2 string `json:"group_name_2"`
	GroupID3   int    `json:"group_id_3"`
	GroupName3 string `json:"group_name_3"`
	Timestamp  int    `json:"timestamp"`

	// The following fields are only filled when the chapter is requested
	// through the Client.Chapter function.
	Hash       string   `json:"hash"`
	Server     string   `json:"server"`
	Pages      []string `json:"pages"`
	LongString int      `json:"long_string"`
	Status     string   `json:"status"`
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
	return chapter, nil
}
