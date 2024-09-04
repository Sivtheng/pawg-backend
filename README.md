# Setup for Golang API for Pawg

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

- Create User: http://localhost:8080/users
> curl -X POST "Endpoint" \ -H "Content-Type: application/json" \ -d '{"name": "John Doe", "password": "securepassword"}'
- Get User by ID: http://localhost:8080/users/{id}
> curl -X GET "Endpoint"
- Update User by ID http://localhost:8080/users/{id}
> curl -X PUT "Endpoint" \ -H "Content-Type: application/json" \ -d '{"name": "John Doe", "password": "newpassword"}'
- Delete User by ID: http://localhost:8080/users/{id}
> curl -X DELETE "Endpoint"

### Get In Touch

- Create Contact: http://localhost:8080/get_in_touch
> curl -X POST "Endpoint"h \ -H "Content-Type: application/json" \ -d '{"name": "Jane Doe", "email": "jane@example.com", "message": "Hello there!"}'
- Get Contact by ID: http://localhost:8080/get_in_touch/{id}
> curl -X GET "Endpoint"
- Update Contact by ID: http://localhost:8080/get_in_touch/{id}
> curl -X PUT "Endpoint"} \ -H "Content-Type: application/json" \ -d '{"name": "Jane Doe", "email": "jane.doe@example.com", "message": "Updated message"}'
- Delete Contact by ID: http://localhost:8080/get_in_touch/{id}
> curl -X DELETE "Endpoint"

### Appointments

- Create Appointments: http://localhost:8080/appointments
> curl -X POST "Endpoint" \ -H "Content-Type: application/json" \ -d '{"name": "John Doe", "email": "john@example.com", "phone_number": "555-1234", "appointment_date": "2024-09-10", "appointment_time": "14:00:00"}'
- Get Appointment by ID: http://localhost:8080/appointments/{id}
> curl -X GET "Endpoint"
- Update Appointment by ID: http://localhost:8080/appointments/{id}
> curl -X PUT "Endpoint" \ -H "Content-Type: application/json" \ -d '{"name": "Jane Doe", "email": "jane@example.com", "phone_number": "555-5678", "appointment_date": "2024-09-11", "appointment_time": "15:00:00"}'
- Delete Appointment by ID: http://localhost:8080/appointments/{id}
> curl -X DELETE "Endpoint"

### Adoption Application

- Create Adoption Application: http://localhost:8080/adoption_applications
> curl -X POST "Endpoint" \ -H "Content-Type: application/json" \ -d '{ "name": "Alice Smith", "email": "alice@example.com", "phone_number": "555-9876", "address": "123 Elm Street", "interest_in_adopting": "Dog", "type_of_animal": "Dog", "special_needs_animal": "No", "own_pet_before": "Yes", "working_time": "9am - 5pm", "living_situation": "House", "other_animals": "None", "animal_access": "Indoor", "travel": "Yes", "leave_cambodia": "No", "feed": "Twice a day", "anything_else": "N/A"}'
- Get Adoption Application by ID: http://localhost:8080/adoption_applications/{id}
> curl -X GET "Endpoint"
- Update Adoption Application by ID: http://localhost:8080/adoption_applications/{id}
> curl -X PUT "Endpoint" \ -H "Content-Type: application/json" \ -d '{"name": "Alice Johnson", "email": "alice.johnson@example.com", "phone_number": "555-4321", "address": "456 Oak Avenue", "interest_in_adopting": "Cat", "type_of_animal": "Cat", "special_needs_animal": "No", "own_pet_before": "No", "working_time": "8am - 4pm", "living_situation": "Apartment", "other_animals": "One dog", "animal_access": "Indoor and outdoor", "travel": "No", "leave_cambodia": "Yes", "feed": "Once a day", "anything_else": "Looking forward to adopting"}'
- Delete Adoption Application by ID: http://localhost:8080/adoption_applications/{id}
> curl -X DELETE "Endpoint"

## Authentication

JWT Token: The API uses JWT tokens for authentication. Tokens are issued upon successful login and must be included in the Authorization header for protected routes.
To access protected routes, first log in to receive a JWT token:
- Login: http://localhost:8080/login
>    curl -X POST "Endpoint" \ -H "Content-Type: application/json" \ -d '{"name": "John Doe", "password": "securepassword"}'
- Test 
> curl -X GET http://localhost:8080/api/users \ -H "Authorization: Bearer JWT_TOKEN"
