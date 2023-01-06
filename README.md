# CRUD with fiber

## Installation

```bash
git clone git@github.com:amm834/fiber-crud-users.git

```

## Usage

To serve a the server, run the following command:

```bash
go run main.go

```

Sever will be running on port 8000

## Hot Reload

To enable hot reload, run the following command:

```bash
go install github.com/cosmtrek/air
```

```bash
air
```

Above command will start the server on port 8000 and will reload the server on any changes in the code.

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