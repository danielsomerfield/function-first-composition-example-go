package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DataSource struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func Connect(datasource DataSource) *sql.DB {

	connectInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		datasource.Host, datasource.Port, datasource.User, datasource.Password, datasource.DbName)

	database, err := sql.Open("postgres", connectInfo)

	if err != nil {
		panic(err)
	}

	return database
}
