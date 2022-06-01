package main

import (
	"database/sql"
	"log"

	"github.com/NaokiYazawa/simplebank/api"
	db "github.com/NaokiYazawa/simplebank/db/sqlc"
	"github.com/NaokiYazawa/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn) // 新しい store を作成できる
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
}
