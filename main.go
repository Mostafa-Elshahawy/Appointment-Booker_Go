package main

import (
	"log"
	"mostafaba29/intialization"
	"mostafaba29/middleware"
	"mostafaba29/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func init() {
	intialization.ConnectToDB()
	intialization.MigrateDB()
}

func main() {
	app := fiber.New()
	middleware.Store = session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSameSite: fiber.CookieSameSiteLaxMode,
		Expiration:     time.Hour * 24,
	})
	app.Use(middleware.NewMiddleware(), cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Acess-Control-Allow-Orgin, Content-type, Origin, Accept",
	}))

	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
