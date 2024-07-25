package usecases

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/gateways"
)

type DoctorUseCase interface {
	GetAllDoctors(filter *GetDoctorsFilter) ([]entities.Doctor, error)
	CreateDoctor(doctor entities.Doctor) error
	UpdateDoctor(id string, doctor entities.Doctor) error
	DeleteDoctor(id string) error
}

type doctorUserCase struct {
	doctorRepository gateways.DoctorRepository
}

func NewDoctorUseCase(doctorRepository gateways.DoctorRepository) DoctorUseCase {
	return &doctorUserCase{
		doctorRepository: doctorRepository,
	}
}

type GetDoctorsFilter struct {
	Coordinates [2]int
	MaxDistance int
}

func (u *doctorUserCase) GetAllDoctors(filter *GetDoctorsFilter) ([]entities.Doctor, error) {
	if filter != nil {
		return u.doctorRepository.FindDoctorsByNear(filter.Coordinates, filter.MaxDistance)
	}

	return u.doctorRepository.FindDoctors()
}

func (u *doctorUserCase) CreateDoctor(doctor entities.Doctor) error {
	return u.doctorRepository.InsertDoctor(doctor)
}

func (u *doctorUserCase) UpdateDoctor(id string, doctor entities.Doctor) error {
	_, err := u.doctorRepository.UpdateDoctor(id, doctor)
	return err
}

func (u *doctorUserCase) DeleteDoctor(id string) error {
	return u.doctorRepository.DeleteDoctor(id)
}
