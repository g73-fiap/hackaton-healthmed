package entities

import "time"

type Schedule struct {
	DoctorLicenseNumber string    `json:"doctorLicenseNumber"`
	MeetingLink         string    `json:"meetingLink"`
	StartTime           time.Time `json:"startTime"`
	EndTime             time.Time `json:"endTime"`
	ReservedBy          string    `json:"reservedBy"`
	ReservedAt          time.Time `json:"reservedAt"`
	CanceledAt          time.Time `json:"canceledAt"`
}

func (s Schedule) GetHashKey() string {
	return s.DoctorLicenseNumber
}

func (s Schedule) GetSortKey() string {
	return s.StartTime.String()
}
