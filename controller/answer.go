package controller

import (
	"be-kelaskita/entity"
	"be-kelaskita/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type answerHandler struct {
	answerService service.AnswerService
}

func NewAnswerHandler(answerService service.AnswerService) *answerHandler {
	return &answerHandler{answerService}
}

func (a *answerHandler) GetAnswer(c *gin.Context) {
	var result gin.H

	answer, err := a.answerService.GetAnswer()
	if err != nil {
		result = gin.H{
			"message": err,
		}
	} else {
		result = gin.H{
			"result": answer,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (a *answerHandler) InsertAnswer(c *gin.Context) {
	var inputAnswer entity.Answer

	err := c.ShouldBindJSON(&inputAnswer)
	if err != nil {
		panic(err)
	}

	newAnswer, err := a.answerService.InsertAnswer(inputAnswer)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success Insert Answer",
		"result":  newAnswer,
	})
}

func (a *answerHandler) UpdateAnswer(c *gin.Context) {
	var inputAnswer entity.Answer
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	err = c.ShouldBindJSON(&inputAnswer)

	if err != nil {
		panic(err)
	}

	answer, err := a.answerService.UpdateAnswer(inputAnswer, id)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Update Answer",
		"result":  answer,
	})
}

func (a *answerHandler) DeleteAnswer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	err = a.answerService.DeleteAnswer(id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Answer",
	})
}

func (a *answerHandler) GetAnswerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	answer, err := a.answerService.GetAnswerById(id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": answer,
	})
}
