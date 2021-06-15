package infra

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

var DB *sql.DB

const (
	MaxConns    = 25
	MaxLifeTime = 5 * time.Minute
)

func InitMysql(dbTCPHost, port, user, password, database, instanceConnectionName string) {
	if dbTCPHost != "" {
		db, err := initTCPConnectionPool(user, password, dbTCPHost, port, database)
		if err != nil {
			log.Fatalf("initTCPConnectionPool: unable to connect: %v", err)
		}
		DB = db
	} else {
		db, err := initSocketConnectionPool(user, password, instanceConnectionName, database)
		if err != nil {
			log.Fatalf("initSocketConnectionPool: unable to connect: %v", err)
		}
		DB = db
	}
}

func initSocketConnectionPool(dbUser, dbPwd, instanceConnectionName, dbName string) (*sql.DB, error) {
	socketDir := "/cloudsql"
	dbURI := fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)

	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}
	configureConnectionPool(dbPool)
	return dbPool, nil
}

func initTCPConnectionPool(dbUser, dbPwd, dbTCPHost, dbPort, dbName string) (*sql.DB, error) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPwd, dbTCPHost, dbPort, dbName)

	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}
	configureConnectionPool(dbPool)
	return dbPool, nil
}

func configureConnectionPool(dbPool *sql.DB) {
	dbPool.SetMaxIdleConns(MaxConns)
	dbPool.SetMaxOpenConns(MaxConns)
	dbPool.SetConnMaxLifetime(MaxLifeTime)
}
