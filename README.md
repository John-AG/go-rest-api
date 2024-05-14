# go-rest-api

Currently this is a simple REST API built with Go, which I am using for personal development.

## Setup

1. Clone the repository:

```sh
git clone https://github.com/John-AG/go-rest-api.git
cd go-rest-api
```

2. Install dependencies:
``` sh
go mod tidy
```

3. Run the application:
```sh
go run main.go
```

4. The server will start on 'localhost:8080'. You can test it by navigating to 'http://localhost:8080/'. I'd recommend using Postman for this.

## Prerequisites

Go 1.16+
PostgreSQL
'golang-migrate' for database migrations

## Environment Variables

Create a '.env' file in the root of your project with the following content:

DATABASE_URL=postgres://username:password@hostname:port/database?sslmode=disable

Replace 'username', 'password', 'hostname', 'port' and 'database>' with your actual PostgreSQL credentials and database details.

## Endpoints

'GET /': Returns a welcome message.

'GET /users' Returns a list of all users.

'GET /users/{id}' Returns a specific user based off the ID provided.

'POST /users' Create a new user.

'PUT /users/{id}' Update an existing user by ID.

'DELETE /users/{id}' Delete a user by ID.