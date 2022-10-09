package main

import (
	"es-in-action/common"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{common.Endpoint},
	})
	if err != nil {
		log.Println(err)
		return
	}
	//log.Println(elasticsearch.Version)
	log.Println(es.Info())
}
