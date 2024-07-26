package main

import (
	"g73-fiap/hackaton-healthmed/internal/api"
	"g73-fiap/hackaton-healthmed/internal/controllers"
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/core/usecases"
	"g73-fiap/hackaton-healthmed/internal/infra/drivers"
	"g73-fiap/hackaton-healthmed/internal/infra/gateways"
)

func main() {
	port := "8080"

	doctorMongoDBDriver := drivers.NewMongoDriver[entities.Doctor]()
	s3Driver := drivers.NewS3Driver()
	queue := drivers.NewQueue()
	sns := drivers.NewSNS()
	google := drivers.NewGoogleAPI()

	clientDynamoDBDriver := drivers.NewDynamoDBDriver[entities.Client]()
	appointmentDynamoDBDriver := drivers.NewDynamoDBDriver[entities.Appointment]()
	schedulerDynamoDBDriver := drivers.NewDynamoDBDriver[entities.Schedule]()

	doctorRepository := gateways.NewDoctorRepository(doctorMongoDBDriver)
	clientRepository := gateways.NewClientRepository(clientDynamoDBDriver, s3Driver)
	appointmentRepository := gateways.NewAppointmentRepository(appointmentDynamoDBDriver)
	appointmentNotifier := gateways.NewAppointmentNotifier(queue)
	schedulerRepository := gateways.NewScheduleRepository(schedulerDynamoDBDriver)
	schedulerNotifier := gateways.NewScheduleNotifier(sns)

	doctorUseCase := usecases.NewDoctorUseCase(doctorRepository)
	clientUseCase := usecases.NewClientUseCase(clientRepository)
	appointmentUseCase := usecases.NewAppointmentUseCase(appointmentRepository, appointmentNotifier)
	schedulerUseCase := usecases.NewScheduleUseCase(schedulerRepository, schedulerNotifier, google)

	doctorController := controllers.NewDoctorController(doctorUseCase)
	clientController := controllers.NewClientController(clientUseCase)
	appointmentController := controllers.NewAppointmentController(appointmentUseCase)
	schedulerController := controllers.NewSchedulerController(schedulerUseCase)

	apiParams := api.APIParams{
		DoctorController:       doctorController,
		ClientController:       clientController,
		AppointmentsController: appointmentController,
		SchedulerController:    schedulerController,
	}
	api := api.NewApi(apiParams)
	api.Run(":" + port)
}
