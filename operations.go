package set

import (
	mm "github.com/kycklingar/MinMax"
)

// Produce the diff of sets A and B
func Diff[T Less[T]](a, b Sorted[T]) Sorted[T] {
	var res = make(Sorted[T], 0, len(a))

	n := func(T) {}
	s := func(v T) { res = append(res, v) }

	i, _ := sync(a, b, s, n, n)

	for ; i < len(a); i++ {
		res = append(res, a[i])
	}

	return res
}

// Produce the union of sets A and B
func Union[T Less[T]](a, b Sorted[T]) Sorted[T] {
	var res = New[T]()

	s := func(v T) {
		res = append(res, v)
	}

	i, j := sync(a, b, s, s, s)

	if i < len(a) {
		res = append(res, a[i:]...)
	}

	if j < len(b) {
		res = append(res, b[j:]...)
	}

	return res
}

// Produce the intersection of sets A and B
func Intersection[T Less[T]](a, b Sorted[T]) Sorted[T] {
	var res = make(Sorted[T], 0, mm.Max(len(a), len(b)))

	n := func(T) {}
	s := func(v T) { res = append(res, v) }

	sync(a, b, n, n, s)

	return res
}

func sync[T Less[T]](a, b Sorted[T], less, greater, equal func(T)) (int, int) {
	var i, j int

	for i < len(a) && j < len(b) {
		if a[i].Less(b[j]) {
			less(a[i])
			i++
		} else if b[j].Less(a[i]) {
			greater(b[j])
			j++
		} else {
			equal(a[i])
			i++
			j++
		}
	}

	return i, j
}
