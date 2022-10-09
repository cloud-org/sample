### ex4

测试一下长时间运行，并通过 context 获取到对应的一些值

```log
FC Invoke Start RequestId: 097abb00-d295-42ee-a52b-aa42a5d53d7f
2022-05-31T16:37:50.029Z 097abb00-d295-42ee-a52b-aa42a5d53d7f [INFO] main.go:34: 认证 {AccessKeyId:STS.xxx AccessKeySecret:xxx SecurityToken:xxx}
2022-05-31T16:37:50.029Z 097abb00-d295-42ee-a52b-aa42a5d53d7f [INFO] main.go:35: 服务配置 {Name:xd LogProject:aliyun-fc-cn-shenzhen-xxx LogStore:function-log Qualifier:LATEST VersionId:}
2022-05-31T16:38:00.029Z 097abb00-d295-42ee-a52b-aa42a5d53d7f [INFO] main.go:37: 休眠 10s
FC Invoke End RequestId: 097abb00-d295-42ee-a52b-aa42a5d53d7f
```

发现还是会占用对应的执行时间，也就是说，这一部分会 计入超时时间的范畴 还有进行计费。

那么也许我应该还是需要实现定时 ping 心跳包。然后异步执行对应的函数。