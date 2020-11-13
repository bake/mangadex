package mangadex_test

import (
	"context"
	"testing"
)

func TestGroupChapters(t *testing.T) {
	tt := []struct {
		id  string
		err bool
	}{
		{"0", true},
		{"27", false},
	}
	for _, tc := range tt {
		ctx := context.Background()
		cs, err := md.GroupChapters(ctx, tc.id, nil)
		if !tc.err && err != nil {
			t.Fatalf("expected group %s to have chapters, got %q", tc.id, err)
		}
		if tc.err {
			continue
		}
		if len(cs) == 0 {
			t.Fatalf("expected group %s to have more than 0 chapters, got %d", tc.id, len(cs))
		}
	}
}
