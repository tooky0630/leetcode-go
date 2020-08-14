package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	assert.True(t, isValid(""), "empty string not passed")
	assert.True(t, isValid("()"), "() not passed")
	assert.True(t, isValid("()[]{}"), "()[]{} not passed")
	assert.False(t, isValid("(]"), "(] not passed")
	assert.False(t, isValid("([)]"), "([)] not passed")
	assert.True(t, isValid("{[]}"), "{[]} not passed")
	assert.False(t, isValid("{"), "{ not passed")
}

func TestIsValid2(t *testing.T) {
	assert.True(t, isValid2(""), "empty string not passed")
	assert.True(t, isValid2("()"), "() not passed")
	assert.True(t, isValid2("()[]{}"), "()[]{} not passed")
	assert.False(t, isValid2("(]"), "(] not passed")
	assert.False(t, isValid2("([)]"), "([)] not passed")
	assert.True(t, isValid2("{[]}"), "{[]} not passed")
	assert.False(t, isValid2("{"), "{ not passed")
}
