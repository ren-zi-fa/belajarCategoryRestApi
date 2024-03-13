package app

import (
	"database/sql"
	"ren-zi-fa/BelajarRestfullApi/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "renzi:renzi@tcp(localhost:3306)/golangapi")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}