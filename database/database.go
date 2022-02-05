package database

import (
	"database/sql"
	"golang-restful/helper"
	"time"
)

func NewDB() *sql.DB {
	conn := "user=myuser password=user dbname=golang_restful"
	db, err := sql.Open("postgres", conn)
	helper.PanicHandler(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
