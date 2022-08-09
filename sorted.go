package set

import "golang.org/x/exp/constraints"

// NewSorted will return a Sorted set of type T using the supplied less function
func NewSorted[T any](fn Less[T]) Sorted[T] {
	return Sorted[T]{less: fn}
}

// NewOrdered is a Sorted set of basic types with a provided less function
// use NewSorted to supply a different less function
func NewOrdered[T ordered]() Sorted[T] {
	return Sorted[T]{less: orderedLess[T]}
}

type (
	// Less is used for sorting in the Sorted set
	Less[T any] func(T, T) bool
	ordered     constraints.Ordered

	// Sorted is a sorted set of type T, providing a set of common functions
	Sorted[T any] struct {
		Slice []T
		less  func(T, T) bool
	}
)

func orderedLess[T ordered](a, b T) bool {
	return a < b
}

// Append adds the value to the set, returning unmodified state of s if it already exist
func (s *Sorted[T]) Append(values ...T) {
	for _, value := range values {
		index := s.bsearch(value)
		if index > -1 {
			continue
		}

		s.insert(-index-1, value)
	}
}

// Set sets or adds the value in the set
func (s *Sorted[T]) Set(values ...T) {
	for _, value := range values {
		index := s.bsearch(value)
		if index > -1 {
			s.Slice[index] = value
		} else {
			s.insert(-index-1, value)
		}
	}
}

// returns true if the value is in the set
func (s Sorted[T]) Has(value T) bool {
	return s.bsearch(value) > -1
}

// returns the value, and true if the value is in the set
// otherwise the provided T value, and false is returned
func (s Sorted[T]) Get(value T) (T, bool) {
	index := s.bsearch(value)
	if index > -1 {
		return s.Slice[index], true
	}

	return value, false
}

func (s *Sorted[T]) insert(index int, value T) {
	s.Slice = append(s.Slice, value)

	if s.less(s.Slice[index], value) {
		index++
	}

	for i := index; i < len(s.Slice); i++ {
		s.Slice[i], value = value, s.Slice[i]
	}
}

// return the index of b or 0< if not found
// less than 0 is the negative index of m - 1 before returning
func (s Sorted[T]) bsearch(value T) int {
	var (
		l, m int
		r    = len(s.Slice) - 1
	)

	for l <= r {
		m = (l + r) / 2
		if s.less(s.Slice[m], value) {
			l = m + 1
		} else if s.less(value, s.Slice[m]) {
			r = m - 1
		} else {
			return m
		}
	}

	return -m - 1
}
