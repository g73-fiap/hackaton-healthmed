package usecases

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/gateways"
)

type ScheduleUseCase interface {
	GetDoctorSchedules(doctorLicenseNumber string) ([]entities.Schedule, error)
	UpdateSchedule(entities.Schedule) error
}

type scheduleUseCase struct {
	scheduleRepository gateways.ScheduleRepository
}

func NewScheduleUseCase(scheduleRepository gateways.ScheduleRepository) ScheduleUseCase {
	return &scheduleUseCase{
		scheduleRepository: scheduleRepository,
	}
}

func (u *scheduleUseCase) GetDoctorSchedules(doctorLicenseNumber string) ([]entities.Schedule, error) {
	return u.scheduleRepository.FindSchedulesByDoctor(doctorLicenseNumber)
}

func (u *scheduleUseCase) UpdateSchedule(schedule entities.Schedule) error {
	return u.scheduleRepository.UpdateSchedule(schedule)
}
