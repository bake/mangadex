package mangadex_test

import (
	"context"
	"testing"
)

func TestMangaCovers(t *testing.T) {
	tt := []struct {
		id          int
		covers      int
		volume, url string
		err         bool
	}{
		{0, 0, "", "", true},
		{23279, 12, "1", "https://mangadex.org/images/covers/23279v1.png?1588597138", false},
	}
	for _, tc := range tt {
		ctx := context.Background()
		cs, err := md.MangaCovers(ctx, tc.id, nil)
		if !tc.err && err != nil {
			t.Fatalf("expected manga %d to have covers, got %q", tc.id, err)
		}
		if tc.err {
			continue
		}
		if len(cs) < tc.covers {
			t.Fatalf("expected manga %d to have %d covers, got 0", tc.id, tc.covers)
		}
		c := cs[0]
		if c.Volume != tc.volume || c.String() != tc.url {
			t.Fatalf("expected the first cover of manga %d to be (%q, %q), got (%q, %q)", tc.id, tc.volume, tc.url, c.Volume, c.String())
		}
	}
}
