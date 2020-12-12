package mangadex_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/bake/mangadex/v2"
)

func ExampleChapter() {
	ctx := context.TODO()
	opts := mangadex.ChapterOptions{
		Saver: true, // Use low quality images.
	}
	c, err := md.Chapter(ctx, 517244, &opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s (Volume %s, Chapter %s)", c, c.Volume, c.Chapter)
	// Output: Cool Day (Volume 3, Chapter 253)
}

func TestChapter(t *testing.T) {
	tt := []struct {
		id    int
		title string
		pages int
		err   bool
	}{
		{0, "", 0, true},
		{517244, "Cool Day", 1, false},
	}
	for _, tc := range tt {
		ctx := context.Background()
		c, err := md.Chapter(ctx, tc.id, nil)
		if !tc.err && err != nil {
			t.Fatalf("expected chapter to exist, got %q", err)
		}
		if tc.err {
			continue
		}
		if c.Title != tc.title {
			t.Fatalf("expected title to be %s, got %s", tc.title, c.Title)
		}
		if len(c.Pages) != tc.pages {
			t.Fatalf("expected chapter to have %d pages, not %d", tc.pages, len(c.Pages))
		}
	}
}
