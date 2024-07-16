package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	//Fields need to be uppercase to export to other packages
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// embedded struct similar to  inheritance
type Admin struct {
	email    string
	password string
	User
}

// receiver methord with a pointer of a struct
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.birthdate)
}

// mutates data in struct, use pointer to update values otherwise in this example the user data won't be updated if not targeting pointer
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "Admin",
			birthdate: "----",
			createdAt: time.Now(),
		},
	}
}
func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("please enter required info")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
