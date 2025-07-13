测试当前项目 GOPATH 模式下,是否能访问 gomod4a 项目的`v1`,`v2`版本

答案: 只能引用最新的 main 版本,结果是 503, `gomod5a_v1.AddA(1,2)= 503`

- v1

```go
func AddA(a, b int) int {
	return a + b
}
```

- v2

```go
func AddA(a, b int) int {
	return a + b + 200
}
```

- latest main

```go
func AddA(a, b int) int {
	return a + b + 500
}
```
