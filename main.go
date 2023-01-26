package main

import (
	"gin-gonic-api/app/router"
	"gin-gonic-api/config"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	godotenv.Load()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
