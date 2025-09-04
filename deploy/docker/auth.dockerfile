FROM golang:1.24 AS builder

# Word directory inside Docker's image
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY pkg ./pkg
COPY internal/auth ./internal/auth/
COPY cmd/auth-service ./cmd/auth-service/

RUN go build -v -o auth-service ./cmd/auth-service/main.go
#CMD ["./auth-service"]

FROM debian

WORKDIR /app
# TEMP
COPY configs/auth/local.yaml ./configs/auth/local.yaml
ENV AUTH_CONFIG_PATH="configs/auth/local.yaml"

COPY --from=builder /usr/src/app/auth-service .
CMD ["./auth-service"]
