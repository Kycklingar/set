package set

import "golang.org/x/exp/constraints"

func NewOrdered[T constraints.Ordered](values ...T) Ordered[T] {
	var set Ordered[T]

	for _, v := range values {
		set.Append(v)
	}

	return set
}

// Sorted wrapper for basic types
type ordered[T constraints.Ordered] struct{ value T }
type Ordered[T constraints.Ordered] Sorted[ordered[T]]

func (a ordered[T]) Less(b ordered[T]) bool {
	return a.value < b.value
}

func convertValues[T constraints.Ordered](values ...T) []ordered[T] {
	var conv = make([]ordered[T], len(values))
	for i := range conv {
		conv[i] = ordered[T]{values[i]}
	}

	return conv
}

func (s Ordered[T]) sorted() Sorted[ordered[T]] {
	return Sorted[ordered[T]](s)
}

// Implement Sorted methods

func (s *Ordered[T]) Append(values ...T) {
	(*Sorted[ordered[T]])(s).Append(convertValues[T](values...)...)
}

func (s *Ordered[T]) Set(values ...T) {
	(*Sorted[ordered[T]])(s).Set(convertValues[T](values...)...)
}

func (s Ordered[T]) Get(v T) (T, bool) {
	val, ok := Sorted[ordered[T]](s).Get(ordered[T]{v})
	return val.value, ok
}

func (s Ordered[T]) Has(v T) bool {
	return Sorted[ordered[T]](s).Has(ordered[T]{v})
}
