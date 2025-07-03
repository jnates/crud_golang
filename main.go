package main

import (
	"github.com/jnates/crud_golang/internal/infrastructure"
	http "github.com/jnates/crud_golang/internal/infrastructure/http"
	"github.com/jnates/crud_golang/internal/infrastructure/kit/enum"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
)

// @title CRUD Golang API
// @version 1.0
// @description This is a sample server for managing users.
// @host localhost:8081
// @BasePath /
func main() {
	log.Info().Msg("Starting API CMD")
	infrastructure.InitLogger()

	port := os.Getenv(enum.APIPort)
	http.Start(port)
}
