package veryfi

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
				"path": []string{"api", "v7"},
			},
			expected: "https://api.veryfi.com/api/v7",
		},
		{
			in: map[string]interface{}{
				"host": "api.veryfi.com",
				"path": []string{"api", "v7", "foo"},
			},
			expected: "https://api.veryfi.com/api/v7/foo",
		},
	}

	for _, tt := range tests {
		out := buildURL(tt.in["host"].(string), tt.in["path"].([]string)...)
		assert.Equal(t, tt.expected, out)
	}
}
