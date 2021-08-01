NAME = svc-fizzbuzz
VERSION = $(shell cat VERSION)

GO_PACKAGE_BASE = github.com/hugdubois
GO_PACKAGE_NAME = $(GO_PACKAGE_BASE)/$(NAME)

DOCKER_TAG = $(shell cat VERSION | tr +- __)
DOCKER_IMAGE_NAME = hugdubois/$(NAME)
DOCKER_REGISTRY?=docker.io
DOCKER_RUN_PORT?=8080

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

test-cover-profile:
	@echo "$(NAME): test-cover-profile task"
	@mkdir -p _build
	@go test ./... -coverprofile=_build/coverage.out

test-cover-report: test-cover-profile
	@echo "$(NAME): test-cover-report task"
	@go tool cover -html=_build/coverage.out

test-cover-func: test-cover-profile
	@echo "$(NAME): test-cover-total task"
	@go tool cover -func=_build/coverage.out

serve: build
	@echo "$(NAME): serve task"
	@_build/svc-fizzbuzz serve -d

clean:
	@echo "$(NAME): clean task"
	@touch svc-fizzbuzz
	@-rm svc-fizzbuzz
	@mkdir -p _build
	@-rm -rf _build

docker:
	@echo "$(NAME): docker task"
	@echo "TAG=$(DOCKER_TAG)" > .env
	@docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) .

docker-push: docker
	@echo "$(NAME): docker-push task"
	@docker tag $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_TAG)
	@docker push $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

docker-run: docker
	@echo "$(NAME): docker-run task"
	@docker run -d --name=svc-fizzbuzz -p $(DOCKER_RUN_PORT):13000 -it $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

docker-rm:
	@echo "$(NAME): docker-rm task"
	-docker stop svc-fizzbuzz
	-docker rm svc-fizzbuzz

update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
		go get $(GO_PACKAGE_NAME)@v$(VERSION)
