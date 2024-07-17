package notification

import "fmt"

type NotificationError struct {
	errors []notificationProps
}

func (n NotificationError) Error() string {
	message := ""
	for _, prop := range n.errors {
		message += fmt.Sprintf("%s: %s, ", prop.context, prop.message)
	}
	return message
}
