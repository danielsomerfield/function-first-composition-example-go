package ratings

import (
	"github.com/gin-gonic/gin"
)

func Initialize(engine *gin.Engine) {
	dependencies := Dependencies{
		GetTopRestaurants: func() []Restaurant {
			return []Restaurant{}
		},
	}
	engine.GET("/:city/restaurants/recommended", createController(&dependencies))
}
