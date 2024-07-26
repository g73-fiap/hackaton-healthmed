package entities

type Appoitment struct {
	ClientEmail         string
	DoctorLicenseNumber string
	Confirmed           bool
	Canceled            bool
	Schedule            Schedule
}

func (a Appoitment) GetHashKey() string {
	return a.ClientEmail
}

func (a Appoitment) GetSortKey() string {
	return a.Schedule.StartTime.String()
}
