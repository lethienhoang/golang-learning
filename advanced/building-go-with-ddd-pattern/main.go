package main

import (
	"building-go-with-ddd-pattern/middlewares"
	"fmt"
	"log"
	"net/http"
	"os"

	"building-go-with-ddd-pattern/application/controllers"
	"building-go-with-ddd-pattern/infrastructure"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func HealthGET(r *gin.Context) {
	r.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

func main() {
	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dbContext, err := infrastructure.NewDbContext(dbdriver, user, password, port, host, dbname)

	db := dbContext.GetDbContext()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	defer dbContext.Close()
	dbContext.Automigrate()

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	//health check
	r.GET("/healthcheck", HealthGET)

	// food routes
	foodController := controllers.NewFoodController(db)
	foodController.FoodControllerRoute(r)

	//Starting the application
	appport := os.Getenv("PORT")
	if appport == "" {
		appport = "8888" //localhost
	}

	log.Fatal(r.Run(":" + appport))
}
