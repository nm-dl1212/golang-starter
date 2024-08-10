package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	/* DBを開く関数
	 */

	// 環境変数を.envファイルから読み込む処理
	// dev環境のみ実行する。prod環境では，環境ごとに設定したenvを使用するため実行不要。
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}
	// 環境変数の情報からpostgresのurlを生成
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s", // 埋め込み先の文字列
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	// GO言語のORM(Object-relational mapping)であるgormを用いて，postgresを開く。
	// 第2引数でDBを開く際のConfig情報を設定することができるが，ここでは空でOK。
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// 接続成功すると，db(中身はgorm.DBオブジェクトのアドレス)を返す。
	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	/* DBを閉じる関数
	 */

	// gorm.DB(ORMでラップしたDB)からsql.DB(実体としてのDB)を受け取る。
	sqlDB, _ := db.DB()

	// dbを閉じる。
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
