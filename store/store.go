package store

import (
	"github.com/go-pg/pg"
	"github.com/yicat/p2s/config"
	"log"
	"time"
)

type Store struct {
	DB *pg.DB
}

func New(debug bool) *Store {
	db := pg.Connect(&pg.Options{
		User:     config.DB_USER,
		Database: config.DB_NAME,
		Password: config.DB_PASS,
	})

	if _, err := db.Exec("SELECT 1"); err != nil {
		panic(err)
	}

	if debug {
		db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			query, err := event.FormattedQuery()
			if err != nil {
				panic(err)
			}
			log.Printf("%s %s", time.Since(event.StartTime), query)
		})
	}

	return &Store{DB: db}
}
