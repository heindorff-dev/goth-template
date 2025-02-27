package repository

import "database/sql"

type UserRepository struct {
	databaseConnection *sql.DB
}

func CreateUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{databaseConnection: db}
}
