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

func getOrDefault(name string, defaultValue string) string {
	value, found := os.LookupEnv(name)
	if !found {
		value = defaultValue
	}
	return value
}

func FromEnv() *Configuration {
	dbPortString := getOrDefault("REVIEW_DATABASE_PORT", "5432")
	dbPort, err := strconv.Atoi(dbPortString)
	if err != nil {
		log.Fatalf("Invalid database port value %v", dbPortString)
	}

	serverPortString := getOrDefault("SERVER_PORT", "8080")
	serverPort, err := strconv.Atoi(serverPortString)
	if err != nil {
		log.Fatalf("Invalid server port value %v", dbPortString)
	}

	return &Configuration{
		DataSource: db.DataSource{
			Host:     getRequired("REVIEW_DATABASE_HOST"),
			Port:     dbPort,
			User:     getRequired("REVIEW_DATABASE_USER"),
			Password: getRequired("REVIEW_DATABASE_PASSWORD"),
			DbName:   getRequired("REVIEW_DATABASE_DATABASE"),
		},
		ServerPort: serverPort,
	}
}
