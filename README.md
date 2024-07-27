# NerdStore 🤓

NerdStore is a simple URL management API that allows you to store and manage your favorite URLs. This is also an attempt to do tasks of [SRE Bootcamp](https://one2n.io/sre-bootcamp)

## Features

* Add a new URL
* Edit an existing URL
* Delete an existing URL
* View all URLs
* Filter URLs by tags
* Search URLs by title
* Sort URLs by title or date added
* TODO: Search for a URL

## Tech Stack

* Golang
* PostgresDB
* Docker/Docker-compose

## Running NerdStore

### Production Environment

1. Clone the repository
2. Create a `.env` file in the root directory with the following variables:
	* `POSTGRES_USER`
	* `POSTGRES_PASSWORD`
	* `POSTGRES_DB`
3. Run `docker-compose up -d` to start the containers in detached mode
4. The API will be available at `http://localhost:8080`

### Local Development Environment

1. Clone the repository
2. Create a `.env` file in the root directory with the following variables:
	* `POSTGRES_USER`
	* `POSTGRES_PASSWORD`
	* `POSTGRES_DB`
3. Run `docker-compose up -d resourcedb migrate` to start the database and migration containers
4. Run `go run cmd/api/main.go` to start the API server
5. The API will be available at `http://localhost:8080`

**Note**: Make sure to update the `POSTGRES_HOST` variable in the `api` service of the `docker-compose.yml` file to `resourcedb` if you want to connect to the database container.

### Building and Running with Docker (Alternative)

1. Clone the repository
2. Create a `.env` file in the root directory with the following variables:
	* `POSTGRES_USER`
	* `POSTGRES_PASSWORD`
	* `POSTGRES_DB`
3. Run `docker-compose build` to build the Docker image
4. Run `docker-compose up` to start the containers
5. The API will be available at `http://localhost:8080`

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

NerdStore is licensed under the MIT License.