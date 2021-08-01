NAME = svc-fizzbuzz
VERSION = $(shell cat VERSION)

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

update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
		go get github.com/hugdubois/svc-fizzbuzz@v$(VERSION)
