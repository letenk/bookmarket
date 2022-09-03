API_GATEWAY_BINARY=apiGatewayApp
AUTHENTICATION_BINARY=authenticationApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_db_dev: start docker compose postgresSql for development
up_db_dev:
	@echo "Starting Database Book Market Development"
	docker-compose -f docker-compose.dev.yml up -d
	@echo "Database Book Market Development started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_api_gateway build_authentication
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## down_with_volume: stop docker compose with remove volume
down_with_volume:
	@echo "Stopping docker compose..."
	docker-compose down -v
	@echo "Done!"

## down_db_dev: stop docker compose postgresSql for development
down_db_dev:
	@echo "Stopping Database Book Market Development"
	docker-compose -f docker-compose.dev.yml down
	@echo "Database Book Market Development stopped!"

## down_db_dev_with_volume: stop docker compose postgresSql for development with volume
down_db_dev_with_volume:
	@echo "Stopping Database Book Market Development"
	docker-compose -f docker-compose.dev.yml down -v
	@echo "Database Book Market Development stopped!"

## build_api_gateway: builds the api gateway binary as a linux executable
build_api_gateway:
	@echo "Building api gateway binary..."
	cd ../api-gateway && env GOOS=linux CGO_ENABLED=0 go build -o ${API_GATEWAY_BINARY} ./cmd
	@echo "Done!"

## build_authentication: builds the authentication binary as a linux executable
build_authentication:
	@echo "Building authentication binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTHENTICATION_BINARY} ./cmd
	@echo "Done!"