package mymock

import (
	"github.com/stretchr/testify/mock"
)

// ------mock bcrypt----------

type MockBcrypt struct {
	mock.Mock
}

func (b MockBcrypt) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	args := b.Called(password, cost)

	return args.Get(0).([]byte), args.Error(1) // type cast
}
func (b MockBcrypt) CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	args := b.Called(hashedPassword, password)

	return args.Error(0)
}

// ------end mock bcrypt----------
