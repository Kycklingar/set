package set

import "golang.org/x/exp/constraints"

func ODiff[T constraints.Ordered](a, b Ordered[T]) Ordered[T] {
	return Ordered[T](Diff(a.sorted(), b.sorted()))
}

func OUnion[T constraints.Ordered](a, b Ordered[T]) Ordered[T] {
	return Ordered[T](Union(a.sorted(), b.sorted()))
}

func OIntersection[T constraints.Ordered](a, b Ordered[T]) Ordered[T] {
	return Ordered[T](Intersection(a.sorted(), b.sorted()))
}
