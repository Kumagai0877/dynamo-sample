package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"math/rand"
	"time"
)

type Track struct {
	DeviceID  string `dynamo:"deviceID,hash"`
	Timestamp int64  `dynamo:"timestamp, range"`
	Value     int    `dynamo:"value"`
}

const AWS_REGION = "ap-northeast-1"
const DYNAMO_ENDPOINT = "http://localhost:8000"

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

	// ランダムで100以下の整数の生成
	timestamp := time.Now().UnixNano()
	rand.Seed(timestamp)
	value := rand.Intn(100)

	err = table.Put(&Track{DeviceID: *deviceID, Timestamp: timestamp, Value: value}).Run()
	if err != nil {
		panic(err)
	}

	var track Track
	// InsertしたデータをDBからGet
	err = table.Get("deviceID", deviceID).Range("timestamp", dynamo.Equal, timestamp).One(&track)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetDB Result %+v\n", track)
}
