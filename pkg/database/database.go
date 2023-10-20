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

	if err := InitTables(db); err != nil {
		return &res, err
	}

	return &Database{
		db: db,
	}, nil
}

func InitTables(db *sqlx.DB) error {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users(
    	id SERIAL PRIMARY KEY,
    	first_name VARCHAR(200),
    	last_name VARCHAR(200),
    	patronymic VARCHAR(200),
    	age INT,
    	gender VARCHAR(10),
    	nationality VARCHAR(10),
    	nationality_probability DOUBLE PRECISION
	)`
	_, err := db.Exec(createUsersTable)
	if err != nil {
		return err
	}
	return nil
}
