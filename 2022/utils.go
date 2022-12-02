package main

type slice[E, V any] []E

func (s slice[E, V]) Map(iteratee func(E) V) []V {
	result := []V{}
	for _, item := range s {
		result = append(result, iteratee(item))
	}

	return result
}
