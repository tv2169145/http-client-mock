package services

import (
	"encoding/json"
	"errors"
	"github.com/tv2169145/golang-http-client-mock/domain/users"
	"github.com/tv2169145/golang-http-client-mock/restclient"
	"net/http"
)

type Service interface {
	GetUser(users.GetUserRequest) (*users.User, error)
}

type userService struct {

}

func NewUserService() Service {
	return &userService{}
}

func (s *userService) GetUser(request users.GetUserRequest) (*users.User, error) {
	header := http.Header{}

	res, err := restclient.Post("http://localhost:8081/users/login", request, header)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)
	if res.StatusCode > 299 {
		var getUserErrorResponse users.GetUserErrorResponse
		if err := decoder.Decode(&getUserErrorResponse); err != nil {
			return nil, err
		}
		return nil, errors.New(getUserErrorResponse.Message)
	}

	var getUser users.User
	if err := decoder.Decode(&getUser); err != nil {
		return nil, err
	}

	return &getUser, nil
}
