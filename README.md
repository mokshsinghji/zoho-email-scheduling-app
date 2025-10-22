# zoho-email-scheduling-app

A basic Go application for email scheduling with Zoho integration.

## Prerequisites

- Go 1.16 or higher

## Getting Started

### Build the application

```bash
go build -o zoho-email-scheduling-app
```

### Run the application

```bash
go run main.go
```

Or run the built binary:

```bash
./zoho-email-scheduling-app
```

The server will start on port 8080.

## Endpoints

- `GET /` - Home page
- `GET /health` - Health check endpoint

## Testing

Test the application using curl:

```bash
# Home endpoint
curl http://localhost:8080/

# Health check
curl http://localhost:8080/health
```