package ratings

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestCalculateHandlesEmptyRatings(t *testing.T) {
	overallRating := Calculate(&RatingsByRestaurant{RestaurantId: "restaurant1", Ratings: []RestaurantRating{}})
	assert.Equal(t, overallRating, 0)
}

func TestCalculatePassesThroughRatingOfUntrusted(t *testing.T) {
	overallRating := Calculate(&RatingsByRestaurant{RestaurantId: "restaurant1", Ratings: []RestaurantRating{
		{
			User: User{
				Id:        "user1",
				IsTrusted: false,
			},
			Rating: Excellent,
		},
	}})

	assert.Equal(t, overallRating, Excellent.value())
}

func TestCalculateWeightsTrusted(t *testing.T) {
	overallRating := Calculate(&RatingsByRestaurant{RestaurantId: "restaurant1", Ratings: []RestaurantRating{
		{
			User: User{
				Id:        "user1",
				IsTrusted: true,
			},
			Rating: Excellent,
		},
		{
			User: User{
				Id:        "user1",
				IsTrusted: false,
			},
			Rating: Excellent,
		},
	}})

	assert.Equal(t, overallRating, Excellent.value()*5)
}
