package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

type LogRecord struct {
	JobName   string    `bson:"jobName"`
	Command   string    `bson:"command"`
	Err       string    `bson:"err"`
	Content   string    `bson:"content"`
	TimePoint TimePoint `bson:"timePoint"`
}

type FindByJobName struct {
	JobName string `bson:"jobName"`
}

func main() {

	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		cond       *FindByJobName
		findOpt    *options.FindOptions
		cur        *mongo.Cursor
		record     *LogRecord
	)

	// 1、建立连接
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017")); err != nil {
		fmt.Println(err)
		return
	}
	// 2、选择数据库
	database = client.Database("cron")
	// 3、选择表
	collection = database.Collection("log")
	// 4、按照 jobName 字段过滤，找出 jobName=job10，找出 5 条
	cond = &FindByJobName{JobName: "job10"}

	findOpt = options.Find().SetSkip(0).SetLimit(2)
	if cur, err = collection.Find(context.TODO(), cond, findOpt); err != nil {
		fmt.Println(err)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		record = &LogRecord{}
		//反序列化
		if err = cur.Decode(record); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(*record)
	}

}
