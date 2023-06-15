package ratings

import (
	"errors"
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_topRatedFromProprietaryAlgorithm(t *testing.T) {
	restaurantsById := map[string]Restaurant{
		"restaurant1": {
			Id:   "restaurant1",
			Name: "Restaurant 1",
		},
		"restaurant2": {
			Id:   "restaurant2",
			Name: "Restaurant 2",
		},
	}

	getRestaurantByIdStub := func(id string) Restaurant {
		return restaurantsById[id]
	}

	ratings := []RatingsByRestaurant{
		{
			RestaurantId: "restaurant1",
			Ratings: []RestaurantRating{
				{
					User: User{
						Id:        "user1",
						IsTrusted: true,
					},
					Rating: Excellent,
				},
			},
		},
		{
			RestaurantId: "restaurant2",
			Ratings: []RestaurantRating{
				{
					User: User{
						Id:        "user2",
						IsTrusted: false,
					},
					Rating: Excellent,
				},
			},
		},
	}

	ratingsByCity := []struct {
		City    string
		Ratings []RatingsByRestaurant
	}{
		{
			City:    "vancouverbc",
			Ratings: ratings,
		},
	}

	findRatingsByRestaurantStub := func(city string) (returnedRatings []RatingsByRestaurant) {

		for _, s := range ratingsByCity {
			if s.City == city {
				returnedRatings = append(returnedRatings, s.Ratings...)
			}
		}
		return returnedRatings
	}

	calculateRatingForRestaurantStub := func(ratings RatingsByRestaurant) (int, error) {
		if ratings.RestaurantId == "restaurant1" {
			return 10, nil
		} else if ratings.RestaurantId == "restaurant2" {
			return 5, nil
		} else {
			return 0, errors.New("unknown restaurant")
		}
	}

	dependencies := TopRatedDependencies{
		getRestaurantById:            getRestaurantByIdStub,
		findRatingsByRestaurant:      findRatingsByRestaurantStub,
		calculateRatingForRestaurant: calculateRatingForRestaurantStub,
	}

	topRated := createTopRated(&dependencies)
	topRestaurants := topRated("vancouverbc")
	assert.Equal(t, len(topRestaurants), 2)
	assert.Equal(t, topRestaurants[0].Id, "restaurant1")
	assert.Equal(t, topRestaurants[0].Name, "Restaurant 1")
	assert.Equal(t, topRestaurants[1].Id, "restaurant2")
	assert.Equal(t, topRestaurants[1].Name, "Restaurant 2")
}
