package main

import "github.com/go-pg/pg/v10"

func PostgresConnect() *pg.DB {
	options := &pg.Options{
		User:     "postgre",
		Password: "ptpit",
		Addr:     "localhost:5432",
		Database: "cats",
	}
	return pg.Connect(options)

}
