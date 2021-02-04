package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bagus2x/fainmi-be/api/routes"
	"github.com/bagus2x/fainmi-be/config"
	"github.com/bagus2x/fainmi-be/pkg/profile"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Environment variable
	port := os.Getenv("PORT")
	databaseURI := os.Getenv("DATABASE_URI")
	accessTokenKey := os.Getenv("ACCESS_TOKEN_KEY")

	// Init router
	app := fiber.New()

	// Global Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // di ganti besok2
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Database connection
	database, err := config.DatabaseConnection(databaseURI)
	if err != nil {
		panic(err)
	}

	// profile layers
	profileRepository := profile.NewRepo(database)
	profileService := profile.NewService(profileRepository, accessTokenKey)
	routes.Profile(app, profileService)

	//

	//

	//

	//

	//

	//

	//

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	go func() {
		log.Printf("%s: Gracefully shutting down\n", <-c)
		app.Shutdown()
	}()

	log.Println("Server started on PORT", port)
	err = app.Listen(":" + port)
	log.Fatalf("err: %v", err)
}
