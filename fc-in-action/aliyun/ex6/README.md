<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [ex6](#ex6)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### ex6

测试 client 传数据到云函数

通过 json 序列化成 `[]byte` 即可

```go
// 封装请求体
request := Req{
    Event: "ping",
}
body, err := json.Marshal(&request)
if err != nil {
    log.Println("json marshal err:", err)
    return err
}
```

PS: 定时触发器的问题

定时触发器的默认格式如下：

```json
{
    "triggerTime":"2018-02-09T05:49:00Z",
    "triggerName":"timer-trigger",
    "payload":"awesome-fc"
}
```

注意这个格式就是传到云函数的 json 格式。

而不是 payload 可以写 json 字符串，会报错(如果云函数没有进行处理的话)。

例如我的函数处理的是

```json
{
    "event": "ping"
}
```

这种格式就无法使用定时触发器进行处理，如果要兼容也不是不可以，

改成对应的 key 即可，然后 payload 赋值一下。