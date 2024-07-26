package controllers

import (
	"errors"
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/core/usecases"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type DoctorController struct {
	doctorUseCase usecases.DoctorUseCase
}

func NewDoctorController(doctorUseCase usecases.DoctorUseCase) DoctorController {
	return DoctorController{
		doctorUseCase: doctorUseCase,
	}
}

func (d DoctorController) GetDoctors(c *gin.Context) {
	filter, err := extractQueryParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}
	doctors, err := d.doctorUseCase.GetAllDoctors(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, doctors)
}

func (d DoctorController) CreateDoctor(c *gin.Context) {
	var doctor entities.Doctor
	err := c.ShouldBindJSON(&doctor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = d.doctorUseCase.CreateDoctor(doctor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (d DoctorController) UpdateDoctor(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	var doctor entities.Doctor
	err := c.ShouldBindJSON(&doctor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = d.doctorUseCase.UpdateDoctor(id, doctor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (d DoctorController) DeleteDoctor(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	err := d.doctorUseCase.DeleteDoctor(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func extractQueryParams(c *gin.Context) (*usecases.GetDoctorsFilter, error) {
	coordinatesParam := c.Query("coordinates")
	maxDistanceParam := c.Query("maxDistance")

	if coordinatesParam == "" || maxDistanceParam == "" {
		return nil, nil
	}

	coordinates := strings.Split(coordinatesParam, ",")
	if len(coordinates) != 0 {
		return nil, errors.New("invalid coordinates")
	}

	lat, err := strconv.Atoi(coordinates[0])
	if err != nil {
		return nil, err
	}

	long, err := strconv.Atoi(coordinates[1])
	if err != nil {
		return nil, err
	}

	maxDistance, err := strconv.Atoi(maxDistanceParam)
	if err != nil {
		return nil, err
	}

	return &usecases.GetDoctorsFilter{
		Coordinates: [2]int{lat, long},
		MaxDistance: maxDistance,
	}, nil
}
