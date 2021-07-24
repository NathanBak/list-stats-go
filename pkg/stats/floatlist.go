package stats

import (
	"math"
	"sort"
)

// FloatList provides various stats for a list of floating point number values.
type FloatList struct {
	sorted []float64
	size   int
	total  float64
}

// NewFloatList creates and returns an FloatList for the provide list of floating point number
// values.  Values from the provided list are copied so it is safe to modify the list argument
// after the FloatList has been created.
func NewFloatList(lst []float64) FloatList {
	cpy := []float64{}
	total := 0.0

	for _, val := range lst {
		cpy = append(cpy, val)
		total += val
	}

	sort.Float64s(cpy)

	return FloatList{cpy, len(cpy), total}
}

// Min returns the list item with the smallest mathematical value.
func (s FloatList) Min() float64 {
	if s.size == 0 {
		return 0
	}
	return s.sorted[0]
}

// Max returns the list item with the largest mathematical value.
func (s FloatList) Max() float64 {
	if s.size == 0 {
		return 0
	}
	return s.sorted[s.size-1]
}

// Mean returns the average value of all items in the list.
func (s FloatList) Mean() float64 {
	if s.size == 0 {
		return 0
	}
	return s.total / float64(s.size)
}

// Median returns the middle value of all items in the list.  If the list has an even number of
// items then the average of the two middle values is returned.
func (s FloatList) Median() float64 {
	if s.size == 0 {
		return 0
	}

	half := s.size / 2

	if s.size%2 != 0 { // check if size is odd number
		return s.sorted[half]
	}

	return s.sorted[half] + s.sorted[half-1]/2
}

// Mode returns the value which appears most frequently in the list.  If all values in the list are
// unique than a value from the list will be selected nondeterministically and returned.  If there
// is a tie for mode value (two or more values have the same number of occurrences) then one of the
// values will be selected nondeterministacally and returned.
func (s FloatList) Mode() float64 {
	if s.size == 0 {
		return 0
	}

	mp := make(map[float64]int)

	for _, v := range s.sorted {
		mp[v]++
	}

	maxCount := 0
	maxVal := 0.0

	for val, count := range mp {
		if count > maxCount {
			maxCount = count
			maxVal = val
		}
	}

	return maxVal
}

// Range returns the difference between the min and the max.
func (s FloatList) Range() float64 {
	if s.size <= 1 {
		return 0
	}

	return s.sorted[s.size-1] - s.sorted[0]
}

// Sum returns the total of all list items added together.
func (s FloatList) Sum() float64 {
	return s.total
}

// Variance returns the statistical variance of the list items.
func (s FloatList) Variance() float64 {
	if s.size <= 1 {
		return 0
	}

	mean := s.Mean()
	total := 0.0

	for _, v := range s.sorted {
		diff := mean - float64(v)
		total += diff * diff
	}

	return total / float64(s.size)
}

// StandardDeviation returns the statistical standard deviation of the list items.  It is calculated
// by taking the square root of the statistical variance.
func (s FloatList) StandardDeviation() float64 {
	return math.Sqrt(s.Variance())
}
