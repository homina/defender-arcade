package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEachLine(t *testing.T) {
	layout := "1504"
	separator := " "
	// test with invalid total time
	err := ValidateTimeRange(TimeRangeParams{
		Layout:    layout,
		Separator: separator,
		Value:     "900",
	})
	assert.NotNil(t, err, "should return error if input by invalid value")
	err = ValidateTimeRange(TimeRangeParams{
		Layout:    layout,
		Separator: separator,
		Value:     "900 1000 900",
	})
	assert.NotNil(t, err, "should return error if input by invalid value")

	// test with invalid time format
	err = ValidateTimeRange(TimeRangeParams{
		Layout:    layout,
		Separator: separator,
		Value:     "90a 1000",
	})
	assert.NotNil(t, err, "should return error if input by invalid format")

	err = ValidateTimeRange(TimeRangeParams{
		Layout:    layout,
		Separator: separator,
		Value:     "10000 1000",
	})
	assert.NotNil(t, err, "should return error if input by invalid format")

	// test with second is smaller than first
	err = ValidateTimeRange(TimeRangeParams{
		Layout:    layout,
		Separator: separator,
		Value:     "1000 900",
	})
	assert.NotNil(t, err, "should return error if second is smaller than first")

	err = ValidateTimeRange(TimeRangeParams{
		Layout:    layout,
		Separator: separator,
		Value:     "900 1000",
	})
	assert.Nil(t, err, "should not return error ")
}

func TestMaxSliceInTimeRange(t *testing.T) {
	layout := "1504"
	separator := " "

	// test with invalid value
	data := []string{
		"1000 900",
	}

	_, err := MaxSliceInTimeRange(data, TimeRangeParams{
		Layout:    layout,
		Separator: separator,
	})
	assert.NotNil(t, err, "should return error if second is smaller than first")

	// test with valid value
	data = []string{
		"900 910",
		"940 1200",
		"950 1120",
		"1100 1130",
		"1000 1130",
		"900 1120",
		"1300 1400",
		"1350 1420",
	}

	maxSlice, err := MaxSliceInTimeRange(data, TimeRangeParams{
		Layout:    layout,
		Separator: separator,
	})
	assert.Nil(t, err, "should not return error")
	assert.Equal(t, maxSlice, 5)
}
