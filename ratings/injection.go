package ratings

import (
	"function-first-composition-example-go/review-server/configuration"
	"function-first-composition-example-go/review-server/db"
	"function-first-composition-example-go/review-server/restaurants"
	"github.com/gin-gonic/gin"
)

func Initialize(engine *gin.Engine, config *configuration.Configuration) {

	database := db.Connect(config.DataSource)

	getRestaurantById := restaurants.CreateGetRestaurantById(database)
	findRatingsByRestaurant := CreateFindRatingsByRestaurant(database)

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
