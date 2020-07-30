package isodd

import "math"

// IsOdd returns if the given integer is an odd number or not.
func IsOdd(num int) bool {
	test := int(math.Abs(float64(num)))
	return test%2 == 1
}
