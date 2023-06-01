package ratings

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Initialize(engine *gin.Engine) {
	engine.GET("/:city/restaurants/recommended", func(c *gin.Context) {
		log.Printf("City = %v", c.Param("city"))
		c.Status(200)
	})
}
