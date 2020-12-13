package mangadex_test

import (
	"context"
	"testing"
)

func TestRelations(t *testing.T) {
	tt := []struct {
		id   int
		name string
	}{
		{1, "Prequel"},
		{5, "Side story"},
	}
	ctx := context.Background()
	relations, err := md.Relations(ctx, nil)
	if err != nil {
		t.Fatalf("could not get relations: %v", err)
	}
	for _, tc := range tt {
		var found bool
		for _, r := range relations {
			if r.ID == tc.id && r.String() == tc.name {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected relation %d (%s) to exist", tc.id, tc.name)
		}
	}
}
