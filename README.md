# golang 性能对比

* 测试平台: M1 Macbook Pro
* golang版本: 1.16.2

### 字符串拼接
```shell
cd string_join
go golang_benchmark -bench=. -benchmem -run=none
```

* 结果
```text
goos: darwin
goarch: arm64
pkg: test/string_join
BenchmarkForPlusShortText-8             1000000000               0.3198 ns/op          0 B/op          0 allocs/op
BenchmarkForFmtText-8                    6477994               184.7 ns/op            48 B/op          1 allocs/op
BenchmarkForStringsJoinShortText-8      13201954                91.08 ns/op           48 B/op          1 allocs/op
BenchmarkForPlusLongText-8                     2         961963916 ns/op        20384446536 B/op          100507 allocs/op
BenchmarkForStringsJoinLongText-8            372           3218276 ns/op         9648757 B/op         31 allocs/op
PASS
ok      test/string_join        7.986s
 ```

* 结论：
    * 短文本：使用"+"拼接 或 strings.Join()拼接皆可
    * 长文本：使用strings.Join()拼接

### 遍历列表
```shell
cd for_range
go golang_benchmark -bench=. -benchmem -run=none
```
* 结果
```text
goos: darwin
goarch: arm64
pkg: test/for_range
BenchmarkForSlice-8                         3404            317767 ns/op               0 B/op          0 allocs/op
BenchmarkForRangeKeySlice-8                 3674            314000 ns/op               0 B/op          0 allocs/op
BenchmarkForRangeSlice-8                     471           2537942 ns/op               0 B/op          0 allocs/op
BenchmarkForSlicePtr-8                      3818            314358 ns/op               0 B/op          0 allocs/op
BenchmarkForRangeKeySlicePtr-8              3817            314096 ns/op               0 B/op          0 allocs/op
BenchmarkForRangeSlicePtr-8                 3816            314184 ns/op               0 B/op          0 allocs/op
PASS
ok      test/for_range  8.525s
```

* 结论：
    * 列表中为对象：不建议直接获取其中的value，而是使用index去间接获取
    * 列表中存指针：性能基本无差异
  
### sync.pool
```shell
cd sync_pool
go golang_benchmark -bench=. -benchmem -run=none
```

* 结果
```text
goos: darwin
goarch: arm64
pkg: test/sync_pool
BenchmarkBufferWithPool-8        6725568               151.1 ns/op             0 B/op          0 allocs/op
BenchmarkBufferWithoutPool-8     1951632               621.9 ns/op         10240 B/op          1 allocs/op
PASS
ok      test/sync_pool  3.800s
```

* 结论
  * sync.pool能有效减少内存调用，复用对象减少GC压力
  
### json序列化
```shell
cd json_marshal
go mod tidy
go golang_benchmark -bench=. -benchmem -run=none
```

* 结果
```text
goos: darwin
goarch: arm64
pkg: test/json_marshal
BenchmarkUseJson-8                746625              1553 ns/op             800 B/op         18 allocs/op
BenchmarkUseJsoniter-8           1936796               620.8 ns/op           560 B/op         14 allocs/op
PASS
ok      test/json_marshal       4.381s
```

* 结论
  * 使用第三方包[Jsoniter](http://github.com/json-iterator/go)能有效加快序列化的速度、减少内存压力
  
### String 与 Byte 转换
```shell
cd byte_string_conversion
go golang_benchmark -bench=. -benchmem -run=none
```

* 结果
```text
goos: darwin
goarch: arm64
pkg: test/byte_string_conversion
BenchmarkForNewByteToString-8           1000000000               0.3222 ns/op          0 B/op          0 allocs/op
BenchmarkForByteToString-8              16624393                72.54 ns/op          704 B/op          1 allocs/op
BenchmarkForNewStringToByte-8           1000000000               0.3131 ns/op          0 B/op          0 allocs/op
BenchmarkForStringToByte-8              14617010                80.47 ns/op          704 B/op          1 allocs/op
PASS
ok      test/byte_string_conversion     3.830s
```

* 结论
  * 避免对象的复制，有效加快了转换的效率，但会缺失string不可变的属性
  
  
... 未完待续