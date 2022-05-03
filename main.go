package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/dheroefic/gin-boilerplate/database"
	"github.com/dheroefic/gin-boilerplate/routes"
	"github.com/dheroefic/gin-boilerplate/utils/helpers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Init the env
	err := godotenv.Load()
	if err != nil {
		helpers.Logger("MAIN PROCESS", "Error loading .env file", true)
	}

	// Set gin mode
	gin.SetMode(os.Getenv("APP_ENVIRONMENT"))
	helpers.Logger("MAIN PROCESS", fmt.Sprintf("Server is initiated in %v mode", os.Getenv("APP_ENVIRONMENT")), false)

	// Init app
	app := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	app.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	app.Use(gin.Recovery())

	// Init DB
	// Load Database
	database.InitConnection()

	serverAddr := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	// Set server configuration
	server := &http.Server{
		Addr:    serverAddr,
		Handler: app,
	}

	helpers.Logger("MAIN PROCESS", fmt.Sprintf("Server URL: %s Port: %s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT")), false)

	// Load all routes
	routes.Load(app)

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		isTLS := os.Getenv("APP_TLS")
		if strings.ToLower(isTLS) == "true" {
			helpers.Logger("MAIN PROCESS", "Server running in TLS mode", false)

			appCertLocEnv := os.Getenv("APP_TLS_CERT_LOCATION")
			appKeyLocEnv := os.Getenv("APP_TLS_KEY_LOCATION")

			if appCertLocEnv == "" && appKeyLocEnv == "" {
				appCertLocEnv, appKeyLocEnv = helpers.CreateOrLoadCerts()
			}

			if err := server.ListenAndServeTLS(appCertLocEnv, appKeyLocEnv); err != nil && errors.Is(err, http.ErrServerClosed) {
				helpers.Logger("MAIN PROCESS", "Closing the server...", false)
			}
		} else {
			helpers.Logger("MAIN PROCESS", "Server running in non TLS mode", false)
			if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
				helpers.Logger("MAIN PROCESS", "Closing the server...", false)
			}
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	helpers.Logger("MAIN PROCESS", "Shutting down server...", false)

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		helpers.Logger("MAIN PROCESS", fmt.Sprintf("Server forced to shutdown: %v", err), true)
	}

	helpers.Logger("MAIN PROCESS", "Server is closed.", false)
}
