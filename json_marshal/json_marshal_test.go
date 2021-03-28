package json_marshal

import (
	"encoding/json"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

var new_json = jsoniter.ConfigCompatibleWithStandardLibrary

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

func useJson() {
	var t Tmp
	_ = json.Unmarshal([]byte(data), &t)
}

func useJsoniter() {
	var t Tmp
	_ = new_json.Unmarshal([]byte(data), &t)
}

func BenchmarkUseJson(b *testing.B) {
	for n := 0; n < b.N; n++ {
		useJson()
	}
}

func BenchmarkUseJsoniter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		useJsoniter()
	}
}
