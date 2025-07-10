# Student Management REST API

A simple REST API built with Go that performs CRUD operations on student data and integrates with Ollama for AI-powered student profile summaries.

## Features

- **CRUD Operations**: Create, Read, Update, Delete students
- **AI Integration**: Generate student profile summaries using Ollama with Llama3
- **Concurrent Safe**: Thread-safe operations with mutex locks
- **Input Validation**: Comprehensive validation for student data
- **Error Handling**: Proper HTTP status codes and error messages
- **Unit Tests**: Comprehensive test coverage for all components

## Project Structure

```
fealtyx_assignment/
├── cmd/
│   └── server/
│       ├── main.go           # Application entry point
│       └── main_test.go      # Integration tests
├── internal/
│   ├── handlers/
│   │   ├── student.go        # Student HTTP handlers
│   │   ├── student_test.go   # Student handler tests
│   │   ├── ollama.go         # Ollama HTTP handlers
│   │   └── ollama_test.go    # Ollama handler tests
│   ├── models/
│   │   ├── student.go        # Student data model
│   │   └── student_test.go   # Model validation tests
│   ├── services/
│   │   ├── student.go        # Student business logic
│   │   ├── student_test.go   # Service layer tests
│   │   ├── ollama.go         # Ollama service integration
│   │   └── ollama_test.go    # Ollama service tests
│   └── middleware/
│       └── cors.go           # CORS middleware
├── pkg/
│   └── utils/
│       └── response.go       # HTTP response utilities
├── go.mod                    # Go module definition
├── go.sum                    # Go module checksums
└── README.md                 # This file
```

## Prerequisites

- Go 1.19 or higher
- Ollama installed and running
- Llama3 model pulled in Ollama

## Setup Instructions

### 1. Install Ollama

```bash
# Install Ollama
curl -fsSL https://ollama.ai/install.sh | sh
```

### 2. Start Ollama and Install Llama3

```bash
# Start Ollama server (if not already running)
ollama serve

# In another terminal, pull the Llama3 model
ollama pull llama3

# Verify installation
ollama list
```

### 3. Clone and Setup the Project

```bash
# Navigate to your project directory
cd fealtyx_assignment

# Initialize Go module (if not done)
go mod init student-api

# Install dependencies
go mod tidy
```

## Running the API

### Start the Server

```bash
# Run the server
go run cmd/server/main.go
```

The server will start on `http://localhost:8080`

You should see the output:
```
2024/07/10 12:00:00 Server starting on :8080
```

### Verify Setup

Test if everything is working:

```bash
# Test the API
curl http://localhost:8080/students

# Test Ollama connection
curl http://localhost:11434/api/version
```

## Running Tests

### Run All Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage report
go test -cover ./...

# Run tests with detailed coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Run Specific Test Files

```bash
# Test student handlers
go test ./internal/handlers/student_test.go

# Test student services
go test ./internal/services/student_test.go

# Test Ollama integration
go test ./internal/handlers/ollama_test.go

# Test models
go test ./internal/models/student_test.go
```

## API Endpoints

### Base URL
```
http://localhost:8080
```

### 1. Create Student
- **Method**: `POST`
- **Endpoint**: `/students`
- **Content-Type**: `application/json`

**Request Body**:
```json
{
    "name": "John Doe",
    "age": 20,
    "email": "john.doe@example.com"
}
```

**Success Response** (201 Created):
```json
{
    "id": 1,
    "name": "John Doe",
    "age": 20,
    "email": "john.doe@example.com"
}
```

**Error Responses**:
- `400 Bad Request`: Invalid JSON or missing required fields
- `400 Bad Request`: Invalid email format or age <= 0

### 2. Get All Students
- **Method**: `GET`
- **Endpoint**: `/students`

**Success Response** (200 OK):
```json
[
    {
        "id": 1,
        "name": "John Doe",
        "age": 20,
        "email": "john.doe@example.com"
    },
    {
        "id": 2,
        "name": "Jane Smith",
        "age": 22,
        "email": "jane.smith@example.com"
    }
]
```

### 3. Get Student by ID
- **Method**: `GET`
- **Endpoint**: `/students/{id}`

**Success Response** (200 OK):
```json
{
    "id": 1,
    "name": "John Doe",
    "age": 20,
    "email": "john.doe@example.com"
}
```

**Error Responses**:
- `400 Bad Request`: Invalid ID format
- `404 Not Found`: Student not found

### 4. Update Student
- **Method**: `PUT`
- **Endpoint**: `/students/{id}`
- **Content-Type**: `application/json`

**Request Body**:
```json
{
    "name": "John Updated",
    "age": 25,
    "email": "john.updated@example.com"
}
```

**Success Response** (200 OK):
```json
{
    "id": 1,
    "name": "John Updated",
    "age": 25,
    "email": "john.updated@example.com"
}
```

**Error Responses**:
- `400 Bad Request`: Invalid JSON, ID format, or student data
- `404 Not Found`: Student not found

### 5. Delete Student
- **Method**: `DELETE`
- **Endpoint**: `/students/{id}`

**Success Response** (204 No Content):
- Empty response body

**Error Responses**:
- `400 Bad Request`: Invalid ID format
- `404 Not Found`: Student not found

### 6. Generate Student Summary (AI-Powered)
- **Method**: `GET`
- **Endpoint**: `/students/{id}/summary`

**Success Response** (200 OK):
```json
{
    "summary": "John Updated is a motivated and ambitious individual with a strong academic background. At 25 years old, he brings a youthful energy and enthusiasm to his studies, with a keen eye for detail and a commitment to excellence. With a focus on personal growth and development, John is poised to make a positive impact in his chosen field."
}
```

**Error Responses**:
- `400 Bad Request`: Invalid student ID format
- `404 Not Found`: Student not found
- `500 Internal Server Error`: Ollama service error

## Sample API Usage

### Complete Workflow Example

```bash
# 1. Create a student
curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Johnson","age":21,"email":"alice@university.edu"}'

# 2. Get all students
curl http://localhost:8080/students

# 3. Get specific student
curl http://localhost:8080/students/1

# 4. Update student
curl -X PUT http://localhost:8080/students/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Johnson","age":22,"email":"alice.johnson@university.edu"}'

# 5. Generate AI summary
curl http://localhost:8080/students/1/summary

# 6. Delete student
curl -X DELETE http://localhost:8080/students/1

# 7. Verify deletion
curl http://localhost:8080/students
```

## Error Handling

The API returns appropriate HTTP status codes:

- `200 OK`: Successful GET/PUT requests
- `201 Created`: Successful POST requests
- `204 No Content`: Successful DELETE requests
- `400 Bad Request`: Invalid input data or malformed requests
- `404 Not Found`: Resource not found
- `405 Method Not Allowed`: Unsupported HTTP method
- `500 Internal Server Error`: Server-side errors

## Data Validation

### Student Model Validation Rules:
- **Name**: Required, non-empty string
- **Age**: Required, positive integer (> 0)
- **Email**: Required, valid email format (regex validated)

### Example Validation Errors:
```bash
# Missing name
curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -d '{"age":20,"email":"test@example.com"}'
# Response: 400 Bad Request - Invalid student data

# Invalid email
curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","age":20,"email":"invalid-email"}'
# Response: 400 Bad Request - Invalid student data

# Invalid age
curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -d '{"name":"Test","age":-5,"email":"test@example.com"}'
# Response: 400 Bad Request - Invalid student data
```

## Concurrency Safety

The API is designed to handle concurrent requests safely:
- Thread-safe operations using `sync.RWMutex`
- Separate read and write locks for optimal performance
- Atomic operations for ID generation

## Ollama Integration

The API integrates with Ollama to provide AI-generated student summaries:

### Configuration:
- **Ollama URL**: `http://localhost:11434`
- **Model**: `llama3`
- **Endpoint**: `/api/generate`

### Features:
- Automated prompt engineering for student profile summaries
- Response cleaning (removes escape characters and formatting)
- Error handling for Ollama service failures
- Non-streaming responses for consistent output

## Testing

The project includes comprehensive tests:

### Test Coverage:
- **Unit Tests**: All service functions and models
- **Integration Tests**: HTTP handlers and endpoints
- **Mock Tests**: Ollama service integration
- **Concurrency Tests**: Thread safety validation
- **Validation Tests**: Input validation and error handling

### Test Files:
- `internal/handlers/student_test.go` - Handler tests
- `internal/handlers/ollama_test.go` - Ollama handler tests
- `internal/services/student_test.go` - Service layer tests
- `internal/services/ollama_test.go` - Ollama service tests
- `internal/models/student_test.go` - Model validation tests
- `cmd/server/main_test.go` - Integration tests

## Troubleshooting

### Common Issues:

1. **"Address already in use" error**:
   ```bash
   # Check if Ollama is already running
   ps aux | grep ollama
   # If running, you don't need to start it again
   ```

2. **"Student not found" for summary endpoint**:
   - Ensure the student exists before generating summary
   - Check the correct student ID

3. **Ollama connection errors**:
   ```bash
   # Verify Ollama is running
   curl http://localhost:11434/api/version
   
   # Check if llama3 is installed
   ollama list
   ```

4. **Test failures**:
   ```bash
   # Clean and rebuild
   go clean -testcache
   go test ./...
   ```