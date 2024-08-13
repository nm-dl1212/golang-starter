# 概要
goで簡易なAPIサーバーを立てる


# コマンド
## マイグレーション
```bash
GO_ENV=dev go run migrate/migrate.go
```
postgresに入って，テーブルを確認
```bash
psql -h localhost -U postgres
>> SELECT * FROM users;
```


## APIサーバーの起動
```bash
GO_ENV=dev go run main.go
```

## 操作
ユーザーを作成する
```bash
curl -X POST http://localhost:8080/signup  -H "Content-Type: application/json" -d '{"email": "test-user@example.com", "password": "dummy"}' 
```

作成したユーザーでログインし，クッキーをテキストファイルに書き込む
```bash
curl -c cookie.txt -X POST http://localhost:8080/login  -H "Content-Type: application/json" -d '{"email": "test-user@example.com", "password": "dummy"}' 
```

タスクを追加
```bash
curl -b cookie.txt -X POST http://localhost:8080/tasks  -H "Content-Type: application/json" -d '{"title": "hugahuga"}' 
```

タスクを確認
```bash
curl -b cookie.txt http://localhost:8080/tasks | jq
```

ログアウト
```bash
curl -c cookie.txt -X POST http://localhost:8080/logout -H "Content-Type: application/json" -d '{"email": "test-user@example.com", "password": "dummy"}' 
```