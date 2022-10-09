<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [调用函数文档示例](#%E8%B0%83%E7%94%A8%E5%87%BD%E6%95%B0%E6%96%87%E6%A1%A3%E7%A4%BA%E4%BE%8B)
  - [执行步骤](#%E6%89%A7%E8%A1%8C%E6%AD%A5%E9%AA%A4)
  - [使用的 API](#%E4%BD%BF%E7%94%A8%E7%9A%84-api)
  - [报错](#%E6%8A%A5%E9%94%99)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 调用函数文档示例

该项目为调用InvokeFunction接口调用执行函数。文档示例，该示例**无法在线调试**，如需调试可下载到本地后替换 [AK](https://usercenter.console.aliyun.com/#/manage/ak) 以及参数后进行调试。

### 执行步骤
下载的代码包，在根据自己需要更改代码中的参数和 AK 以后，可以在**解压代码所在目录下**按如下的步骤执行

- Go
*Golang 版本要求 1.13 及以上*
```sh
GOPROXY=https://goproxy.cn,direct go run ./main
```

- Python
*Python 版本要求 Python3*
```sh
python3 setup.py install && python ./alibabacloud_sample/sample.py
```

### 使用的 API

-  InvokeFunction 调用InvokeFunction接口调用执行函数。文档示例，可以参考：[文档](https://next.api.aliyun.com/document/FC-Open/2021-04-06/InvokeFunction)


### 报错

```
   StatusCode: 0
   Code: InvalidArgument
   Message: code: 400, Function with http trigger(name: defaultTrigger) can only be invoked with http trigger URL request id: <nil>
   Data: {"Code":"InvalidArgument","Message":"Function with http trigger(name: defaultTrigger) can only be invoked with http trigger URL","RequestID":"ab40531c-46c1-413b-b1ee-cd84fdfa1199"}
```

注意如果是 http trigger，需要直接指定 endpoint 为对应 http 触发器服务的 url 即可。

而不是用账号 id + region 服务地址组成，会报错 