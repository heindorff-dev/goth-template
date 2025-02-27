package database

import (
	"database/sql"
	"errors"
	"fmt"
	"goth-template/server/helper"
)

type DatabaseConnection struct {
	databaseQueries  *Queries
	connectionString string
	DatabaseName     string
}

func NewDatabaseConnection() (*DatabaseConnection, error) {
	username, err := helper.GetEnv("DB_USERNAME")
	if err != nil {
		return nil, err
	}

	password, err := helper.GetEnv("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	dbName, err := helper.GetEnv("DB_DATABASE")
	if err != nil {
		return nil, err
	}

	connStr := username + ":" + password + "@/" + dbName + "parseTime=true"

	return &DatabaseConnection{connectionString: connStr, DatabaseName: dbName}, nil
}

func (d *DatabaseConnection) Init() error {
	db, err := sql.Open("mysql", d.connectionString)
	if err != nil {
		return err
	}
	d.databaseQueries = New(db)
	fmt.Println("Successfully connected to db: ", d.DatabaseName)
	return nil
}

func (d *DatabaseConnection) MustGetQuery() (*Queries, error) {
	if d.databaseQueries == nil {
		return nil, errors.New("Queries do not exist for database connection")
	}
	return d.databaseQueries, nil
}
