package controller

import (
	"be-kelaskita/entity"
	"be-kelaskita/helper"
	"be-kelaskita/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type subjectHandler struct {
	subjectService service.SubjectService
}

func NewSubjectHandler(subjectService service.SubjectService) *subjectHandler {
	return &subjectHandler{subjectService}
}

func (s *subjectHandler) GetSubject(c *gin.Context) {
	var result gin.H

	subjects, err := s.subjectService.GetSubject()
	if err != nil {
		result = gin.H{
			"message": err,
		}
	} else {
		result = gin.H{
			"result": subjects,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (s *subjectHandler) InsertSubject(c *gin.Context) {
	var inputSubject entity.Subject

	err := c.ShouldBindJSON(&inputSubject)
	if err != nil {
		panic(err)
	}

	newSubject, err := s.subjectService.InsertSubject(inputSubject)

	if err != nil {
		helper.ErrorHandler(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success Insert Subject",
		"result":  newSubject,
	})
}

func (s *subjectHandler) UpdateSubject(c *gin.Context) {
	var inputSubject entity.Subject
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	err = c.ShouldBindJSON(&inputSubject)

	if err != nil {
		panic(err)
	}

	subject, err := s.subjectService.UpdateSubject(inputSubject, id)

	if err != nil {
		helper.ErrorHandler(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Update Subject",
		"result":  subject,
	})
}

func (s *subjectHandler) DeleteSubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	err = s.subjectService.DeleteSubject(id)
	if err != nil {
		helper.ErrorHandler(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Subject",
	})
}

func (s *subjectHandler) GetQuestionBySubjectId(c *gin.Context) {
	var result gin.H
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	questions, err := s.subjectService.GetQuestionBySubjectId(id)
	if err != nil {
		panic(err)
	} else {
		result = gin.H{
			"result": questions,
		}
	}

	c.JSON(http.StatusOK, result)
}
