package services

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/tv2169145/golang-http-client-mock/domain/users"
	"testing"
)

func TestNewUserServiceInvalidRequest(t *testing.T) {
	request := users.GetUserRequest{
		Email: "jimmy@gmail.com",
		Password: "1234a",
	}
	user, err := NewUserService().GetUser(request)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestNewUserServiceNotFoundUserAndDecodeError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	request := users.GetUserRequest{
		Email: "jimmy@gmail.com",
		Password: "1234",
	}
	httpmock.RegisterResponder("POST", "http://localhost:8081/users/login",
		httpmock.NewStringResponder(404, `{"message": "not found user", "status": "404", "error":"not_found"}`))
	user, err := NewUserService().GetUser(request)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestUserServiceGetUserNotFoundUser(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	request := users.GetUserRequest{
		Email: "jimmy@gmail.com",
		Password: "1234",
	}
	httpmock.RegisterResponder("POST", "http://localhost:8081/users/login",
		httpmock.NewStringResponder(404, `{"message": "not found user", "status": 404, "error":"not_found"}`))

	user, err := NewUserService().GetUser(request)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, "not found user", err.Error())
}

func TestNewUserServiceGetUserWithDecodeError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	request := users.GetUserRequest{
		Email: "jimmy@gmail.com",
		Password: "1234",
	}

	httpmock.RegisterResponder("POST", "http://localhost:8081/users/login",
		httpmock.NewStringResponder(200, `{"id": "3", "first_name": "lin", "last_name": "jimmy", "email": "jimmy@gmail.com"}`))

	user, err := NewUserService().GetUser(request)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestNewUserServiceGetUserSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	request := users.GetUserRequest{
		Email: "jimmy@gmail.com",
		Password: "1234",
	}

	httpmock.RegisterResponder("POST", "http://localhost:8081/users/login",
		httpmock.NewStringResponder(200, `{"id": 3, "first_name": "lin", "last_name": "jimmy", "email": "jimmy@gmail.com"}`))

	user, err := NewUserService().GetUser(request)
	assert.Nil(t, err)
	assert.NotNil(t, user)
}
