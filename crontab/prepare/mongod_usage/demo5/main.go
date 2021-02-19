package main

import (
	"context"
	"crontab/prepare/mongod_usage/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

//更新实体
type UpdateByJobName struct {
	Command string `bson:"command"` //shell命令
	Content string `bson:"content"` //脚本输出
}

func main() {
	var (
		client     = util.GetMgoCli()
		collection *mongo.Collection
		err        error
		uResult    *mongo.UpdateResult
		//upsertedID model.LogRecord
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("test")
	filter := bson.M{"jobName": "job10"}
	//update := bson.M{"$set": bson.M{"command": "ByBsonM",}}
	update := bson.M{"$push": bson.M{ "interests": "Golang" }}
	//update := bson.M{"$set": model.LogRecord{JobName:"job10",Command:"byModel"}}
	if uResult, err = collection.UpdateMany(context.TODO(), filter, update); err != nil {
		log.Fatal(err)
	}
	//uResult.MatchedCount表示符合过滤条件的记录数，即更新了多少条数据。
	log.Println(uResult.MatchedCount)
}
