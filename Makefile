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
