package gateways

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/drivers"
)

type DoctorRepository interface {
	FindDoctors() ([]entities.Doctor, error)
	FindDoctorsByNear(cordinates [2]int, maxDistanceInMeters int) ([]entities.Doctor, error)
	FindDoctorByID(id string) (entities.Doctor, error)
	InsertDoctor(doctor entities.Doctor) error
	UpdateDoctor(id string, doctor entities.Doctor) (entities.Doctor, error)
	DeleteDoctor(id string) error
}

type doctorRepository struct {
	mongoDriver drivers.MongoDriver[entities.Doctor]
}

func NewDoctorRepository(mongoDriver drivers.MongoDriver[entities.Doctor]) DoctorRepository {
	return &doctorRepository{
		mongoDriver: mongoDriver,
	}
}

func (r *doctorRepository) FindDoctors() ([]entities.Doctor, error) {
	return r.mongoDriver.FindAll(), nil
}

func (r *doctorRepository) FindDoctorsByNear(cordinates [2]int, maxDistanceInMeters int) ([]entities.Doctor, error) {
	return r.mongoDriver.FindByNear(cordinates, maxDistanceInMeters), nil
}

func (r *doctorRepository) FindDoctorByID(id string) (entities.Doctor, error) {
	return r.mongoDriver.FindByID(id)
}

func (r *doctorRepository) InsertDoctor(doctor entities.Doctor) error {
	_ = r.mongoDriver.InsertOne(doctor)
	return nil
}

func (r *doctorRepository) UpdateDoctor(id string, doctor entities.Doctor) (entities.Doctor, error) {
	return r.mongoDriver.UpdateOne(id, doctor)
}

func (r *doctorRepository) DeleteDoctor(id string) error {
	return r.mongoDriver.DeleteOne(id)
}
