package ratings

import (
	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	GetTopRestaurants func() ([]Restaurant, error)
}

func createController(dependencies *Dependencies) func(c *gin.Context) {

	return func(c *gin.Context) {
		restaurants, err := (*dependencies).GetTopRestaurants()
		if err == nil {
			body := ResponseBody{
				Restaurants: restaurants,
			}
			c.JSON(200, body)
		} else {
			c.JSON(500, struct{}{})
		}
	}
}

type ResponseBody struct {
	Restaurants []Restaurant `json:"restaurants"`
}
