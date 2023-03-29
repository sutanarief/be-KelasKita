package main

import (
	"be-kelaskita/controller"
	"be-kelaskita/database"
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

	// user Route
	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepository)
	userHandler := controller.NewUserHandler(userService)
	router.GET("/users", userHandler.GetUser)
	router.POST("/users", userHandler.InsertUser)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)
	router.GET("/user/:id", userHandler.GetUserById)

	// class Route
	classRepository := repository.NewClassRepository(DB)
	classService := service.NewClassService(classRepository)
	classHandler := controller.NewClassHandler(classService)
	router.GET("/classes", classHandler.GetClass)
	router.POST("/classes", classHandler.InsertClass)
	router.PUT("/classes/:id", classHandler.UpdateClass)
	router.DELETE("/classes/:id", classHandler.DeleteClass)
	router.GET("/classes/:id/users", classHandler.GetUserByClassId)

	// subject Route
	subjectRepository := repository.NewSubjectRepository(DB)
	subjectService := service.NewSubjectService(subjectRepository)
	subjectHandler := controller.NewSubjectHandler(subjectService)
	router.GET("/subjects", subjectHandler.GetSubject)
	router.POST("/subjects", subjectHandler.InsertSubject)
	router.PUT("/subjects/:id", subjectHandler.UpdateSubject)
	router.DELETE("/subjects/:id", subjectHandler.DeleteSubject)

	// question Route
	questionRepository := repository.NewQuestionRepository(DB)
	questionService := service.NewQuestionService(questionRepository)
	questionHandler := controller.NewQuestionHandler(questionService)
	router.GET("/questions", questionHandler.GetQuestion)
	router.POST("/questions", questionHandler.InsertQuestion)
	router.PUT("/questions/:id", questionHandler.UpdateQuestion)
	router.DELETE("/questions/:id", questionHandler.DeleteQuestion)

	router.Run(":8080")
}
