package main

import (
	"be-kelaskita/controller"
	"be-kelaskita/database"
	"be-kelaskita/middleware"
	"be-kelaskita/repository"
	"be-kelaskita/service"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

var router = gin.Default()

func main() {
	err = godotenv.Load("config/.env")

	if err != nil {
		log.Fatalf("Error load env")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)
	DB, err = sql.Open("postgres", psqlInfo)

	err = DB.Ping()

	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	authorized := router.Group("/")
	authorized.Use(middleware.Auth(DB))
	// user Route
	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepository)
	userHandler := controller.NewUserHandler(userService)
	authorized.GET("/user", userHandler.GetUser)
	authorized.PUT("/user/:id", userHandler.UpdateUser)
	authorized.DELETE("/user/:id", userHandler.DeleteUser)
	authorized.GET("/user/:id", userHandler.GetUserById)
	router.POST("/user/register", userHandler.InsertUser)
	router.POST("/user/login", userHandler.UserLogin)

	// class Route
	classRepository := repository.NewClassRepository(DB)
	classService := service.NewClassService(classRepository)
	classHandler := controller.NewClassHandler(classService)
	router.GET("/class", classHandler.GetClass)
	authorized.POST("/class", classHandler.InsertClass)
	authorized.PUT("/class/:id", classHandler.UpdateClass)
	authorized.DELETE("/class/:id", classHandler.DeleteClass)
	authorized.GET("/class/:id/users", classHandler.GetUserByClassId)

	// subject Route
	subjectRepository := repository.NewSubjectRepository(DB)
	subjectService := service.NewSubjectService(subjectRepository)
	subjectHandler := controller.NewSubjectHandler(subjectService)
	router.GET("/subject", subjectHandler.GetSubject)
	authorized.POST("/subject", subjectHandler.InsertSubject)
	authorized.PUT("/subject/:id", subjectHandler.UpdateSubject)
	authorized.DELETE("/subject/:id", subjectHandler.DeleteSubject)

	// question Route
	questionRepository := repository.NewQuestionRepository(DB)
	questionService := service.NewQuestionService(questionRepository)
	questionHandler := controller.NewQuestionHandler(questionService)
	router.GET("/question", questionHandler.GetQuestion)
	authorized.POST("/question", questionHandler.InsertQuestion)
	authorized.PUT("/question/:id", questionHandler.UpdateQuestion)
	authorized.DELETE("/question/:id", questionHandler.DeleteQuestion)
	router.GET("/question/:id", questionHandler.GetQuestionById)

	// answer Route
	answerRepository := repository.NewAnswerRepository(DB)
	answerService := service.NewAnswerService(answerRepository)
	answerHandler := controller.NewAnswerHandler(answerService)
	router.GET("/answer", answerHandler.GetAnswer)
	authorized.POST("/answer", answerHandler.InsertAnswer)
	authorized.PUT("/answer/:id", answerHandler.UpdateAnswer)
	authorized.DELETE("/answer/:id", answerHandler.DeleteAnswer)
	router.GET("/answer/:id", answerHandler.GetAnswerById)

	router.Run(":8080")
}
