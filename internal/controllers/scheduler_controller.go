package controllers

import (
	"errors"
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/core/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SchedulerController struct {
	schedulerUseCase usecases.ScheduleUseCase
}

func NewSchedulerController(schedulerUseCase usecases.ScheduleUseCase) SchedulerController {
	return SchedulerController{
		schedulerUseCase: schedulerUseCase,
	}
}

func (s SchedulerController) GetDoctorSchedules(c *gin.Context) {
	doctorLicenseNumber := c.Query("doctorLicenseNumber")
	if doctorLicenseNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("doctorLicenseNumber is required")})
		return
	}

	schedules, err := s.schedulerUseCase.GetDoctorSchedules(doctorLicenseNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

func (s SchedulerController) UpdateSchedule(c *gin.Context) {
	var schedule entities.Schedule
	err := c.ShouldBindJSON(&schedule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = s.schedulerUseCase.UpdateSchedule(schedule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
