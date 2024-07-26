package drivers

type GoogleAPI interface {
	CreateMeeting() (string, error)
}

type googleAPI struct {
}

func NewGoogleAPI() GoogleAPI {
	return googleAPI{}
}

func (g googleAPI) CreateMeeting() (string, error) {
	return "meet.google/aaa-ccc", nil
}
