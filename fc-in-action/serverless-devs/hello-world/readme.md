<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [hello-world](#hello-world)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### hello-world

s 工具部署 hello-world

- 编译

```
make build
```

- 部署

```
make deploy
```

- 调用

```
s invoke -f helloworld.json
```

output:

```log
$ s invoke -f helloworld.json                         
Reading event file content:
{
    "name": "panda"
}

========= FC invoke Logs begin =========
FC Invoke Start RequestId: 409626e4-b47b-xxxx-bffb-8a14d3249c51
2022-06-05T02:38:06.488Z: 409626e4-b47b-xxxx-bffb-8a14d3249c51 [INFO]  hello golang! from hook
2022-06-05T02:38:06.488Z: 409626e4-b47b-xxxx-bffb-8a14d3249c51 [INFO]  name is panda
FC Invoke End RequestId: 409626e4-b47b-xxxx-bffb-8a14d3249c51

Duration: 1.28 ms, Billed Duration: 2 ms, Memory Size: 128 MB, Max Memory Used: 7.95 MB
========= FC invoke Logs end =========

FC Invoke instanceId: c-xxxx-0e4a541c26bb42c6a1ec

FC Invoke Result:
{    "name": "panda"}


End of method: invoke
```