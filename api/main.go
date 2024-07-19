package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/tiandingyi/shorten-url-fiber-redis/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Printf("err")
	}

	app := fiber.New()

	// 将日志中间件添加到 Fiber 应用程序中。
	// 这行代码的作用是告诉 Fiber 在处理每个请求之前都要使用这个日志中间件，以记录请求的详细信息（如请求方法、URL、响应状态码等）。
	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
