# Fiber User API

A user management API built with the Fiber framework using Clean Architecture.

## Features

- Create users (automatically generates UUID7)
- Query all users
- Automatic timestamp management (created_at, updated_at)

## API Endpoints

### Create User

```
POST /users
Content-Type: application/json

{
    "name": "me",
    "email": "test@test.com"
}
```

Response:

```json
{
  "id": "01983c1d-a853-7e84-b42f-3d834cf35e5d",
  "name": "me",
  "email": "test@test.com",
  "created_at": "2025-07-24T19:07:13.3639515+08:00",
  "updated_at": "2025-07-24T19:07:13.3639515+08:00"
}
```

### Get All Users

```
GET /users
```

Response:

```json
[
  {
    "id": "01983c1d-a853-7e84-b42f-3d834cf35e5d",
    "name": "me",
    "email": "test@test.com",
    "created_at": "2025-07-24T19:07:13.3639515+08:00",
    "updated_at": "2025-07-24T19:07:13.3639515+08:00"
  }
]
```

## Project Structure

```
test-fiber/
├── config/
│   └── db.go          # Database configuration
├── internal/
│   ├── domain/
│   │   └── user.go     # User model and interfaces
│   ├── handler/
│   │   └── user_handler.go  # HTTP handlers
│   ├── repository/
│   │   └── user_repo.go     # Data access layer
│   └── usecase/
│       └── user_usecase.go  # Business logic
├── routes/
│   └── routes.go       # Route definitions
├── main.go            # Application entry point
└── README.md
```

## Architecture

This project follows Clean Architecture principles:

- **Domain Layer**: Contains business entities and interfaces
- **Use Case Layer**: Implements business logic
- **Repository Layer**: Handles data persistence
- **Handler Layer**: Manages HTTP requests and responses

## Database

### Current Setup (Development)

Uses SQLite database with GORM's AutoMigrate for development convenience. The database file is `test.db`.

**Note**: AutoMigrate is used here for simplicity during development, but I'm not recommend it for production environments.

### Production Recommendations

For production environments, consider using:

1. **Database**: MySQL or PostgreSQL instead of SQLite

   - Better concurrency handling
   - More robust transaction support
   - Better performance for complex queries
   - Built-in replication and clustering

2. **Migration Tool**: [golang-migrate](https://github.com/golang-migrate/migrate) instead of AutoMigrate
   - Version-controlled schema changes
   - Rollback capabilities
   - Better for team collaboration
   - More reliable in production environments

### Database Schema

The `users` table includes:

- `id` (UUID, Primary Key) - Automatically generated UUID7
- `name` (String) - User's name
- `email` (String) - User's email
- `created_at` (Timestamp) - Automatically managed by GORM
- `updated_at` (Timestamp) - Automatically managed by GORM

## Running the Application

```bash
go run main.go
```

The server will start on `http://localhost:3000`.

## Dependencies

- **Fiber**: Web framework
- **GORM**: ORM for database operations
- **UUID**: For generating unique identifiers
- **SQLite**: Database (development only)

## Development

The application includes:

- Request validation
- Error handling
- Logging middleware
- Automatic database migration (development only)

## Production Considerations

- Replace SQLite with MySQL/PostgreSQL
- Implement proper database migrations using golang-migrate
- Add database connection pooling
- Implement proper error handling and logging
- Add authentication and authorization
- Set up monitoring and health checks
