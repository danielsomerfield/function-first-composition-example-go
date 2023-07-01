package restaurantRatings

import (
	ratings2 "function-first-composition-example-go/review-server/restaurantRatings/ratings"
	"function-first-composition-example-go/review-server/restaurantRatings/restaurants"
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_topRatedFromProprietaryAlgorithm(t *testing.T) {
	restaurantsById := map[string]restaurants.Restaurant{
		"restaurant1": {
			Id:   "restaurant1",
			Name: "Restaurant 1",
		},
		"restaurant2": {
			Id:   "restaurant2",
			Name: "Restaurant 2",
		},
	}

	getRestaurantByIdStub := func(id string) (*restaurants.Restaurant, error) {
		restaurant := restaurantsById[id]
		return &restaurant, nil
	}

	ratings := []ratings2.RatingsByRestaurant{
		{
			RestaurantId: "restaurant1",
			Ratings: []ratings2.RestaurantRating{
				{
					User: ratings2.User{
						Id:        "user1",
						IsTrusted: true,
					},
					Rating: ratings2.Terrible,
				},
			},
		},
		{
			RestaurantId: "restaurant2",
			Ratings: []ratings2.RestaurantRating{
				{
					User: ratings2.User{
						Id:        "user2",
						IsTrusted: false,
					},
					Rating: ratings2.Excellent,
				},
			},
		},
	}

	ratingsByCity := []struct {
		City    string
		Ratings []ratings2.RatingsByRestaurant
	}{
		{
			City:    "vancouverbc",
			Ratings: ratings,
		},
	}

	findRatingsByRestaurantStub := func(city string) (returnedRatings []ratings2.RatingsByRestaurant, err error) {
		for _, s := range ratingsByCity {
			if s.City == city {
				returnedRatings = append(returnedRatings, s.Ratings...)
			}
		}
		return returnedRatings, nil
	}

	calculateRatingForRestaurantStub := func(ratings *ratings2.RatingsByRestaurant) int {
		if ratings.RestaurantId == "restaurant1" {
			return 10
		} else if ratings.RestaurantId == "restaurant2" {
			return 5
		} else {
			t.Fatalf("Unknown restaurant %v\n", ratings.RestaurantId)
			return 0
		}
	}

	dependencies := TopRatedDependencies{
		getRestaurantById:            getRestaurantByIdStub,
		findRatingsByRestaurant:      findRatingsByRestaurantStub,
		calculateRatingForRestaurant: calculateRatingForRestaurantStub,
	}

	topRated := createTopRated(&dependencies)
	topRestaurants, err := topRated("vancouverbc")
	assert.Equal(t, err == nil, true)
	assert.Equal(t, len(topRestaurants), 2)
	assert.Equal(t, topRestaurants[0].Id, "restaurant1")
	assert.Equal(t, topRestaurants[0].Name, "Restaurant 1")
	assert.Equal(t, topRestaurants[1].Id, "restaurant2")
	assert.Equal(t, topRestaurants[1].Name, "Restaurant 2")
}
