package controller

import (
	"be-kelaskita/entity"
	"be-kelaskita/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type classHandler struct {
	classService service.ClassService
}

func NewClassHandler(classService service.ClassService) *classHandler {
	return &classHandler{classService}
}

func (cl *classHandler) GetClass(c *gin.Context) {
	var result gin.H

	classes, err := cl.classService.GetClass()
	if err != nil {
		result = gin.H{
			"message": err,
		}
	} else {
		result = gin.H{
			"result": classes,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (cl *classHandler) InsertClass(c *gin.Context) {
	var inputClass entity.Class

	err := c.ShouldBindJSON(&inputClass)
	if err != nil {
		panic(err)
	}

	newClass, err := cl.classService.InsertClass(inputClass)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success Insert Class",
		"result":  newClass,
	})
}

func (cl *classHandler) GetUserByClassId(c *gin.Context) {
	var result gin.H
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	users, err := cl.classService.GetUserByClassId(id)
	if err != nil {
		panic(err)
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}
