package utils

import "math"

type Slice[E, V any] []E

func (s Slice[E, V]) Map(iteratee func(E) V) []V {
	result := []V{}
	for _, item := range s {
		result = append(result, iteratee(item))
	}

	return result
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Max(arr []int) int {
	max := math.MinInt32
	for _, item := range arr {
		if item > max {
			max = item
		}
	}
	return max
}
