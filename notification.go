package notification

type notificationProps struct {
	message string
	context string
}
type NotificationInterface interface {
	AddError(err notificationProps)
	HasError() bool
	GetError() []notificationProps
}

type Notification struct {
	errors []notificationProps
}

func (n *Notification) AddError(err notificationProps) {
	n.errors = append(n.errors, err)
}

func (n *Notification) HasError() bool {
	return len(n.errors) > 0
}

func (n *Notification) GetError() []notificationProps {
	return n.errors
}
