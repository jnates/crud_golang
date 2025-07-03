package db

import (
	"database/sql"
	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/jnates/crud_golang/internal/domain/ports"
	queryVar "github.com/jnates/crud_golang/internal/infrastructure/db/queries"
	"github.com/jnates/crud_golang/internal/infrastructure/kit/enum"
	"github.com/jnates/crud_golang/internal/infrastructure/kit/tool/dbutils"
	"github.com/rs/zerolog/log"
)

// userRepository implementa el puerto UserRepository con una fuente de datos SQL.
type userRepository struct {
	db *sql.DB
}

// NewUserRepository crea una nueva instancia de userRepository.
func NewUserRepository(db *sql.DB) ports.UserRepository {
	return &userRepository{db: db}
}

// GetByID obtiene un usuario por su ID.
// Devuelve un puntero al modelo de usuario o un error si no se encuentra o hay problemas en la base de datos.
func (r *userRepository) GetByID(id int64) (*model.User, error) {
	log.Debug().Int64(enum.ID, id).Msg("🟢 Buscando usuario por ID")

	row := r.db.QueryRow(queryVar.QueryGetUserByID, id)

	var user model.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		log.Error().Err(err).Int64(enum.ID, id).Msg("🔴 Error al escanear usuario por ID")
		return nil, err
	}

	log.Debug().Int64(enum.ID, user.ID).Msg("✅ Usuario encontrado")
	return &user, nil
}

// Create inserta un nuevo usuario en la base de datos.
// Devuelve el ID del nuevo usuario o un error si ocurre un fallo.
func (r *userRepository) Create(user *model.User) (int64, error) {
	log.Debug().Str(enum.Name, user.Name).Str(enum.Email, user.Email).Msg("🟢 Creando nuevo usuario")

	var id int64
	err := r.db.QueryRow(queryVar.QueryInsertUser, user.Name, user.Email).Scan(&id)
	if err != nil {
		log.Error().Err(err).Msg("🔴 Error al crear usuario")
		return 0, err
	}

	log.Info().Int64(enum.ID, id).Msg("✅ Usuario creado exitosamente")
	return id, nil
}

// Update actualiza los datos de un usuario existente por su ID.
// Devuelve un error si la operación falla.
func (r *userRepository) Update(user *model.User) error {
	log.Debug().Int64(enum.ID, user.ID).Msg("🟡 Actualizando usuario")

	_, err := r.db.Exec(queryVar.QueryUpdateUser, user.Name, user.Email, user.ID)
	if err != nil {
		log.Error().Err(err).Int64(enum.ID, user.ID).Msg("🔴 Error al actualizar usuario")
		return err
	}

	log.Info().Int64(enum.ID, user.ID).Msg("✅ Usuario actualizado correctamente")
	return nil
}

// Delete elimina un usuario de la base de datos por su ID.
// Devuelve un error si ocurre un fallo.
func (r *userRepository) Delete(id int64) error {
	log.Debug().Int64(enum.ID, id).Msg("🟠 Eliminando usuario")

	_, err := r.db.Exec(queryVar.QueryDeleteUser, id)
	if err != nil {
		log.Error().Err(err).Int64(enum.ID, id).Msg("🔴 Error al eliminar usuario")
		return err
	}

	log.Info().Int64(enum.ID, id).Msg("✅ Usuario eliminado correctamente")
	return nil
}

// List obtiene una lista paginada de usuarios con filtros dinámicos opcionales.
// Recibe offset, limit y un mapa de filtros (por nombre, email, etc.).
// Devuelve un slice de punteros a modelo User o un error.
func (r *userRepository) List(offset int, limit int, filters map[string]interface{}) ([]*model.User, error) {
	log.Debug().
		Int(enum.Offset, offset).
		Int(enum.Limit, limit).
		Interface(enum.Filters, filters).
		Msg("🔍 Listando usuarios con filtros")

	query, args := dbutils.BuildDynamicQuery(queryVar.QuerySelectUserBase, filters, 1)
	query, args = dbutils.AddPagination(query, args, len(args)+1, limit, offset)

	log.Debug().Str(enum.Query, query).Interface(enum.Args, args).Msg("📄 Query final construida")

	rows, err := r.db.Query(query, args...)
	if err != nil {
		log.Error().Err(err).Msg("🔴 Error ejecutando query de listado")
		return nil, err
	}
	defer rows.Close()

	users, scanErr := dbutils.ScanRows(rows, func(row *sql.Rows) (*model.User, error) {
		var user model.User
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Error().Err(err).Msg("🔴 Error al escanear fila de usuario")
			return nil, err
		}
		return &user, nil
	})

	if scanErr != nil {
		log.Error().Err(scanErr).Msg("🔴 Error al escanear resultados del listado")
		return nil, scanErr
	}

	log.Info().Int(enum.Total, len(users)).Msg("✅ Usuarios listados exitosamente")
	return users, nil
}
