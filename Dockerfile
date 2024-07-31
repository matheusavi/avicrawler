# syntax=docker/dockerfile:1
# A simple webcrawler, which stores the data in a PostgreSQL database

FROM golang:1.22.5 AS api

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api/main.go

EXPOSE 3457

CMD ["/api"]

FROM api AS migration
RUN CGO_ENABLED=0 GOOS=linux go build -o /migrate ./cmd/migrate/main.go

CMD ["/migrate"]