package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type dbConnectionConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func ConnectDatabase(configFilename string) (*sql.DB, error) {
	dbConnConfig, err := loadDatabaseConfig(configFilename)
	if err != nil {
		return nil, err
	}

	connectionString := createConnectionString(dbConnConfig)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func loadDatabaseConfig(filename string) (dbConnectionConfig, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return dbConnectionConfig{}, err
	}

	var dbConnConfig dbConnectionConfig
	err = json.Unmarshal(fileContent, &dbConnConfig)
	if err != nil {
		return dbConnectionConfig{}, err
	}

	return dbConnConfig, nil
}

func createConnectionString(dbConnConfig dbConnectionConfig) string {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbConnConfig.Username,
		dbConnConfig.Password,
		dbConnConfig.Host,
		dbConnConfig.Port,
		dbConnConfig.DBName,
	)

	return connectionString
}
