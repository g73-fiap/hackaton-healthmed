package api

import (
	"g73-fiap/hackaton-healthmed/internal/controllers"

	"github.com/gin-gonic/gin"
)

type APIParams struct {
	DoctorsAPIParams
	AppointmentsAPIParams
	SchedulerAPIParams
	ClientsAPIParams
}

type DoctorsAPIParams struct {
	doctorController controllers.DoctorController
}

type AppointmentsAPIParams struct {
}

type SchedulerAPIParams struct {
}

type ClientsAPIParams struct {
}

func NewApi(params APIParams) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		createDoctorsAPIRoutes(v1, params)
		createAppointmentsAPIRoutes(v1)
		createScheduleAPIRoutes(v1)
		createClientsAPIRoutes(v1)
	}

	return router
}

func createDoctorsAPIRoutes(router *gin.RouterGroup, params APIParams) {
	router.GET("/doctors", params.doctorController.GetDoctors)
	router.POST("/doctors", params.doctorController.CreateDoctor)
	router.PUT("/doctors/:id", params.doctorController.UpdateDoctor)
}

func createAppointmentsAPIRoutes(router *gin.RouterGroup) {
	router.GET("/appointments")
	router.POST("/appointments")
	router.PUT("/appointments/:id")
	router.PUT("/appointments/:id/cancel")
	router.PUT("/appointments/:id/confirm")
}

func createScheduleAPIRoutes(router *gin.RouterGroup) {
	router.POST("/schedule")
}

func createClientsAPIRoutes(router *gin.RouterGroup) {
	router.GET("/clients")
	router.POST("/clients")
	router.PUT("/clients/:id")
	router.PUT("/clients/:id/medicalReport")
	router.DELETE("/clients/:id/medicalReport")
}
