# Description: Makefile for the API server

# Migration directory
MIGRATIONS_DIR=./migrations

# Setup the database and run the migrations
setup-db:
	docker-compose up -d

# Run the API server
run: setup-db
	go run ./cmd/api/

build:
	pwd
	ls
	go build -ldflags="-s -w" -o ./bin/nerdstore ./cmd/api/


