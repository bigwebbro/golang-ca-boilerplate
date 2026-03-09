package age

import (
	"example.com/pkg/errors"
	"strconv"
)

var (
	MaxAge    uint64 = 200
	ErrMaxAge        = errors.Newf("Age must be less than or equal to %d", MaxAge)
)

type Age uint16

func New(age uint64) (Age, error) {
	a := Age(age)

	if a.Gte(Age(MaxAge)) {
		return Age(0), ErrMaxAge
	}

	return a, nil
}

func (a Age) String() string {
	return strconv.FormatUint(uint64(a), 10)
}

func (a Age) Value() uint16 {
	return uint16(a)
}

func (a Age) Gte(age Age) bool {
	return a >= age
}
