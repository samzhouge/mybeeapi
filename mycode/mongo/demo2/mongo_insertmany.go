package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTime   int64 `bson:"endTime"`
}

type LogRecord struct {
	JobName   string    `bson:"jobName"` // 任务名
	Command   string    `bson:"command"`
	Err       string    `bson:"err"`
	Content   string    `bson:"content"`
	TimePoint TimePoint `bson:"timePoint"`
}

func main() {
	var (
		client   *mongo.Client
		err      error
		record   *LogRecord
		result   *mongo.InsertManyResult
		logArr   []interface{}
		insertId interface{}
		docId    primitive.ObjectID
	)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		fmt.Println(err)
		return
	}

	database := client.Database("cron")

	collect := database.Collection("log")

	record = &LogRecord{
		JobName:   "job10",
		Command:   "echo hello",
		Err:       "",
		Content:   "hello",
		TimePoint: TimePoint{StartTime: time.Now().Unix(), EndTime: time.Now().Unix() + 10},
	}

	logArr = []interface{}{record, record, record}

	if result, err = collect.InsertMany(context.TODO(), logArr); err != nil {
		fmt.Println(err)
		return
	}

	for _, insertId = range result.InsertedIDs {
		fmt.Println(insertId)
		docId = insertId.(primitive.ObjectID)
		fmt.Println("自增ID:", docId.Hex())
	}
}
