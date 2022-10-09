<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [qrcode-decode](#qrcode-decode)
- [使用](#%E4%BD%BF%E7%94%A8)
- [致谢](#%E8%87%B4%E8%B0%A2)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

### qrcode-decode

二维码识别工具

### 使用

```sh
# ronething @ ashings-macbook-pro in ~/Documents/qrcode-decode on git:master x [17:17:17] 
$ ./qrcode-decode -h               
qrcode-decode tool
Usage: qrcode [-h help]
Options:
  -h    帮助
  -p string
        图片文件路径 (default "./qrcode.png")

# ronething @ ashings-macbook-pro in ~/Documents/qrcode-decode on git:master x [17:17:21] 
$ ./qrcode-decode -p ./qrcode.jpeg 
2021/02/10 17:17:27 识别链接为 https://support.weixin.qq.com/cgi-bin/mmsupport-bin/showredpacket?receiveuri=xCvTj3TtrzF&check_type=1#wechat_redirect
```

### 致谢

- github.com/tuotoo/qrcode 