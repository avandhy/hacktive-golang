package routers

import (
	"final-project/pkg/controllers"
	"final-project/pkg/middleware"
	"final-project/pkg/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(gorm *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, world!")
	})

	// inisiasi controller
	userService := services.NewUserService(gorm)
	userController := controllers.NewUserController(userService)
	tableService := services.NewTableService(gorm)
	tableController := controllers.NewTableController(tableService)
	reservationService := services.NewReservationService(gorm)
	reservationController := controllers.NewReservationController(reservationService, userController)
	
	//	routing
	r.GET("/reservation", middleware.VerifyToken, reservationController.GetAllReservation)
	r.POST("/reservation", middleware.VerifyToken, reservationController.CreateReservation)
	r.GET("/reservation/:id", middleware.VerifyToken, reservationController.GetReservationByID)
	r.PUT("/reservation/:id", middleware.VerifyToken, reservationController.UpdateReservation)
	r.DELETE("/reservation/:id", middleware.VerifyToken, reservationController.CancelReservation)

	r.GET("/table", tableController.GetAllTable)
	r.POST("/table", tableController.CreateTable)
	r.GET("/table/:id", tableController.GetTableByID)
	r.PUT("/table/:id", tableController.UpdateTable)
	r.DELETE("/table/:id", tableController.DeleteTable)

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)

	return r
}
