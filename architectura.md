# Arquitectura Hexagonal

├── cmd
│   └── main.go
├── internal
│   ├── core
│   │   ├── domain
│   │   │   └── user.go
│   │   ├── ports
│   │   │   └── user_ports.go
│   │   └── services
│   │       └── user_services.go
│   ├── handlers
│   │   └── user_handlers.go
│   ├── repositories
│   │   └── user_repositories.go
│   └── server
│       └── server.go

- cmd/main.go: Es el punto de entrada de tu aplicación.
- internal/core/domain/user.go: Aquí defines tus entidades o modelos.
- internal/core/ports/user_ports.go: Define las interfaces para los servicios y repositorios.
- internal/core/services/user_services.go: Implementa la lógica de negocio principal de tu aplicación.
- internal/handlers/user_handlers.go: Maneja las solicitudes HTTP y las respuestas.
- internal/repositories/user_repositories.go: Implementa la lógica para interactuar con tu base de datos.
- internal/server/server.go: Configura y arranca tu servidor.

## Otra manera de hacer la arquitectura hexagonal

├── cmd
│   └── myapp
│       └── main.go # Punto de entrada de la aplicación
├── pkg
│   ├── adapter
│   │   ├── http
│   │   │   ├── handler.go # Handlers HTTP
│   │   │   └── router.go # Configuración del router HTTP
│   │   └── persistence
│   │       ├── repository.go # Implementación del repositorio
│   │       └── database.go # Configuración de la conexión a la base de datos
│   ├── domain
│   │   ├── model.go # Definición de los modelos de dominio
│   │   └── service.go # Interfaces de los servicios de dominio
│   └── service
│       └── service.go # Implementación de los servicios de dominio
└── README.md
