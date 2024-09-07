package main

import (
	"log"

	"github.com/temur-shamshidinov/task_app/api"
	"github.com/temur-shamshidinov/task_app/config"
	"github.com/temur-shamshidinov/task_app/pkg/db"
	"github.com/temur-shamshidinov/task_app/storage"
)

func main() {
	cfg := config.Load()

	db, err := db.ConnectToDb(cfg.PgConfig)
	if err != nil {
		log.Println("error on connect to ConToDb:", err)
		return
	}
	storage := storage.NewStorage(db)

	api.Api(storage)

}
