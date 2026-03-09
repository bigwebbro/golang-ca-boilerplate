package email

import (
	"example.com/pkg/errors"
)

var (
	MaxLength = 250
	ErrLength = errors.Newf("Email must be less than or equal to %d characters", MaxLength)
)

type Email struct {
	value string
}

func New(name string) (Email, error) {
	if len([]rune(name)) >= MaxLength {
		return Email{}, ErrLength
	}
	return Email{value: name}, nil
}

func (e Email) Value() string {
	return e.value
}
