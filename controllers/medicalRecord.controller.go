package controllers

import (
	"medical_api/dto"
	"medical_api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MedicalRecordController interface {
	GetAllMedicalRecords(ctx *gin.Context)
	CreateMedicalRecord(ctx *gin.Context)
	FindMedicalRecordByID(ctx *gin.Context)
	UpdateMedicalRecord(ctx *gin.Context)
	DeleteMedicalRecord(ctx *gin.Context)
}

type medicalRecordController struct {
	medicalRecordService services.MedicalRecordService
}

func NewMedicalRecordController(service services.MedicalRecordService) MedicalRecordController {
	return &medicalRecordController{
		medicalRecordService: service,
	}
}

func (c *medicalRecordController) GetAllMedicalRecords(ctx *gin.Context) {
	records := c.medicalRecordService.GetAllMedicalRecords()
	ctx.JSON(http.StatusOK, records)
}

func (c *medicalRecordController) CreateMedicalRecord(ctx *gin.Context) {
	var recordDTO dto.CreateMedicalRecordDTO
	if err := ctx.ShouldBindJSON(&recordDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record := c.medicalRecordService.CreateMedicalRecord(recordDTO)
	ctx.JSON(http.StatusCreated, record)
}

func (c *medicalRecordController) FindMedicalRecordByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	record := c.medicalRecordService.FindMedicalRecordByID(id)
	if record.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	ctx.JSON(http.StatusOK, record)
}

func (c *medicalRecordController) UpdateMedicalRecord(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var recordDTO dto.UpdateMedicalRecordDTO
	if err := ctx.ShouldBindJSON(&recordDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedRecord := c.medicalRecordService.UpdateMedicalRecord(recordDTO, id)
	ctx.JSON(http.StatusOK, updatedRecord)
}

func (c *medicalRecordController) DeleteMedicalRecord(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	deletedRecord := c.medicalRecordService.DeleteMedicalRecord(id)
	ctx.JSON(http.StatusOK, deletedRecord)
}
