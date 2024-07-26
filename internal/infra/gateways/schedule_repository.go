package gateways

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/drivers"
)

type ScheduleRepository interface {
	FindSchedulesByDoctor(doctorLicenseNumber string) ([]entities.Schedule, error)
	UpdateSchedule(schedule entities.Schedule) error
}

type scheduleRepository struct {
	dynamodbDriver drivers.DynamoDBDriver[entities.Schedule]
}

func NewScheduleRepository(dynamodbDriver drivers.DynamoDBDriver[entities.Schedule]) ScheduleRepository {
	return &scheduleRepository{
		dynamodbDriver: dynamodbDriver,
	}
}

func (r *scheduleRepository) FindSchedulesByDoctor(doctorLicenseNumber string) ([]entities.Schedule, error) {
	return r.dynamodbDriver.FindMany(doctorLicenseNumber)
}

func (r *scheduleRepository) UpdateSchedule(schedule entities.Schedule) error {
	_ = r.dynamodbDriver.InsertOne(schedule)
	return nil
}
