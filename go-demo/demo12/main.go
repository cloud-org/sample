package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// 序列化之后 {"$lt":timestamp}
type TimeBeforeCond struct {
	Before int64 `bson:"$lt"`
}

// {"timePoint.startTime":{"$lt":timestamp}}
type DeleteCond struct {
	BeforeCond TimeBeforeCond `bson:"timePoint.startTime"`
}

func main() {

	var (
		client     *mongo.Client
		err        error
		database   *mongo.Database
		collection *mongo.Collection
		delCond    *DeleteCond
		delResult  *mongo.DeleteResult
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

	//	4、要删除开始时间早于当前时间的所有日志
	delCond = &DeleteCond{BeforeCond: TimeBeforeCond{Before: time.Now().Unix()}}

	if delResult, err = collection.DeleteMany(context.TODO(), delCond); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("删除的行数", delResult.DeletedCount)
}
