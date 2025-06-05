package veryfi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/veryfi/veryfi-go/v3/veryfi/scheme"
)

func TestUnitBuildURL(t *testing.T) {
	tests := []struct {
		in       map[string]interface{}
		expected string
	}{
		{
			in: map[string]interface{}{
				"host": "api.veryfi.com",
				"path": []string{""},
			},
			expected: "https://api.veryfi.com",
		},
		{
			in: map[string]interface{}{
				"host": "api.veryfi.com",
				"path": []string{"api"},
			},
			expected: "https://api.veryfi.com/api",
		},
		{
			in: map[string]interface{}{
				"host": "api.veryfi.com",
				"path": []string{"api", "v8"},
			},
			expected: "https://api.veryfi.com/api/v8",
		},
		{
			in: map[string]interface{}{
				"host": "api.veryfi.com",
				"path": []string{"api", "v8", "foo"},
			},
			expected: "https://api.veryfi.com/api/v8/foo",
		},
	}

	for _, tt := range tests {
		out := buildURL(tt.in["host"].(string), tt.in["path"].([]string)...)
		assert.Equal(t, tt.expected, out)
	}
}

func TestUnitStructToMap(t *testing.T) {
	tests := []struct {
		in       scheme.DocumentSearchOptions
		expected map[string]string
	}{
		{
			in:       scheme.DocumentSearchOptions{},
			expected: map[string]string{},
		},
		{
			in: scheme.DocumentSearchOptions{
				CreatedGT: "foo",
			},
			expected: map[string]string{
				"created__gt": "foo",
			},
		},
		{
			in: scheme.DocumentSearchOptions{
				Q:         "",
				CreatedGT: "foo",
			},
			expected: map[string]string{
				"created__gt": "foo",
			},
		},
		{
			in: scheme.DocumentSearchOptions{
				Q:         "foo",
				CreatedGT: "foo",
			},
			expected: map[string]string{
				"q":           "foo",
				"created__gt": "foo",
			},
		},
		{
			in: scheme.DocumentSearchOptions{
				Q:          "foo",
				ExternalID: "foo",
				Tag:        "foo",
				CreatedGT:  "foo",
				CreatedGTE: "foo",
				CreatedLT:  "foo",
				CreatedLTE: "foo",
			},
			expected: map[string]string{
				"q":            "foo",
				"external_id":  "foo",
				"tag":          "foo",
				"created__gt":  "foo",
				"created__gte": "foo",
				"created__lt":  "foo",
				"created__lte": "foo",
			},
		},
	}

	for _, tt := range tests {
		out := structToMap(tt.in)
		assert.Equal(t, tt.expected, out)
	}
}
