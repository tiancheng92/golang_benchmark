package string_deduplicate

import (
	"testing"

	"github.com/wxnacy/wgo/arrays"
)

var (
	list = [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "a", "b", "z", "k", "i", "p", "q", "a"}
)

func deduplicateBeforeAppend() {
	var res []string
	for i := range list {
		if arrays.Contains(res, list[i]) == -1 {
			res = append(res, list[i])
		}
	}
	_ = res
}

func deduplicateAfterAppend() {
	var res []string
	for i := range list {
		res = append(res, list[i])
	}
	res = arrays.StringsDeduplicate(res)
	_ = res
}

func BenchmarkForDeduplicateBeforeAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deduplicateBeforeAppend()
	}
}

func BenchmarkForDeduplicateAfterAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		deduplicateAfterAppend()
	}
}
