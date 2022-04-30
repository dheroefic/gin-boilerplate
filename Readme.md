# Gin boilerplate
Personal boiler plate for starting gin project


## Folder Structure

```
.
└── gin-boilerplate/
    ├── .vscode
    ├── database/
    │   ├── database.go
    │   └── migrations/
    │       └── // Migrations file goes here
    ├── dist/
    │   └── // Mainly to hold compiled application
    ├── domains/
    │   └── // Your API usecase goes here
    ├── models/
    │   ├── // All models goes here
    │   └── structs
    ├── routes/
    │   └── routes.go
    ├── utils/
    │   ├── constants/
    │   │   └── // All constants data goes here
    │   └── helpers/
    │       └── // Helper function goes here
    ├── .env.example
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── Makefile
    └── Readme.md
```

## How to set up
- Clone the repo
- Run

```sh
go mod tidy
go mod vendor
```

## How to run on local for testing

- Run `make start`
- or Run VSCode Debugger through delve (https://github.com/go-delve/delve)

## How to test

- Run `make test`
- Please create another test suite before adding new feature