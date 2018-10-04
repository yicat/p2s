package main

import (
	"flag"
	"fmt"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
	"github.com/yicat/p2s/config"
	"log"
	"os"
	"time"
)

const usageText = `
This program runs command on the db. Support commands are:
	- up [target] - runs all available migrations by default or up to target one if argument is provided.		
	- down - reverts last migration.
	- reset - revers all migrations.
	- version - prints current db version.
	- set_version [version] - sets db version without running migrations

Usage:
	go run *.go <command> [args]
`

func main() {
	flag.Usage = usage
	flag.Parse()

	db := pg.Connect(&pg.Options{
		User:     config.DB_USER,
		Password: config.DB_PASS,
		Database: config.DB_NAME,
	})
	// view log
	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}
		log.Printf("%s %s", time.Since(event.StartTime), query)
	})

	// 整个 migrate 过程是在一个事务中
	// 参见 https://github.com/go-pg/migrations
	var oldVersion, newVersion int64
	err := db.RunInTransaction(func(tx *pg.Tx) (err error) {
		oldVersion, newVersion, err = migrations.Run(db, flag.Args()...)
		return
	})
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func exitf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
	os.Exit(1)
}
