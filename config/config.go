package config

import (
	"database/sql"
	"fmt"
	"go_inven_ctrl/utils"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "1w551474"
// 	dbname   = "testing"
// )

// config db
var (
	dbHost = utils.DotEnv("DB_HOST")
	dbPort = utils.DotEnv("DB_PORT")
	dbUser = utils.DotEnv("DB_USER")
	dbPassword = utils.DotEnv("DB_PASSWORD")
	dbName = utils.DotEnv("DB_NAME")
	sslMode = utils.DotEnv("SSL_MODE")
)

// db connection
var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)
// db, err := sql.Open("postgres", dataSourceName)

// koneksi database ================================================================================

// var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 	host, port, user, password, dbname)

func ConnectDB() *sql.DB {
	// db, err := sql.Open("postgres", psqlInfo)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully connected")
	}
	return db
}
