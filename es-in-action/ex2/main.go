package main

import (
	"context"
	"es-in-action/common"
	"github.com/olivere/elastic/v7"
	"log"
)

func main() {

	client, err := elastic.NewClient(
		elastic.SetURL(common.Endpoint),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Println(err)
		return
	}

	version, err := client.ElasticsearchVersion(common.Endpoint)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("version: %v\n", version)

	item := common.JobRet{
		HostId: "host1",
		CronId: "cron1",
		Ctime:  100,
	}
	resp, err := client.Index().Index("cron").BodyJson(&item).Do(context.Background())
	if err != nil {
		log.Println("do err:", err)
		return
	}

	log.Printf("resp is %+v\n", resp)

	return
}
