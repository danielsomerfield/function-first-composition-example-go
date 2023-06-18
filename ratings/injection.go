package ratings

import (
	"github.com/gin-gonic/gin"
)

func Initialize(engine *gin.Engine) {
	dependencies := Dependencies{
		GetTopRestaurants: func() ([]Restaurant, error) {
			return []Restaurant{}, nil
		},
	}
	engine.GET("/:city/restaurants/recommended", createController(&dependencies))
}
