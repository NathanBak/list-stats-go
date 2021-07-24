package stats

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntList01(t *testing.T) {
	lst := []int{1, 2, 4, 7}
	s := NewIntList(lst)

	assert.Equal(t, 1, s.Min())
	assert.Equal(t, 7, s.Max())

	assert.Equal(t, 3.5, s.Mean())
	assert.Equal(t, 3.0, s.Median())
	// assert.Equal(t, 1, s.Mode())  // No mode so response is nondeterministic

	assert.Equal(t, 6, s.Range())
	assert.Equal(t, 14, s.Sum())

	roundEqual(t, 5.25, s.Variance())
	roundEqual(t, 2.2913, s.StandardDeviation())
}

func TestIntList02(t *testing.T) {
	lst := []int{8, 9, 10, 10, 10, 11, 11, 11, 12, 13}
	s := NewIntList(lst)

	assert.Equal(t, 8, s.Min())
	assert.Equal(t, 13, s.Max())

	assert.Equal(t, 10.5, s.Mean())
	assert.Equal(t, 10.5, s.Median())
	assert.Equal(t, 10, s.Mode())

	assert.Equal(t, 5, s.Range())
	assert.Equal(t, 105, s.Sum())

	roundEqual(t, 1.85, s.Variance())
	roundEqual(t, 1.3601, s.StandardDeviation())
}

func TestIntList03(t *testing.T) {
	lst := []int{13, 18, 13, 14, 13, 16, 14, 21, 13}
	s := NewIntList(lst)

	assert.Equal(t, 13, s.Min())
	assert.Equal(t, 21, s.Max())

	assert.Equal(t, 15.0, s.Mean())
	assert.Equal(t, 14.0, s.Median())
	assert.Equal(t, 13, s.Mode())

	assert.Equal(t, 8, s.Range())
	assert.Equal(t, 135, s.Sum())

	roundEqual(t, 7.1111, s.Variance())
	roundEqual(t, 2.6667, s.StandardDeviation())
}

func TestIntList04(t *testing.T) {
	var lst []int
	s := NewIntList(lst)

	assert.Equal(t, 0, s.Min())
	assert.Equal(t, 0, s.Max())

	assert.Equal(t, 0.0, s.Mean())
	assert.Equal(t, 0.0, s.Median())
	assert.Equal(t, 0, s.Mode())

	assert.Equal(t, 0, s.Range())
	assert.Equal(t, 0, s.Sum())

	roundEqual(t, 0.0, s.Variance())
	roundEqual(t, 0.0, s.StandardDeviation())
}

func TestIntList05(t *testing.T) {
	lst := []int{5}
	s := NewIntList(lst)

	assert.Equal(t, 5, s.Min())
	assert.Equal(t, 5, s.Max())

	assert.Equal(t, 5.0, s.Mean())
	assert.Equal(t, 5.0, s.Median())
	assert.Equal(t, 5, s.Mode())

	assert.Equal(t, 0, s.Range())
	assert.Equal(t, 5, s.Sum())

	roundEqual(t, 0.0, s.Variance())
	roundEqual(t, 0.0, s.StandardDeviation())
}

// roundEqual compares two floats rounded to four decimal places.
func roundEqual(t *testing.T, expected, actual float64) {
	expectedStr := fmt.Sprintf("%.4f", expected)
	actualStr := fmt.Sprintf("%.4f", actual)
	assert.Equal(t, expectedStr, actualStr)
}
