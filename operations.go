package set

import (
	mm "github.com/kycklingar/MinMax"
)

// Produce the diff of sets A and B
func Diff[T any](a, b Sorted[T]) Sorted[T] {
	var res = Sorted[T]{
		Slice: make([]T, 0, len(a.Slice)),
		less:  a.less,
	}

	n := func(T) {}
	s := func(v T) { res.Slice = append(res.Slice, v) }

	i, _ := sync(a, b, s, n, n)

	for ; i < len(a.Slice); i++ {
		res.Slice = append(res.Slice, a.Slice[i])
	}

	return res
}

// Produce the union of sets A and B
func Union[T any](a, b Sorted[T]) Sorted[T] {
	var res = NewSorted[T](a.less)

	s := func(v T) {
		res.Slice = append(res.Slice, v)
	}

	i, j := sync(a, b, s, s, s)

	if i < len(a.Slice) {
		res.Slice = append(res.Slice, a.Slice[i:]...)
	}

	if j < len(b.Slice) {
		res.Slice = append(res.Slice, b.Slice[j:]...)
	}

	return res
}

// Produce the intersection of sets A and B
func Intersection[T any](a, b Sorted[T]) Sorted[T] {
	var res = Sorted[T]{
		Slice: make([]T, 0, mm.Max(len(a.Slice), len(b.Slice))),
		less:  a.less,
	}

	n := func(T) {}
	s := func(v T) { res.Slice = append(res.Slice, v) }

	sync(a, b, n, n, s)

	return res
}

func sync[T any](a, b Sorted[T], less, greater, equal func(T)) (int, int) {
	var (
		l    = a.less
		i, j int
	)

	for i < len(a.Slice) && j < len(b.Slice) {
		if l(a.Slice[i], b.Slice[j]) {
			less(a.Slice[i])
			i++
		} else if l(b.Slice[j], a.Slice[i]) {
			greater(b.Slice[j])
			j++
		} else {
			equal(a.Slice[i])
			i++
			j++
		}
	}

	return i, j
}
