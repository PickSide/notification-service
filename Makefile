run: build
	@./bin/notification-service

build:
	@echo "Building Notification service..."
	@go build -o ./bin/notification-service ./cmd/main.go
	@echo "Done"