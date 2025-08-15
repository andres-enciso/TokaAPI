# TokaAPI – Gestión de Tareas (Go + Gin + GORM)

API RESTful para **Gestión de Tareas** con operaciones CRUD y **autenticación basica**.  
Stack: **Go (Gin)**, **GORM**, **SQL Server**.

---

## Características

- CRUD de tareas (`titulo`, `completada` + timestamps).
- Autenticación Basica para todos los endpoints bajo `/tasks/*`.
- Migraciones automáticas con GORM (crea tablas y usuario admin si no existe).
- Listado paginado (`page`, `limit`).
- Salud del servicio en `/health`.

---

## Endpoints

- `GET  /health` → `{ "status": "ok" }`
- `POST /tasks/` → Crea tarea. Body: `{"titulo":"...", "completada":false}`
- `GET  /tasks/` → Lista paginada: `?page=1&limit=20` → `{items, page, limit, total}`
- `GET  /tasks/:id` → Obtiene una tarea por ID
- `PUT  /tasks/:id` → Actualiza (p. ej. `{"completada":true}`)
- `DELETE /tasks/:id` → Elimina

> **Auth:** Todos los endpoints bajo `/tasks/*` requieren ** Auth**.  
> Usuario/Password iniciales se generan desde variables de entorno `ADMIN_USER` / `ADMIN_PASS`.

---

## Requisitos

- Docker + Docker Compose
- (Opcional) Go 1.25+ si vas a correr **tests** localmente

---

## Variables de entorno

> En `docker-compose.yml` ya están provistas; para referencia:

| Variable              | Ejemplo         | Descripción                              |
|----------------------|-----------------|------------------------------------------|
| `PORT`               | `8080`          | Puerto HTTP del API                      |
| `ADMIN_USER`         | `admin`         | Usuario admin inicial                    |
| `ADMIN_PASS`         | `Admin!2025`    | Password admin inicial (se guarda con bcrypt) |
| `DB_HOST`            | `mssql`         | Host SQL Server (nombre del servicio en Compose) |
| `DB_PORT`            | `1433`          | Puerto SQL Server                        |
| `DB_USER`            | `sa`            | Usuario DB                               |
| `DB_PASS`            | `juE03W~3]$@o`  | Password DB                              |
| `DB_NAME`            | `TokaTasks`     | Base de datos                            |
| `DB_ENCRYPT`         | `disable`       | Encriptación (ajusta según tu SQL Server) |
| `DB_TRUSTSERVERCERT` | `true`          | Confiar en el certificado del servidor   |
| `GIN_MODE`           | `release`       | `release` para prod, `debug` para dev    |

---

## Levantar con Docker Compose
```powershell
# 1) Construir e iniciar TODO
docker compose up -d --build

# 2) Verificar el estatus de la api
curl http://localhost:8080/health     
# -> {"status":"ok"}

# 3) Probar endpoints protegidos ( Auth)
curl -u admin:Admin!2025 "http://localhost:8080/tasks/?page=1&limit=5"

# (Opcional) Reconstruir solo la API
docker compose build --no-cache api
docker compose up -d --force-recreate api

# (Opcional) Ver logs
docker logs -f go-api

# (Opcional) Apagar todo
docker compose down

# (Opcional) Reset TOTAL
docker compose down -v
docker compose up -d --build
```
---

## Ejemplos de uso

### Windows PowerShell

> Usa **`curl.exe`** y here-strings para JSON multilínea.

```powershell
# Listar
curl.exe -u admin:Admin!2025 "http://localhost:8080/tasks/?page=1&limit=5"

# Crear
$body = @"
{"titulo":"Primera tarea","completada":false}
"@
curl.exe -u admin:Admin!2025 -H "Content-Type: application/json" -d $body http://localhost:8080/tasks/

# Obtener por ID
curl.exe -u admin:Admin!2025 http://localhost:8080/tasks/1

# Actualizar
$update = @"
{"titulo":"Primera tarea (editada)","completada":true}
"@
curl.exe -u admin:Admin!2025 -H "Content-Type: application/json" -X PUT -d $update http://localhost:8080/tasks/1

# Eliminar
curl.exe -i -u admin:Admin!2025 -X DELETE http://localhost:8080/tasks/1

# 401 (sin auth)
curl.exe -i http://localhost:8080/tasks/
```

---

## Estructura del proyecto (API)

```
api/
  main.go                 # arranca el servidor, rutas base, health
internal/
  auth/
    middleware.go         # Auth y EnsureAdmin
  db/
    db.go                 # conexión SQL Server y migraciones
  models/
    task.go               # modelo Task 
    user.go               # modelo User
  tasks/
    handlers.go           # CRUD y paginación
Dockerfile
docker-compose.yml
README.md
```

---

## Solución de problemas

- **401 Unauthorized:** revisa usuario/contraseña (`ADMIN_USER`, `ADMIN_PASS`).  
- **“.env file not found” en logs:** es informativo si ya pasas variables por Compose.  
- Y cualquier otro de docker :)

---

## Notas

- **Seguridad:** `ADMIN_PASS` se almacena hasheado con **bcrypt**.  
- **Producción:** configura `GIN_MODE=release`, recursos de SQL Server y secretos a través de un vault/secret manager.

---

# Contacto
Nombre: Andres Enciso

Correo Electrónico: andresenciso20@gmail.com

GitHub: andres-enciso
