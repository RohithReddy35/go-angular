# go-angular
a small web application

# User Management Backend

This is the backend implementation of the User Management system built using Go, Echo framework, and PostgreSQL. It provides APIs to perform CRUD operations on user data.

---

## Features

- View all users
- Create a new user
- Update an existing user
- Delete a user
- Includes validation for required fields and email structure
- Handles unique username validation

---

## Prerequisites

1. Install [Go 1.22](https://golang.org/dl/).
2. Install [PostgreSQL](https://www.postgresql.org/download/).
3. Install [Docker](https://docs.docker.com/get-docker/) (optional, for running PostgreSQL in a container).

---

## Setup Instructions

### 1. Clone the Repository
```bash
git clone <repository-url>
cd <repository-folder>
```

### 2. Install Dependencies
Navigate to the backend folder and run:
```bash
go mod tidy
```

### 3. Database Setup
#### Option 1: Using Docker
Run PostgreSQL using Docker:
```bash
docker run --name usermgmt-db -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=usermanagement -p 5432:5432 -d postgres
```

#### Option 2: Local Installation
1. Create a database named `usermanagement`.
2. Create a user with username `admin` and password `admin`.
3. Grant all privileges to the `admin` user on the `usermanagement` database.

### 4. Configure Environment Variables
Create a `.env` file in the project root and add the following:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=usermanagement
DB_SSLMODE=disable
```

### 5. Run Database Migrations
If migrations are included in the project:
```bash
migrate -path ./migrations -database "postgres://admin:admin@localhost:5432/usermanagement?sslmode=disable" up
```

### 6. Start the Server
Run the backend server:
```bash
go run main.go
```
The server will be available at `http://localhost:8080`.

---

## API Endpoints

### Base URL
`http://localhost:8080`

### Endpoints

#### 1. Get All Users
```http
GET /users
```
Response:
```json
[
  {
    "id": 1,
    "user_name": "JohnDoe",
    "email": "john.doe@example.com"
  }
]
```

#### 2. Create a User
```http
POST /users
```
Request Body:
```json
{
  "user_name": "JaneDoe",
  "email": "jane.doe@example.com"
}
```
Response:
```json
{
  "id": 2,
  "user_name": "JaneDoe",
  "email": "jane.doe@example.com"
}
```

#### 3. Update a User
```http
PUT /users/{id}
```
Request Body:
```json
{
  "user_name": "JaneDoeUpdated",
  "email": "jane.updated@example.com"
}
```
Response:
```json
{
  "id": 2,
  "user_name": "JaneDoeUpdated",
  "email": "jane.updated@example.com"
}
```

#### 4. Delete a User
```http
DELETE /users/{id}
```
Response:
```json
{
  "message": "User deleted successfully"
}
```

---

## Testing

### Unit Tests
Run unit tests using [Ginkgo](https://github.com/onsi/ginkgo):
```bash
ginkgo ./...
```

### Swagger Documentation
1. Generate the Swagger documentation if not already included:
   ```bash
   swag init -g main.go
   ```
2. View Swagger UI:
   Open `http://localhost:8080/swagger/index.html` in your browser.

---

## Frontend Integration
Refer to the [Frontend Repository](#) for instructions on how to integrate this backend with the Angular application.

---

## Troubleshooting

### CORS Issues
If you encounter CORS issues, ensure that the Echo server includes the following middleware:
```go
import (
  "github.com/labstack/echo/v4/middleware"
)

e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
  AllowOrigins: []string{"http://localhost:4200"},
  AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
}))
```

---

## Contributing
1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Push to the branch.
5. Create a Pull Request.

---

## License
This project is licensed under the MIT License.

