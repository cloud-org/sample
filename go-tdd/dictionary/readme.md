## map

由于 nil 指针异常，你永远不应该初始化一个空的 map 变量：

```go
var m map[string]string
```

相反，你可以像我们上面那样初始化空 map，或使用 make 关键字创建 map：

```go
dictionary = map[string]string{}
// OR
dictionary = make(map[string]string)
```

这两种方法都可以创建一个空的 hash map 并指向 dictionary。这确保永远不会获得 nil 指针异常。