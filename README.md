# REST API in Go

## Introduction

This project is a REST API built in Go that performs CRUD (Create, Read, Update, Delete) operations and provides data from a persistent backend MySQL database. The API returns responses in JSON format, making it easy to integrate with various front-end applications and services.

The project utilizes the following libraries:
- **go-sql-driver/mysql**: A MySQL driver for Go's database/sql package, used for database operations.
- **gorilla/mux**: A powerful URL router and dispatcher for matching incoming requests to their respective handler functions.

With this setup, the API ensures efficient and reliable data handling and routing, making it a robust solution for backend services.

## Table of Contents

- [Requirements](#requirements)
- [Setup](#setup)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Running Tests](#running-tests)
- [Code Overview](#code-overview)
- [Continuous Integration](#continuous-integration)
- [Contributions](#contributions)
- [License](#license)



## Requirements

- Go 1.22.2 or later
- MySQL

## Setup

1. **Clone the repository:**

    ```sh
    git clone https://github.com/deviant101/REST-API.git
    cd REST-API
    ```

2. **Install dependencies:**

    ```sh
    go mod download
    ```

3. **Set up MySQL:**

    Ensure MySQL is running and create a database for the application. Update the database credentials in the environment variables.

4. **Set Environment Variables:**

    Create a `.env` file in the root directory of the project and add the following environment variables:

    ```sh
    DB_HOST=localhost
    DB_PORT=3306
    DB_USER=root
    DB_PASSWORD=yourpassword
    DB_NAME=yourdatabase
    ```

5. **Docker Setup:**

    This project includes Docker files to set up the entire stack easily. Ensure you have Docker installed on your machine.

    - **Build the Docker image:**

        ```sh
        docker build -t rest-api .
        ```

    - **Run the Docker container:**

        ```sh
        docker run -d -p 8080:8080 --env-file .env rest-api
        ```

6. **Docker Compose Setup:**

    This project includes a `docker-compose.yml` file to set up the entire stack, including MySQL, easily. Ensure you have Docker and Docker Compose installed on your machine.

    - **Update the `.env` file to use the MySQL container:**

        ```sh
        DB_HOST=db
        DB_PORT=3306
        DB_USER=root
        DB_PASSWORD=yourpassword
        DB_NAME=yourdatabase
        ```

    - **Start the services using Docker Compose:**

        ```sh
        docker-compose up -d
        ```

## Running the Application

To run the application, use the following command:

```sh
go run main.go
```


## API Endpoints

- **GET/products:** Retrieve all products.
- **GET/product/{id}:** Retrieve a product by ID
- **POST/product:** Create a new product.
- **PUT/product/{id}:** Update a product by ID.
- **DELETE/product/{id}:** Delete a product by ID.

### Example Requests

### GET /products
```sh
curl -X GET http://localhost:8080/products
```

### POST /products
```sh
curl -X POST http://localhost:8080/product -d '{"name": "New Product", "price": 100.00}'
```
### PUT /product/{id}
```sh
curl -X PUT http://localhost:8080/product/1 -d '{"name": "Updated Product", "price": 150.00}'
```

### DELETE /product/{id}
```sh
curl -X DELETE http://localhost:8080/product/1
```

## Running Tests

To ensure the functionality and reliability of the API, a suite of tests has been included. These tests can be run using the Go testing framework.

1. **Run all tests:**

    To run all the tests in the project, use the following command:

    ```sh
    go test ./...
    ```

2. **View detailed test output:**

    To see detailed output from the tests, use the `-v` (verbose) flag:

    ```sh
    go test -v ./...
    ```

These commands will help you verify that your API is functioning correctly and that any changes you make do not introduce new issues.

## Code Overview
`main.go` The entry point of the application. It initializes the application and starts the server.

`app.go` Defines the application structure and routes:
```sh
func (app *App) handleRoutes() {
    app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
    app.Router.HandleFunc("/product/{id}", app.getProduct).Methods("GET")
    app.Router.HandleFunc("/product", app.createProduct).Methods("POST")
    app.Router.HandleFunc("/product/{id}", app.updateProduct).Methods("PUT")
    app.Router.HandleFunc("/product/{id}", app.deleteProduct).Methods("DELETE")
}
```

`app_test.go` Contains tests for the API:
```sh
func TestGetProduct(t *testing.T) {
    clearTable()
    addProduct("Product1", 10, 100.00)
    request, _ := http.NewRequest("GET", "/products", nil)
    response := sendRequest(request)
    checkStatusCode(t, http.StatusOK, response.Code)
}
```

`go.mod` Module dependencies:
```sh
module github.com/deviant101/REST-API

go 1.22.2

require (
    github.com/go-sql-driver/mysql v1.8.1
    github.com/gorilla/mux v1.8.1
)
require filippo.io/edwards25519 v1.1.0 // indirect
```

## Continuous Integration

This project includes a Continuous Integration (CI) pipeline configured using GitHub Actions. The CI pipeline ensures that the codebase remains stable and that all tests pass before any changes are merged into the main branch.

### CI Pipeline Features

- **Automated Testing:** The pipeline automatically runs all tests in the project whenever a new commit is pushed or a pull request is opened. This helps catch any issues early in the development process.
- **Dependency Installation:** The pipeline installs all necessary dependencies, ensuring that the environment is correctly set up for testing.

## Contributions

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

Feel free to customize this [`README.md`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2Fhome%2Fdeviant%2FData%2FREST-API%2FREADME.md%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/home/deviant/Data/REST-API/README.md") file further based on your specific project details and requirements.
