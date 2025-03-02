package database

import (
	"context"
	"database/sql"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
)

type ConnectionManager struct {
	configuration    ConnectionConfiguration
	credentials      ConnectionCredential
	connectionString string
	database         *sql.DB
	logger           *slog.Logger
}

type ConnectionCredential struct {
	username string
	password string
}

type ConnectionConfiguration struct {
	databaseName string
	port         string
	host         string
	protocol     string
}

type DatabaseClient struct {
	connectionManager *ConnectionManager
}

func NewConnectionManager(config ConnectionConfiguration, cred ConnectionCredential, logger *slog.Logger) (*ConnectionManager, error) {
	connectionManager := &ConnectionManager{
		configuration: config,
		credentials:   cred,
		logger:        logger,
	}

	// Create connectionstring. Better way??
	connectionManager.connectionString = connectionManager.credentials.username + ":" + connectionManager.credentials.password + "@" + connectionManager.configuration.protocol + "(" +
		connectionManager.configuration.host + ":" + connectionManager.configuration.port + ")/" + connectionManager.configuration.databaseName

	db, err := sql.Open("mysql", connectionManager.connectionString)
	if err != nil {
		return nil, err
	}

	connectionManager.database = db
	return connectionManager, nil
}

func NewConnectionCredentials(username string, password string) *ConnectionCredential {
	creds := &ConnectionCredential{
		username: username,
		password: password,
	}
	return creds
}

func NewConnectionConfiguration(databaseName string, host string, port string, protocol string) *ConnectionConfiguration {
	config := &ConnectionConfiguration{
		databaseName: databaseName,
		host:         host,
		port:         port,
		protocol:     protocol,
	}
	return config
}

func (c *ConnectionManager) NewClient() *DatabaseClient {
	client := &DatabaseClient{
		connectionManager: c,
	}
	return client
}

func (c *DatabaseClient) Query(ctx context.Context) *Queries {
	conn, err := c.connectionManager.database.Conn(ctx)

	if err != nil {
		return nil
	}
	defer conn.Close()
	query := New(c.connectionManager.database)
	return query
}

func (c *ConnectionManager) Ping() error {
	err := c.database.Ping()
	if err != nil {
		return err
	}
	return nil
}
