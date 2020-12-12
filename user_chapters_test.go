package mangadex_test

import (
	"context"
	"testing"
)

func TestUserChapters(t *testing.T) {
	tt := []struct {
		id  int
		err bool
	}{
		{0, true},
		{2, false},
	}
	for _, tc := range tt {
		ctx := context.Background()
		cs, err := md.UserChapters(ctx, tc.id, nil)
		if !tc.err && err != nil {
			t.Fatalf("expected user %d to have chapters, got %q", tc.id, err)
		}
		if tc.err {
			continue
		}
		if len(cs) == 0 {
			t.Fatalf("expected user %d to have more than 0 chapters", tc.id)
		}
	}
}
