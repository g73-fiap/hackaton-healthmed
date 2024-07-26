package controllers

import (
	"errors"
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/core/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppointmentController struct {
	appointmentUseCase usecases.AppointmentUseCase
}

func (a AppointmentController) GetClientAppointments(c *gin.Context) {
	clientEmail := c.Query("clientEmail")
	if clientEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("clientEmail is required")})
		return
	}

	client, err := a.appointmentUseCase.FindClientAppointments(clientEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

func (a AppointmentController) CreateAppointment(c *gin.Context) {
	var appointment entities.Appoitment
	err := c.ShouldBindJSON(&appointment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = a.appointmentUseCase.CreateAppointment(appointment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (a AppointmentController) ConfirmAppointment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	var appointment entities.Appoitment
	err := c.ShouldBindJSON(&appointment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = a.appointmentUseCase.ConfirmAppointment(appointment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (a AppointmentController) CancelAppointment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("id is required")})
		return
	}

	var appointment entities.Appoitment
	err := c.ShouldBindJSON(&appointment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = a.appointmentUseCase.CancelAppointment(appointment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
