package master

import (
	"context"
	"crontab/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	G_logMgr *LogMgr
)

type LogMgr struct {
	client        *mongo.Client
	logCollection *mongo.Collection
}

func InitLogMgr() (err error) {
	var (
		client *mongo.Client
	)

	clientOptions := options.Client().ApplyURI(G_config.MongodbUri)

	// 建立mongodb连接
	if client, err = mongo.Connect(
		context.TODO(),
		clientOptions,
	); err != nil {
		return
	}

	G_logMgr = &LogMgr{
		client:        client,
		logCollection: client.Database("cron").Collection("log"),
	}
	return
}

func (logMgr *LogMgr) ListLog(name string, skip, limit int64) (logArr []*common.JobLog, err error) {

	var (
		filter  *common.JobLogFilter
		logSort *common.SortLogByStartTime
		findopt *options.FindOptions
		cursor  *mongo.Cursor
		jobLog  *common.JobLog
	)
	logArr = make([]*common.JobLog, 0)

	// 过滤条件
	filter = &common.JobLogFilter{JobName: name}

	// 按照任务开始时间倒叙
	logSort = &common.SortLogByStartTime{SortOrder: -1}
	findopt = options.Find()

	// 查询
	if cursor, err = logMgr.logCollection.Find(context.TODO(), filter, findopt.SetSort(logSort), findopt.SetSkip(skip), findopt.SetLimit(limit)); err != nil {
		return
	}

	// 遍历游标
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		jobLog = &common.JobLog{}

		// 反序列化Bson
		if err = cursor.Decode(jobLog); err != nil {
			continue
		}
		logArr = append(logArr, jobLog)

	}

	return
}
