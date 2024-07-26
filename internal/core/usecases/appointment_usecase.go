package usecases

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/gateways"
)

type AppointmentUseCase interface {
	FindClientAppointments(clientEmail string) ([]entities.Appoitment, error)
	CreateAppointment(appointment entities.Appoitment) error
	ConfirmAppointment(appointmentId entities.Appoitment) error
	CancelAppointment(appointment entities.Appoitment) error
}

type appointmentUseCase struct {
	appointmentRepository gateways.AppointmentRepository
}

func NewAppointmentUseCase(appointmentRepository gateways.AppointmentRepository) AppointmentUseCase {
	return &appointmentUseCase{
		appointmentRepository: appointmentRepository,
	}
}

func (u *appointmentUseCase) FindClientAppointments(clientEmail string) ([]entities.Appoitment, error) {
	return u.appointmentRepository.FindAppointmentsByClient(clientEmail)
}

func (u *appointmentUseCase) CreateAppointment(appointment entities.Appoitment) error {
	return u.appointmentRepository.InsertAppointment(appointment)
}

func (u *appointmentUseCase) ConfirmAppointment(appointment entities.Appoitment) error {
	_, err := u.appointmentRepository.UpdateAppointment(appointment)
	return err
}

func (u *appointmentUseCase) CancelAppointment(appointment entities.Appoitment) error {
	_, err := u.appointmentRepository.UpdateAppointment(appointment)
	return err
}
