package sync_pool

import (
	"bytes"
	"sync"
	"testing"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

func bufferWithPool() {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		bufferPool.Put(buf)
	}()
	buf.Write(data)
}

func bufferWithoutPool() {
	var buf bytes.Buffer
	buf.Write(data)
}

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bufferWithPool()
	}
}

func BenchmarkBufferWithoutPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bufferWithoutPool()
	}
}
