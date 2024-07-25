package entities

import "time"

type Doctor struct {
	ID                string    `json:"id"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Specialization    string    `json:"specialization"`
	Email             string    `json:"email"`
	PhoneNumber       string    `json:"phoneNumber"`
	Address           string    `json:"address"`
	Coordinates       [2]int    `json:"coordinates"`
	DateOfBirth       time.Time `json:"dateOfBirth"`
	Gender            string    `json:"gender"`
	LicenseNumber     string    `json:"licenseNumber"`
	YearsOfExperience int       `json:"yearsOfExperience"`
	Available         bool      `json:"available"`
}

func (d Doctor) GetID() string {
	return d.ID
}

func (d Doctor) SetID(id string) {
	d.ID = id
}
