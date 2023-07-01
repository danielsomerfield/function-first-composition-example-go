package restaurantRatings

import "function-first-composition-example-go/review-server/restaurantRatings/ratings"

func Calculate(r *ratings.RatingsByRestaurant) (overallRating int) {
	for _, r := range r.Ratings {
		if r.User.IsTrusted {
			overallRating += r.Rating.Value() * 4
		} else {
			overallRating += r.Rating.Value()
		}
	}
	return overallRating
}
