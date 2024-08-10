package main

/*
マイグレーション処理をプログラムのエントリーポイントで行いたい
→ mainパッケージに所属させる
*/

import (
	"fmt"
	"rest-api/db"
	"rest-api/model"
)

func main() {
	// DBを開く
	dbConn := db.NewDB()

	// 遅延実行(defer)を用いて，DBへの接続が確立できたら以下を実行する
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)

	// マイグレーション
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
