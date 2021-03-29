package number_string_conversion

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	Int   = 123456
	Float = 123.456
)

func useFmtInt() {
	i := fmt.Sprint(Int)
	_ = i
}

func useFmtFloat() {
	f := fmt.Sprintf("%.2f", Float)
	_ = f
}

func useStrconvInt() {
	i := strconv.Itoa(Int)
	_ = i
}

func useStrconvFloat() {
	f := strconv.FormatFloat(Float, 'f', 2, 32)
	_ = f
}

func BenchmarkForUseFmtInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useFmtInt()
	}
}

func BenchmarkForUseFmtFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useFmtFloat()
	}
}

func BenchmarkForUseStrconvInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useStrconvInt()
	}
}

func BenchmarkForUseStrconvFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		useStrconvFloat()
	}
}
