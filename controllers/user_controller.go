package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tv2169145/golang-http-client-mock/domain/users"
	"github.com/tv2169145/golang-http-client-mock/services"
	"net/http"
)

func GetUser(c *gin.Context) {
	var request users.GetUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.String(http.StatusBadRequest, "invalid json request")
		return
	}
	user, err := services.NewUserService().GetUser(request)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
