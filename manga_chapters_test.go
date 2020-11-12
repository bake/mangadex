package mangadex_test

import (
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
		cs, err := md.MangaChapters(tc.mid)
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
