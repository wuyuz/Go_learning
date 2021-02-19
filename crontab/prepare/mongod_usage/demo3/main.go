package main

import (
	"context"
	"crontab/prepare/mongod_usage/util"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//查询实体
type FindByJobName struct {
	JobName string `bson:"jobName"` //任务名
}

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
		cursor     *mongo.Cursor
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("test")

	filter := bson.M{"jobName":"job10"}
	if cursor, err = collection.Find(context.TODO(), filter,options.Find().SetSkip(0), options.Find().SetLimit(2) ); err != nil {
		log.Fatal(err)
	}

	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	//遍历游标获取结果数据
	for cursor.Next(context.TODO()) {
		var lr LogRecord
		//反序列化Bson到对象
		if cursor.Decode(&lr) != nil {
			fmt.Print(err)
			return
		}
		//打印结果数据
		fmt.Println(lr)
	}

	//这里的结果遍历可以使用另外一种更方便的方式：
	var results []LogRecord
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
