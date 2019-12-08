package mangadex_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/bake/httpcache"
	"github.com/bake/httpcache/diskcache"
	"github.com/bake/mangadex"
)

var md *mangadex.Client

func init() {
	c := httpcache.New(diskcache.New("testdata", diskcache.NoExpiration))
	md = mangadex.New(
		mangadex.WithHTTPClient(c.Client()),
		mangadex.WithBase("https://mangadex.org/"),
		mangadex.WithPath("api/"),
	)
}

func ExampleManga() {
	m, _, err := md.Manga("23279")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s by %s", m, m.Author)
	// Output: Wonder Cat Kyuu-chan by Nitori Sasami
}

func TestManga(t *testing.T) {
	tt := []struct {
		id, title, cid string
		err            bool
	}{
		{"23279", "Wonder Cat Kyuu-chan", "26457", false},
		{"0", "", "", true},
	}
	for _, tc := range tt {
		m, cs, err := md.Manga(tc.id)
		if !tc.err && err != nil {
			t.Fatalf("expected manga to exist, got %q", err)
		}
		if tc.err {
			continue
		}
		if m.Title != tc.title {
			t.Fatalf("expected title to be %s, got %s", tc.title, m.Title)
		}
		if cs[0].ID.String() != tc.cid {
			t.Fatalf("expected first chapter id to be %s, got %s", tc.cid, cs[0].ID)
		}
	}
}

func ExampleChapter() {
	c, err := md.Chapter("517244")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s (Volume %s, Chapter %s)", c, c.Volume, c.Chapter)
	// Output: Cool Day (Volume 3, Chapter 253)
}

func TestChapter(t *testing.T) {
	tt := []struct {
		id, title string
		images    int
		err       bool
	}{
		{"517244", "Cool Day", 1, false},
		{"0", "", 0, true},
	}
	for _, tc := range tt {
		c, err := md.Chapter(tc.id)
		if !tc.err && err != nil {
			t.Fatalf("expected chapter to exist, got %q", err)
		}
		if tc.err {
			continue
		}
		if c.Title != tc.title {
			t.Fatalf("expected title to be %s, got %s", tc.title, c.Title)
		}
		if len(c.Images()) != tc.images {
			t.Fatalf("expected chapter to have %d images, not %d", tc.images, len(c.Images()))
		}
	}
}
