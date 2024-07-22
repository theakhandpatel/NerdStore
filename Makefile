# Makefile to run migrations using the migrate/migrate Docker container

# Database URL
DATABASE_URL=postgresql://resourcer:pa55word@resourcedb:7000/resourcedb?sslmode=disable

# Migration directory
MIGRATIONS_DIR=./migrations

# 
setup-db:
	docker-compose up -d

# Run the API server
run: setup-db migrate-up
	go run ./cmd/api/main.go

# Run migrations up
migrate-up:
	docker run --rm  --net="host" -v $(MIGRATIONS_DIR):/migrations migrate/migrate -path=/migrations/ -database "$(DATABASE_URL)" up
