package services

import (
	"final-project/pkg/models"

	"gorm.io/gorm"
)

type TableService struct {
	gorm *gorm.DB
}

func NewTableService(gorm *gorm.DB) *TableService {
	return &TableService{gorm: gorm}
}

func (t *TableService) CreateTable(request models.CreateTableRequest) (*models.Table, error) {
	table := models.Table{
		Number:       request.Number,
		Capacity:     request.Capacity,
		RestaurantID: request.RestaurantID,
	}

	err := t.gorm.Create(&table).Error

	if err != nil {
		return nil, err
	}

	return &table, nil
}

func (t *TableService) GetAllTable() ([]models.Table, error) {
	var tables []models.Table

	err := t.gorm.Preload("Reservations").Find(&tables).Error

	if err != nil {
		return nil, err
	}

	return tables, nil
}

func (t *TableService) GetTableByID(id int) (*models.Table, error) {
	var table models.Table

	err := t.gorm.Where("id = ?", id).Preload("Reservations").First(&table).Error

	if err != nil {
		return nil, err
	}

	return &table, nil
}

func (t *TableService) UpdateTable(id int, request models.UpdateTableRequest) (*models.Table, error) {
	table, err := t.GetTableByID(id)

	if err != nil {
		return nil, err
	}

	table.Number = request.Number
	table.Capacity = request.Capacity
	table.RestaurantID = request.RestaurantID

	err = t.gorm.Save(&table).Error

	if err != nil {
		return nil, err
	}

	return table, nil
}

func (t *TableService) DeleteTable(id int) error {
	err := t.gorm.Delete(&models.Table{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
