package restaurantRatings

import (
	"function-first-composition-example-go/review-server/restaurantRatings/ratings"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCalculateHandlesEmptyRatings(t *testing.T) {
	overallRating := Calculate(&ratings.RatingsByRestaurant{RestaurantId: "restaurant1", Ratings: []ratings.RestaurantRating{}})
	assert.Equal(t, overallRating, 0)
}

func TestCalculatePassesThroughRatingOfUntrusted(t *testing.T) {
	overallRating := Calculate(&ratings.RatingsByRestaurant{RestaurantId: "restaurant1", Ratings: []ratings.RestaurantRating{
		{
			User: ratings.User{
				Id:        "user1",
				IsTrusted: false,
			},
			Rating: ratings.Excellent,
		},
	}})

	assert.Equal(t, overallRating, ratings.Excellent.Value())
}

func TestCalculateWeightsTrusted(t *testing.T) {
	overallRating := Calculate(&ratings.RatingsByRestaurant{RestaurantId: "restaurant1", Ratings: []ratings.RestaurantRating{
		{
			User: ratings.User{
				Id:        "user1",
				IsTrusted: true,
			},
			Rating: ratings.Excellent,
		},
		{
			User: ratings.User{
				Id:        "user1",
				IsTrusted: false,
			},
			Rating: ratings.Excellent,
		},
	}})

	assert.Equal(t, overallRating, ratings.Excellent.Value()*5)
}
