package drivers

type Queue interface {
	PushItem(queueName string, item any) error
}

type queue struct {
}

func NewQueue() Queue {
	return &queue{}
}

func (q *queue) PushItem(queueName string, item any) error {
	return nil
}
