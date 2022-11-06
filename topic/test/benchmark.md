
## 性能测试

性能测试函数以 Benchmark 开头，参数类型是 *testing.B，  
可与 Test 函数 放在 同个文件中。  
默认情况下，go test 不执行 Benchmark 测试，必须用 `-bench <pattern>` 指定 性能测试 函数。
`go test -bench=.`  






B 类型 也有 以下参数：

- benchmem：输出 内存 分配统计
- benchtime：指定 测试时间
- cpu：指定 GOMAXPROCS
- timeout：超时限制


`go test -v -bench=. -cpu=8 -benchtime="3s" -timeout="5s" -benchmem`  


- Benchmark-8：-cpu 参数指定，-8 表示 8 个 CPU 线程执行
- 5000000000：表示总共执行了 5000000000次
- 0.34 ns/op：表示每次执行耗时 0.34 纳秒
- 0 B/op: 表示 每次 执行 分配的 内存（字节）
- 0 allocs/op：表示 每次执行 分配了 多少次对象