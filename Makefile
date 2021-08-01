NAME = svc-fizzbuzz
VERSION = $(shell cat VERSION)
GO_PACKAGE_BASE = github.com/hugdubois
GO_PACKAGE_NAME = $(GO_PACKAGE_BASE)/$(NAME)

build:
	@echo "$(NAME): build task"
	-mkdir -p _build
	CGO_ENABLED=0 go build \
		-ldflags '-extldflags "-lm -lstdc++ -static"' \
		-ldflags "-X $(GO_PACKAGE_NAME)/service.version=v$(VERSION)" \
		-o _build/$(NAME) \
	main.go

test:
	@echo "$(NAME): test task"
	@go test ./... -v -race

test-cover:
	@echo "$(NAME): test-cover task"
	@go test ./... -cover

test-cover-report:
	@echo "$(NAME): test-cover-report task"
	@mkdir -p _build
	@go test ./... -coverprofile=_build/coverage.out
	@go tool cover -html=_build/coverage.out

serve: build
	@echo "$(NAME): serve task"
	@_build/svc-fizzbuzz serve

clean:
	@echo "$(NAME): clean task"
	@touch svc-fizzbuzz
	@-rm svc-fizzbuzz
	@mkdir -p _build
	@-rm -rf _build

update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
		go get $(GO_PACKAGE_NAME)@v$(VERSION)
