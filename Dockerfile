FROM golang:1.21.6-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/notification-service ./cmd/main.go

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /bin/notification-service /app/

EXPOSE 8084

CMD [ "/app/notification-service" ]
