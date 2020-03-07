# MangaDex

[![GoDoc](https://godoc.org/github.com/bake/mangadex?status.svg)](https://pkg.go.dev/github.com/bake/mangadex)
[![Go Report Card](https://goreportcard.com/badge/github.com/bake/mangadex)](https://goreportcard.com/report/github.com/bake/mangadex)
[![codecov](https://codecov.io/gh/bake/mangadex/branch/master/graph/badge.svg)](https://codecov.io/gh/bake/mangadex)

A Go client for the MangaDex API.

```go
func main() {
  md := mangadex.New()
  m, _, err := md.Manga("23279")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(m.Title)
  // Wonder Cat Kyuu-chan
}
```

```go
func main() {
  md := mangadex.New()
  c, err := md.Chapter("517244")
  if err != nil {
      log.Fatal(err)
  }
  fmt.Printf("%s (Volume %s, Chapter %s)", c, c.Volume, c.Chapter)
  // Cool Day (Volume 3, Chapter 253)
}
```
