# Project agendahora

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```

DDD
```bash
├── cmd
│   ├── api
│   │   └── main.go
│   └── web
│       ├── assets
│       │   ├── css
│       │   │   ├── input.css
│       │   │   └── output.css
│       │   └── js
│       │       └── htmx.min.js
│       ├── base.templ
│       ├── base_templ.go
│       ├── handlers
│       │   ├── auth_handler.go
│       │   ├── dashboard_handler.go
│       │   ├── professional_handler.go
│       │   ├── client_handler.go
│       │   ├── service_handler.go
│       │   ├── appointment_handler.go
│       │   └── payment_handler.go
│       ├── middleware
│       │   ├── auth_middleware.go
│       │   └── payment_middleware.go
│       └── templates
│           ├── index.html
│           ├── dashboard.html
│           ├── client_form.html
│           ├── service_list.html
│           ├── appointment_calendar.html
│           └── components
│               ├── header.html
│               ├── sidebar.html
│               ├── footer.html
│               └── modal.html
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── database
│   │   └── database.go
│   └── domain
│       ├── entities
│       │   ├── cliente.go
│       │   ├── servico.go
│       │   ├── agendamento.go
│       │   └── usuario.go
│       ├── repositories
│       │   ├── cliente_repository.go
│       │   ├── servico_repository.go
│       │   ├── agendamento_repository.go
│       │   └── usuario_repository.go
│       └── services
│           ├── cliente_service.go
│           ├── servico_service.go
│           ├── agendamento_service.go
│           └── usuario_service.go
├── main
├── Makefile
├── README.md
├── tailwind.config.js
├── tests
│   └── handler_test.go
└── tmp
    └── build-errors.log
```
