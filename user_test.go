package notification

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

var (
	shouldHaveReturnedNilError = errors.New("should return nil when an error is returned")
	name                       = "John Doe"
	email                      = "john.doe@example.com"
)

func TestShouldReturnErrorWhenIdIsInvalid(t *testing.T) {
	user, err := NewUser(0, name, email)

	if err == nil {
		log.Println(err)
		t.Error("must return an error when an id is lesser or equal to zero")
	}
	if user != nil {
		fmt.Println(user)
		t.Error(shouldHaveReturnedNilError)
	}
	log.Println(err)
	log.Println(user)
}
func TestShouldReturnErrorWhenEmailIsInvalid(t *testing.T) {
	user, err := NewUser(25, name, "")
	if err == nil {
		log.Println(err)
		t.Error("must return an error when an email is not valid")
	}
	if user != nil {
		t.Error(shouldHaveReturnedNilError)
	}
}
func TestShouldReturnErrorWhenNameIsInvalid(t *testing.T) {
	user, err := NewUser(25, "", email)
	if err == nil {
		log.Println(err)
		t.Error("must return an error when a name is not valid")
	}
	if user != nil {
		t.Error(shouldHaveReturnedNilError)
	}
}
func TestShouldReturnAllErrors(t *testing.T) {
	notification := Notification{}
	user := new(User)
	user.id = -1
	user.name = ""
	user.email = ""
	user.notification = &notification
	user.Validate()

	_, err := NewUser(-1, "", "")
	if err == nil {
		t.Error(shouldHaveReturnedNilError)
	}

	allErrors := NotificationError{
		errors: user.notification.GetError(),
	}

	if !errors.As(err, &allErrors) {
		log.Println(err)
		log.Println(allErrors)
		t.Error("must return all all errors")
	}
}
