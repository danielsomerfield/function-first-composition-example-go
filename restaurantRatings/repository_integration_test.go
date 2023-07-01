package restaurantRatings

import (
	"context"
	ratings2 "function-first-composition-example-go/review-server/restaurantRatings/ratings"
	"function-first-composition-example-go/review-server/testutil"
	"github.com/magiconair/properties/assert"
	"sort"
	"testing"
)

func TestFindRatingsByRestaurant(t *testing.T) {

	ctx := context.Background()

	db := testutil.StartDB(ctx, "../db/init.sql")
	t.Cleanup(func() {
		if err := db.Stop(ctx); err != nil {
			t.Fatalf("failed to terminate db: %s", err.Error())
		}
	})

	restaurant1 := testutil.Restaurant{Id: "restaurant1", Name: "Restaurant 1"}
	restaurant2 := testutil.Restaurant{Id: "restaurant2", Name: "Restaurant 2"}
	user1 := testutil.User{
		Id:      "user1",
		Name:    "User 1",
		Trusted: false,
	}

	user2 := testutil.User{
		Id:      "user2",
		Name:    "User 2",
		Trusted: false,
	}

	rating1 := testutil.RatingForRestaurant{
		Id:         "rating1",
		User:       user1,
		Restaurant: restaurant1,
		Rating:     "Average",
	}
	rating2 := testutil.RatingForRestaurant{
		Id:         "rating2",
		User:       user2,
		Restaurant: restaurant1,
		Rating:     "Excellent",
	}
	rating3 := testutil.RatingForRestaurant{
		Id:         "rating3",
		User:       user1,
		Restaurant: restaurant2,
		Rating:     "Excellent",
	}

	testutil.CreateRestaurant(db, restaurant1)
	testutil.CreateRestaurant(db, restaurant2)

	testutil.CreateUser(db, user1)
	testutil.CreateUser(db, user2)

	testutil.CreateRatingByUserForRestaurant(db, rating1)
	testutil.CreateRatingByUserForRestaurant(db, rating2)
	testutil.CreateRatingByUserForRestaurant(db, rating3)

	findRatingsByRestaurantId := ratings2.CreateFindRatingsByRestaurant(db.(testutil.DBImplementation).GetDB())
	ratings, err := findRatingsByRestaurantId("vancouverbc")
	if err != nil {
		t.Fatalf("Received error %v", err)
	}
	sort.Slice(ratings, func(i, j int) bool {
		return ratings[i].RestaurantId < ratings[j].RestaurantId
	})

	assert.Equal(t, len(ratings), 2, "restaurantRatings by restaurant")
	assert.Equal(t, ratings[0].RestaurantId, restaurant1.Id)
	assert.Equal(t, ratings[1].RestaurantId, restaurant2.Id)
	assert.Equal(t, len(ratings[0].Ratings), 2)
	assert.Equal(t, len(ratings[1].Ratings), 1)
}
