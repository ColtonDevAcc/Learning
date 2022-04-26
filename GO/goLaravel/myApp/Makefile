BINARY_NAME=vooApp.exe

## build: builds all binaries
build:
	@go mod vendor 
	@go build -o tmp\${BINARY_NAME} .
	@echo VOO built!

run: build
	@echo Staring Celeritas...
	@.\tmp\${BINARY_NAME} &
	@echo VOO started!

clean:
	@echo Cleaning...
	@go clean
	@del .\tmp\${BINARY_NAME}
	@echo VOO Cleaned!

test:
	@echo Testing...
	@go test ./...
	@echo Done!

start: run
    
stop:
	@echo "Starting the front end..."
	@taskkill \IM ${BINARY_NAME} \F
	@echo Stopped VOO

restart: stop start