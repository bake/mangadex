package mangadex_test

import (
	"context"
	"testing"
)

func TestFollows(t *testing.T) {
	tt := []struct {
		id   int
		name string
	}{
		{1, "Reading"},
		{4, "Plan to read"},
	}
	ctx := context.Background()
	follows, err := md.Follows(ctx, nil)
	if err != nil {
		t.Fatalf("could not get follows: %v", err)
	}
	for _, tc := range tt {
		var found bool
		for _, f := range follows {
			if f.ID == tc.id && f.String() == tc.name {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected follow %d (%s) to exist", tc.id, tc.name)
		}
	}
}
