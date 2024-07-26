package api

import (
	"g73-fiap/hackaton-healthmed/internal/controllers"

	"github.com/gin-gonic/gin"
)

type APIParams struct {
	DoctorController       controllers.DoctorController
	SchedulerController    controllers.SchedulerController
	ClientController       controllers.ClientController
	AppointmentsController controllers.AppointmentController
}

func NewApi(params APIParams) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		createDoctorsAPIRoutes(v1, params)
		createClientsAPIRoutes(v1, params)
		createAppointmentsAPIRoutes(v1, params)
		createScheduleAPIRoutes(v1, params)
	}

	return router
}

func createDoctorsAPIRoutes(router *gin.RouterGroup, params APIParams) {
	router.GET("/doctors", params.DoctorController.GetDoctors)
	router.POST("/doctors", params.DoctorController.CreateDoctor)
	router.PUT("/doctors/:id", params.DoctorController.UpdateDoctor)
	router.DELETE("/doctors/:id", params.DoctorController.DeleteDoctor)
}

func createAppointmentsAPIRoutes(router *gin.RouterGroup, params APIParams) {
	router.GET("/appointments", params.AppointmentsController.GetClientAppointments)
	router.POST("/appointments", params.AppointmentsController.CreateAppointment)
	router.PUT("/appointments/:id/cancel", params.AppointmentsController.ConfirmAppointment)
	router.PUT("/appointments/:id/confirm", params.AppointmentsController.CancelAppointment)
}

func createScheduleAPIRoutes(router *gin.RouterGroup, params APIParams) {
	router.GET("/schedule", params.SchedulerController.GetDoctorSchedules)
	router.PUT("/schedule", params.SchedulerController.UpdateSchedule)
}

func createClientsAPIRoutes(router *gin.RouterGroup, params APIParams) {
	router.GET("/clients/:id", params.ClientController.GetClient)
	router.POST("/clients", params.ClientController.CreateClient)
	router.PUT("/clients/:id", params.ClientController.UpdateClient)
	router.DELETE("/clients/:id", params.ClientController.DeleteClient)
	router.PUT("/clients/:id/medicalReport", params.ClientController.SaveMedicalReport)
	router.DELETE("/clients/:id/medicalReport", params.ClientController.DeleteMedicalReport)
	router.GET("/clients/:id/medicalReport/:fileName", params.ClientController.GetMedicalReport)
	router.PUT("/clients/:id/medicalReport/:fileName/share", params.ClientController.ShareMedicalReport)
}
