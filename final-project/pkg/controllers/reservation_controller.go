package controllers

import (
	"final-project/pkg/models"
	"final-project/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReservationController struct {
	service *services.ReservationService
	userController *UserController
}

func NewReservationController(service *services.ReservationService, userController *UserController) *ReservationController {
	return &ReservationController{service: service, userController: userController}
}

func (r *ReservationController) CreateReservation(ctx *gin.Context) {
	var request models.CreateReservationRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := r.userController.GetIdUserActive(ctx)

	reservation, err := r.service.CreateReservation(request, userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, reservation)
}

func (r *ReservationController) GetAllReservation(ctx *gin.Context) {
	
	userId := r.userController.GetIdUserActive(ctx)

	reservations, err := r.service.GetAllReservation(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}

func (r *ReservationController) GetReservationByID(ctx *gin.Context) {
	id := ctx.Param("id")

	reservationID, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid reservation ID"})
		return
	}

	reservation, err := r.service.GetReservationByID(reservationID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reservation)
}

func (r *ReservationController) UpdateReservation(ctx *gin.Context) {
	id := ctx.Param("id")

	reservationID, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid reservation ID"})
		return
	}

	var request models.UpdateReservationRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := r.userController.GetIdUserActive(ctx)

	reservation, err := r.service.UpdateReservation(reservationID, request, userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if reservation == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
	}

	ctx.JSON(http.StatusOK, reservation)
}

func (r *ReservationController) CancelReservation(ctx *gin.Context) {
	id := ctx.Param("id")

	reservationID, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid reservation ID"})
		return
	}

	userId := r.userController.GetIdUserActive(ctx)

	err = r.service.DeleteReservation(reservationID, userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reservation canceled successfully"})
}
