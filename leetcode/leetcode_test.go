package leetcode

import (
	"gotest.tools/v3/assert"
	"testing"
)

func Test_isAnagram(t *testing.T) {
	assert.Equal(t, isAnagram("first", "fisrt"), true)

	assert.Equal(t, isAnagram("firrst", "fisrtr"), true)

	assert.Equal(t, isAnagram("afirst", "fisrt"), false)

	assert.Equal(t, isAnagram("afirst", "fisrtb"), false)
}
