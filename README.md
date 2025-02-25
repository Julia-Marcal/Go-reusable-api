## EN
# Go-reusable-api

This README is available in multiple languages:
- [English](README.md)
- [Portuguese](docs_readme/README.pt.md)

## Overview
Go-reusable-api is a highly performant, reusable, and easy-to-understand API built with Golang, Gin, and Gorm. This API is designed to offer a robust feature set while ensuring optimized performance and security.

### Features
- **Cache Support**: Utilizes Redis for caching.
- **Database**: PostgreSQL.
- **Rate Limiting**: Built-in rate limiter to prevent spam and service crashes.
- **Docker Support**: Includes Dockerfile for containerization.
- **Monitoring**: Prometheus integration for performance monitoring.
- **Authentication**: Bearer token and hashed passwords for security.

## Getting Started

### Prerequisites
- Golang
- Docker
- PostgreSQL
- Redis

### Running with Docker
Execute the following commands to build and run the Docker container:

```bash
docker build -t go-reusable-api .
docker run -p 8080:8080 go-reusable-api
```

## Configuration

### `config/env/env.go`
This file should contain one function called `SetSalt` that returns the salt in byte type.

```go
func SetSalt() []byte {
	your_salt := "Hey"
	return []byte(your_salt)
}
```

It also should contain a function called  `envGo`, so that you can set your database information.

```go 
func envGo() {
    os.Setenv("database", "host=example-db server user=fakeuser password=fakepassword dbname=fakedb port=1234 sslmode=require")
}

```

## Performance
To ensure the highest level of performance, the API has:
- Cache optimization through Redis.
- Database indexes and optimized queries in PostgreSQL.
- Clean and efficient codebase.

## Security
- Passwords are hashed for secure storage.
- Bearer token is used for system authorization.

## Monitoring with Prometheus
The API includes a Prometheus setup within the Docker container for real-time performance monitoring.

## API Requests

Examples for making API requests can be found in the `config/rest` folder. This should guide you on how to properly make requests to the API.

### Required Parameters
- `Email`
- `Password`
- `Bearer Token`

Remember to include your email, password, and the appropriate bearer token when making API calls.

## New Features:
- [ ] **Adding Swagger docummentation with swaggo** 
- [ ] **Adding govalidator**
- [ ] **Load-Balancer**
- [ ] **Add Feature Flags**

