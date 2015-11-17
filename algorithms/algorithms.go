package algorithms

import (
"math"
)

// Fac return factorial(n) where n is an int32.
func Fac(n int64) int64 {
	if n == 0 {
		return 1
	}
	
	return n * Fac(n - 1)
}

// MinInt32 return the value and index of the smallest int in vals.
// The first return is the value, and the second is the argument.
func MinInt(vals []int) (int, int) {
	return MinIntAfter(vals, -1)
}

// MinIntAfter return the value and index of the smallest int in vals
// after the given starting index.
// The first return is the value, and the second is the argument.
func MinIntAfter(vals []int, after int) (int, int) {
	if len(vals) == 0 {
		return math.MinInt32, -1
	}
	
	min := vals[after + 1]
	pos := after + 1
	
	for i := after + 1; i < len(vals); i++ {
		if vals[i] < min {
			min = vals[i]
			pos = i
		}
	}
	
	return min, pos
}

// SortInt takes a source and returns its values sortedVals
// in ascending order.
func SortInt(vals []int) []int {
	length := len(vals)
	buf := make([]int, length)
	copy(buf, vals)
	index := -1
	
	for index < length - 1{
		_, pos := MinIntAfter(buf, index)
		index++
		temp := buf[index]
		buf[index] = buf[pos]
		buf[pos] = temp
	}
	
	return buf
}