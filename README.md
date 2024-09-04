# Setup for Golang API for Pawg

## Prerequisites

Before setting up the project, ensure that the following are installed:
1. Golang
2. Postgresql

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
> \du
> \l
6. Exit
> \q

## Create schema with golang-migrate\

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




