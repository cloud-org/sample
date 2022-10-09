### pointers and errors

在 Go 中，当调用一个函数或方法时，参数会被复制。

Go 允许从现有的类型创建新的类型。

```go
type Stringer interface {
	String() string
}
```

实现 String() 方法，fmt %s 的时候可以进行替换

### link

https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/pointers-and-errors

