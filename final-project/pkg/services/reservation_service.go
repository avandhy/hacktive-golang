package services

import (
	"errors"
	"final-project/pkg/models"

	"gorm.io/gorm"
)

type ReservationService struct {
	gorm *gorm.DB
}

func NewReservationService(gorm *gorm.DB) *ReservationService {
	return &ReservationService{gorm: gorm}
}

func (r *ReservationService) CreateReservation(request models.CreateReservationRequest, userId uint) (*models.Reservation, error) {
	user, err := r.GetUserActive(userId)

	if err != nil {
		return nil, err
	}

	reservation := models.Reservation{
		Name:    user.FullName,
		Email:   user.Email,
		Phone:   request.Phone,
		Date:    request.Date,
		TableID: request.TableID,
		UserID:  user.ID,
	}

	avail, err := r.isTableAvailable(request.TableID)

	if err != nil {
		return nil, err
	}
	
	
	if avail {
		err = r.gorm.Create(&reservation).Error
	
		if err != nil {
			return nil, err
		}
	
		err = r.UpdateTableStatus(reservation.TableID, "booked")

		if err != nil {
			return nil, err
		}
	} else{
		return nil, errors.New("table not available")
	}
	

	return &reservation, nil
}

func (r *ReservationService) GetAllReservation(userId uint) ([]models.Reservation, error) {
	var reservations []models.Reservation

	err := r.gorm.Where("user_id = ?", userId).Find(&reservations).Error

	if err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *ReservationService) GetReservationByID(id int) (*models.Reservation, error) {
	var reservation models.Reservation

	err := r.gorm.Where("id = ?", id).First(&reservation).Error

	if err != nil {
		return nil, err
	}

	return &reservation, nil
}

func (r *ReservationService) UpdateReservation(id int, request models.UpdateReservationRequest, userId uint) (*models.Reservation, error) {
	reservation, err := r.GetReservationByID(id)

	if err != nil {
		return nil, err
	}

	if reservation.UserID != userId {
		return nil, errors.New("Unauthorized")
	}

	err = r.UpdateTableStatus(reservation.TableID, "available")

	if err != nil {
		return nil, err
	}

	reservation.Date = request.Date
	reservation.TableID = request.TableID

	err = r.gorm.Save(&reservation).Error

	if err != nil {
		return nil, err
	}

	err = r.UpdateTableStatus(request.TableID, "booked")

	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (r *ReservationService) DeleteReservation(id int, userId uint) error {
	reservation, err := r.GetReservationByID(id)

	if err != nil {
		return err
	}

	if reservation.UserID != userId {
		return errors.New("Unauthorized")
	}

	err = r.gorm.Delete(&models.Reservation{}, id).Error

	if err != nil {
		return err
	}

	err = r.UpdateTableStatus(reservation.TableID, "available")

	if err != nil {
		return err
	}

	return nil
}

func (r *ReservationService) UpdateTableStatus(tableId int, status string) error {

	err := r.gorm.Model(&models.Table{}).Where("id = ?", tableId).Update("status", status).Error

	if err != nil {
		return err
	}

	return err
}

func (r *ReservationService) GetUserActive(userId uint) (*models.User, error) {
	var user models.User

	err := r.gorm.Model(&models.User{}).First(&user, userId).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *ReservationService) isTableAvailable(tableId int) (bool, error) {
	var table models.Table

	err := r.gorm.Model(&models.Table{}).First(&table, tableId).Error

	if err != nil {
		return false, err
	}

	if table.Status == "available" {
		return true, nil
	}

	return false, nil

}
