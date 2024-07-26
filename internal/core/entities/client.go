package entities

type Client struct {
	FirstName      string
	LastName       string
	Email          string
	PhoneNumber    string
	Address        string
	DateOfBirth    string
	Gender         string
	MedicalReports []MedicalReport
}

type MedicalReport struct {
	Client       string
	SharedWith   []string
	FileName     string
	FileLocation string
}

func (c Client) GetHashKey() string {
	return c.Email
}

func (c Client) GetSortKey() string {
	return ""
}
