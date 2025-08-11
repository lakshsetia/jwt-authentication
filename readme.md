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
