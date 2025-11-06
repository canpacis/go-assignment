# Greeter Service

A simple HTTP service that greets users based on their name with specific validation rules.

## Prerequisites

- Go 1.24.5 or higher
- [Task](https://taskfile.dev) (optional, for using Taskfile commands)

## Running the Application

### Using Task (recommended)

```bash
task dev
```

### Using Go directly

```bash
go run ./cmd
```

The server will start on `http://localhost:8080`

## Running Tests

### Using Task

```bash
task test
```

### Using Go directly

```bash
go test -v ./...
```

## API Usage

### Endpoint

```
GET /hello-world?name=<name>
```

### Example Request

```bash
curl "http://localhost:8080/hello-world?name=Alice"
```

### Example Response (Success)

```json
{
  "message": "Hello Alice"
}
```

### Example Response (Error)

```json
{
  "error": "Invalid Input"
}
```

## Validation Rules

The service validates input names with the following rules

1. **Name cannot be empty** - The name parameter must contain at least one non-whiespace character
2. **Name must start with A-M** - The first letter of the name must be between A-M (case insensitive)

Names starting with letters N-Z, non-letter characters or non English characters will be rejected with a 400 Bad Request status.

## Assumptions

- The applications runs on port 8080 by default (hardcoded)
- Case is preserved in the greeting message (e.g., "alice" returns "Hello alice")
- The validation is character-based, not letter-based, so special characters at the start (e.g, @Alice, 'Alice' etc.) will fail validation
- No authentication or rate limiting is implemented
- No logging middleware is included in this minimal implementation