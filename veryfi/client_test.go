package veryfi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnitNewClientV7_NilConfig(t *testing.T) {
	client, err := NewClientV7(nil)
	assert.Nil(t, client)
	assert.Error(t, err)
}
