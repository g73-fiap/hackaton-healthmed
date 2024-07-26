package drivers

type SNS interface {
	NotifyItem(topic string, item any) error
}

type sns struct {
}

func NewSNS() SNS {
	return &sns{}
}

func (s *sns) NotifyItem(topic string, item any) error {
	return nil
}
