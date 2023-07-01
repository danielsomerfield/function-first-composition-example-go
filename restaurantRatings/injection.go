package restaurantRatings

import (
	"function-first-composition-example-go/review-server/configuration"
	"function-first-composition-example-go/review-server/db"
	ratings2 "function-first-composition-example-go/review-server/restaurantRatings/ratings"
	"function-first-composition-example-go/review-server/restaurantRatings/restaurants"
	"github.com/gin-gonic/gin"
)

func Initialize(engine *gin.Engine, config *configuration.Configuration) {

	database := db.Connect(config.DataSource)

	getRestaurantById := restaurants.CreateGetRestaurantById(database)
	findRatingsByRestaurant := ratings2.CreateFindRatingsByRestaurant(database)

	topRatedDependencies := &TopRatedDependencies{
		getRestaurantById:            getRestaurantById,
		findRatingsByRestaurant:      findRatingsByRestaurant,
		calculateRatingForRestaurant: Calculate,
	}
	dependencies := ControllerDependencies{
		GetTopRestaurants: createTopRated(topRatedDependencies),
	}
	engine.GET("/:city/restaurants/recommended", createController(&dependencies))
}
