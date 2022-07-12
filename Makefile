ifeq ($(PREFIX),)
	PREFIX := /usr/local
endif

.PHONY: build
build:
	go build -o togolist cmd/server/togolist/main.go

.PHONY: clean
clean:
	rm -f togolist

.PHONY: dev
dev:
	go run cmd/server/togolist/main.go

PHONY: test
test:
	go test ./...

.PHONY: docs
docs:
	swag fmt -d ./internal/pkg/router/,./internal/pkg/list/,./internal/pkg/healthz/ -g swagger.go
	swag init -d ./internal/pkg/router/,./internal/pkg/list/,./internal/pkg/healthz/ -g swagger.go

.PHONY: protoc
protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/pkg/list/grpc/list_service.proto
