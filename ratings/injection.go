package ratings

import (
	"function-first-composition-example-go/review-server/configuration"
	"function-first-composition-example-go/review-server/db"
	"function-first-composition-example-go/review-server/restaurants"
	"github.com/gin-gonic/gin"
)

func Initialize(engine *gin.Engine) {

	config := configuration.FromEnv()

	database := db.Connect(config.DataSource)

	getRestaurantById := restaurants.CreateGetRestaurantById(database)

	topRatedDependencies := &TopRatedDependencies{
		getRestaurantById: getRestaurantById,
		findRatingsByRestaurant: func(city string) ([]RatingsByRestaurant, error) {
			return nil, nil
		},
		calculateRatingForRestaurant: func(ratings *RatingsByRestaurant) (int, error) {
			return 0, nil
		},
	}
	dependencies := ControllerDependencies{
		GetTopRestaurants: createTopRated(topRatedDependencies),
	}
	engine.GET("/:city/restaurants/recommended", createController(&dependencies))
}
