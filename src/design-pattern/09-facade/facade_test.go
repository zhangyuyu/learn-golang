package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Login(t *testing.T) {
	service := UserService{}
	user, err := service.Login(13112341234, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}

func TestUserService_Register(t *testing.T) {
	service := UserService{}
	user, err := service.Register(13112341234, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test register"}, user)
}

func TestUserService_LoginOrRegister(t *testing.T) {
	service := UserService{}
	user, err := service.LoginOrRegister(13112341234, 1234)
	assert.NoError(t, err)
	assert.Equal(t, &User{Name: "test login"}, user)
}
