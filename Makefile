NAME = svc-fizzbuzz
VERSION = $(shell cat VERSION)
GO_PACKAGE_BASE = github.com/hugdubois
GO_PACKAGE_NAME = $(GO_PACKAGE_BASE)/$(NAME)

build:
	@echo "$(NAME): build task"
	-mkdir -p _build
	CGO_ENABLED=0 go build \
		-ldflags '-extldflags "-lm -lstdc++ -static"' \
		-ldflags "-X $(GO_PACKAGE_NAME)/service.version=$(VERSION)" \
		-o _build/$(NAME) \
	main.go

test:
	@echo "$(NAME): test task"
	@go test ./...

update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
		go get $(GO_PACKAGE_NAME)@v$(VERSION)
