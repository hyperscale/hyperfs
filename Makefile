BUILD_DIR ?= target
COMMIT = $(shell git rev-parse HEAD)
VERSION ?= $(shell git describe --always --tags --dirty)
ORG := github.com/hyperscale
PROJECT := hyperfs
REPOPATH ?= $(ORG)/$(PROJECT)
VERSION_PACKAGE = $(REPOPATH)/pkg/hyperfs/version

GO_LDFLAGS :="
GO_LDFLAGS += -X $(VERSION_PACKAGE).version=$(VERSION)
GO_LDFLAGS += -X $(VERSION_PACKAGE).buildDate=$(shell date +'%Y-%m-%dT%H:%M:%SZ')
GO_LDFLAGS += -X $(VERSION_PACKAGE).gitCommit=$(COMMIT)
GO_LDFLAGS += -X $(VERSION_PACKAGE).gitTreeState=$(if $(shell git status --porcelain),dirty,clean)
GO_LDFLAGS +="

GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: all
all: deps build test

.PHONY: deps
deps:
	@go mod download

.PHONY: clean
clean:
	@go clean -i ./...

$(BUILD_DIR)/coverage.out: $(GO_FILES)
	@go test -cover -coverprofile $(BUILD_DIR)/coverage.out.tmp ./...
	@cat $(BUILD_DIR)/coverage.out.tmp | grep -v '.pb.go' | grep -v 'mock_' > $(BUILD_DIR)/coverage.out
	@rm $(BUILD_DIR)/coverage.out.tmp

ci-test:
	@go test -race -cover -coverprofile ./coverage.out.tmp -v ./... | go2xunit -fail -output tests.xml
	@cat ./coverage.out.tmp | grep -v '.pb.go' | grep -v 'mock_' > ./coverage.out
	@rm ./coverage.out.tmp
	@echo ""
	@go tool cover -func ./coverage.out

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: test
test: $(BUILD_DIR)/coverage.out

.PHONY: coverage
coverage: $(BUILD_DIR)/coverage.out
	@echo ""
	@go tool cover -func ./$(BUILD_DIR)/coverage.out

.PHONY: coverage-html
coverage-html: $(BUILD_DIR)/coverage.out
	@go tool cover -html ./$(BUILD_DIR)/coverage.out


${BUILD_DIR}/hyperfs: $(GO_FILES)
	@echo "Building $@..."
	@go generate ./cmd/$(subst ${BUILD_DIR}/,,$@)/
	@CGO_ENABLED=0 go build -ldflags $(GO_LDFLAGS) -o $@ ./cmd/$(subst ${BUILD_DIR}/,,$@)/

.PHONY: run-hyperfs
run-hyperfs: ${BUILD_DIR}/hyperfs
	@echo "Running $<..."
	@./$<

${BUILD_DIR}/hyperfs-api: $(GO_FILES)
	@echo "Building $@..."
	@go generate ./cmd/$(subst ${BUILD_DIR}/,,$@)/
	@CGO_ENABLED=0 go build -ldflags $(GO_LDFLAGS) -o $@ ./cmd/$(subst ${BUILD_DIR}/,,$@)/

.PHONY: run-hyperfs-api
run-hyperfs-api: ${BUILD_DIR}/hyperfs-api
	@echo "Running $<..."
	@./$<

${BUILD_DIR}/hyperfs-index: $(GO_FILES)
	@echo "Building $@..."
	@go generate ./cmd/$(subst ${BUILD_DIR}/,,$@)/
	@CGO_ENABLED=0 go build -ldflags $(GO_LDFLAGS) -o $@ ./cmd/$(subst ${BUILD_DIR}/,,$@)/

hyperfs-index-docker: $(GO_FILES) ./cmd/hyperfs-index/Dockerfile
	@docker build -f cmd/hyperfs-index/Dockerfile -t 127.0.0.1:5000/hyperfs-index:latest .
	@docker image push 127.0.0.1:5000/hyperfs-index

.PHONY: run-hyperfs-index
run-hyperfs-index: ${BUILD_DIR}/hyperfs-index
	@echo "Running $<..."
	@./$<

${BUILD_DIR}/hyperfs-storage: $(GO_FILES)
	@echo "Building $@..."
	@go generate ./cmd/$(subst ${BUILD_DIR}/,,$@)/
	@CGO_ENABLED=0 go build -ldflags $(GO_LDFLAGS) -o $@ ./cmd/$(subst ${BUILD_DIR}/,,$@)/

.PHONY: run-hyperfs-storage
run-hyperfs-storage: ${BUILD_DIR}/hyperfs-storage
	@echo "Running $<..."
	@./$<

.PHONY: build
build: ${BUILD_DIR}/hyperfs ${BUILD_DIR}/hyperfs-api ${BUILD_DIR}/hyperfs-index ${BUILD_DIR}/hyperfs-storage

swarm: hyperfs-index-docker
	@docker stack deploy --compose-file docker-compose.yml hyperfs-cluster

swarm-update:
	@docker stack deploy --compose-file docker-compose.yml hyperfs-cluster
