package string_join

import (
	"fmt"
	"strings"
	"testing"
)

const TestText = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var (
	TenTestTextList      []string
	ThousandTestTextList []string
)

func init() {
	for i := 0; i < 10; i++ {
		TenTestTextList = append(TenTestTextList, TestText)
	}
	for i := 0; i < 1000; i++ {
		ThousandTestTextList = append(ThousandTestTextList, TestText)
	}
}

func BenchmarkForPlusBy10TestText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var res string
		for i := 0; i < 10; i++ {
			res += TestText
		}
		_ = res
	}
}

func BenchmarkForFmtBy10TestText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText)
		_ = res
	}
}

func BenchmarkForStringsJoinBy10TestText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := strings.Join(TenTestTextList, "")
		_ = res
	}
}

func BenchmarkForStringsBuildBy10TestText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := 0
		for j := range TenTestTextList {
			n += len(TenTestTextList[j])
		}
		var res strings.Builder
		res.Grow(n)
		for j := range TenTestTextList {
			res.WriteString(TenTestTextList[j])
		}
		_ = res.String()
	}
}

func BenchmarkForPlusBy1000TestText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var res string
		for i := 0; i < 1000; i++ {
			res += TestText
		}
		_ = res
	}
}

func BenchmarkForFmtBy1000TestText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s", TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText, TestText)
		_ = res
	}
}

func BenchmarkForStringsJoinBy1000TestText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := strings.Join(ThousandTestTextList, "")
		_ = res
	}
}

func BenchmarkForStringsBuildBy1000TestText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n := 0
		for j := range ThousandTestTextList {
			n += len(ThousandTestTextList[j])
		}
		var res strings.Builder
		res.Grow(n)
		for j := range ThousandTestTextList {
			res.WriteString(ThousandTestTextList[j])
		}
		_ = res.String()
	}
}
