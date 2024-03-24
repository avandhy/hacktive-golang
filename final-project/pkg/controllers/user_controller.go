package controllers

import (
	"final-project/pkg/models"
	"final-project/pkg/services"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) Login(ctx *gin.Context) {
	var request models.LoginRequest

	err := ctx.ShouldBind(&request)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := c.userService.Login(request)

	if err != nil {
		ctx.JSON(401, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})

}

func (c *UserController) Register(ctx *gin.Context) {
	var request models.RegisterRequest

	err := ctx.ShouldBind(&request)

	fmt.Printf("%+v", request)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := c.userService.Register(request)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})

}

// func (u *UserController) Logout(ctx *gin.Context) {
// 	ctx.Set("user", "")
// 	ctx.JSON(200, gin.H{"success": "user logged out"})
// }

func (u *UserController) GetIdUserActive(ctx *gin.Context) uint {

	user, _ := ctx.Get("user")

	userData := user.(jwt.MapClaims)
	userfloatID := userData["id"].(float64)
	
	userID := uint(userfloatID)

	fmt.Printf("username %v", userID)
	return userID

}
