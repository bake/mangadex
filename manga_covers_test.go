package mangadex_test

import (
	"context"
	"testing"
)

func TestMangaCovers(t *testing.T) {
	tt := []struct {
		mid         string
		covers      int
		volume, url string
		err         bool
	}{
		{"23279", 12, "1", "https://mangadex.org/images/covers/23279v1.png?1588597138", false},
		{"0", 0, "", "", true},
	}
	for _, tc := range tt {
		ctx := context.Background()
		cs, err := md.MangaCovers(ctx, tc.mid)
		if !tc.err && err != nil {
			t.Fatalf("expected manga %s to have covers, got %q", tc.mid, err)
		}
		if tc.err {
			continue
		}
		if len(cs) < tc.covers {
			t.Fatalf("expected manga %s to have %d covers, got 0", tc.mid, tc.covers)
		}
		c := cs[0]
		if c.Volume != tc.volume || c.URL != tc.url {
			t.Fatalf("expected the first cover of manga %s to be (%q, %q), got (%q, %q)", tc.mid, tc.volume, tc.url, c.Volume, c.URL)
		}
	}
}
