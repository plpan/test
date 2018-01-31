package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// user-defined test suite
type UserSuite struct {
	suite.Suite
	AgeInitial int
}

// prehook before any test function will be called
func (u *UserSuite) SetupTest() {
	fmt.Println("setup")
	u.AgeInitial = 2
}

// afterhook after all test functions are called
func (u *UserSuite) TearDownTest() {
	fmt.Println("teardown")
}

// all test functions of UserSuite receiver will be called when running `go test`
func (u *UserSuite) TestUser() {
	fmt.Println("user-defined tests")
	assert.Equal(u.T(), 2, u.AgeInitial)
	u.Equal(2, u.AgeInitial)
}

// when run `go test -testify.m User`, this test function will not be called
func (u *UserSuite) TestAnother() {
	fmt.Println("another user-defined tests")
	assert.Equal(u.T(), 2, u.AgeInitial)
	u.Equal(2, u.AgeInitial)
}

// used by `go test`, otherwise, no test will be run
func TestUserSuite(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
