# Authentication Service (Bookmarket)
Authentication service is part of service Bookmarket App.
The instruction below is a instruction for running only this authentication service.
If want to run all services at a time, please enter to folder [project](https://github.com/rizkydarmawan-letenk/bookmarket/tree/master/project) and follow the instruction.

# Documentation
[Swagger](https://app.swaggerhub.com/apis-docs/DARMAWANRIZKY43/authentication_service_bookmarket/1.0.0#/Auth/post_api_v1_register)

[Postman](https://documenter.getpostman.com/view/12132212/2s83eypwRy#f0f61476-dfa0-4626-bd67-15faef7d1eaf)

# How To Run
1. Entered into folder project and run `PostgreSQL` with use docker-compose for database development
```
docker-compose -f docker-compose.dev.yml up -d --build
```

2. Export environment in terminal
```
export DSN="host=localhost port=5433 user=postgres password=password dbname=users sslmode=disable timezone=UTC connect_timeout=5"
SECRET_JWT=1e9kBs3BKOLOo6dcEavuBJS2aGW8COeuMKL5MgOuaFAHWk1idEltgbNyfLj6
```

3. Run service authentication in to folder `authentication-service`
```
go run cmd/main.go
```

# Run Tests
Before run test, please follow instructuin [How To Run](#how-to-run).
- Run all tests
```go
go test -v ./...
```

- Run all tests with coverage
```go
go test ./... -v -coverpkg=./...
```

- Run all test with coverage and print on CLI
```go
go test ./... -v -coverpkg=./... -coverprofile=coverage.out
go tool cover -func=coverage.out
```

- Run all test with coverage and print on browser
```go
go test ./... -v -coverpkg=./...
go tool cover -html=coverage.out
```