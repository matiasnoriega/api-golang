package main

import (
	"api-golang/internal/controller"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Loads Dotenv library
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf("Error loading dotenv: %s", err))
	}

	r := gin.Default()

	// Defines a Group of routes for V1 versioning
	group := r.Group("/v1")
	{
		group.GET("/insulina/:value", controller.AddRegistry)
	}

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

}
