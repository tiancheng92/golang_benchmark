package for_range

import "testing"

const N = 1000000

type Tmp struct {
	A string
	B string
	C int
	D bool
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
