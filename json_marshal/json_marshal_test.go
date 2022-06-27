package json_marshal

import (
	"encoding/json"
	"testing"

	go_json "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
)

var data = `{
"a":"a",
"b":1,
"c":true,
"d":{"text":1,"text1":"text"}
}`

type Tmp struct {
	A string                 `json:"a"`
	B int                    `json:"b"`
	C bool                   `json:"c"`
	D map[string]interface{} `json:"d"`
}

func BenchmarkUseEncodingJson(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t Tmp
		_ = json.Unmarshal([]byte(data), &t)
	}
}

func BenchmarkUseJsoniter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t Tmp
		_ = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(data), &t)
	}
}

func BenchmarkUseGoJson(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t Tmp
		_ = go_json.Unmarshal([]byte(data), &t)
	}
}
