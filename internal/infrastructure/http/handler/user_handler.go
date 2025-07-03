package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/jnates/crud_golang/internal/application"
	"github.com/jnates/crud_golang/internal/domain/model"
	"github.com/jnates/crud_golang/internal/infrastructure/kit/enum"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type UserHandler struct {
	Service *application.UserService
}

func NewUserHandler(svc *application.UserService) *UserHandler {
	return &UserHandler{Service: svc}
}

// Get godoc
// @Summary      Get user by ID
// @Description  Retrieve a user using their ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  model.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func (h *UserHandler) Get(c echo.Context) error {
	id, err := parseID(c.Param(enum.ID))
	if err != nil {
		log.Error().Err(err).Int(enum.Status, http.StatusBadRequest).Msg("❌ ID inválido")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user ID"})
	}

	u, err := h.Service.Get(id)
	if err != nil {
		log.Warn().Err(err).Int(enum.Status, http.StatusNotFound).Msg("⚠️ Usuario no encontrado")
		return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
	}

	log.Info().Int(enum.Status, http.StatusOK).Int64("userID", u.ID).Msg("✅ Usuario encontrado")
	return c.JSON(http.StatusOK, u)
}

// Create godoc
// @Summary      Create new user
// @Description  Create a new user with name and email
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "User data"
// @Success      201   {object}  model.User
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users [post]
func (h *UserHandler) Create(c echo.Context) error {
	var u model.User
	if err := c.Bind(&u); err != nil {
		log.Error().Err(err).Int(enum.Status, http.StatusBadRequest).Msg("❌ Error al parsear body")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	if err := c.Validate(&u); err != nil {
		log.Error().Err(err).Int(enum.Status, http.StatusBadRequest).Msg("❌ Validación fallida")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	id, err := h.Service.Create(&u)
	if err != nil {
		log.Error().Err(err).Int(enum.Status, http.StatusInternalServerError).Msg("❌ Error al crear usuario")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	u.ID = id
	log.Info().Int64(enum.ID, id).Int(enum.Status, http.StatusCreated).Msg("✅ Usuario creado")
	return c.JSON(http.StatusCreated, u)
}

// Update godoc
// @Summary      Update user
// @Description  Update user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "User ID"
// @Param        user  body      model.User  true  "Updated user"
// @Success      200   "No Content"
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /users/{id} [put]
func (h *UserHandler) Update(c echo.Context) error {
	id, err := parseID(c.Param(enum.ID))
	if err != nil {
		log.Error().Err(err).Int(enum.Status, http.StatusBadRequest).Msg("❌ ID inválido")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user ID"})
	}

	var u model.User
	if err := c.Bind(&u); err != nil {
		log.Error().Err(err).Int(enum.Status, http.StatusBadRequest).Msg("❌ Error al parsear body")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}

	if err := c.Validate(&u); err != nil {
		log.Error().Err(err).Int(enum.Status, http.StatusBadRequest).Msg("❌ Validación fallida")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	u.ID = id
	if err := h.Service.Update(&u); err != nil {
		log.Error().Err(err).Int(enum.Status, http.StatusInternalServerError).Msg("❌ Error al actualizar usuario")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	log.Info().Int64(enum.ID, id).Int(enum.Status, http.StatusOK).Msg("✅ Usuario actualizado")
	return c.NoContent(http.StatusOK)
}

// Delete godoc
// @Summary      Delete user
// @Description  Delete a user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /users/{id} [delete]
func (h *UserHandler) Delete(c echo.Context) error {
	id, err := parseID(c.Param(enum.ID))
	if err != nil {
		log.Error().Err(err).Int(enum.Status, http.StatusBadRequest).Msg("❌ ID inválido")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user ID"})
	}

	if err := h.Service.Delete(id); err != nil {
		log.Error().Err(err).Int(enum.Status, http.StatusInternalServerError).Msg("❌ Error al eliminar usuario")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	log.Info().Int64(enum.ID, id).Int(enum.Status, http.StatusNoContent).Msg("✅ Usuario eliminado")
	return c.NoContent(http.StatusNoContent)
}

// List godoc
// @Summary      List users
// @Description  Retrieve paginated and filtered list of users
// @Tags         users
// @Produce      json
// @Param        name   query     string  false  "Filter by name"
// @Param        email  query     string  false  "Filter by email"
// @Param        page   query     int     false  "Page number"
// @Param        limit  query     int     false  "Items per page"
// @Success      200    {array}   model.User
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /users [get]
func (h *UserHandler) List(c echo.Context) error {
	filters := make(map[string]interface{})
	if name := c.QueryParam(enum.Name); name != enum.EmptyString {
		filters[enum.Name] = name
	}
	if email := c.QueryParam(enum.Email); email != enum.EmptyString {
		filters[enum.Email] = email
	}

	page, err := parseIntOrDefault(c.QueryParam(enum.Page), 1)
	if err != nil {
		log.Error().Err(err).Int(enum.Status, http.StatusBadRequest).Msg("❌ Página inválida")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid page number"})
	}

	limit, err := parseIntOrDefault(c.QueryParam(enum.Limit), 10)
	if err != nil {
		log.Error().Err(err).Int(enum.Status, http.StatusBadRequest).Msg("❌ Límite inválido")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid limit"})
	}

	offset := (page - 1) * limit
	users, err := h.Service.List(offset, limit, filters)
	if err != nil {
		log.Error().Err(err).Int(enum.Status, http.StatusInternalServerError).Msg("❌ Error al listar usuarios")
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	log.Info().Int(enum.Status, http.StatusOK).Int(enum.Total, len(users)).Msg("✅ Usuarios listados")
	return c.JSON(http.StatusOK, users)
}

// --- helpers ---

func parseID(idStr string) (int64, error) {
	if strings.EqualFold(idStr, enum.EmptyString) {
		return 0, errors.New("missing ID")
	}
	return strconv.ParseInt(idStr, 10, 64)
}

func parseIntOrDefault(value string, def int) (int, error) {
	value = strings.TrimSpace(value)
	if strings.EqualFold(value, enum.EmptyString) {
		return def, nil
	}
	return strconv.Atoi(value)
}
