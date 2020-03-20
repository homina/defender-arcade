package internal

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

var (
	ErrInvalidRange = errors.New("invalid time range!")
)

type TimeRangeParams struct {
	Layout    string
	Separator string
	Value     string
}

// find max slice in a time range
func MaxSliceInTimeRange(data []string, params TimeRangeParams) (int, error) {
	sot := [][]int{}
	maxSlice := 0

	for _, timeRange := range data {
		if err := ValidateTimeRange(TimeRangeParams{
			Layout:    params.Layout,
			Separator: params.Separator,
			Value:     timeRange,
		}); err != nil {
			return 0, err
		}
		strs := strings.Split(timeRange, params.Separator)
		is := []int{}
		for _, str := range strs {
			s, _ := strconv.Atoi(str)
			is = append(is, s)
		}

		sot = append(sot, is)
	}

	for _, i := range sot {
		total := 0
		for _, j := range sot {
			if i[0] >= j[0] && i[0] < j[1] {
				total++
			}
		}
		if total > maxSlice {
			maxSlice = total
		}
	}

	return maxSlice, nil
}

// validate time range
func ValidateTimeRange(params TimeRangeParams) error {
	timeString := strings.Split(params.Value, params.Separator)
	if len(timeString) != 2 {
		return ErrInvalidRange
	}

	times := []time.Time{}
	for _, str := range timeString {
		if len(str) == 3 {
			str = "0" + str
		}
		t, err := time.Parse(params.Layout, str)
		if err != nil {
			log.Println(err)
			return err
		}
		times = append(times, t)
	}

	if times[1].Unix() < times[0].Unix() {
		return ErrInvalidRange
	}

	return nil
}
