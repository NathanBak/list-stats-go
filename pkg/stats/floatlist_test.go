package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloatList01(t *testing.T) {
	lst := []float64{1.1, 2.2, 4.4, 7.7, 2.2}
	s := NewFloatList(lst)

	assert.Equal(t, 1.1, s.Min())
	assert.Equal(t, 7.7, s.Max())

	roundEqual(t, 3.52, s.Mean())
	assert.Equal(t, 2.2, s.Median())
	assert.Equal(t, 2.2, s.Mode())

	assert.Equal(t, 6.6, s.Range())
	assert.Equal(t, 17.6, s.Sum())

	roundEqual(t, 5.5176, s.Variance())
	roundEqual(t, 2.3490, s.StandardDeviation())
}

func TestFloatList02(t *testing.T) {
	var lst []float64
	s := NewFloatList(lst)

	assert.Equal(t, 0.0, s.Min())
	assert.Equal(t, 0.0, s.Max())

	assert.Equal(t, 0.0, s.Mean())
	assert.Equal(t, 0.0, s.Median())
	assert.Equal(t, 0.0, s.Mode())

	assert.Equal(t, 0.0, s.Range())
	assert.Equal(t, 0.0, s.Sum())

	assert.Equal(t, 0.0, s.Variance())
	assert.Equal(t, 0.0, s.StandardDeviation())
}

func TestFloatList03(t *testing.T) {
	lst := []float64{5}
	s := NewFloatList(lst)

	assert.Equal(t, 5.0, s.Min())
	assert.Equal(t, 5.0, s.Max())

	assert.Equal(t, 5.0, s.Mean())
	assert.Equal(t, 5.0, s.Median())
	assert.Equal(t, 5.0, s.Mode())

	assert.Equal(t, 0.0, s.Range())
	assert.Equal(t, 5.0, s.Sum())

	assert.Equal(t, 0.0, s.Variance())
	assert.Equal(t, 0.0, s.StandardDeviation())
}
