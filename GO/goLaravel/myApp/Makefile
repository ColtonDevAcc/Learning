BINARY_NAME=VooApp

build:
	@go mod vendor
	@echo "Building Voo..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "Voo built!"

run: build
	@echo "Starting Voo..."
	@./tmp/${BINARY_NAME} &
	@echo "Voo started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping Voo..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped Voo!"

restart: stop start