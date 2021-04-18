
# Go PostgreSQL CI/CD

Project to test CI/CD pipeline

## Run

`/!\ WARN` Requires PostgreSQL database

## Test

`/!\ WARN` Requires Docker-in-Docker if run from CI/CD

## Endpoints

Root URL: `localhost:8080/`

| Method | URL | Description
| --- | --- | --- |
| GET | /products | Fetch list of products |
| GET | /products/{id} | Fetch a product by ID |
| POST | /products | Create a new product |
| PUT | /products/{id} | Update an existing product retrieved by ID |
| DELETE | /products/{id} | Delete a product by ID |
