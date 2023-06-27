package ratings

import (
	. "function-first-composition-example-go/review-server/domain"
	"log"
	"sort"
)

type TopRatedDependencies struct {
	getRestaurantById            func(id string) (*Restaurant, error)
	findRatingsByRestaurant      func(city string) ([]RatingsByRestaurant, error)
	calculateRatingForRestaurant func(ratings *RatingsByRestaurant) int
}

type OverallRating struct {
	restaurantId string
	rating       int32
}

func createTopRated(dependencies *TopRatedDependencies) func(string) ([]Restaurant, error) {

	calculateRatingForRestaurant := func(ratings []RestaurantRating) (total int32) {
		for _, r := range ratings {
			total = total + int32(r.Rating)
		}
		return total
	}

	calculateRatings := func(ratings *[]RatingsByRestaurant) (overall OverallRatings) {
		for _, r := range *ratings {
			overall = append(overall, OverallRating{
				restaurantId: r.RestaurantId,
				rating:       calculateRatingForRestaurant(r.Ratings),
			})
		}
		return overall
	}

	toRestaurants := func(ratings OverallRatings) (restaurants []Restaurant) {
		for _, rating := range ratings {
			restaurant, err := dependencies.getRestaurantById(rating.restaurantId)
			if err != nil {
				log.Printf("Could not load restaurant with from %v. Err: %v", rating, err)
			} else {
				restaurants = append(restaurants, *restaurant)
			}
		}

		return restaurants
	}

	return func(city string) ([]Restaurant, error) {
		ratings, err := dependencies.findRatingsByRestaurant(city)

		if err != nil {
			return nil, err
		}

		overallRatings := calculateRatings(&ratings)

		sort.Sort(overallRatings)

		restaurants := toRestaurants(overallRatings)

		return restaurants, nil
	}
}

type OverallRatings []OverallRating

func (ratings OverallRatings) Len() int {
	return len(ratings)
}
func (ratings OverallRatings) Less(i int, j int) bool {
	return ratings[i].rating < ratings[j].rating
}

func (ratings OverallRatings) Swap(i int, j int) {
	first := ratings[i]
	ratings[i] = ratings[j]
	ratings[j] = first
}
