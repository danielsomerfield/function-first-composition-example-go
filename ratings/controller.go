package ratings

import (
	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	GetTopRestaurants func() []Restaurant
}

func createController(dependencies *Dependencies) func(c *gin.Context) {

	return func(c *gin.Context) {
		restaurants := (*dependencies).GetTopRestaurants()
		body := ResponseBody{
			Restaurants: restaurants,
		}
		c.JSON(200, body)
	}
}

type ResponseBody struct {
	Restaurants []Restaurant `json:"restaurants"`
}
