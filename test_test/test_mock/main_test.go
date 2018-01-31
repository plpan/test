package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	mock.Mock
}

func (m *UserMock) Up() int {
	args := m.Called()
	return args.Int(0)
}

func TestUp(t *testing.T) {
	// mock means that the real method will not be called, instead of calling a fake
	m := &UserMock{}
	m.On("Up").Return(2)

	r := Up(m)
	// parameters and return values should be the same, otherwise it'll fail
	m.AssertExpectations(t)

	assert := assert.New(t)
	// the return value should be equal with Return
	assert.Equal(r, 2)
}
