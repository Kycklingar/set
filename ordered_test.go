package set

import (
	"math/rand"
	"testing"

	"golang.org/x/exp/constraints"
)

func testOrdHas[T constraints.Ordered](t *testing.T, s Ordered[T], value T, expect bool) {
	if has := s.Has(value); has != expect {
		t.Errorf(
			"Ordered.Has resulted in incorrect return value: %t, expected %t for value  %v\n",
			has,
			expect,
			value,
		)
	}
}

func TestOrderedAppend(t *testing.T) {
	var set = NewOrdered[int]()
	set.Append(3, 4, 5, 6, 7)
	set.Append(1, 2, 3, 4, 5)
	set.Append(6, 7, 8, 9, 10)

	for i := 1; i < 11; i++ {
		testOrdHas(t, set, i, true)
	}
}

func TestOrderedSet(t *testing.T) {
	var set = NewOrdered[string]("a", "b")
	set.Set("a")
	set.Set("c")

	testOrdHas(t, set, "a", true)
	testOrdHas(t, set, "b", true)
	testOrdHas(t, set, "c", true)
	testOrdHas(t, set, "d", false)
}

func TestOrderedGet(t *testing.T) {
	var set = NewOrdered[rune]('a', 'b', 'c')
	v, ok := set.Get('b')
	if !ok {
		t.Errorf("Ordered.Get is not ok")
	}
	if v != 'b' {
		t.Errorf("Ordered.Get returned wrong value")
	}
}

func BenchmarkOrderedSetInt10k(b *testing.B) {
	benchmarkOrderedSetInt(b, 10000)
}

func BenchmarkOrderedSetInt100k(b *testing.B) {
	benchmarkOrderedSetInt(b, 100000)
}

func benchmarkOrderedSetInt(b *testing.B, numItems int) {
	var (
		set     = NewOrdered[int]()
		ints    = generateRandomInts(numItems)
		indices = generateRandomIndices(b.N, len(ints))
	)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		set.Set(ints[indices[n]])
	}
}

func generateRandomStrings(size int) []string {
	var runes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	var strings []string
	for i := 0; i < size; i++ {
		var str []rune
		for n := 0; n < 8; n++ {
			str = append(str, runes[rand.Int()%len(runes)])
		}
		strings = append(strings, string(str))
	}

	return strings
}

func BenchmarkOrderedSetString10k(b *testing.B) {
	benchmarkOrderedSetString(b, 10000)
}

func BenchmarkOrderedSetString100k(b *testing.B) {
	benchmarkOrderedSetString(b, 100000)
}

func benchmarkOrderedSetString(b *testing.B, numItems int) {
	var (
		set     = NewOrdered[string]()
		strings = generateRandomStrings(numItems)
		indices = generateRandomIndices(b.N, len(strings))
	)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		set.Set(strings[indices[n]])
	}
}
