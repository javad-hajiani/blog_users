package users

import (
	"github.com/javad-hajiani/blog_users/utils/errors"
	"strings"
)

type User struct {
	Id int64
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestError {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	if user.Email == "" {
		return errors.NewBadRequestError("Email should not be empty")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("Password should not be empty")
	}
	return nil
}