package mangadex

import (
	"encoding/json"
	"strconv"
)

// Group contains information about a group.
type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (g Group) String() string { return g.Name }

// UnmarshalJSON implements the json.Unmarshaler interface allowing for
// interpreting []int as []Group while ignoring names. This is necessary since
// groups in /manga/{id}/chapters only contain IDs.
func (g *Group) UnmarshalJSON(d []byte) error {
	if id, err := strconv.Atoi(string(d)); err == nil {
		g.ID = id
		return nil
	}
	type group Group
	if err := json.Unmarshal(d, (*group)(g)); err != nil {
		return err
	}
	return nil
}
