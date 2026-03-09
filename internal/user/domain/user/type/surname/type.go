package surname

import "example.com/pkg/errors"

var (
	MaxLength      = 100
	ErrWrongLength = errors.Newf("surname must be less than or equal to %d characters", MaxLength)
)

type Surname string

func (s Surname) Value() string {
	return string(s)
}

func New(surname string) (Surname, error) {
	if len([]rune(surname)) > MaxLength {
		return "", ErrWrongLength
	}
	s := Surname(surname)

	return s, nil
}
