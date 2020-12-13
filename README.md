# MangaDex

[![GoDoc](https://godoc.org/github.com/bake/mangadex?status.svg)](https://pkg.go.dev/github.com/bake/mangadex/v2)
[![codecov](https://codecov.io/gh/bake/mangadex/branch/v2/graph/badge.svg)](https://codecov.io/gh/bake/mangadex/branch/v2)
<!-- [![Go Report Card](https://goreportcard.com/badge/github.com/bake/mangadex)](https://goreportcard.com/report/github.com/bake/mangadex) -->

A Go client for the MangaDex API v2.

```bash
$ go get github.com/bake/mangadex/v2
```

```go
func main() {
  md := mangadex.New()
  ctx := context.TODO()
  m, err := md.Manga(ctx, 23279, nil)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s by %s", m, m.Author[0])
  // Output: Wonder Cat Kyuu-chan by Nitori Sasami
}
```

```go
func main() {
  md := mangadex.New()
  ctx := context.TODO()
  c, err := md.Chapter(ctx, 517244, nil)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s (Volume %s, Chapter %s)", c, c.Volume, c.Chapter)
  // Output: Cool Day (Volume 3, Chapter 253)
}
```
