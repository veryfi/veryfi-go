package veryfi

import (
	"net/url"
	"strings"
)

// buildURL builds up a complete URL from given scheme, host and path.
func buildURL(host string, path ...string) string {
	u := &url.URL{
		Scheme: "https",
		Host:   host,
		Path:   strings.Join(path, "/"),
	}

	return u.String()
}
