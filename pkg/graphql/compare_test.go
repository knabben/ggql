package graphql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestGraphqlCompare(t *testing.T) {
	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(t, a, b, "The two words should be the same.")
}