FROM golang:1.21.6-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/notification-service ./cmd/main.go

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /bin/notification-service /app/

ARG VERSION

ENV SERVICE_NAME="notification-service" \
    SERVICE_VERSION=${VERSION}

EXPOSE 8084

CMD [ "/app/notification-service" ]
