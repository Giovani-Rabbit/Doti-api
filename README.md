<h1 align="center">
   Doti - API
</h1>

<h3 align="center">
    This is the REST API written in Go to communicate with <a href="https://github.com/Giovani-Coelho/Doti">DOTI</a>.
</h3>

### Built With
- [Golang](https://go.dev/)
- [Zap](https://github.com/uber-go/zap)
- [Migrate](https://github.com/golang-migrate/migrate)
- [godotenv](github.com/joho/godotenv)

### Relationship
<img src="./public/relationship.png">

## Running

- Execute the container.
- Run the migrations.
- Run the server.

### Docker
Run docker container:
```
docker compose up -d
```

### Migrations
Run Migrations:
```
make migrate-up
```
Revert Migrations:
```
make migrate-down
```

### Running the application
Execute the server:
```
go run cmd/api/main.go
```
