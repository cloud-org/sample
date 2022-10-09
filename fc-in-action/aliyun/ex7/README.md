<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [ex7](#ex7)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### ex7

测试一下 esc 通过 vpc 地址访问云函数

```log
$ DEBUG=tea ./ex7 
2022/06/02 12:09:36 args is []
> POST https://accountId.cn-shenzhen-internal.fc.aliyuncs.com/2021-04-06/services/xd/functions/func-15i9eih2/invocations
> user-agent: AlibabaCloud (linux; amd64) Golang/1.16.2 Core/0.01 TeaDSL/1
> x-acs-date: 2022-06-02T04:09:36Z
> content-type: application/json; charset=utf-8
> x-acs-version: 2021-04-06
> x-acs-signature-nonce: 58b69344xxxxxx0eb715d9d9b5a62837
> accept: application/json
> host: accountId.cn-shenzhen-internal.fc.aliyuncs.com
> X-Fc-Invocation-Type: Sync
```

开启 debug 可以看到 post 地址带了 internal 表示走内网

- aliyun 服务地址(endpoint): https://help.aliyun.com/document_detail/52984.html