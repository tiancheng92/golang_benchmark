package string_join

import (
	"fmt"
	"strings"
	"testing"
)

const text = "text"

// 使用 + 拼接 短文本
func plusShortText() {
	res := text + text + text + text + text + text + text + text + text + text
	_ = res
}

// 使用 + 拼接 长文本
func plusLongText() {
	var res string
	for i := 0; i < 100000; i++ {
		res += text
	}
	_ = res
}

// 使用 fmt包 拼接文本
func fmtText() {
	res := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", text, text, text, text, text, text, text, text, text, text)
	_ = res
}

// 使用 strings.join 拼接 短文本
func stringsJoinShortText() {
	res := strings.Join([]string{text, text, text, text, text, text, text, text, text, text}, "")
	_ = res
}

// 使用 strings.join 拼接 长文本
func stringsJoinLongText() {
	var list []string
	for i := 0; i < 100000; i++ {
		list = append(list, text)
	}
	res := strings.Join(list, "")
	_ = res
}

func BenchmarkForPlusShortText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusShortText()
	}
}

func BenchmarkForFmtText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmtText()
	}
}

func BenchmarkForStringsJoinShortText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsJoinShortText()
	}
}

func BenchmarkForPlusLongText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusLongText()
	}
}

func BenchmarkForStringsJoinLongText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringsJoinLongText()
	}
}
