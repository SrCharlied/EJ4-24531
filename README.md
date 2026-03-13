# 🎮 RPG Games API — RESTful JSON API with Go

API REST desarrollada en Go utilizando únicamente la librería estándar (`net/http`).  
Permite gestionar una colección de videojuegos RPG (inspirados en títulos como Elden Ring, Dark Souls, Persona, etc.) con persistencia en archivo JSON y ejecución mediante Docker.

---

## 🌟 Descripción General

Este proyecto implementa una API REST completa con soporte para:

- CRUD completo (GET, POST, PUT, PATCH, DELETE)
- Query parameters
- Path parameters
- Filtros combinados
- Validaciones robustas
- Manejo consistente de errores en formato JSON
- Persistencia real en archivo JSON
- Ejecución en contenedor Docker
- Uso de volúmenes para persistencia en contenedor

El objetivo principal es demostrar comprensión real del funcionamiento de HTTP, métodos, parámetros y manejo de datos en Go sin frameworks externos.

---

## 🧠 Características Implementadas

- ✅ GET todos los elementos
- ✅ GET por query parameter (`?id=`)
- ✅ GET por path parameter (`/api/games/{id}`)
- ✅ POST (crear recurso)
- ✅ PUT (actualización completa)
- ✅ PATCH (actualización parcial)
- ✅ DELETE (eliminar recurso)
- ✅ Filtros combinados dinámicos
- ✅ Validación de datos (tipos, rangos, campos requeridos)
- ✅ Manejo estructurado de errores en JSON
- ✅ Persistencia real en `games.json`
- ✅ Docker multi-stage build
- ✅ Soporte de volumen para persistencia en contenedor

---

## 🏗 Arquitectura del Proyecto

```Text
EJ4-24531/
│
├── cmd/
│ └── server/
│ └── main.go
│
├── internal/
│ ├── handlers/
│ ├── models/
│ ├── storage/
│ └── utils/
│
├── data/
│ └── games.json
│
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
```

### Separación de responsabilidades:

- `handlers/` → Manejo de endpoints HTTP
- `storage/` → Persistencia en JSON
- `models/` → Estructuras de datos
- `utils/` → Respuestas JSON y manejo de errores
- `cmd/server/` → Punto de entrada de la aplicación

---

## 🚀 Ejecución Local

### 1️⃣ Requisitos

- Go 1.22+

### 2️⃣ Ejecutar servidor

```bash
go run ./cmd/server

servidor escuchando en:

http://localhost:24531
```

---

## 🐳 Ejecución con Docker

```bash
docker build -t games-api .

docker run -p 24531:24531 games-api
```

Si se desea persistencia con volumen:

```bash
docker build -t games-api .

docker run -p 24531:24531 -v ${PWD}/data:/app/data games-api
```
Esto permite que los cambios en games.json persistan fuera del contenedor.

---

## 📡 Endpoints Disponibles

### GET todos los juegos

```bash
GET /api/games
```

### GET por query parameter

```bash
GET /api/games?id=1
```

### GET por path parameter

```bash
GET /api/games/1
```

### POST crear juego

```bash
POST /api/games

Body:
{
  "title": "Persona 5 Royal",
  "developer": "Atlus",
  "genre": "JRPG",
  "release_year": 2020,
  "difficulty": 6,
  "platform": "PS5",
  "boss_count": 12
}
```

### PUT actualizar completamente

```bash
PUT /api/games/{id}
```

### PATCH actualizar parcialmente

```bash
PATCH /api/games/{id}

Body:
{
  "difficulty": 9
}
```

### DELETE eliminar recurso

```bash
DELETE /api/games/{id}
```

### Filtros Combinados

Se soportan varios filtros

```bash
GET /api/games?genre=Action RPG
GET /api/games?platform=PC
GET /api/games?difficulty=9
GET /api/games?genre=Action RPG&platform=PC
GET /api/games?genre=Action RPG&platform=PC&difficulty=9
```
Nota: tomar en cuenta los espacios :)

### Filtros Combinados

Manejo de errores

```bash
{
  "error": "mensaje descriptivo"
}
```
---
## Persistencia

Los datos se almacenan en:
```bash
data/games.json
```
Cada operación que modifica datos guarda automáticamente el archivo.

---

## 🛠 Tecnologías Utilizadas

- Go (net/http estándar)
- JSON como almacenamiento
- Docker (multi-stage build)
- Manejo manual de routing y parámetros