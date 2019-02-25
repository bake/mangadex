package mangadex

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// Client is a MangaRock Client.
type Client struct {
	base   string
	client *http.Client
}

// WithBase sets the API base.
func WithBase(base string) func(*Client) {
	return func(md *Client) {
		md.base = base
	}
}

// WithHTTPClient makes the manga client use a given http.Client to make
// requests.
func WithHTTPClient(c *http.Client) func(*Client) {
	return func(md *Client) {
		md.client = c
	}
}

// New returns a new MangaRock Client.
func New(options ...func(*Client)) *Client {
	c := &Client{
		base:   "https://mangadex.org/api/",
		client: &http.Client{},
	}
	for _, option := range options {
		option(c)
	}
	return c
}

// get sends a HTTP GET request.
func (c *Client) get(id, t string) (json.RawMessage, error) {
	req, err := http.NewRequest(http.MethodGet, c.base, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not create get request")
	}
	query := url.Values{"id": []string{id}, "type": []string{t}}
	req.URL.RawQuery = query.Encode()
	res, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "cound not get %s", req.URL)
	}
	defer res.Body.Close()
	var raw json.RawMessage
	if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
		return nil, errors.Wrap(err, "could not decode response")
	}
	return raw, nil
}