FROM golang:1.23.2-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./

RUN go mod download

RUN chmod -R 755 /app

EXPOSE 5000

CMD ["air", "-c", ".air.toml"]