package main

import (
	"context"
	"es-in-action/common"
	"github.com/olivere/elastic/v7"
	"log"
)

// bulk write records
func main() {

	esClient, err := elastic.NewClient(
		elastic.SetURL(common.Endpoint),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Println(err)
		return
	}

	mockData := []common.JobRet{
		{
			Id:     "1",
			HostId: "host1",
			CronId: "cron1",
			Ctime:  100,
		},
		{
			Id:     "2",
			HostId: "host2",
			CronId: "cron1",
			Ctime:  101,
		},
		{
			Id:     "3",
			HostId: "host3",
			CronId: "cron1",
			Ctime:  102,
		},
	}

	bulkService := esClient.Bulk().Index(common.JobRetIndex)

	for _, item := range mockData {
		// need set id
		item := item // important
		req := elastic.NewBulkCreateRequest().Index(common.JobRetIndex).Doc(&item).Id(item.Id)
		bulkService.Add(req)
	}

	resp, err := bulkService.Do(context.TODO())
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v %v\n", resp.Took, resp.Errors)

	flushResp, err := esClient.Flush().Index(common.JobRetIndex).Do(context.TODO())
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("flush resp: %+v\n", flushResp.Shards)

	return
}
