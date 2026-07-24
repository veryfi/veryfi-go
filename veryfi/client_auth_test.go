package veryfi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizationHeader(t *testing.T) {
	// Legacy keys use the `apikey <username>:<key>` format.
	legacy := authorizationHeader(&Options{Username: "user", APIKey: "legacykey"})
	assert.Equal(t, "apikey user:legacykey", legacy)

	// Client-scoped keys (vrfk_ prefix) authenticate as a Bearer token, no username needed.
	bearer := authorizationHeader(&Options{APIKey: "vrfk_abc123"})
	assert.Equal(t, "Bearer vrfk_abc123", bearer)

	// A vrfk_ key still uses Bearer even if a username happens to be set (it's ignored).
	bearerWithUser := authorizationHeader(&Options{Username: "user", APIKey: "vrfk_xyz"})
	assert.Equal(t, "Bearer vrfk_xyz", bearerWithUser)
}
