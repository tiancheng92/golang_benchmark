package byte_string_conversion

import (
	"testing"
	"unsafe"
)

const text = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func stringToByte(s string) []byte {
	return []byte(s)
}

func byteToString(b []byte) string {
	return string(b)
}

func unsafeStringToByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func unsafeByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func BenchmarkForByteToString(b *testing.B) {
	tmpByte := []byte(text)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := byteToString(tmpByte)
		_ = res
	}
}

func BenchmarkForUnsafeByteToString(b *testing.B) {
	tmpByte := []byte(text)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res := unsafeByteToString(tmpByte)
		_ = res
	}
}

func BenchmarkForStringToByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := stringToByte(text)
		_ = res
	}
}

func BenchmarkForUnsafeStringToByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := unsafeStringToByte(text)
		_ = res
	}
}
