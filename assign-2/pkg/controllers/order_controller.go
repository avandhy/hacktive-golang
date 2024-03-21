package controllers

import (
	"assign-2/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	gorm *gorm.DB
}

func NewOrderController(gorm *gorm.DB) *OrderController {
	return &OrderController{gorm: gorm}
}

func (o *OrderController) Routes(r *gin.RouterGroup) {
	routeGroup := r.Group("/orders")

	routeGroup.GET("", o.GetAllOrder)
	routeGroup.POST("", o.CreateOrder)
	routeGroup.GET("/:id", o.GetOrderbyID)
	routeGroup.PUT("/:id", o.UpdateOrder)
	routeGroup.DELETE("/:id", o.DeleteOrder)
}

func (o *OrderController) GetAllOrder(ctx *gin.Context) {
	var orders []models.Order

	err := o.gorm.Model(&models.Order{}).Preload("Item").Find(&orders).Error

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, orders)

}

func (o *OrderController) CreateOrder(ctx *gin.Context) {
	var request models.CreateOrderRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		Name:  request.Name,
		Items: request.Items,
	}

	err := o.gorm.Save(&order).Error

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, order)
}

func (o *OrderController) GetOrderbyID(ctx *gin.Context) {
	var order models.Order

	id := ctx.Param("id")

	ID, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = o.gorm.Where("id = ?", ID).First(&order).Error

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, order)
}

func (o *OrderController) UpdateOrder(ctx *gin.Context) {
	var order models.Order

	id := ctx.Param("id")

	ID, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = o.gorm.Where("id = ?", ID).First(&order).Error

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var request models.UpdateOrderRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	order.Name = request.Name
	order.Items = request.Items

	err = o.gorm.Save(&order).Error

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, order)

}

func (o *OrderController) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")

	ID, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = o.gorm.Delete(&models.Order{}, ID).Error

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Order deleted"})

}
