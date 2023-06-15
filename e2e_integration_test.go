package main

import (
	"context"
	"encoding/json"
	"function-first-composition-example-go/review-server/server"
	"function-first-composition-example-go/review-server/testutil"
	"github.com/magiconair/properties/assert"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
)

func TestRestaurantEndPointRanksRestaurant(t *testing.T) {
	t.Skipf("NYI")
	ctx := context.Background()

	db := testutil.StartDB(ctx)

	users := []testutil.User{
		{"user1", "User 1", true},
		{"user2", "User 2", false},
		{"user3", "User 3", false},
	}

	restaurants := []testutil.Restaurant{
		{"cafegloucesterid", "Cafe Gloucester"},
		{"burgerkingid", "Burger King"},
	}

	ratings := []testutil.RatingForRestaurant{
		{"rating1", users[0], restaurants[0], "EXCELLENT"},
		{"rating2", users[1], restaurants[0], "TERRIBLE"},
		{"rating3", users[2], restaurants[0], "AVERAGE"},
		{"rating4", users[2], restaurants[1], "ABOVE_AVERAGE"},
	}

	for _, user := range users {
		testutil.CreateUser(db, user)
	}

	for _, r := range restaurants {
		testutil.CreateRestaurant(db, r)
	}

	for _, r := range ratings {
		testutil.CreateRatingByUserForRestaurant(db, r)
	}

	t.Cleanup(func() {
		if err := db.Stop(ctx); err != nil {
			t.Fatalf("failed to terminate db: %s", err.Error())
		}
	})

	randomPort := rand.Intn(65535-1024) + 1024
	reviewServer := server.NewServer("127.0.0.1", randomPort)
	if err := reviewServer.Start(); err != nil {
		t.Fatalf("failed to start server: %s", err.Error())
	}

	t.Cleanup(func() {
		if err := reviewServer.Stop(); err != nil {
			t.Fatalf("failed to stop server: %s", err.Error())
		}
	})

	response, err := http.Get("http://localhost:" + strconv.Itoa(randomPort) + "/vancouverbc/restaurants/recommended")
	if err != nil {
		t.Fatalf("Failed to make request: %s", err.Error())
	}

	assert.Equal(t, response.StatusCode, 200)

	var body Response
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Failed to read body: %v", err)
	}

	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		t.Fatalf("failed to unmarshall request body: %s", err.Error())
	}

	_ = json.Unmarshal(bodyBytes, &body)
	assert.Equal(t, len(body.Restaurants), 2)
	assert.Equal(t, body.Restaurants[0].Id, "cafegloucesterid")
	assert.Equal(t, body.Restaurants[0].Id, "burgerkingid")
}

type Restaurant struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Restaurants []Restaurant `json:"restaurants"`
}
