package slicer

import (
	"errors"
	"strings"
)

type UserAccount struct {
	Username string
	Domain   string
}

var ErrMessageInvalidEmail = errors.New("invalid email")

func EmailSlicer(email string) (UserAccount, error) {
	components := strings.Split(email, "@")
	if len(components) != 2 || components[1] == "" {
		return UserAccount{}, ErrMessageInvalidEmail
	}
	username, domain := components[0], components[1]
	userData := UserAccount{
		Username: username,
		Domain:   domain,
	}

	return userData, nil
}
