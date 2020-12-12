package mangadex_test

import (
	"context"
	"testing"
)

func TestMangaChapters(t *testing.T) {
	tt := []struct {
		id  int
		err bool
	}{
		{0, true},
		{23279, false},
	}
	for _, tc := range tt {
		ctx := context.Background()
		cs, err := md.MangaChapters(ctx, tc.id, nil)
		if !tc.err && err != nil {
			t.Fatalf("expected manga %d to have chapters, got %q", tc.id, err)
		}
		if tc.err {
			continue
		}
		if len(cs) == 0 {
			t.Fatalf("expected manga %d to have more than 0 chapters", tc.id)
		}
	}
}
