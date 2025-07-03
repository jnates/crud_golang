package infrastructure

import (
	"flag"
	"github.com/jnates/crud_golang/internal/infrastructure/kit/enum"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	zerolog.MessageFieldName = "message"
	zerolog.TimestampFieldName = "date"
	zerolog.LevelFieldName = "type"
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05 Z0700 UTC"

	log.Logger = log.With().Str("app", enum.App).Logger()
	envDebug := parseBool(os.Getenv("LOGGER_DEBUG"))
	debug := flag.Bool("debug", envDebug, "sets log level to debug")

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if e := log.Debug(); e.Enabled() {
		e.Msg("Debug mode enabled")
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Error().Msgf("Error loading file .env: %v", err)
	}
}

func parseBool(boolString string) bool {
	boolVal, err := strconv.ParseBool(boolString)
	if err != nil {
		log.Error().Msgf("Error in conversion: [error] %s", err)
	} else {
		log.Error().Msgf("Converted Boolean value - %t", boolVal)
	}

	return boolVal
}
