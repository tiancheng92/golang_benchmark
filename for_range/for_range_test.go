package for_range

import "testing"

const N = 1000000

type Tmp struct {
	A string
	B string
	C int
	D bool
}

func initMap() map[int]string {
	s := make(map[int]string, N)
	for i := 0; i < N; i++ {
		s[i] = "1"
	}
	return s
}

func initLagerMap() map[int]string {
	st := ""
	for j := 0; j < 100000; j++ {
		st += "1"
	}
	s := make(map[int]string, N)
	for i := 0; i < N; i++ {
		s[i] = st
	}
	return s
}

func initSlice() []Tmp {
	s := make([]Tmp, N)
	for i := 0; i < N; i++ {
		s[i] = Tmp{
			A: "1",
			B: "2",
			C: i,
			D: true,
		}
	}
	return s
}

func initSlicePtr() []*Tmp {
	s := make([]*Tmp, N)
	for i := 0; i < N; i++ {
		s[i] = &Tmp{
			A: "1",
			B: "2",
			C: i,
			D: true,
		}
	}
	return s
}

func ForSlice(s []Tmp) {
	for i := 0; i < len(s); i++ {
		a, b := i, s[i]
		_, _ = a, b
	}
}

func ForRangeKeySlice(s []Tmp) {
	for i := range s {
		a, b := i, s[i]
		_, _ = a, b
	}
}

func ForRangeSlice(s []Tmp) {
	for i, v := range s {
		a, b := i, v
		_, _ = a, b
	}
}

func ForSlicePtr(s []*Tmp) {
	for i := 0; i < len(s); i++ {
		a, b := i, s[i]
		_, _ = a, b
	}
}

func ForRangeKeySlicePtr(s []*Tmp) {
	for i := range s {
		a, b := i, s[i]
		_, _ = a, b
	}
}

func ForRangeSlicePtr(s []*Tmp) {
	for i, v := range s {
		a, b := i, v
		_, _ = a, b
	}
}

func ForRangeMapWithValue(s map[int]string) {
	for i, v := range s {
		a, b := i, v
		_, _ = a, b
	}
}

func ForRangeMapWithoutValue(s map[int]string) {
	for i := range s {
		a, b := i, s[i]
		_, _ = a, b
	}
}

func BenchmarkForMapWithValue(b *testing.B) {
	s := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForRangeMapWithValue(s)
	}
}

func BenchmarkForMapWithoutValue(b *testing.B) {
	s := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForRangeMapWithoutValue(s)
	}
}

func BenchmarkForLargeMapWithValue(b *testing.B) {
	s := initLagerMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForRangeMapWithValue(s)
	}
}

func BenchmarkForLargeMapWithoutValue(b *testing.B) {
	s := initLagerMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForRangeMapWithoutValue(s)
	}
}

func BenchmarkForSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForSlice(s)
	}
}

func BenchmarkForRangeKeySlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForRangeKeySlice(s)
	}
}

func BenchmarkForRangeSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForRangeSlice(s)
	}
}

func BenchmarkForSlicePtr(b *testing.B) {
	s := initSlicePtr()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForSlicePtr(s)
	}
}

func BenchmarkForRangeKeySlicePtr(b *testing.B) {
	s := initSlicePtr()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForRangeKeySlicePtr(s)
	}
}

func BenchmarkForRangeSlicePtr(b *testing.B) {
	s := initSlicePtr()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForRangeSlicePtr(s)
	}
}
