package mangadex_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/bake/mangadex/v2"
)

func TestGroup(t *testing.T) {
	tt := []struct {
		id   int
		name string
		err  bool
	}{
		{0, "", true},
		{27, "Helvetica Scans", false},
	}
	for _, tc := range tt {
		ctx := context.Background()
		g, err := md.Group(ctx, tc.id, nil)
		if !tc.err && err != nil {
			t.Fatalf("expected group %d to not exist, got %q", tc.id, err)
		}
		if tc.err {
			continue
		}
		if g.String() != tc.name {
			t.Fatalf("expected group %d to have the name %q, got %q", tc.id, tc.name, g.String())
		}
	}
}

func TestPreviewGroup(t *testing.T) {
	tt := []struct {
		data  []byte
		group mangadex.PreviewGroup
		err   bool
	}{
		{[]byte(""), mangadex.PreviewGroup{}, true},
		{[]byte("314"), mangadex.PreviewGroup{314, ""}, false},
		{[]byte("{\"id\":314,\"name\":\"Pi\"}"), mangadex.PreviewGroup{314, "Pi"}, false},
	}
	for _, tc := range tt {
		var g mangadex.PreviewGroup
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
