package mangadex_test

import (
	"context"
	"testing"
)

func TestMangaChapters(t *testing.T) {
	tt := []struct {
		mid string
		err bool
	}{
		{"23279", false},
		{"0", true},
	}
	for _, tc := range tt {
		ctx := context.Background()
		cs, err := md.MangaChapters(ctx, tc.mid)
		if !tc.err && err != nil {
			t.Fatalf("expected manga %s to have chapters, got %q", tc.mid, err)
		}
		if tc.err {
			continue
		}
		if len(cs) == 0 {
			t.Fatalf("expected manga %s to have chapters", tc.mid)
		}
	}
}
