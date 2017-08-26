.PHONY: all
all: push

BUILDTAGS=

# Use the 0.0.0 tag for testing, it shouldn't clobber any release builds
APP=codegen
PROJECT=github.com/k8s-community/codegen
RELEASE?=0.0.0
GOOS?=darwin

CODEGEN_LEFT_DELIM?="{[("
CODEGEN_RIGHT_DELIM?=")]}"

REPO_INFO=$(shell git config --get remote.origin.url)

ifndef COMMIT
	COMMIT := git-$(shell git rev-parse --short HEAD)
endif

.PHONY: vendor
vendor: clean bootstrap
	glide install --strip-vendor

.PHONY: build
build: vendor
	CGO_ENABLED=0 GOOS=${GOOS} go build -a -installsuffix cgo \
		-ldflags "-s -w -X ${PROJECT}/version.RELEASE=${RELEASE} -X ${PROJECT}/version.COMMIT=${COMMIT} -X ${PROJECT}/version.REPO=${REPO_INFO}" \
		-o ./bin/${GOOS}/${APP} ${PROJECT}/cmd

.PHONY: run
run:
	@echo "+ $@"
	@env CODEGEN_LEFT_DELIM=${CODEGEN_LEFT_DELIM} \
		CODEGEN_RIGHT_DELIM=${CODEGEN_RIGHT_DELIM} \
		./bin/${GOOS}/${APP} 

GO_LIST_FILES=$(shell go list ${PROJECT}/... | grep -v vendor | grep -v templates)

.PHONY: fmt
fmt:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"gofmt -s -l {{.Dir}}"{{end}}' ${GO_LIST_FILES} | xargs -L 1 sh -c

.PHONY: lint
lint:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"golint {{.Dir}}/..."{{end}}' ${GO_LIST_FILES} | xargs -L 1 sh -c

.PHONY: vet
vet:
	@echo "+ $@"
	@go vet ${GO_LIST_FILES}

.PHONY: test
test: vendor fmt lint vet
	@echo "+ $@"
	@go test -v -race -tags "$(BUILDTAGS) cgo" ${GO_LIST_FILES}

q
.PHONY: clean
clean:
	rm -f ./bin/${GOOS}/${APP}

HAS_GLIDE := $(shell command -v glide;)
HAS_LINT := $(shell command -v golint;)

.PHONY: bootstrap
bootstrap:
ifndef HAS_GLIDE
	go get -u github.com/Masterminds/glide
endif
ifndef HAS_LINT
	go get -u github.com/golang/lint/golint
endif
