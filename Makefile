BINARY=run

build:
	go build -o ./bin/admin ./cmd/admin.go
	go build -o ./bin/public ./cmd/public.go
	go build -o ./bin/app_gallery ./cmd/app_gallery.go

package:
	docker build -t ghcr.io/jiramot/go-oauth2:latest -f Dockerfile .

public:
	 go run cmd/public.go

admin:
	 go run cmd/admin.go

test:
	go test -short  ./...

lint:
	golangci-lint run --fix --fast ./...
	
.PHONY: admin