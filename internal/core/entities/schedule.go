package entities

import "time"

type Schedule struct {
	DoctorLicenseNumber string
	StartTime           time.Time
	EndTime             time.Time
	ReservedBy          string
	ReservedAt          time.Time
	CanceledAt          time.Time
}

func (s Schedule) GetHashKey() string {
	return s.DoctorLicenseNumber
}

func (s Schedule) GetSortKey() string {
	return s.StartTime.String()
}
