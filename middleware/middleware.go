package middleware

import (
	"be-kelaskita/auth"
	"be-kelaskita/repository"
	"be-kelaskita/service"
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

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
		return []byte("secret"), nil
	})

	if err != nil {
		panic(err)
	}

	data := credential.Claims.(jwt.MapClaims)
	role := data["user_role"].(string)

	if strings.Contains(string(path), "class") {
		if method == "POST" || method == "DELETE" {
			if role != "Teacher" {
				return "Student cannot create or delete a class", false
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

	return "", true
}
