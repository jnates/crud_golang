package di

import (
	"database/sql"

	"github.com/jnates/crud_golang/internal/application"
	"github.com/jnates/crud_golang/internal/domain/ports"
	"github.com/jnates/crud_golang/internal/infrastructure/db"
	"github.com/jnates/crud_golang/internal/infrastructure/http/handler"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

func BuildContainer(conn *sql.DB) *dig.Container {
	log.Debug().Msg("🧱 Iniciando construcción del contenedor de dependencias")

	container := dig.New()

	if err := container.Provide(func() ports.UserRepository {
		log.Debug().Msg("🔌 Registrando UserRepository")
		return db.NewUserRepository(conn)
	}); err != nil {
		log.Error().Err(err).Msg("❌ Error registrando UserRepository")
		return nil
	}

	if err := container.Provide(func(repo ports.UserRepository) *application.UserService {
		log.Debug().Msg("🔌 Registrando UserService")
		return application.NewUserService(repo)
	}); err != nil {
		log.Error().Err(err).Msg("❌ Error registrando UserService")
		return nil
	}

	if err := container.Provide(func(svc *application.UserService) *handler.UserHandler {
		log.Debug().Msg("🔌 Registrando UserHandler")
		return handler.NewUserHandler(svc)
	}); err != nil {
		log.Error().Err(err).Msg("❌ Error registrando UserHandler")
		return nil
	}

	log.Debug().Msg("✅ Contenedor construido exitosamente")
	return container
}
