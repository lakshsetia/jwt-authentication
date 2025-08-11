# jwt-authentication

Project for learning JWT authentication

## Project Structure

```
jwt-authentication
│   .env
│   compose.yml
│   Dockerfile
│   go.mod
│   go.sum
├───cmd
│   └───app
│           main.go
├───config
│       local.yml
└───internal
    ├───config
    │       config.go
    ├───db
    │   │   db.go
    │   └───pg
    │           pg.go
    ├───handlers
    │       auth_handler.go
    │       user_handler.go
    ├───jwt
    │       jwt.go
    ├───models
    │       login.go
    │       user.go
    ├───server
    │       routes.go
    │       server.go
    └───utils
        ├───json
        │       json.go
        ├───password
        │       password.go
        └───response
                error.go
                user.go
```

## Project setup

1. Declare `config/local.yml` file:

   ```
   env: "dev"
   http_server:
   address: ":8000"
   database:
   pg:
       user: "postgres"
       password: "secret"
       dbname: "db"
       port: "5432"
       host: "db"
   key:
       hmac_key: "secret key"
   ```

2. Declare `.env` file:
   ```
   CONFIG_PATH=config/local.yml
   POSTGRES_USER=postgres
   POSTGRES_PASSWORD=secret
   POSTGRES_DB=db
   ```

## Run project

Run `docker compose up --build -d`
