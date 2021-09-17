BINARY=run

public:
	 go run cmd/public.go

admin:
	 go run cmd/admin.go

build-admin:
	docker build -t ghcr.io/jiramot/go-oauth2/admin:latest -f Dockerfile.admin . --no-cache

build-public:
	docker build -t ghcr.io/jiramot/go-oauth2/public:latest -f Dockerfile.public . --no-cache

push-admin:
	docker push ghcr.io/jiramot/go-oauth2/admin:latest

push-public:
	docker push ghcr.io/jiramot/go-oauth2/public:latest

test:	
	go test -short  ./...

lint:
	golangci-lint run --fix --fast ./...
	
.PHONY: admin