package ratings

type TopRatedDependencies struct {
	getRestaurantById            func(id string) Restaurant
	findRatingsByRestaurant      func(city string) (returnedRatings []RatingsByRestaurant)
	calculateRatingForRestaurant func(ratings RatingsByRestaurant) (int, error)
}

func createTopRated(dependencies *TopRatedDependencies) func(string) []Restaurant {
	return func(city string) []Restaurant {
		return []Restaurant{}
	}
}
