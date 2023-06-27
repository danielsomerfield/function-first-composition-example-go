package ratings

func Calculate(r *RatingsByRestaurant) (overallRating int) {
	for _, r := range r.Ratings {
		if r.User.IsTrusted {
			overallRating += r.Rating.value() * 4
		} else {
			overallRating += r.Rating.value()
		}
	}
	return overallRating
}
