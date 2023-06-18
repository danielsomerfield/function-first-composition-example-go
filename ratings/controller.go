package ratings

import (
	"github.com/gin-gonic/gin"
)

type ControllerDependencies struct {
	GetTopRestaurants func(city string) ([]Restaurant, error)
}

func createController(dependencies *ControllerDependencies) func(c *gin.Context) {

	return func(c *gin.Context) {
		city := c.Param("city")
		if city == "" {
			c.JSON(400, struct{}{})
		} else {
			restaurants, err := (*dependencies).GetTopRestaurants(city)
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
}

type ResponseBody struct {
	Restaurants []Restaurant `json:"restaurants"`
}
