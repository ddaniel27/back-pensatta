FROM golang:1.20-bullseye AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download -x

COPY . .
CMD ["go", "run", "."]
