package number_string_conversion

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	Int   = 123456789
	Float = 1.23456789
)

func BenchmarkForUseFmtInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := fmt.Sprint(Int)
		_ = j
	}
}

func BenchmarkForUseStrconvInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		j := strconv.Itoa(Int)
		_ = j
	}
}

func BenchmarkForUseFmtFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := fmt.Sprintf("%.3f", Float)
		_ = f
	}
}

func BenchmarkForUseStrconvFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := strconv.FormatFloat(Float, 'f', 3, 64)
		_ = f
	}
}
