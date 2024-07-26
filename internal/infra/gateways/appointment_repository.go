package gateways

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/drivers"
)

type AppointmentRepository interface {
	FindAppointmentsByClient(clientEmail string) ([]entities.Appoitment, error)
	InsertAppointment(appointment entities.Appoitment) error
	UpdateAppointment(appointment entities.Appoitment) (entities.Appoitment, error)
}

type appointmentRepository struct {
	dynamodbDriver drivers.DynamoDBDriver[entities.Appoitment]
}

func NewAppointmentRepository(dynamodbDriver drivers.DynamoDBDriver[entities.Appoitment]) AppointmentRepository {
	return &appointmentRepository{
		dynamodbDriver: dynamodbDriver,
	}
}

func (r *appointmentRepository) FindAppointmentsByClient(clientEmail string) ([]entities.Appoitment, error) {
	return r.dynamodbDriver.FindMany(clientEmail)
}

func (r *appointmentRepository) InsertAppointment(appointment entities.Appoitment) error {
	_ = r.dynamodbDriver.InsertOne(appointment)
	return nil
}

func (r *appointmentRepository) UpdateAppointment(appointment entities.Appoitment) (entities.Appoitment, error) {
	return r.dynamodbDriver.UpdateOne(appointment)
}