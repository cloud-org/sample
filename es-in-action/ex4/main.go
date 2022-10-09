package main

import (
	"context"
	"encoding/json"
	"es-in-action/common"
	"log"
	"time"

	"github.com/olivere/elastic/v7"
)

// search group by and top hits sort ctime.desc see issue #2
func main() {
	esClient, err := elastic.NewClient(
		elastic.SetURL(common.Endpoint),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Println(err)
		return
	}

	start := time.Now()
	query := elastic.NewTermQuery("cronId", "cron1")
	// agg2 size is top hit size
	agg2 := elastic.NewTopHitsAggregation().Size(1).Sort("ctime", false)
	// agg1.size set hostIds length
	agg1 := elastic.NewTermsAggregation().Size(100).Field("hostId.keyword").
		SubAggregation("cron_docs", agg2)
	resp, err := esClient.Search().Index(common.JobRetIndex).Query(query).Size(0).
		Aggregation("cron", agg1).Do(context.TODO())
	if err != nil {
		log.Println("search err:", err)
		return
	}

	//log.Printf("total hits: %v\n", resp.TotalHits())

	//log.Printf("%+v\n", resp.Aggregations)
	message, ok := resp.Aggregations["cron"]
	if !ok {
		log.Println("get cron not ok")
		return
	}

	var Agg common.Cron
	if err = json.Unmarshal(message, &Agg); err != nil {
		log.Println(err)
		return
	}

	log.Println("time duration: ", time.Since(start))

	for _, value := range Agg.Buckets {
		for _, item := range value.CronDocs.Hits.Hits {
			log.Printf("%+v\n", item.Source)
		}
	}

}
