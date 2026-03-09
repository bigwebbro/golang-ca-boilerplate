package patronymic

import "example.com/pkg/errors"

var (
	MaxLength      = 100
	ErrWrongLength = errors.Newf("patronymic must be less than or equal to %d characters", MaxLength)
)

type Patronymic string

func (p Patronymic) Value() string {
	return string(p)
}

func New(patronymic string) (Patronymic, error) {
	if len([]rune(patronymic)) > MaxLength {
		return "", ErrWrongLength
	}
	p := Patronymic(patronymic)
	return p, nil
}
