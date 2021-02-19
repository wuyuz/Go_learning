package main

import (
	"context"
	"crontab/prepare/mongod_usage/util"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type TimePorint struct {
	StartTime int64 `bson:"startTime"` //开始时间
	EndTime   int64 `bson:"endTime"`   //结束时间
}

type LogRecord struct {
	JobName string     `bson:"jobName"` //任务名
	Command string     `bson:"command"` //shell命令
	Err     string     `bson:"err"`     //脚本错误
	Content string     `bson:"content"` //脚本输出
	Tp      TimePorint //执行时间
}

func main() {
	var (
		client = util.GetMgoCli()
		err        error
		collection *mongo.Collection
		result     *mongo.InsertManyResult
		id         primitive.ObjectID
	)
	collection = client.Database("my_db").Collection("test")

	//批量插入
	result, err = collection.InsertMany(context.TODO(), []interface{}{
		LogRecord{
			JobName: "job10",
			Command: "echo 1",
			Err:     "",
			Content: "1",
			Tp: TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		LogRecord{
			JobName: "job10",
			Command: "echo 2",
			Err:     "",
			Content: "2",
			Tp: TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		LogRecord{
			JobName: "job10",
			Command: "echo 3",
			Err:     "",
			Content: "3",
			Tp: TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
	})
	if err != nil{
		log.Fatal(err)
	}
	if result == nil {
		log.Fatal("result nil")
	}
	for _, v := range result.InsertedIDs {
		id = v.(primitive.ObjectID)
		fmt.Println("自增ID", id.Hex())
	}
}
