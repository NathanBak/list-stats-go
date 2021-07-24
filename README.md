# list-stats-go

## Introduction
The list-stats-go project provides a way to get basic statistics such as max, min, mean, mode, and median from lists of numbers.  

## Design Drivers
There are similar libraries out there.  This library had certain requirements driving its creation instead of using a known existing solution.
- Support for ints.  For good reasons most stats packages are built around float64 slices.  This library will work with those, but also has a solution for regular int slices.
- No errors.  All responses are either ints or float64s.  For cases where an error might be appropriate (such as an average of an empty list), zero is returned.
- Front-loading of likely used calculations.  A sorted copy of all items and the total of all items are cached since they will likely be used one or more times.
- Enough functionality, but not too much.  Some solutions have mean and median, but not mode or standard deviation.  Others have dozens of statistical functions that may be useful to experts but can clutter the API for a novice in stats.

## Sample Use
```golang
package main

import (
	"fmt"
	
	"github.com/NathanBak/list-stats-go/pkg/stats"
)

func main() {
	lst := []int{8, 9, 10, 10, 10, 11, 11, 11, 12, 13}
	s := NewIntList(lst)

	fmt.Println("Min", s.Min())
	fmt.Println("Max", s.Max())

	fmt.Println("Mean", s.Mean())
	fmt.Println("Median", s.Median())
	fmt.Println("Mode", s.Mode())

	fmt.Println("Range", s.Range())
	fmt.Println("Sum", s.Sum())

	fmt.Println("Variance", s.Variance())
	fmt.Println("Standard Deviation", fmt.Sprintf("%.4f", s.StandardDeviation()))

}
```

## Sample Output

```
Min 8
Max 13
Mean 10.5
Median 10.5
Mode 11
Range 5
Sum 105
Variance 1.85
Standard Deviation 1.3601
```

