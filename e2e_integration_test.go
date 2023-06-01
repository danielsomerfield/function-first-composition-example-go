package main

import (
	"context"
	server2 "function-first-composition-example-go/review-server/server"
	"github.com/magiconair/properties/assert"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
)

func TestRestaurantEndPointRanksRestaurant(t *testing.T) {
	ctx := context.Background()

	dbContainer, err := postgres.RunContainer(
		ctx,
		postgres.WithDatabase("postgres"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		postgres.WithInitScripts("db/init.sql"),
	)

	if err != nil {
		panic(err)
	}

	err = dbContainer.Start(ctx)

	if err != nil {
		panic(err)
	}

	t.Cleanup(func() {
		if err := dbContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate db: %s", err.Error())
		}
	})

	randomPort := rand.Intn(65535-1024) + 1024
	server := server2.NewServer("127.0.0.1", randomPort)
	if err := server.Start(); err != nil {
		t.Fatalf("failed to start server: %s", err.Error())
	}

	t.Cleanup(func() {
		if err := server.Stop(); err != nil {
			t.Fatalf("failed to stop server: %s", err.Error())
		}
	})

	response, err := http.Get("http://localhost:" + strconv.Itoa(randomPort) + "/vancouverbc/restaurants/recommended")
	if err != nil {
		t.Fatalf("Failed to make request: %s", err.Error())
	}

	assert.Equal(t, response.StatusCode, 200)

	//var body Response
	//bodyBytes, err := io.ReadAll(response.Body)
	//if err != nil {
	//	t.Fatalf("Failed to read body: %v", err)
	//}
	//
	//if err := json.Unmarshal(bodyBytes, &body); err != nil {
	//	t.Fatalf("failed to unmarshall request body: %s", err.Error())
	//}
}

type Restaurant struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Restaurants []Restaurant `json:"restaurants"`
}
