package set

import (
	"fmt"
	"math/rand"
	"testing"
)

func lessInt(a, b int) bool {
	return a < b
}

func TestSortedHas(t *testing.T) {
	var s = Sorted[int]{
		Slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 20, 21, 22, 23, 24, 25, 49, 123123},
		less:  lessInt,
	}

	testHas(t, s, 1, true)
	testHas(t, s, 2, true)
	testHas(t, s, 3, true)
	testHas(t, s, 4, true)
	testHas(t, s, 5, true)
	testHas(t, s, 6, true)
	testHas(t, s, 7, true)
	testHas(t, s, 8, true)
	testHas(t, s, 24, true)
	testHas(t, s, 49, true)
	testHas(t, s, 123123, true)
	testHas(t, s, 9, false)
	testHas(t, s, 123, false)
	testHas(t, s, 54982, false)
	testHas(t, s, 58439754, false)
	testHas(t, s, 483924, false)
	testHas(t, s, 0, false)
}

type combinedString struct {
	a, b string
}

func (cs combinedString) String() string {
	return fmt.Sprint(cs.a, " ", cs.b)
}

func lessCombinedString(a, b combinedString) bool {
	return a.String() < b.String()
}

func greatCombinedString(a, b combinedString) bool {
	return a.String() > b.String()
}

func TestSortedTypes(t *testing.T) {
	var s = NewSorted[combinedString](lessCombinedString)

	s.Append(combinedString{"hello", "world"})
	s.Append(combinedString{"hello", "world"})
	s.Append(combinedString{"hello", "earth"})
	s.Append(combinedString{"hello", "ocean"})
	s.Append(combinedString{"hello", "forest"})
	s.Append(combinedString{"goodbye", "forest"})
	s.Append(combinedString{"goodbye", "earth"})
	s.Append(combinedString{"goodbye", "world"})
	s.Append(combinedString{"goodbye", "ocean"})

	testHas[combinedString](t, s, combinedString{"hello", "world"}, true)
	testHas[combinedString](t, s, combinedString{"hello", "earth"}, true)
	testHas[combinedString](t, s, combinedString{"hello", "ocean"}, true)
	testHas[combinedString](t, s, combinedString{"hello", "forest"}, true)
	testHas[combinedString](t, s, combinedString{"goodbye", "forest"}, true)
	testHas[combinedString](t, s, combinedString{"goodbye", "earth"}, true)
	testHas[combinedString](t, s, combinedString{"goodbye", "world"}, true)
	testHas[combinedString](t, s, combinedString{"goodbye", "ocean"}, true)
	testHas[combinedString](t, s, combinedString{"howdy", "partner"}, false)
}

type wval struct {
	int
	string
}

func lessWvalInt(a, b wval) bool {
	return a.int < b.int
}

func TestSortedGet(t *testing.T) {
	var set = NewSorted[wval](lessWvalInt)
	set.Append(wval{1, "yes"})
	set.Append(wval{2, "no"})
	set.Append(wval{3, "maybe"})
	set.Append(wval{1, "no"})

	val, ok := set.Get(wval{int: 3})
	if !ok {
		t.Errorf("Sorted.Get returned !ok when ok is expected")
	}

	if val.string != "maybe" {
		t.Errorf("Sorted.Get returned incorrect value")
	}

	val, ok = set.Get(wval{int: 1})
	if !ok {
		t.Errorf("Sorted.Get returned !ok when ok is expected")
	}
	if val.string != "yes" {
		t.Errorf("Sorted.Get returned incorrect value")
	}
}

func TestSortedAppend(t *testing.T) {
	var s = Sorted[int]{
		less: lessInt,
	}

	var ints = generateRandomInts(10000)

	for i := 0; i < len(ints); i++ {
		s.Append(ints[i])
	}
}

func TestSortedSet(t *testing.T) {
	var (
		s    = NewSorted[wval](lessWvalInt)
		nwal = wval{1, "false"}
		ok   bool
	)

	s.Set(wval{1, "true"})
	s.Set(wval{2, "true"})
	s.Set(wval{3, "true"})
	s.Set(wval{4, "true"})
	s.Set(wval{5, "true"})
	s.Set(wval{6, "true"})
	s.Set(nwal)
	nwal, ok = s.Get(nwal)
	if !ok {
		t.Errorf("Sorted.Get failed")
	}
	if nwal.string != "false" {
		t.Errorf("Sorted.Set failed")
	}
}

func generateRandomInts(size int) []int {
	var ints = make([]int, size)
	for i := 0; i < size; i++ {
		ints[i] = rand.Int()
	}

	return ints
}

func generateRandomIndices(size, max int) []int {
	var indices = make([]int, size)
	for i := 0; i < size; i++ {
		indices[i] = rand.Int() % max
	}

	return indices
}

func BenchmarkSortedAppend10k(b *testing.B) {
	benchmarkSortedAppend(b, 10000)
}

func BenchmarkSortedAppend100k(b *testing.B) {
	benchmarkSortedAppend(b, 100000)
}

func benchmarkSortedAppend(b *testing.B, numItems int) {
	var (
		set     = NewOrdered[int]()
		ints    = generateRandomInts(numItems)
		indices = generateRandomIndices(b.N, len(ints))
	)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		set.Append(ints[indices[n]])
	}
}

func BenchmarkSortedHasKnown10k(b *testing.B) {
	benchmarkSortedHasKnown(b, 10000)
}

func BenchmarkSortedHasKnown100k(b *testing.B) {
	benchmarkSortedHasKnown(b, 100000)
}

func benchmarkSortedHasKnown(b *testing.B, numItems int) {
	var (
		set  = NewOrdered[int]()
		ints = generateRandomInts(numItems)

		indices = generateRandomIndices(b.N, len(ints))
	)

	for i := 0; i < len(ints); i++ {
		set.Append(ints[i])
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		set.Has(ints[indices[n]])
	}
}

func BenchmarkSortedHasUnknown10k(b *testing.B) {
	benchmarkSortedHasUnknown(b, 10000)
}
func BenchmarkSortedHasUnknown100k(b *testing.B) {
	benchmarkSortedHasUnknown(b, 100000)
}

func benchmarkSortedHasUnknown(b *testing.B, numItems int) {
	var (
		set     = NewOrdered[int]()
		ints    = generateRandomInts(numItems)
		indices = generateRandomIndices(b.N, len(ints))
	)

	for i := 0; i < len(ints); i++ {
		set.Append(ints[i])
	}

	ints = generateRandomInts(len(ints))

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		set.Has(ints[indices[n]])
	}
}

//func BenchmarkSortedStrings(b *testing.B) {
//	var (
//		set = OrderedSet[string]()
//	)
//}

func testHas[T any](t *testing.T, s Sorted[T], value T, expect bool) {
	if has := s.Has(value); has != expect {
		t.Errorf(
			"Sorted.Has resulted in incorrect return value: %t, expected %t for value  %v\n",
			has,
			expect,
			value,
		)
	}
}
