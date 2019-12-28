BINARY = gsd
COMMIT := $(shell git rev-parse HEAD)
BRANCH := $(shell git symbolic-ref --short -q HEAD || echo HEAD)
DATE := $(shell date -u +%Y%m%d-%H:%M:%S)
VERSION_PKG = github.com/seanmalloy/gsd/pkg/version
LDFLAGS := "-X ${VERSION_PKG}.Branch=${BRANCH} -X ${VERSION_PKG}.BuildDate=${DATE} \
	-X ${VERSION_PKG}.GitSHA1=${COMMIT}"

.PHONY: all
all: build

.PHONY: clean
clean:
	rm -rf $(BINARY)

.PHONY: build
build:
	go build -o $(BINARY) -ldflags $(LDFLAGS)

.PHONY: test
test: fmt lint vet test-unit

.PHONY: test-unit
test-unit:
	go test -race ./...

.PHONY: fmt
fmt:
	test -z "$(shell gofmt -l .)"

.PHONY: lint
lint:
	LINT_INPUT="$(shell go list ./...)"; golint -set_exit_status $$LINT_INPUT

.PHONY: vet
vet:
	VET_INPUT="$(shell go list ./...)"; go vet $$VET_INPUT