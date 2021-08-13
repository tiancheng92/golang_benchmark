package byte_string_conversion

import (
	"testing"
	"unsafe"
)

const text = "text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text text"

func stringToByte(s string) {
	res := []byte(s)
	_ = res
}

func byteToString(b []byte) {
	res := string(b)
	_ = res
}

func newStringToByte(s string) {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	res := *(*[]byte)(unsafe.Pointer(&h))
	_ = res
}

func newByteToString(b []byte) {
	res := *(*string)(unsafe.Pointer(&b))
	_ = res
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func BenchmarkForNewByteToString(b *testing.B) {
	tmpByte := []byte(text)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		newByteToString(tmpByte)
	}
}

func BenchmarkForByteToString(b *testing.B) {
	tmpByte := []byte(text)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byteToString(tmpByte)
	}
}

func BenchmarkForNewStringToByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newStringToByte(text)
	}
}

func BenchmarkForStringToByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringToByte(text)
	}
}

func BenchmarkForStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytes(text)
	}
}
