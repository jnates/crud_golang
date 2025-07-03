package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jnates/crud_golang/internal/infrastructure/kit/enum"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func NewPostgresConnection() *sql.DB {
	host := os.Getenv(enum.DBHost)
	port := os.Getenv(enum.DBPort)
	user := os.Getenv(enum.DBUser)
	password := os.Getenv(enum.DBPassword)
	dbname := os.Getenv(enum.DBName)
	sslmode := os.Getenv(enum.SSLMode)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	log.Debug().Str("dsn", dsn).Msg("Construyendo conexión a la DB")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Error al abrir conexión a DB")
	}

	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("No se pudo conectar a la base de datos")
	}

	log.Info().Msg("✅ Conexión a PostgreSQL establecida")
	return db
}
