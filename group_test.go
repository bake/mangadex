package mangadex_test

import (
	"encoding/json"
	"testing"

	"github.com/bake/mangadex/v2"
)

func TestGroup(t *testing.T) {
	tt := []struct {
		data  []byte
		group mangadex.Group
		err   bool
	}{
		{[]byte(""), mangadex.Group{}, true},
		{[]byte("314"), mangadex.Group{314, ""}, false},
		{[]byte("{\"id\":314,\"name\":\"Pi\"}"), mangadex.Group{314, "Pi"}, false},
	}
	for _, tc := range tt {
		var g mangadex.Group
		err := json.Unmarshal(tc.data, &g)
		if !tc.err && err != nil {
			t.Fatalf("expected group unmarshal to %v, got %q", tc.err, err)
		}
		if tc.err {
			continue
		}
		if g != tc.group {
			t.Fatalf("expected group to be %#v, got %#v", tc.group, g)
		}
	}
}
