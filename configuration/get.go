package configuration

import (
	"function-first-composition-example-go/review-server/db"
	"log"
	"os"
	"strconv"
)

type Configuration struct {
	DataSource db.DataSource
	ServerPort int
}

func getRequired(name string) string {
	value, found := os.LookupEnv(name)
	if !found {
		log.Fatalf("Missing required value %v", name)
	}
	return value
}

func FromEnv() *Configuration {
	portString := getRequired("PORT")
	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatalf("Invalid port value %v", portString)
	}
	return &Configuration{DataSource: db.DataSource{
		Host:     getRequired("DB_HOST"),
		Port:     port,
		User:     getRequired("DB_USER"),
		Password: getRequired("DB_PASSWORD"),
		DbName:   getRequired("DB_NAME"),
	}}
}
