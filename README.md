# CRUD with fiber

## Installation

```bash
go get github.com/gofiber/fiber
```

## Usage

To serve a the server, run the following command:

```bash
```go
go run main.go
```

Sever will be running on port 8000

## API Documentation

### Create a new user

```bash 
curl --request POST \
  --url http://127.0.0.1:8000/users \
  --header 'Content-Type: application/json' \
  --data '{
	"usename": "Mg Mg",
	"age": 30
}'
```

### Get all users

```bash
curl --request GET \
  --url http://127.0.0.1:8000/users
```

### Get a user

```bash
curl --request POST \
  --url http://127.0.0.1:8000/users \
  --header 'Content-Type: application/json' \
  --data '{
	"usename": "Mg Mg",
	"age": 30
}'
``` 

## Update a user

``` bash
curl --request POST \
  --url http://127.0.0.1:8000/users \
  --header 'Content-Type: application/json' \
  --data '{
	"usename": "Mg Mg",
	"age": 30
}'
```

## Delete a user

```bash
curl --request DELETE \
  --url http://127.0.0.1:8000/users/63b88848d11b8d9a12868c71 \
  --header 'Content-Type: application/json' \
  --data '{
	"usename": "Mg Mg",
	"age": 1002
}'
```