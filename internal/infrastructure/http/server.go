package infrastructure

import (
	_ "github.com/jnates/crud_golang/docs"
	"github.com/jnates/crud_golang/internal/infrastructure/db"
	"github.com/jnates/crud_golang/internal/infrastructure/di"
	"github.com/jnates/crud_golang/internal/infrastructure/http/handler"
	validatorPackage "github.com/jnates/crud_golang/internal/infrastructure/http/validetor"
	"github.com/jnates/crud_golang/internal/infrastructure/kit/enum"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Start(port string) {
	conn := db.NewPostgresConnection()
	container := di.BuildContainer(conn)
	if container == nil {
		log.Fatal().Msg("Error al construir contenedor DI")
		return
	}

	err := container.Invoke(func(userHandler *handler.UserHandler) {
		e := echo.New()
		e.HideBanner = true
		e.Logger.SetOutput(log.Logger)

		e.Validator = validatorPackage.NewValidator()

		// Swagger docs
		e.GET("/swagger/*", echoSwagger.WrapHandler)

		// Rutas de API
		api := e.Group("/users")
		api.GET("", userHandler.List)
		api.GET("/:id", userHandler.Get)
		api.POST("", userHandler.Create)
		api.PUT("/:id", userHandler.Update)
		api.DELETE("/:id", userHandler.Delete)

		log.Info().Str(enum.APIPort, port).Msg("🚀 Servidor escuchando")
		if err := e.Start(":" + port); err != nil {
			log.Fatal().Err(err).Msg("Error al iniciar servidor")
		}
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Error al inicializar dependencias con dig")
	}
}
