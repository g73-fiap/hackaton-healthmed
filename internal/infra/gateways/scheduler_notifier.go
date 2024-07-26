package gateways

import (
	"g73-fiap/hackaton-healthmed/internal/core/entities"
	"g73-fiap/hackaton-healthmed/internal/infra/drivers"
)

type SchedulerNotifier interface {
	NotifyScheduler(schedule entities.Schedule) error
}

type schedulerNotifier struct {
	sns drivers.SNS
}

func NewScheduleNotifier(sns drivers.SNS) SchedulerNotifier {
	return &schedulerNotifier{
		sns: sns,
	}
}

func (n *schedulerNotifier) NotifyScheduler(schedule entities.Schedule) error {
	return n.sns.NotifyItem("t1", schedule)
}
