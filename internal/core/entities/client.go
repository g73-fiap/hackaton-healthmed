package entities

type Client struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Address     string
	DateOfBirth string
	Gender      string
}

func (c Client) GetHashKey() string {
	return c.Email
}

func (c Client) GetSortKey() string {
	return ""
}
