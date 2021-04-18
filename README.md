
# Go PostgreSQL CI/CD

Project to test CI/CD pipeline

## Endpoints

Root URL: `localhost:8080/`

| Method | URL | Description
| --- | --- | --- |
| GET | /products | Fetch list of products |
| GET | /products/{id} | Fetch a product by ID |
| POST | /products | Create a new product |
| PUT | /products/{id} | Update an existing product retrieved by ID |
| DELETE | /products/{id} | Delete a product by ID |

---

## Build

```bash
make build
```

## Test

`/!\ WARN` Requires Docker-in-Docker if run from CI/CD

### All tests

```bash
make test
```

### Unit tests

```bash
make unit-test
```

### Integration tests

`/!\ WARN` Requires Docker up and running

```bash
make integr-test
```

### Show coverage

```bash
# cli
make coverage-cli

# browser
make coverage-browser
```

## Run

`/!\ WARN` Requires Docker up and running

`/!\ WARN` Requires PostgreSQL database

```bash
make start-postgres

make run
```

## Container

`/!\ WARN` Requires Docker up and running

### Build

```bash
make container-build CONTAINER_TAG=dev
```

### Push

```bash
make container-push CONTAINER_TAG=dev
```
