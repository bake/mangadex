package mangadex_test

import (
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
