# syntax=docker/dockerfile:1
# A simple webcrawler, which stores the data in a PostgreSQL database

FROM golang:1.22.5

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api/main.go

EXPOSE 3457

CMD ["/api"]