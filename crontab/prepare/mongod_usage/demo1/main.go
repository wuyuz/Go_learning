package main

import (
	"context"
	"crontab/prepare/mongod_usage/util"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		client     = util.GetMgoCli()
		err        error
		collection *mongo.Collection
		lr         LogRecord
		iResult    *mongo.InsertOneResult
		id         primitive.ObjectID
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("my_collection")

	//插入某一条数据
	if iResult, err = collection.InsertOne(context.TODO(), lr); err != nil {
		fmt.Print(err)
		return
	}
	//_id:默认生成一个全局唯一ID
	id = iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID", id.Hex())
}
