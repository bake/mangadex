package mangadex_test

import (
	"context"
	"testing"
)

func TestUser(t *testing.T) {
	tt := []struct {
		id       string
		username string
		joined   string
		err      bool
	}{
		{"0", "", "", true},
		{"1", "MangaDex", "2012", false},
		{"2", "Holo", "2012", false},
	}
	for _, tc := range tt {
		ctx := context.Background()
		u, err := md.User(ctx, tc.id, nil)
		if !tc.err && err != nil {
			t.Fatalf("expected user to exist, got %q", err)
		}
		if tc.err {
			continue
		}
		if u.Username != tc.username {
			t.Fatalf("expected username of %s to be %s, got %s", tc.id, tc.username, u.Username)
		}
		if u.Joined.Format("2006") != tc.joined {
			t.Fatalf("expected user %s to have joined %s, got %s", tc.id, tc.joined, u.Joined.Format("2006"))
		}
	}
}
