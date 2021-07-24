package stats

import (
	"math"
	"sort"
)

// IntList provides various stats for a list of integers.
type IntList struct {
	sorted []int
	size   int
	total  int
}

// NewIntList creates and returns an IntList for the provide list of integers.  Values from the
// provided list are copied so it is safe to modify the list argument after the IntList has been
// created.
func NewIntList(lst []int) IntList {
	cpy := []int{}
	total := 0

	for _, val := range lst {
		cpy = append(cpy, val)
		total += val
	}

	sort.Ints(cpy)

	return IntList{cpy, len(cpy), total}
}

// Min returns the list item with the smallest mathematical value.
func (s IntList) Min() int {
	if s.size == 0 {
		return 0
	}
	return s.sorted[0]
}

// Max returns the list item with the largest mathematical value.
func (s IntList) Max() int {
	if s.size == 0 {
		return 0
	}
	return s.sorted[s.size-1]
}

// Mean returns the average value of all items in the list.
func (s IntList) Mean() float64 {
	if s.size == 0 {
		return 0
	}
	return float64(s.total) / float64(s.size)
}

// Median returns the middle value of all items in the list.  If the list has an even number of
// items then the average of the two middle values is returned.
func (s IntList) Median() float64 {
	if s.size == 0 {
		return 0
	}

	half := s.size / 2

	if s.size%2 != 0 { // check if size is odd number
		return float64(s.sorted[half])
	}

	return float64(s.sorted[half]+s.sorted[half-1]) / 2
}

// Mode returns the value which appears most frequently in the list.  If all values in the list are
// unique than a value from the list will be selected nondeterministically and returned.  If there
// is a tie for mode value (two or more values have the same number of occurrences) then one of the
// values will be selected nondeterministacally and returned.
func (s IntList) Mode() int {
	if s.size == 0 {
		return 0
	}

	mp := make(map[int]int)

	for _, v := range s.sorted {
		mp[v]++
	}

	maxCount := 0
	maxVal := 0

	for val, count := range mp {
		if count > maxCount {
			maxCount = count
			maxVal = val
		}
	}

	return maxVal
}

// Range returns the difference between the min and the max.
func (s IntList) Range() int {
	if s.size <= 1 {
		return 0
	}

	return s.sorted[s.size-1] - s.sorted[0]
}

// Sum returns the total of all list items added together.
func (s IntList) Sum() int {
	return s.total
}

// Variance returns the statistical variance of the list items.
func (s IntList) Variance() float64 {
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
func (s IntList) StandardDeviation() float64 {
	return math.Sqrt(s.Variance())
}
