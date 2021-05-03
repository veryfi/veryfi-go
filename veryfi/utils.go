package veryfi

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

// Base64EncodeFile encodes a file using base64.
func Base64EncodeFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(content), nil
}

// buildURL builds up a complete URL from given scheme, host and path.
func buildURL(host string, path ...string) string {
	u := &url.URL{
		Scheme: "https",
		Host:   host,
		Path:   strings.Join(path, "/"),
	}

	return u.String()
}
