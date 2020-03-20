package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEachLine(t *testing.T) {
	// test with invalid file
	_, err := GetEachLine("dummy.txt")
	assert.NotNil(t, err, "should return error if input by invalid value")
	// test with valid file
	lines, err := GetEachLine("input1.txt")
	assert.Nil(t, err, "should not return error if input by valid value")
	assert.Equal(t, len(lines), 9, "should return with 3 lines")
	assert.Equal(t, lines[0], "Louis IX", "should return with match value")
}
