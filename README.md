# 🧹 CRUD Golang (Arquitectura Hexagonal)

Este proyecto es un CRUD de usuarios desarrollado en Go, siguiendo los principios de arquitectura hexagonal, con buenas prácticas de validación, logging, documentación y manejo de dependencias.

---

## 🚀 Tecnologías Utilizadas

* **Go** 1.21+
* **Echo**: Framework web
* **PostgreSQL**: Base de datos relacional
* **Zerolog**: Logging estructurado
* **Swaggo**: Generador de documentación Swagger
* **Go-playground/validator**: Validación de structs
* **Uber/dig**: Inyección de dependencias

---

## 📦 Instalación

```bash
git clone https://github.com/tu-usuario/crud_golang.git
cd crud_golang
go mod tidy
```

---

## 🥪 Ejecutar Proyecto

Asegúrate de tener PostgreSQL ejecutando y un archivo `.env` con las credenciales. Luego:

```bash
go run main.go
```

Por defecto, el servicio escuchará en `http://localhost:8081`.

---

## 💃 Endpoints

| Método | Endpoint     | Descripción            |
| ------ | ------------ | ---------------------- |
| GET    | `/users`     | Listar usuarios        |
| GET    | `/users/:id` | Obtener usuario por ID |
| POST   | `/users`     | Crear nuevo usuario    |
| PUT    | `/users/:id` | Actualizar usuario     |
| DELETE | `/users/:id` | Eliminar usuario       |

---

## 📓 Ejemplo de Usuario

```json
{
  "name": "Juan Pérez",
  "email": "juan@example.com"
}
```

---

## 📘 Documentación Swagger

Después de compilar los docs con:

```bash
swag init
```

La documentación estará disponible en:

```
http://localhost:8081/swagger/index.html
```

---

## 🔍 Estructura del Proyecto

```
internal/
├── application/         # Casos de uso (servicios)
├── domain/              # Modelos y puertos
├── infrastructure/
│   ├── db/              # Acceso a datos con SQL
│   ├── http/            # Controladores y middlewares
│   ├── di/              # Inyección de dependencias
│   ├── kit/             # Utilidades y constantes
cmd/                     # Entry point
docs/                    # Archivos Swagger generados
```
