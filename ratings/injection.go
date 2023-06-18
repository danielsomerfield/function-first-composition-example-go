package ratings

import (
	"github.com/gin-gonic/gin"
)

func Initialize(engine *gin.Engine) {

	dependencies := ControllerDependencies{
		GetTopRestaurants: createTopRated(&TopRatedDependencies{
			getRestaurantById:            nil,
			findRatingsByRestaurant:      nil,
			calculateRatingForRestaurant: nil,
		}),
	}
	engine.GET("/:city/restaurants/recommended", createController(&dependencies))
}
