package controllers

import (
	"final-project/pkg/models"
	"final-project/pkg/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TableController struct {
	service *services.TableService
}

func NewTableController(service *services.TableService) *TableController {
	return &TableController{service: service}
}

func (t *TableController) CreateTable(ctx *gin.Context) {
	var request models.CreateTableRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	table, err := t.service.CreateTable(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, table)
}

func (t *TableController) GetAllTable(ctx *gin.Context) {
	tables, err := t.service.GetAllTable()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tables)
}

func (t *TableController) GetTableByID(ctx *gin.Context) {
	id := ctx.Param("id")

	tableID, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid table ID"})
		return
	}

	table, err := t.service.GetTableByID(tableID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, table)
}

func (t *TableController) UpdateTable(ctx *gin.Context) {
	id := ctx.Param("id")

	tableID, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid table ID"})
		return
	}

	var request models.UpdateTableRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	table, err := t.service.UpdateTable(tableID, request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if table == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Table not found"})
	}

	ctx.JSON(http.StatusOK, table)
}

func (t *TableController) DeleteTable(ctx *gin.Context) {
	id := ctx.Param("id")

	tableID, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid table ID"})
		return
	}

	err = t.service.DeleteTable(tableID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Table deleted successfully"})
}
