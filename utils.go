package cloudflare

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

// buildURI assembles the base path and queries.
func buildURI(path string, options interface{}) string {
	v, _ := query.Values(options)
	return (&url.URL{Path: path, RawQuery: v.Encode()}).String()
}
