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