package mangadex_test

import (
	"testing"

	"github.com/bakerolls/httpcache"
	"github.com/bakerolls/httpcache/diskcache"
	"github.com/bakerolls/mangadex"
)

var md *mangadex.Client

func init() {
	c := httpcache.New(diskcache.New("testdata", diskcache.NoExpiration))
	md = mangadex.New(mangadex.WithHTTPClient(c.Client()))
}

func TestManga(t *testing.T) {
	tt := []struct{ id, title string }{
		{"23279", "Wonder Cat Kyuu-chan"},
	}
	for _, tc := range tt {
		m, _, err := md.Manga("23279")
		if err != nil {
			t.Fatal(err)
		}
		if m.Title != tc.title {
			t.Fatalf("expected title to be %s, got %s", tc.title, m.Title)
		}
	}
}

func TestChapter(t *testing.T) {
	tt := []struct{ id, title string }{
		{"517244", "Cool Day"},
	}
	for _, tc := range tt {
		c, err := md.Chapter(tc.id)
		if err != nil {
			t.Fatal(err)
		}
		if c.Title != tc.title {
			t.Fatalf("expected title to be %s, got %s", tc.title, c.Title)
		}
	}
}
