package main

import (
	"fiber-shop/internal/config"
	"fiber-shop/internal/routes"
	"fiber-shop/pkg/logger"
	"fiber-shop/pkg/utils"
	"log"

	_ "fiber-shop/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Fiber Shop API
// @version 1.0
// @description This is a simple (yet) Shop api build using Fiber and Go.
func main() {

	log.Println("Initializing config")

	config := config.GetConfig()

	log.Println("Config initialized")

	log.Println("Initializing logger")

	l := logger.NewLogger("go-fiber-shop: ")

	l.Println("Logger initialized")

	l.Println("Defining fiber app")

	a := fiber.New()

	l.Println("App defined")

	l.Println("Middleware initializing")

	a.Use(cors.New())
	//TODO middlewares

	l.Println("Middleware initialized")

	l.Println("Routes initializing")

	routes.Swagger(a)
	routes.Metrics(a)

	// TODO routes
	l.Println("Routes initialized")

	l.Println("Starting server")

	url := config.Listen.BindIP+":"+config.Listen.Port

	utils.StartServerWithGracefulShutdown(a, url, l)

}