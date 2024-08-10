# 概要
goで簡易なAPIサーバーを立てる


# コマンド
APIサーバーの起動

```bash
GO_ENV=dev go run main.go
```

postgresに入る
```bash
psql -h localhost -U postgres
>> SELECT * FROM users;
```

ログイン
```bash
curl -c cookie.txt -X POST http://localhost:8080/login  -H "Content-Type: application/json" -d '{"email": "user2@example.com", "password": "dummy"}' 
```

タスクを追加
```bash
curl -b cookie.txt -X POST http://localhost:8080/tasks  -H "Content-Type: application/json" -d '{"title": "hugahuga"}' 
```

タスクを確認
```bash
curl -b cookie.txt http://localhost:8080/tasks | jq
```