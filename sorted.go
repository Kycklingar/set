package set

func New[T Less[T]](values ...T) Sorted[T] {
	var s Sorted[T]
	s.Append(values...)
	return s
}

type (
	Sorted[T Less[T]] []T

	// Less is used for sorting in the Sorted set
	Less[T any] interface{ Less(T) bool }
)

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
			(*s)[index] = value
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
		return s[index], true
	}

	return value, false
}

func (s *Sorted[T]) insert(index int, value T) {
	*s = append(*s, value)

	if (*s)[index].Less(value) {
		index++
	}

	for i := index; i < len(*s); i++ {
		(*s)[i], value = value, (*s)[i]
	}
}

// return the index of b or 0< if not found
// less than 0 is the negative index of m - 1 before returning
func (s Sorted[T]) bsearch(value T) int {
	var (
		l, m int
		r    = len(s) - 1
	)

	for l <= r {
		m = (l + r) / 2
		if s[m].Less(value) {
			l = m + 1
		} else if value.Less(s[m]) {
			r = m - 1
		} else {
			return m
		}
	}

	return -m - 1
}
