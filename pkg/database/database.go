package database

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	Conn *sql.DB
}

var counts int64

func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func ConnectSQL(dsn string) (*DB, error) {
	var db = new(sql.DB)

	for {
		conn, err := NewDatabase(dsn)
		if err != nil {
			log.Println("Postgres not ready...")
			counts++
		} else {
			log.Println("Connected to database!")
			db = conn
			break
		}

		if counts > 10 {
			log.Println(err)
			return nil, errors.New("unable to connect to db")
		}

		log.Println("Backing off for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}

	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	return &DB{Conn: db}, nil

}
