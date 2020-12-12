package mangadex

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// Group contains information about a given group.
type Group struct {
	PreviewGroup
	AltNames    string    `json:"altNames"`
	Language    string    `json:"language"`
	Leader      Member    `json:"leader"`
	Members     []Member  `json:"members"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	Discord     string    `json:"discord"`
	IRCServer   string    `json:"ircServer"`
	IRCChannel  string    `json:"ircChannel"`
	Email       string    `json:"email"`
	Founded     string    `json:"founded"`
	Likes       int       `json:"likes"`
	Follows     int       `json:"follows"`
	Views       int       `json:"views"`
	Chapters    int       `json:"chapters"`
	ThreadID    int       `json:"threadId"`
	ThreadPosts int       `json:"threadPosts"`
	IsLocked    bool      `json:"isLocked"`
	IsInactive  bool      `json:"isInactive"`
	Delay       int       `json:"delay"`
	LastUpdated Timestamp `json:"lastUpdated"`
	Banner      string    `json:"banner"`
}

// GroupOptions contains options that can be passed to the endpoint.
type GroupOptions struct{}

// Group returns a group.
func (c *Client) Group(ctx context.Context, id int, opts *GroupOptions) (Group, error) {
	values, err := query.Values(opts)
	if err != nil {
		return Group{}, errors.Wrap(err, "could not encode options")
	}
	raw, err := c.get(ctx, fmt.Sprintf("/group/%d", id), values)
	if err != nil {
		return Group{}, errors.Wrapf(err, "could not get group %d", id)
	}
	var res Group
	if err := json.Unmarshal(raw, &res); err != nil {
		return Group{}, errors.Wrapf(err, "could not unmarshal group %d", id)
	}
	return res, nil
}

// Member belogs to a group.
type Member struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// PreviewGroup contains limited information about a group.
type PreviewGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (g PreviewGroup) String() string { return g.Name }

// UnmarshalJSON implements the json.Unmarshaler interface allowing for
// interpreting []int as []Group while ignoring names. This is necessary since
// groups in /manga/{id}/chapters only contain IDs.
func (g *PreviewGroup) UnmarshalJSON(d []byte) error {
	if id, err := strconv.Atoi(string(d)); err == nil {
		g.ID = id
		return nil
	}
	type group PreviewGroup
	if err := json.Unmarshal(d, (*group)(g)); err != nil {
		return err
	}
	return nil
}
