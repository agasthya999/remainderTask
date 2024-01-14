package main

import (
	"log"
	"remainderTask/migrations"
	"remainderTask/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func main() {
	migrations.Migrate()
	app := fiber.New()
	app.Use(cors.New())
	app.Use(pprof.New())
	router.SetupRoutes(app)
	// cert, _ := tls.LoadX509KeyPair("/app/mysportsfeed.io.pem", "/app/mysportsfeed.io.key")
	// config := tls.Config{
	// 	Certificates:       []tls.Certificate{cert},
	// 	InsecureSkipVerify: true,
	// }
	log.Fatal(app.Listen(":80"))
}
