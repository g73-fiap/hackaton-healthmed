package api

import (
	"g73-fiap/hackaton-healthmed/internal/controllers"

	"github.com/gin-gonic/gin"
)

type APIParams struct {
	doctorController       controllers.DoctorController
	schedulerController    controllers.SchedulerController
	clientController       controllers.ClientController
	appointmentsController controllers.AppointmentController
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
	router.GET("/doctors", params.doctorController.GetDoctors)
	router.POST("/doctors", params.doctorController.CreateDoctor)
	router.PUT("/doctors/:id", params.doctorController.UpdateDoctor)
	router.DELETE("/doctors/:id", params.doctorController.DeleteDoctor)
}

func createAppointmentsAPIRoutes(router *gin.RouterGroup, params APIParams) {
	router.GET("/appointments", params.appointmentsController.GetClientAppointments)
	router.POST("/appointments", params.appointmentsController.CreateAppointment)
	router.PUT("/appointments/:id/cancel", params.appointmentsController.ConfirmAppointment)
	router.PUT("/appointments/:id/confirm", params.appointmentsController.CancelAppointment)
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
	router.PUT("/clients/:id/medicalReport", params.clientController.SaveMedicalReport)
	router.DELETE("/clients/:id/medicalReport", params.clientController.DeleteMedicalReport)
	router.GET("/clients/:id/medicalReport/:fileName", params.clientController.GetMedicalReport)
	router.PUT("/clients/:id/medicalReport/:fileName/share", params.clientController.ShareMedicalReport)
}
