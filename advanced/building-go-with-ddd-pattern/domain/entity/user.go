package entity

import (
	"strings"

	"github.com/badoux/checkmail"
)

// User information
type User struct {
	BaseEntity
	FirstName string `form:"size:100;not null;"`
	LastName  string `gorm:"size:100;not null;"`
	Email     string `gorm:"size:100;not null;unique;"`
	Password  string `gorm:"size:100;not null;"`
}

// Validate will base on action type to execute specific action
func (user *User) Validate(action string) map[string]string {
	var errMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if user.Email == "" {
			errMessages["email_required"] = "email required"
		}

		if user.Email != "" {
			if err = checkmail.ValidateFormat(user.Email); err != nil {
				errMessages["invalid_email"] = "invalid email"
			}
		}
	case "login":
		if user.Password == "" {
			errMessages["password_required"] = "password is required"
		}

		if user.Email == "" {
			errMessages["email_required"] = "email required"
		}

		if user.Email != "" {
			if err = checkmail.ValidateFormat(user.Email); err != nil {
				errMessages["invalid_email"] = "invalid email"
			}
		}
	case "forgotpassword":
		if user.Email == "" {
			errMessages["email_required"] = "email required"
		}
		if user.Email != "" {
			if err = checkmail.ValidateFormat(user.Email); err != nil {
				errMessages["invalid_email"] = "please provide a valid email"
			}
		}
	default:
		if user.FirstName == "" {
			errMessages["firstname_required"] = "first name is required"
		}

		if user.LastName == "" {
			errMessages["lastname_required"] = "last name is required"
		}

		if user.Password == "" {
			errMessages["password_required"] = "password is required"
		}

		if user.Password != "" && len(user.Password) < 6 {
			errMessages["invalid_password"] = "password should be at least 6 characters"
		}

		if user.Email == "" {
			errMessages["email_required"] = "email is required"
		}

		if user.Email != "" {
			if err = checkmail.ValidateFormat(user.Email); err != nil {
				errMessages["invalid_email"] = "please provide a valid email"
			}
		}
	}

	return errMessages
}
