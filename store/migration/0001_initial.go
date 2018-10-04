package main

import (
	"fmt"
	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table app")
		_, err := db.Exec(CreateAppSchema)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table app")
		_, err := db.Exec(DeleteAppSchema)
		return err
	})
}

const CreateAppSchema = `
	CREATE TABLE app (
		id serial PRIMARY KEY,
		name varchar(200) UNIQUE NOT NULL,
		description varchar(200),
		created_at timestamp NOT NULL,
		updated_at timestamp NOT NULL,
		deleted_at timestamp
	)
`

const DeleteAppSchema = `
	DROP TABLE app
`
