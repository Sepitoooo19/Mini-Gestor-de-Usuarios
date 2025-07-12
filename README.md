# Mini-Gestor-de-Usuarios

API REST para la gestión de usuarios desarrollada en Go con Gin framework y MongoDB.

## Requisitos previos

- Go 1.23.0 o superior
- MongoDB ejecutándose en localhost:27017
- Puerto 8080 disponible

## Instalación y ejecución

### 1. Clonar el repositorio
```bash
git clone https://github.com/Sepitoooo19/Mini-Gestor-de-Usuarios.git
cd Mini-Gestor-de-Usuarios
```

### 2. Sincronizar dependencias
```bash
go mod tidy
```
Este comando descarga e instala todas las dependencias necesarias del proyecto.

### 3. Compilar el proyecto
```bash
go build -o mini-gestor.exe
```
Esto genera un ejecutable llamado `mini-gestor.exe`.

### 4. Ejecutar el proyecto

**Opción A: Ejecutar directamente**
```bash
go run main.go
```

**Opción B: Ejecutar el compilado**
```bash
./mini-gestor.exe
```

El servidor se iniciará en `http://localhost:8080`

## Endpoints disponibles

- `GET /ping` - Verificar estado del servidor
- `POST /users` - Crear un nuevo usuario
- `GET /users` - Obtener todos los usuarios
- `GET /users/:id` - Obtener un usuario por ID

## Estructura del proyecto

```
├── config/          # Configuración (CORS)
├── controllers/     # Controladores HTTP
├── models/          # Modelos de datos
├── routes/          # Definición de rutas
├── services/        # Lógica de negocio y conexión DB
├── main.go          # Punto de entrada de la aplicación
├── go.mod           # Dependencias del proyecto
└── go.sum           # Checksums de dependencias
```