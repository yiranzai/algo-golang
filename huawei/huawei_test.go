package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestTest(t *testing.T) {
	test("#222233")
	assert.Equal(t, 1, 1)
}

func TestTest1(t *testing.T) {
	test("2222/22")
	assert.Equal(t, 1, 1)
}
