package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bagus2x/fainmi-be/api/middleware"
	"github.com/bagus2x/fainmi-be/api/routes"
	"github.com/bagus2x/fainmi-be/config"
	"github.com/bagus2x/fainmi-be/pkg/background"
	"github.com/bagus2x/fainmi-be/pkg/button"
	"github.com/bagus2x/fainmi-be/pkg/font"
	"github.com/bagus2x/fainmi-be/pkg/like"
	"github.com/bagus2x/fainmi-be/pkg/link"
	"github.com/bagus2x/fainmi-be/pkg/profile"
	"github.com/bagus2x/fainmi-be/pkg/style"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	// Load Environment variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

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
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Database connection
	database, err := config.DatabaseConnection(databaseURI)
	if err != nil {
		log.Fatal(err)
	}

	// Repository
	profileRepo := profile.NewRepo(database)
	styleRepo := style.NewRepo(database)
	linkRepo := link.NewRepo(database)
	likeRepo := like.NewRepo(database)
	backgroundRepo := background.NewRepo(database)
	buttonRepo := button.NewRepo(database)
	fontRepo := font.NewRepo(database)

	// Service
	profileService := profile.NewService(profileRepo, accessTokenKey)
	styleService := style.NewService(styleRepo)
	linkService := link.NewService(linkRepo)
	likeService := like.NewService(likeRepo)
	backgroundService := background.NewService(backgroundRepo)
	buttonService := button.NewService(buttonRepo)
	fontService := font.NewService(fontRepo)

	// Middleware
	auth := middleware.NewAuth(profileService)

	app.Static("/public", "./public/", fiber.Static{
		Compress:  true,
		ByteRange: true,
	})

	// Routes/Controller
	routes.Test(app, auth)
	routes.Profile(app, profileService, auth)
	routes.Style(app, styleService, auth)
	routes.Link(app, linkService, auth)
	routes.Like(app, likeService, auth)
	routes.Background(app, backgroundService)
	routes.Button(app, buttonService)
	routes.Font(app, fontService)

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
