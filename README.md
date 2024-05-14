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

## 'GET /': Returns a welcome message.
## '/users' Returns a list of all users.
## '/users/{id}' Returns a specific user based off the ID provided.