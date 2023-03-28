package controller

import (
	"be-kelaskita/entity"
	"be-kelaskita/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{userService}
}

func (u *userHandler) GetUser(c *gin.Context) {
	var result gin.H

	users, err := u.userService.GetUser()
	if err != nil {
		result = gin.H{
			"message": err,
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (u *userHandler) InsertUser(c *gin.Context) {
	var inputUser entity.User

	err := c.ShouldBindJSON(&inputUser)
	if err != nil {
		panic(err)
	}

	newUser, err := u.userService.InsertUser(inputUser)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success Insert User",
		"result":  newUser,
	})
}
