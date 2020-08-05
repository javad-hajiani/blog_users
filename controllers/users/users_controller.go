package users

import (
	"github.com/gin-gonic/gin"
	"github.com/javad-hajiani/blog_users/domain/users"
	"github.com/javad-hajiani/blog_users/services"
	"github.com/javad-hajiani/blog_users/utils/errors"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, userError := strconv.ParseInt(c.Param("user_id"),10,64)
	if userError != nil {
		err := errors.NewBadRequestError("user id should be integer")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}
	result, NewErr := services.CreateUser(user)
	if NewErr != nil {
		c.JSON(NewErr.Status, NewErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}