NAME = svc-fizzbuzz

build:
	@echo "$(NAME): build task"
	-mkdir -p _build
	CGO_ENABLED=0 go build \
		-ldflags '-extldflags "-lm -lstdc++ -static"' \
		-o _build/$(NAME) \
	main.go

test:
	@echo "$(NAME): test task"
	@go test ./...
