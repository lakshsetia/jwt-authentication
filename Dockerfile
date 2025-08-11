FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod tidy

EXPOSE "8000"

RUN go build -o app cmd/app/main.go

CMD ["./app"]