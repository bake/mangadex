package mangadex

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

// Timestamp contains a time.
type Timestamp struct{ time.Time }

// UnmarshalJSON interprets int64 as a Unix timestamp and unmarshals it into
// time.Time.
func (t *Timestamp) UnmarshalJSON(d []byte) error {
	ts, err := strconv.ParseInt(string(d), 10, 64)
	if err != nil {
		return errors.Wrap(err, "could not parse timestamp")
	}
	t.Time = time.Unix(ts, 0)
	return nil
}

// User contains information about a given user.
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	LevelID   int       `json:"levelId"`
	Joined    Timestamp `json:"joined"`
	LastSeen  Timestamp `json:"lastSeen"`
	Website   string    `json:"website"`
	Biography string    `json:"biography"`
	Views     int       `json:"views"`
	Uploads   int       `json:"uploads"`
	Premium   bool      `json:"premium"`
	MdAtHome  int       `json:"mdAtHome"`
	Avatar    string    `json:"avatar"`
}

// UserOptions contains options that can be passed to the endpoint.
type UserOptions struct{}

// User returns a user.
func (c *Client) User(ctx context.Context, id int, opts *UserOptions) (User, error) {
	values, err := query.Values(opts)
	if err != nil {
		return User{}, errors.Wrap(err, "could not encode options")
	}
	raw, err := c.get(ctx, fmt.Sprintf("/user/%d", id), values)
	if err != nil {
		return User{}, errors.Wrapf(err, "could not get user %d", id)
	}
	var res User
	if err := json.Unmarshal(raw, &res); err != nil {
		return User{}, errors.Wrapf(err, "could not unmarshal user %d", id)
	}
	return res, nil
}

func (u User) String() string { return u.Username }
