package entities

type Client struct {
	FirstName      string          `json:"firstName"`
	LastName       string          `json:"lastName"`
	Email          string          `json:"email"`
	PhoneNumber    string          `json:"phoneNumber"`
	Address        string          `json:"address"`
	DateOfBirth    string          `json:"dateOfBirth"`
	Gender         string          `json:"gender"`
	MedicalReports []MedicalReport `json:"medicalReports"`
}

type MedicalReport struct {
	Client       string   `json:"client"`
	SharedWith   []string `json:"sharedWith"`
	FileName     string   `json:"fileName"`
	FileLocation string   `json:"fileLocation"`
}

func (c Client) GetHashKey() string {
	return c.Email
}

func (c Client) GetSortKey() string {
	return ""
}
