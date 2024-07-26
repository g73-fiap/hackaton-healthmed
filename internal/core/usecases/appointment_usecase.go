package usecases

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/gateways"
)

type AppointmentUseCase interface {
	FindClientAppointments(clientEmail string) ([]entities.Appointment, error)
	CreateAppointment(appointment entities.Appointment) error
	ConfirmAppointment(appointmentId entities.Appointment) error
	CancelAppointment(appointment entities.Appointment) error
}

type appointmentUseCase struct {
	appointmentRepository gateways.AppointmentRepository
	appointmentNotifier   gateways.AppointmentNotifier
}

func NewAppointmentUseCase(appointmentRepository gateways.AppointmentRepository, appointmentNotifier gateways.AppointmentNotifier) AppointmentUseCase {
	return &appointmentUseCase{
		appointmentRepository: appointmentRepository,
		appointmentNotifier:   appointmentNotifier,
	}
}

func (u *appointmentUseCase) FindClientAppointments(clientEmail string) ([]entities.Appointment, error) {
	return u.appointmentRepository.FindAppointmentsByClient(clientEmail)
}

func (u *appointmentUseCase) CreateAppointment(appointment entities.Appointment) error {
	return u.appointmentRepository.InsertAppointment(appointment)
}

func (u *appointmentUseCase) ConfirmAppointment(appointment entities.Appointment) error {
	_, err := u.appointmentRepository.UpdateAppointment(appointment)
	if err != nil {
		return err
	}

	err = u.appointmentNotifier.NotifyAppointment(appointment)
	if err != nil {
		return err
	}

	return nil
}

func (u *appointmentUseCase) CancelAppointment(appointment entities.Appointment) error {
	_, err := u.appointmentRepository.UpdateAppointment(appointment)
	if err != nil {
		return err
	}

	err = u.appointmentNotifier.NotifyAppointment(appointment)
	if err != nil {
		return err
	}

	return nil
	return err
}
