# Go Backend API

A clean, high-performance backend API built with Go.

## Features

- **Standard Library Only**: Minimal dependencies for maximum performance and security.
- **Modern Routing**: Utilizes Go 1.22+ enhanced `http.ServeMux`.
- **JSON Ready**: Built-in JSON response handling.
- **Health Checks**: Included `/health` endpoint for monitoring.

## Getting Started

### Prerequisites

- Go 1.22 or higher

### Running the Server

To start the API server, run:

```bash
go run main.go
```

The server will start on `http://localhost:8080`.

### Endpoints

- `GET /health` - Check service status
- `GET /api/v1/welcome` - Welcome message

## Project Structure

```text
go-api/
├── main.go        # Entry point and routing
├── go.mod         # Module definition
└── README.md      # Project documentation
```
