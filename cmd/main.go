package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/guregu/dynamo"
	"math/rand"
	"time"
)

type Track struct {
	DeviceID  string `dynamo:"deviceID,hash"`
	Timestamp string  `dynamo:"timestamp, range"`
	Value     int    `dynamo:"value"`
}

const AWS_REGION = "ap-northeast-1"
const DYNAMO_ENDPOINT = "http://localhost:8000"
const TIME_FORMAT = "2006-01-02 15:04:05"

func main() {

	// クライアントの設定
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_REGION),
		Endpoint:    aws.String(DYNAMO_ENDPOINT),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		panic(err)
	}
	db := dynamo.New(sess)
	table := db.Table("tracks")

	// 引数のデバイスIDを変数化
	deviceID := flag.String("deviceID", "1", "Specify the device ID.")
	flag.Parse()

	from := time.Now().Format(TIME_FORMAT)

	for i := 0; i < 5; i++ {
		// ランダムで100以下の整数の生成
		timestamp := time.Now()
		rand.Seed(timestamp.Unix())
		value := rand.Intn(100)

		err = table.Put(&Track{DeviceID: *deviceID, Timestamp: timestamp.Format(TIME_FORMAT), Value: value}).Run()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 1)
	}
	to := time.Now().Format(TIME_FORMAT)

	res, err := QueryTable(sess, *deviceID, from, to)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetDB Result %+v\n", res)
}

func QueryTable(sess *session.Session, deviceID, start, end string) ([]map[string]*dynamodb.AttributeValue, error) {
	var db = dynamodb.New(sess)
	keyCond := expression.Key("deviceID").Equal(expression.Value(deviceID)).
		And(expression.Key("timestamp").Between(
			expression.Value(start),
			expression.Value(end)))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		panic(err)
	}

	result, err := db.QueryWithContext(context.Background(), &dynamodb.QueryInput{
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String("tracks"),
	})
	if err != nil {
		panic(err)
	}
	return result.Items, nil
}
