package utils

import (
	"errors"
	"regexp"
	"telegraph/model"
)

func ValidateNewUser(user *model.User) error {
	rxpUsername := regexp.MustCompile("^[a-z0-9_]{3,16}$")
	if !rxpUsername.MatchString(user.Username) {
		return errors.New("Invalid username format.")
	}

	rxpPhoneNumber := regexp.MustCompile("^[0-9]{11}$")
	if !rxpPhoneNumber.MatchString(user.PhoneNumber) {
		return errors.New("Invalid phone number format.")
	}

	return nil
}
