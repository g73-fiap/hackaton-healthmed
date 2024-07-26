package gateways

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/drivers"
)

type AppointmentNotifier interface {
	NotifyAppointment(appointment entities.Appointment) error
}

type appointmentNotifier struct {
	queue drivers.Queue
}

func NewAppointmentNotifier(queue drivers.Queue) AppointmentNotifier {
	return &appointmentNotifier{
		queue: queue,
	}
}

func (n *appointmentNotifier) NotifyAppointment(appointment entities.Appointment) error {
	return n.queue.PushItem("q1", appointment)
}
