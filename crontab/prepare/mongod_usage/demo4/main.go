package main

import (
	"context"
	"crontab/prepare/mongod_usage/util"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	var (
		client     = util.GetMgoCli()
		collection *mongo.Collection
		err        error
		cursor     *mongo.Cursor
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("test")
	//filter := bson.M{"jobName": "job10"}

	//按照jobName分组,countJob中存储每组的数目
	groupStage := mongo.Pipeline{bson.D{
		{"$group", bson.D{   // 先group分组
			{"_id", "$jobName"},
			{"countJob", bson.D{
				{"$sum", 1},  // 计数
			}},
		}},
	}}
	if cursor, err = collection.Aggregate(context.TODO(), groupStage, ); err != nil {
		log.Fatal(err)
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	//遍历游标
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}

}
