package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TimePoint 任务的执行时间点
type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

// LogRecord 一条日志
type LogRecord struct {
	JobName   string    `bson:"jobName"` // 任务名
	Command   string    `bson:"command"`
	Err       string    `bson:"err"`
	Content   string    `bson:"content"`
	TimePoint TimePoint `bson:"timePoint"`
}

// FindByJobName jobName过滤条件
type FindByJobName struct {
	JobName string `bson:"jobName"`
}

func main() {
	var (
		client *mongo.Client
		err    error
		cond   FindByJobName
		cursor *mongo.Cursor
		record *LogRecord
	)

	// 1.建立连接
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		fmt.Println(err)
		return
	}

	// 2.选择数据库
	database := client.Database("cron")

	// 3.选择表
	collect := database.Collection("log")

	// 4.过滤，按照jobName=job10过滤，找出2条
	cond = FindByJobName{JobName: "job10"} // 相当于{"jobName":"job10"}
	//cond2 := bson.M{
	//	"jobName": "job10",
	//}

	// 5.查询
	findOptions := options.Find()
	findOptions.SetSkip(0)
	findOptions.SetLimit(2)
	if cursor, err = collect.Find(context.TODO(), cond, findOptions); err != nil {
		fmt.Println(err)
		return
	}

	// 延迟释放游标
	defer cursor.Close(context.TODO())

	// 6.遍历结果集
	for cursor.Next(context.TODO()) {
		// 定义一个日志对象
		record = &LogRecord{}

		// 反序列化bson到对象
		if err = cursor.Decode(record); err != nil {
			fmt.Println(err)
			return
		}
		// 把日志打印出来
		fmt.Println(*record)
	}

}
