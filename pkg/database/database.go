package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sqlx.DB
}

type DatabaseLoad struct {
	Host, Port, DBName, User, Password string
}

func NewDatabase(body DatabaseLoad) (*Database, error) {
	var res Database
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		body.User, body.Password, body.Host, body.Port, body.DBName,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return &res, err
	}
	if err := db.Ping(); err != nil {
		return &res, err
	}

	//if err := InitTables(db); err != nil {
	//	return &res, err
	//}

	return &Database{
		db: db,
	}, nil
}
