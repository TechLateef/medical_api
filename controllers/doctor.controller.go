package controllers

import (
	"medical_api/dto"
	"medical_api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DoctorController interface {
	GetAllDoctors(ctx *gin.Context)
	CreateDoctor(ctx *gin.Context)
	FindDoctorByID(ctx *gin.Context)
	UpdateDoctor(ctx *gin.Context)
	DeleteDoctor(ctx *gin.Context)
}

type doctorController struct {
	doctorService services.DoctorService
}

func NewDoctorController(service services.DoctorService) DoctorController {
	return &doctorController{
		doctorService: service,
	}
}

func (c *doctorController) GetAllDoctors(ctx *gin.Context) {
	doctors := c.doctorService.GetAllDoctors()
	ctx.JSON(http.StatusOK, doctors)
}

func (c *doctorController) CreateDoctor(ctx *gin.Context) {
	var doctor dto.CreateDoctorDTO
	if err := ctx.ShouldBindJSON(&doctor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdDoctor := c.doctorService.CreateDoctor(doctor)
	ctx.JSON(http.StatusCreated, createdDoctor)
}

func (c *doctorController) FindDoctorByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	doctor := c.doctorService.GetDoctorByID(id)
	if doctor.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}
	ctx.JSON(http.StatusOK, doctor)
}

func (c *doctorController) UpdateDoctor(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var doctor dto.UpdateDoctorDTO
	if err := ctx.ShouldBindJSON(&doctor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	doctor.Id = id

	updatedDoctor := c.doctorService.UpdateDoctor(doctor, id)
	ctx.JSON(http.StatusOK, updatedDoctor)
}

func (c *doctorController) DeleteDoctor(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	deletedDoctor := c.doctorService.DeleteDoctor(id)
	ctx.JSON(http.StatusOK, deletedDoctor)
}
