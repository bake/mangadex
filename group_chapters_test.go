package mangadex_test

import (
	"context"
	"testing"

	"github.com/bake/mangadex/v2"
)

func TestGroupChapters(t *testing.T) {
	tt := []struct {
		id          string
		page, limit int
		err         bool
	}{
		{"0", 0, 0, true},
		{"27", 10, 10, false},
	}
	for _, tc := range tt {
		ctx := context.Background()
		opts := mangadex.GroupChaptersOptions{Page: tc.page, Limit: tc.limit}
		cs, err := md.GroupChapters(ctx, tc.id, &opts)
		if !tc.err && err != nil {
			t.Fatalf("expected group %s to have chapters, got %q", tc.id, err)
		}
		if tc.err {
			continue
		}
		if len(cs) != tc.limit {
			t.Fatalf("limited number of chapters to %d, got %d", tc.limit, len(cs))
		}
	}
}
