NAME = svc-fizzbuzz
VERSION = $(shell cat VERSION)

GO_PACKAGE_BASE = github.com/hugdubois
GO_PACKAGE_NAME = $(GO_PACKAGE_BASE)/$(NAME)

DOCKER_TAG = $(shell cat VERSION | tr +- __)
DOCKER_IMAGE_NAME?=hugdubois/$(NAME)
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

version:
	@echo $(VERSION)

tools:
	@echo "$(NAME): tools task"
	@cd _tools && make

test:
	@echo "$(NAME): test task"
	@go test ./... -race

test-v:
	@echo "$(NAME): test-v task"
	@go test ./... -race -v

test-live:
	@echo "$(NAME): test-live task"
	@while true; do $(MAKE) test; echo "\nPress [CTRL+C] to exit this loop..." ; sleep 5; done

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

gen-swagger: tools
	@echo "$(NAME): gen-swagger task"
	@_tools/bin/swag init service/service.go

clean:
	@echo "$(NAME): clean task"
	@touch svc-fizzbuzz
	@-rm svc-fizzbuzz
	@mkdir -p _tools/bin
	@-rm -rf _tools/bin
	@mkdir -p _build
	@-rm -rf _build
	@touch .env
	@-rm .env
	@touch coverage.txt
	@-rm coverage.txt
	@touch coverage.out
	@-rm coverage.out
	@touch dump.rdb
	@-rm dump.rdb

docker-tag:
	@echo "$(NAME): docker-tag task"
	@echo "TAG=$(DOCKER_TAG)" > .env

docker: test docker-tag
	@echo "$(NAME): docker task"
	@docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) .

docker-push: docker
	@echo "$(NAME): docker-push task"
	@docker tag $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_TAG)
	@docker push $(DOCKER_REGISTRY)/$(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

docker-run: docker
	@echo "$(NAME): docker-run task"
	@docker run -d --name=svc-fizzbuzz -p $(DOCKER_RUN_PORT):8080 -it $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

docker-rm:
	@echo "$(NAME): docker-rm task"
	-docker stop svc-fizzbuzz
	-docker rm svc-fizzbuzz

compose-up: docker-tag
	@echo "$(NAME): compose-up task"
	@docker-compose up -d

compose-down:
	@echo "$(NAME): compose-down task"
	@docker-compose down

compose-ps:
	@echo "$(NAME): compose-down task"
	@docker-compose ps

k8s-deploy:
	@echo "$(NAME): k8s-deploy task"
	@kubectl apply -f k8s-deployment.yaml
	@echo ''
	@echo 'Note (on a local environment) do :'
	@echo ''
	@echo '    minikube service svc-fizzbuzz'
	@echo ''

k8s-delete:
	@echo "$(NAME): k8s-delete task"
	@kubectl delete -f k8s-deployment.yaml

update-pkg-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on \
		go get $(GO_PACKAGE_NAME)@v$(VERSION)
