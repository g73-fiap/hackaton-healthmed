package usecases

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/drivers"
	"g73-fiap/hackaton-healthmed/internal/infra/gateways"
)

type ScheduleUseCase interface {
	GetDoctorSchedules(doctorLicenseNumber string) ([]entities.Schedule, error)
	UpdateSchedule(entities.Schedule) error
}

type scheduleUseCase struct {
	scheduleRepository gateways.ScheduleRepository
	scheduleNotifier   gateways.SchedulerNotifier
	googleAPI          drivers.GoogleAPI
}

func NewScheduleUseCase(scheduleRepository gateways.ScheduleRepository, scheduleNotifier gateways.SchedulerNotifier, googleAPI drivers.GoogleAPI) ScheduleUseCase {
	return &scheduleUseCase{
		scheduleRepository: scheduleRepository,
		scheduleNotifier:   scheduleNotifier,
		googleAPI:          googleAPI,
	}
}

func (u *scheduleUseCase) GetDoctorSchedules(doctorLicenseNumber string) ([]entities.Schedule, error) {
	return u.scheduleRepository.FindSchedulesByDoctor(doctorLicenseNumber)
}

func (u *scheduleUseCase) UpdateSchedule(schedule entities.Schedule) error {
	meeting, err := u.googleAPI.CreateMeeting()
	if err != nil {
		return err
	}

	schedule.MeetingLink = meeting

	err = u.scheduleRepository.UpdateSchedule(schedule)
	if err != nil {
		return err
	}

	err = u.scheduleNotifier.NotifyScheduler(schedule)
	if err != nil {
		return err
	}

	return nil
}
