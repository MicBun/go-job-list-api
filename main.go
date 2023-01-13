package main

import (
	"github.com/MicBun/go-job-list-api/config"
	"github.com/MicBun/go-job-list-api/docs"
	"github.com/MicBun/go-job-list-api/route"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default env")
	}

	description := "This is a sample server Job List server.\n\n" +
		"To get Bearer Token, first you need to login.\n\n" +
		"Checkout my Github: https://github.com/MicBun\n\n" +
		"Checkout my Linkedin: https://www.linkedin.com/in/MicBun\n\n"

	docs.SwaggerInfo.Title = "Job List API"
	docs.SwaggerInfo.Description = description

	database := config.ConnectDataBase()
	sqlDB, _ := database.DB()
	defer sqlDB.Close()
	r := route.SetupRouter(database)
	r.Run()
}
