package middleware

import (
	"be-kelaskita/auth"
	"be-kelaskita/repository"
	"be-kelaskita/service"
	"database/sql"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = os.Getenv("JWTKEY")

func Auth(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"unauthorized": "Access token is missing",
			})
			c.Abort()
			return
		}

		err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		message, ok := acccessValidator(c, tokenString, db)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"unauthorized": message,
			})
			c.Abort()
		}
		c.Next()
	}
}

func acccessValidator(c *gin.Context, tokenString string, db *sql.DB) (string, bool) {
	method := c.Request.Method
	path := c.Request.RequestURI

	credential, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		panic(err)
	}

	data := credential.Claims.(jwt.MapClaims)
	role := data["user_role"].(string)

	if strings.Contains(string(path), "class") {
		if method == "POST" || method == "DELETE" || method == "PUT" {
			if role != "Teacher" {
				return "Student cannot create, delete, or edit a class", false
			}
		}
	}

	if strings.Contains(string(path), "subject") {
		if method == "POST" || method == "DELETE" || method == "PUT" {
			if role != "Teacher" {
				return "Student cannot create, delete, or edit a subject", false
			}
		}
	}

	if strings.Contains(string(path), "question") {
		if method == "PUT" || method == "DELETE" {
			if role != "Teacher" {
				userId := strconv.FormatFloat(data["id"].(float64), 'f', 0, 64)
				questionId := c.Param("id")
				id, err := strconv.Atoi(questionId)
				if err != nil {
					panic(err)
				}
				questionRepository := repository.NewQuestionRepository(db)
				questionService := service.NewQuestionService(questionRepository)
				question, err := questionService.GetQuestionById(id)
				questionUserId := strconv.Itoa(question.User_id)

				if userId != questionUserId {
					return "As a Student you cannot delete or edit another user's question", false
				}
			}
		}
	}

	if strings.Contains(string(path), "answer") {
		if method == "PUT" || method == "DELETE" {
			if role != "Teacher" {
				userId := strconv.FormatFloat(data["id"].(float64), 'f', 0, 64)
				questionId := c.Param("id")
				id, err := strconv.Atoi(questionId)
				if err != nil {
					panic(err)
				}
				answerRepository := repository.NewAnswerRepository(db)
				answerService := service.NewAnswerService(answerRepository)
				answer, err := answerService.GetAnswerById(id)
				answerUserId := strconv.Itoa(answer.User_id)

				if userId != answerUserId {
					return "As a Student you cannot delete or edit another user's answer", false
				}
			}
		}
	}

	return "", true
}
