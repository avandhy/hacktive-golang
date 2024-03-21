package routers

import (
	// "assign-2/pkg/controllers"
	"assign-2/pkg/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(gorm *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})

	api := r.Group("/api")

	productController := controllers.NewOrderController(gorm)
	productController.Routes(api)

	return r
}
