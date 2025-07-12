# Mini-Gestor-de-Usuarios

API REST para la gestión de usuarios desarrollada en Go con MongoDB.

## Requisitos previos

- Go 1.23.0 o superior
- MongoDB ejecutándose en localhost:27017
- Puerto 8080 disponible
- IDE recomendado: GoLand o VS Code

## Instalación y ejecución

### 1. Clonar el repositorio

**Usando Git desde terminal:**
```bash
git clone https://github.com/Sepitoooo19/Mini-Gestor-de-Usuarios.git
cd Mini-Gestor-de-Usuarios
```

**Usando GitKraken:**
1. Abrir GitKraken
2. Hacer clic en "Clone a repo"
3. Seleccionar "Clone with URL"
4. Pegar la URL: `https://github.com/Sepitoooo19/Mini-Gestor-de-Usuarios.git`
5. Elegir la carpeta de destino
6. Hacer clic en "Clone the repo!"

### 2. Configurar MongoDB

Crear la base de datos y colección necesarias:

**Usando MongoDB Compass:**
1. Conectar a `mongodb://localhost:27017`
2. Crear nueva base de datos llamada `Mongo`
3. Dentro de la base de datos, crear una colección llamada `users`

#### Poblar la base de datos con datos de ejemplo

El proyecto incluye un archivo `sample_users.json` con datos de ejemplo.

**Importar usando MongoDB Compass:**
1. En MongoDB Compass, navegar a la base de datos `Mongo`
2. Hacer clic en la colección `users`
3. Hacer clic en "ADD DATA" → "Import JSON or CSV file"
4. Seleccionar el archivo `sample_users.json`
5. Hacer clic en "Import"




### 3. Sincronizar dependencias
```bash
go mod tidy
```
Este comando descarga e instala todas las dependencias necesarias del proyecto.

### 4. Compilar el proyecto
```bash
go build
```

### 5. Ejecutar el proyecto
```bash
go run main.go
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