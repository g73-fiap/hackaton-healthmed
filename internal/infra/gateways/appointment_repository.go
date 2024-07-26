package gateways

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/drivers"
)

type AppointmentRepository interface {
	FindAppointmentsByClient(clientEmail string) ([]entities.Appointment, error)
	InsertAppointment(appointment entities.Appointment) error
	UpdateAppointment(appointment entities.Appointment) (entities.Appointment, error)
}

type appointmentRepository struct {
	dynamodbDriver drivers.DynamoDBDriver[entities.Appointment]
}

func NewAppointmentRepository(dynamodbDriver drivers.DynamoDBDriver[entities.Appointment]) AppointmentRepository {
	return &appointmentRepository{
		dynamodbDriver: dynamodbDriver,
	}
}

func (r *appointmentRepository) FindAppointmentsByClient(clientEmail string) ([]entities.Appointment, error) {
	return r.dynamodbDriver.FindMany(clientEmail)
}

func (r *appointmentRepository) InsertAppointment(appointment entities.Appointment) error {
	_ = r.dynamodbDriver.InsertOne(appointment)
	return nil
}

func (r *appointmentRepository) UpdateAppointment(appointment entities.Appointment) (entities.Appointment, error) {
	return r.dynamodbDriver.UpdateOne(appointment)
}
