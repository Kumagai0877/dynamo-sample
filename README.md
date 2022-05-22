# GolangでDynamoDB操作する

## 概要

ランダムの数値を保存し、保存したデータを取得して出力します。

## ローカル環境構築

Docker環境起動
```
make up
```

ローカルのDynamoDBの管理コンソールを表示します
```
npm install -g
dynamodb-admin dynamodb-admin
```

下記の設定になるようにテーブルを作成します（コマンドで作れる様にする予定）
```
"AttributeDefinitions": [
    {
      "AttributeName": "deviceID",
      "AttributeType": "S"
    },
    {
      "AttributeName": "timestamp",
      "AttributeType": "N"
    }
  ],
  "TableName": "Tracks",
  "KeySchema": [
    {
      "AttributeName": "deviceID",
      "KeyType": "HASH"
    },
    {
      "AttributeName": "timestamp",
      "KeyType": "RANGE"
    }
  ],
```