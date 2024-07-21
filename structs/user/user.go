package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (user *User) OutputUserData() {
	fmt.Println("User's first name is: ", user.firstName)
	fmt.Println("User's last name is: ", user.lastName)
	fmt.Println("User's birthdate is: ", user.birthDate)
	fmt.Println("User's account created at: ", user.createdAt)
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func New(userFirstName, userLastName, userBirthdate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("First name, last name and birthdate are required")
	}
	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}
