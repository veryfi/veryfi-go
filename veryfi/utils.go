package veryfi

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"reflect"
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

// structToMap converts a struct of string fields to a map[string]string. It is nil-safe: a nil
// interface, a nil pointer, or a non-struct value yields an empty map (some callers, e.g.
// GetLineItems/GetTags, pass nil query params).
func structToMap(s interface{}) map[string]string {
	out := map[string]string{}
	if s == nil {
		return out
	}

	values := reflect.ValueOf(s)
	for values.Kind() == reflect.Ptr {
		if values.IsNil() {
			return out
		}
		values = values.Elem()
	}
	if values.Kind() != reflect.Struct {
		return out
	}

	fields := values.Type()
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i).Tag.Get("json")
		value := fmt.Sprint(values.Field(i))
		if len(value) > 0 {
			out[field] = value
		}
	}

	return out
}
