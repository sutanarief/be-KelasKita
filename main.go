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
	fmt.Println(psqlInfo)
	DB, err = sql.Open("postgres", psqlInfo)
	fmt.Println(DB)

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
	router.Run(":8080")
}
