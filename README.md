# Golang API for Pawg ver: 1.0.0-alpha

## Prerequisites

Before setting up the project, ensure that the following are installed:
1. Golang
2. Postgresql


## Running the API (Set up everything first)

1. Install dependencies
> go mod tidy
2. Run the server
> go run main.go

## Set up the database

1. Access Postgresql (Assuming you are admin user on your pc)
> psql postgres
2. Create user
> CREATE ROLE myuser WITH SUPERUSER CREATEDB CREATEROLE LOGIN PASSWORD 'mypassword';
3. Create a database
> CREATE DATABASE mydb;
4. Grand all privileges
> GRANT ALL PRIVILEGES ON DATABASE mydb TO myuser;
5. Verify
> \du and \l
6. Exit
> \q

## Create schema with golang-migrate

Since I already have the sql schemas, you only have to apply the migration.
1. Install golang-migrate CLI
> brew install golang-migrate
2. Create migration files
> migrate create -ext sql -dir directory_here schema_name_here
3. Edit the generate sql files to include your schema
4. Apply migrations
> migrate -path directory_here -database "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable" up
5. If need to roll back migrations
> migrate -path directory_here -database "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable" down

## API Endpoints

Here are the available API endpoints and their usage:

### Users

- List and Create User: http://localhost:8080/api/users
- Get User by ID: http://localhost:8080/api/users/{id}\
- Update User by ID http://localhost:8080/api/users/{id}
- Delete User by ID: http://localhost:8080/api/users/{id}

### Get In Touch

- List and Create Contact: http://localhost:8080/api/get_in_touch
- Get Contact by ID: http://localhost:8080/api/get_in_touch/{id}
- Update Contact by ID: http://localhost:8080/api/get_in_touch/{id}
- Delete Contact by ID: http://localhost:8080/api/get_in_touch/{id}

### Appointments

- List and Create Appointments: http://localhost:8080/api/appointments
- Get Appointment by ID: http://localhost:8080/api/appointments/{id}
- Update Appointment by ID: http://localhost:8080/api/appointments/{id}
- Delete Appointment by ID: http://localhost:8080/api/appointments/{id}

### Adoption Application

- List and Create Adoption Application: http://localhost:8080/api/adoption_applications
- Get Adoption Application by ID: http://localhost:8080/api/adoption_applications/{id}
- Update Adoption Application by ID: http://localhost:8080/api/adoption_applications/{id}
- Delete Adoption Application by ID: http://localhost:8080/api/adoption_applications/{id}

## Authentication

JWT Token: The API uses JWT tokens for authentication. Tokens are issued upon successful login and must be included in the Authorization header for protected routes.
To access protected routes, first log in to receive a JWT token:
- Login: http://localhost:8080/login

## Possible features and improvements

1. Pagination and Filtering
- Implement pagination for endpoints that return a list of resources
- Allow filtering of results based on query parameters (ex: filter users by creation date, etc)
2. Validation
- Add validation for input data to ensure it meets the expected format and constraints (ex: validate email format, required fields). This also probably can be done in frontend too.
- Custom Validation Messages: Provide user-friendly error messages for validation failures.
3. Rate Limiting
- Implement rate limiting to prevent abuse of the API and manage traffic effectively
4. Logging
- Request/Response Logging: Log API requests and responses for debugging and monitoring.
- Structured Logging: Use a structured logging library to log data in a structured format (e.g., JSON) for easier analysis.
5. Better API Documentation
- Swagger/OpenAPI: Use Swagger or OpenAPI to generate interactive API documentation for easier understanding and testing of your endpoints.
6. Error Handling
- Custom Error Responses: Create a standardized format for error responses to ensure consistency across the API.
- Error Logging: Log errors to track issues and improve debugging.
7. Testing
- Unit Tests: Write comprehensive unit tests for your handlers, models, and middleware.
- Integration Tests: Implement integration tests to verify the interactions between different components of the system.
- CI/CD Pipeline: Set up a Continuous Integration/Continuous Deployment (CI/CD) pipeline to automate testing and deployment processes.
8. Better Security
- HTTPS: Configure your API to use HTTPS to secure data transmission.
- Input Sanitization: Ensure all inputs are sanitized to prevent security vulnerabilities like SQL injection or XSS.
9. Caching 
- In-Memory Caching: Use an in-memory cache (e.g., Redis) to improve performance for frequently accessed data.
- HTTP Caching: Implement HTTP caching mechanisms to reduce server load and improve response times.
10. Background Jobs
- Asynchronous Processing: Implement background job processing for tasks that can be performed asynchronously, such as sending emails