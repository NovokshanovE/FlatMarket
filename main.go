package main

import (
	"log"

	"github.com/NovokshanovE/flatmarket/database"
	"github.com/NovokshanovE/flatmarket/handlers"
	"github.com/NovokshanovE/flatmarket/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()

	database.ConnectDatabase()

	r.POST("/dummyLogin", handlers.DummyLogin)
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	house := r.Group("/house")
	{
		house.POST("/create", middleware.AuthMiddleware("moderator"), handlers.CreateHouse)
		house.GET("/:id", middleware.AuthMiddleware("client", "moderator"), handlers.GetFlats)
		house.POST("/:id/subscribe", middleware.AuthMiddleware("client"), handlers.SubscribeHouse)
	}

	flat := r.Group("/flat")
	{
		flat.POST("/create", middleware.AuthMiddleware("client", "moderator"), handlers.CreateFlat)
		flat.POST("/update", middleware.AuthMiddleware("moderator"), handlers.UpdateFlat)
	}

	r.Run(":8080")
}
