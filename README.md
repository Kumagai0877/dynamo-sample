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
dynamodb-admin
```

テーブル作成
```
make create-table
```