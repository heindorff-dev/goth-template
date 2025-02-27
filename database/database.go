package database

import (
	"database/sql"
	"fmt"
	"goth-template/server/helper"

	"github.com/joho/godotenv"
)

func ConnectDatabase() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

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

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to db:", dbName)
	return db, nil
}
