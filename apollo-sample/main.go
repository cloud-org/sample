package main

import (
	"fmt"
	"github.com/zouyx/agollo/v4"
	"github.com/zouyx/agollo/v4/env/config"
)

func main() {

	configService := "http://192.168.3.36:8080"

	c := &config.AppConfig{
		AppID:         "myapp",
		Cluster:       "default",
		IP:            configService,
		NamespaceName: "application",
		//IsBackupConfig: true,
		//Secret: "c2f72659a11846e88b5e80f21372f358",
	}

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		fmt.Printf("startWithConfig err: %v\n", err)
		return
	}
	fmt.Println("初始化 Apollo 配置成功")

	//Use your apollo key to test
	cache := client.GetConfigCache(c.NamespaceName)
	cache.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	value, err := cache.Get("lease.time")
	if err != nil {
		fmt.Printf("cache.get err: %v\n", err)
		return
	}
	fmt.Println("value is", value)

	return
}
