package mangadex_test

import (
	"context"
	"testing"
)

func TestTags(t *testing.T) {
	tt := []struct {
		id, name string
		err      bool
	}{
		{"-", "", true},
		{"5", "Comedy", false},
		{"19", "Music", false},
	}
	for _, tc := range tt {
		ctx := context.Background()
		tag, err := md.Tag(ctx, tc.id, nil)
		if !tc.err && err != nil {
			t.Fatalf("expected tag %s to exist, got %q", tc.id, err)
		}
		if tc.err {
			continue
		}
		if tag.Name != tc.name {
			t.Fatalf("expected title to be %s, got %s", tc.name, tag.Name)
		}
	}
}
