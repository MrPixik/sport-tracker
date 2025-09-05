FROM golang:1.24 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY pkg ./pkg
COPY internal/auth ./internal/auth/
COPY cmd/auth-service ./cmd/auth-service/

RUN go build -v -o auth-service ./cmd/auth-service/main.go

FROM debian

WORKDIR /app

COPY --from=builder /usr/src/app/auth-service .
CMD ["./auth-service"]
