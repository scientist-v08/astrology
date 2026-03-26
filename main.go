package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/scientist-v08/bphs/initializers"
	"github.com/scientist-v08/bphs/routes"
	"github.com/scientist-v08/bphs/validator"
)

func setGinMode() {
	// Set GIN MODE
	if mode := os.Getenv("GIN_MODE"); mode != "" {
		gin.SetMode(mode)
	}
}

func registerValidators() {
	if err := validator.RegisterRaashiValidator(); err != nil {
		log.Fatalf("Failed to register raashi validator: %v", err)
	}
	if err := validator.RegisterGrahaValidator(); err != nil {
		log.Fatalf("Failed to register graha validator: %v", err)
	}
	if err := validator.RegisterNakshatraValidator(); err != nil {
		log.Fatalf("Failed to register nakshatra validator: %v", err)
	}
}

func init() {
	initializers.LoadEnvVariables()
	registerValidators()
	setGinMode()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{allowedOrigin},
		AllowMethods:  []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type", "Authorization"},
    	ExposeHeaders: []string{"Content-Disposition", "Content-Length"},
		AllowCredentials: true,
		MaxAge: 8 * time.Hour,
	}))

	// Now initialize all the routes
	routes.UserRoutes(r)
	routes.RegisterHouseEffectsRoutes(r)
	routes.RegisterShadbalaEffectsRoutes(r)
	routes.RegisterPairingRoutes(r)
	routes.RegisterUpagrahaRoutes(r)
	
	// Now run the application
	r.Run()
}
