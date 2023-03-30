package controller

import (
	"be-kelaskita/entity"
	"be-kelaskita/helper"
	"be-kelaskita/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type questionHandler struct {
	questionService service.QuestionService
}

func NewQuestionHandler(questionService service.QuestionService) *questionHandler {
	return &questionHandler{questionService}
}

func (q *questionHandler) GetQuestion(c *gin.Context) {
	var result gin.H

	questions, err := q.questionService.GetQuestion()
	if err != nil {
		result = gin.H{
			"message": err,
		}
	} else {
		result = gin.H{
			"result": questions,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (q *questionHandler) InsertQuestion(c *gin.Context) {
	var inputQuestion entity.Question

	err := c.ShouldBindJSON(&inputQuestion)
	if err != nil {
		panic(err)
	}

	newQuestions, err := q.questionService.InsertQuestion(inputQuestion)

	if err != nil {
		helper.ErrorHandler(err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Success Insert Question",
		"result":  newQuestions,
	})
}

func (q *questionHandler) UpdateQuestion(c *gin.Context) {
	var inputQuestion entity.Question
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		panic(err)
	}

	err = c.ShouldBindJSON(&inputQuestion)

	if err != nil {
		panic(err)
	}

	question, err := q.questionService.UpdateQuestion(inputQuestion, id)

	if err != nil {
		helper.ErrorHandler(err, c)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Update Question",
		"result":  question,
	})
}

func (q *questionHandler) DeleteQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	err = q.questionService.DeleteQuestion(id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Question",
	})
}

func (q *questionHandler) GetQuestionById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	question, err := q.questionService.GetQuestionById(id)
	if err != nil {
		helper.ErrorHandler(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": question,
	})
}

func (q *questionHandler) GetQuestionWithAnswer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}

	question, err := q.questionService.GetQuestionWithAnswer(id)
	if err != nil {
		helper.ErrorHandler(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": question,
	})
}
