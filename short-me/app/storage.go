package main

type Storage interface {
	// 生成短地址接口
	Shorten(url string, exp int64) (string, error)
	// 获取短地址信息接口
	ShortLinkInfo(eid string) (interface{}, error)
	// 获取长链接接口
	UnShorten(eid string) (string, error)
}
