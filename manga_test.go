package mangadex_test

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func ExampleManga() {
	ctx := context.TODO()
	m, err := md.Manga(ctx, "23279", nil)
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
		m, err := md.Manga(ctx, tc.mid, nil)
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

func TestMangaError(t *testing.T) {
	tt := []struct {
		mid     string
		message string
		err     bool
	}{
		{"", "", true},
		{"0", "", true},
		{"test", "", true},
	}
	for _, tc := range tt {
		ctx := context.Background()
		m, err := md.Manga(ctx, tc.mid, nil)
		if !tc.err && err != nil {
			t.Fatalf("expected manga %s to not exist, got %q", tc.mid, m)
		}
	}
}
