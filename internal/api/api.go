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
	schedulerController controllers.SchedulerController
}

type ClientsAPIParams struct {
	clientController controllers.ClientController
}

func NewApi(params APIParams) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		createDoctorsAPIRoutes(v1, params)
		createClientsAPIRoutes(v1, params)
		createAppointmentsAPIRoutes(v1)
		createScheduleAPIRoutes(v1, params)
	}

	return router
}

func createDoctorsAPIRoutes(router *gin.RouterGroup, params APIParams) {
	router.GET("/doctors", params.doctorController.GetDoctors)
	router.POST("/doctors", params.doctorController.CreateDoctor)
	router.PUT("/doctors/:id", params.doctorController.UpdateDoctor)
	router.DELETE("/doctors/:id", params.doctorController.DeleteDoctor)
}

func createAppointmentsAPIRoutes(router *gin.RouterGroup) {
	router.GET("/appointments")
	router.POST("/appointments")
	router.PUT("/appointments/:id")
	router.PUT("/appointments/:id/cancel")
	router.PUT("/appointments/:id/confirm")
}

func createScheduleAPIRoutes(router *gin.RouterGroup, params APIParams) {
	router.GET("/schedule", params.schedulerController.GetDoctorSchedules)
	router.PUT("/schedule", params.schedulerController.UpdateSchedule)
}

func createClientsAPIRoutes(router *gin.RouterGroup, params APIParams) {
	router.GET("/clients/:id", params.clientController.GetClient)
	router.POST("/clients", params.clientController.CreateClient)
	router.PUT("/clients/:id", params.clientController.UpdateClient)
	router.DELETE("/clients/:id", params.clientController.DeleteClient)
	router.PUT("/clients/:id/medicalReport")
	router.DELETE("/clients/:id/medicalReport")
}
