# Doti-API
## Executing

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

### Running the Application
Execute the server:
```
go run cmd/api/main.go
```
