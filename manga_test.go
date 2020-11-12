package mangadex_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/bake/httpcache"
	"github.com/bake/httpcache/diskcache"
	"github.com/bake/mangadex/v2"
)

var md *mangadex.Client

func init() {
	c := httpcache.New(diskcache.New("testdata", diskcache.NoExpiration))
	md = mangadex.New(
		mangadex.WithHTTPClient(c.Client()),
		mangadex.WithBase("https://mangadex.org"),
		mangadex.WithPath("/api/v2"),
	)
}

func ExampleManga() {
	ctx := context.TODO()
	m, err := md.Manga(ctx, "23279")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s by %s", m, m.Author[0])
	// Output: Wonder Cat Kyuu-chan by Nitori Sasami
}

func TestManga(t *testing.T) {
	tt := []struct {
		mid   string
		title string
		err   bool
	}{
		{"23279", "Wonder Cat Kyuu-chan", false},
		{"45112", "Tensei Kenja wa Musume to Kurasu", false},
		{"0", "", true},
	}
	for _, tc := range tt {
		ctx := context.Background()
		m, err := md.Manga(ctx, tc.mid)
		if !tc.err && err != nil {
			t.Fatalf("expected manga to exist, got %q", err)
		}
		if tc.err {
			continue
		}
		if m.Title != tc.title {
			t.Fatalf("expected title to be %s, got %s", tc.title, m.Title)
		}
	}
}
