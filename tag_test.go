package mangadex_test

import (
	"context"
	"testing"
)

func TestTags(t *testing.T) {
	tt := []struct {
		id   int
		name string
	}{
		{5, "Comedy"},
		{19, "Music"},
	}
	ctx := context.Background()
	tags, err := md.Tags(ctx, nil)
	if err != nil {
		t.Fatalf("could not get tags: %v", err)
	}
	for _, tc := range tt {
		var found bool
		for _, t := range tags {
			if t.ID == tc.id && t.String() == tc.name {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected tag %d (%s) to exist", tc.id, tc.name)
		}
	}
}

func TestTag(t *testing.T) {
	tt := []struct {
		id   int
		name string
		err  bool
	}{
		{5, "Comedy", false},
		{19, "Music", false},
	}
	for _, tc := range tt {
		ctx := context.Background()
		tag, err := md.Tag(ctx, tc.id, nil)
		if !tc.err && err != nil {
			t.Fatalf("expected tag %d to exist, got %q", tc.id, err)
		}
		if tc.err {
			continue
		}
		if tag.Name != tc.name {
			t.Fatalf("expected title to be %s, got %s", tc.name, tag.Name)
		}
	}
}
