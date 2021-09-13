BINARY=run

public:
	 go run cmd/public.go

admin:
	 go run cmd/admin.go

build:
	go build -o bin/engine cmd/main.go

test:	
	go test -short  ./...

lint:
	golangci-lint run --fix --fast ./...
	
.PHONY: admin