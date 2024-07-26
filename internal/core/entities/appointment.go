package entities

type Appointment struct {
	ClientEmail         string   `json:"clientEmail"`
	DoctorLicenseNumber string   `json:"doctorLicenseNumber"`
	Confirmed           bool     `json:"confirmed"`
	Canceled            bool     `json:"canceled"`
	Schedule            Schedule `json:"schedule"`
}

func (a Appointment) GetHashKey() string {
	return a.ClientEmail
}

func (a Appointment) GetSortKey() string {
	return a.Schedule.StartTime.String()
}
