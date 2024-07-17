package notification

type User struct {
	id           int
	name         string
	email        string
	notification NotificationInterface
}

var (
	idIsRequiredError = notificationProps{
		context: "user",
		message: "id is required",
	}
	nameIsRequiredError = notificationProps{
		context: "user",
		message: "name is required",
	}
	emailIsRequiredError = notificationProps{
		context: "user",
		message: "email is required",
	}
)

func (u *User) Validate() {
	if u.id <= 0 {
		u.notification.AddError(idIsRequiredError)
	}
	if u.name == "" {
		u.notification.AddError(nameIsRequiredError)
	}
	if u.email == "" {
		u.notification.AddError(emailIsRequiredError)
	}
	return
}

func NewUser(id int, name, email string) (*User, error) {
	notification := Notification{}
	user := User{
		id:           id,
		name:         name,
		email:        email,
		notification: &notification,
	}
	user.Validate()

	if user.notification.HasError() {
		return nil, NotificationError{
			errors: user.notification.GetError(),
		}
	}

	return &user, nil
}
