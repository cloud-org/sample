<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [ex5](#ex5)
- [小结](#%E5%B0%8F%E7%BB%93)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### ex5

测试长时间运行+心跳包+执行完成之后进行消息通知

经过测试，`time.Sleep(10*time.Second)` 像是被暂停下来一样。需要通过 ping 不断发心跳，这是有用的。只不过间隔秒数看起来可能是要比较短的。例如 1、2s

如果是这样的，其实发心跳的意义不是很大，因为如果要这样保持的话，相当于每分钟可能要 60 次探活，12 分钟就需要 360 次探活。

如果是同步执行 12 分钟，则 12 * 60 = 720s，GB-s = 128MB/1024 * 720s = 90GB-s

免费额度：

- 请求次数：100w / 360 = 2777 次(如果使用 ping 的话)
- 资源使用：40w / 90 = 4444 次(如果使用同步方法的话)

如果只是使用 2 分钟

- 请求次数：100w/120 = 8333(ping 保持心跳)
- 资源使用：40w/15 = 26666(同步)

计费问题：

- 请求次数：0.0133元/万次 能够使用 27 次
- 资源使用：0.000110592元/GB-秒* 90 = 0.00995328 大概可以使用一次

```log
2022-06-02 00:18:57FC Invoke Start RequestId: 8c54a949-5420-46d9-ae87-b1ee6ded9e8e
2022-06-02 00:18:572022-06-02 00:18:57 8c54a949-5420-46d9-ae87-b1ee6ded9e8e [INFO] main.go:49: 请求 id 8c54a949-5420-46d9-ae87-b1ee6ded9e8e
2022-06-02 00:18:57FC Invoke End RequestId: 8c54a949-5420-46d9-ae87-b1ee6ded9e8e
2022-06-02 00:18:572022-06-02 00:18:57 8c54a949-5420-46d9-ae87-b1ee6ded9e8e [INFO] main.go:64: 服务配置 8c54a949-5420-46d9-ae87-b1ee6ded9e8e
2022-06-02 00:19:01FC Invoke Start RequestId: acb6d569-dc30-4404-bdd3-6530a5b8436d
2022-06-02 00:19:012022-06-02 00:19:01 acb6d569-dc30-4404-bdd3-6530a5b8436d [INFO] main.go:59: 2022-06-01T16:19:01Z
2022-06-02 00:19:01FC Invoke End RequestId: acb6d569-dc30-4404-bdd3-6530a5b8436d
2022-06-02 00:19:02FC Invoke Start RequestId: 6d948034-6929-4086-982d-4b7e0615f163
2022-06-02 00:19:022022-06-02 00:19:02 6d948034-6929-4086-982d-4b7e0615f163 [INFO] main.go:59: 2022-06-01T16:19:02Z
2022-06-02 00:19:02FC Invoke End RequestId: 6d948034-6929-4086-982d-4b7e0615f163
2022-06-02 00:19:04FC Invoke Start RequestId: 062a7e5f-c4e9-4af8-9f1d-029c408ab601
2022-06-02 00:19:042022-06-02 00:19:04 062a7e5f-c4e9-4af8-9f1d-029c408ab601 [INFO] main.go:59: 2022-06-01T16:19:04Z
2022-06-02 00:19:04FC Invoke End RequestId: 062a7e5f-c4e9-4af8-9f1d-029c408ab601
2022-06-02 00:19:05FC Invoke Start RequestId: c1a1cc1e-6e24-4803-ab7e-e2fceb6d71e0
2022-06-02 00:19:052022-06-02 00:19:05 c1a1cc1e-6e24-4803-ab7e-e2fceb6d71e0 [INFO] main.go:59: 2022-06-01T16:19:05Z
2022-06-02 00:19:05FC Invoke End RequestId: c1a1cc1e-6e24-4803-ab7e-e2fceb6d71e0
2022-06-02 00:19:06FC Invoke Start RequestId: e3512ceb-99fc-4c58-9160-c8b931704e11
2022-06-02 00:19:062022-06-02 00:19:06 e3512ceb-99fc-4c58-9160-c8b931704e11 [INFO] main.go:59: 2022-06-01T16:19:06Z
2022-06-02 00:19:06FC Invoke End RequestId: e3512ceb-99fc-4c58-9160-c8b931704e11
2022-06-02 00:19:072022-06-02 00:19:07 8c54a949-5420-46d9-ae87-b1ee6ded9e8e [INFO] main.go:66: 休眠 10s
```

### 小结

可以考虑还是使用同步的方法，然后提供多一个 ping 方法，在发送正式请求前发起 ping，避免冷启动。