RELEASE_VER ?= latest
BUILD_DATE  := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
BASE_DIR    := $(shell git rev-parse --show-toplevel)
GIT_SHA     := $(shell git rev-parse --short HEAD)
BIN         := $(BASE_DIR)/bin

DOCKER_IMAGE_REPO?=portworx
DOCKER_IMAGE_NAME?=kdmp
DOCKER_IMAGE_TAG?=$(RELEASE_VER)

DOCKER_KDMP_UNITTEST_IMAGE?=px-kdmp-unittest
DOCKER_KDMP_TAG?=latest

DOCKER_IMAGE=$(DOCKER_IMAGE_REPO)/$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)
KDMP_UNITTEST_IMG=$(DOCKER_IMAGE_REPO)/$(DOCKER_KDMP_UNITTEST_IMAGE):$(DOCKER_KDMP_TAG)

RESTICEXECUTOR_DOCKER_IMAGE_NAME=resticexecutor
RESTICEXECUTOR_DOCKER_IMAGE=$(DOCKER_IMAGE_REPO)/$(RESTICEXECUTOR_DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

KOPIAEXECUTOR_DOCKER_IMAGE_NAME=kopiaexecutor
KOPIAEXECUTOR_DOCKER_IMAGE=$(DOCKER_IMAGE_REPO)/$(KOPIAEXECUTOR_DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)

NFSEXECUTOR_DOCKER_IMAGE_NAME=nfsexecutor
NFSEXECUTOR_DOCKER_IMAGE=$(DOCKER_IMAGE_REPO)/$(NFSEXECUTOR_DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)


export GO111MODULE=on
export GOFLAGS = -mod=vendor

MAJOR_VERSION := 1
MINOR_VERSION := 2
PATCH_VERSION := 1

ifndef PKGS
	PKGS := $(shell GOFLAGS=-mod=vendor go list ./... 2>&1 | grep -v 'go: ' | grep -v 'github.com/portworx/kdmp/vendor' | grep -v versioned | grep -v 'pkg/apis/v1')
endif

GO_FILES := $(shell find . -name '*.go' | grep -v 'vendor' | \
                                   grep -v '\.pb\.go' | \
                                   grep -v '\.pb\.gw\.go' | \
                                   grep -v 'externalversions' | \
                                   grep -v 'versioned' | \
                                   grep -v 'generated')

.DEFAULT_GOAL: all
.PHONY: test deploy build container

all: do-fmt pretest test-container test build container

test:
	docker run --rm -it -v ${GOPATH}:/go: $(KDMP_UNITTEST_IMG) make unittest

test-container:
	@echo "Building container: docker build --tag $(KDMP_UNITTEST_IMG) -f Dockerfile.unittest ."
	docker build --tag $(KDMP_UNITTEST_IMG) -f Dockerfile.unittest .

pretest: check-fmt lint vet errcheck staticcheck
build: update-deployment build-kdmp build-restic-executor kopia-executor build-pxc-exporter
container: container-kdmp container-restic-executor container-kopia-executor container-pxc-exporter
deploy: deploy-kdmp deploy-restic-executor deploy-kopia-executor deploy-pxc-exporter

### util targets ###
unittest:
	echo "mode: atomic" > coverage.txt
	for pkg in $(PKGS); do \
		go test -v -tags unittest -coverprofile=profile.out -covermode=atomic $(BUILD_OPTIONS) $${pkg} || exit 1; \
		if [ -f profile.out ]; then \
			cat profile.out | grep -v "mode: atomic">> coverage.txt; \
			rm profile.out; \
		fi; \
	done

lint:
	GO111MODULE=off go get -u golang.org/x/lint/golint
	for file in $(GO_FILES); do \
        golint $${file}; \
        if [ -n "$$(golint $${file})" ]; then \
            exit 1; \
        fi; \
        done

vet:
	GO111MODULE=off go vet $(PKGS)
	GO111MODULE=off go vet -tags unittest $(PKGS)
	#go vet -tags integrationtest github.com/portworx/kdmp/test/integration_test


staticcheck:
	GOFLAGS="" go install honnef.co/go/tools/cmd/staticcheck@v0.2.2
	staticcheck $(PKGS)
	#staticcheck -tags integrationtest test/integration_test/*.go
	staticcheck -tags unittest $(PKGS)


errcheck:
	GO111MODULE=off go get -u github.com/kisielk/errcheck
	errcheck -ignoregenerated -ignorepkg fmt -verbose -blank $(PKGS)
	errcheck -ignoregenerated -ignorepkg fmt -verbose -blank -tags unittest $(PKGS)
	#errcheck -ignoregenerated -verbose -blank -tags integrationtest github.com/portworx/kdmp/test/integration_test

check-fmt:
	bash -c "diff -u <(echo -n) <(gofmt -l -d -s -e $(GO_FILES))"

do-fmt:
	gofmt -s -w $(GO_FILES)

gocyclo:
	go get -u github.com/fzipp/gocyclo
	gocyclo -over 15 $(GO_FILES)

codegen:
	@echo "Generating CRD"
	hack/update-codegen.sh

gogenerate:
	go generate ./...

update-deployment:
	@echo "Updating deployment resources"
	hack/generate-operator-deployment.sh

vendor-sync:
	go mod tidy
	go mod vendor

### kdmp-operator targets ###
build-kdmp:
	@echo "Build kdmp"
	go build -o ${BIN}/kdmp -ldflags="-s -w \
	-X github.com/portworx/kdmp/pkg/version.gitVersion=${RELEASE_VER} \
	-X github.com/portworx/kdmp/pkg/version.gitCommit=${GIT_SHA} \
	-X github.com/portworx/kdmp/pkg/version.buildDate=${BUILD_DATE} \
	-X github.com/portworx/kdmp/pkg/version.major=${MAJOR_VERSION} \
	-X github.com/portworx/kdmp/pkg/version.minor=${MINOR_VERSION} \
	-X github.com/portworx/kdmp/pkg/version.patch=${PATCH_VERSION}" \
	-a $(BASE_DIR)/cmd/kdmp

container-kdmp:
	@echo "Build kdmp docker image"
	docker build --tag $(DOCKER_IMAGE) .

deploy-kdmp:
	@echo "Deploy kdmp docker image"
	docker push $(DOCKER_IMAGE)

kdmp: build-kdmp container-kdmp deploy-kdmp
kopia: kopia-executor deploy-kopia-executor
restic: restic-executor deploy-restic-executor
nfs: nfs-executor deploy-nfs-executor

restic-executor: build-restic-executor container-restic-executor

kopia-executor: build-kopia-executor container-kopia-executor

nfs-executor: build-nfs-executor container-nfs-executor

### restic-executor targets ###
build-restic-executor:
	@echo "Build restic-executor"
	go build -o $(BIN)/resticexecutor -ldflags="-s -w \
	-X github.com/portworx/kdmp/pkg/version.gitVersion=${RELEASE_VER} \
	-X github.com/portworx/kdmp/pkg/version.gitCommit=${GIT_SHA} \
	-X github.com/portworx/kdmp/pkg/version.buildDate=${BUILD_DATE}" \
	-a $(BASE_DIR)/cmd/executor/restic

build-kopia-executor:
	@echo "Build kopia-executor"
	go build -o $(BIN)/kopiaexecutor -ldflags="-s -w \
	-X github.com/portworx/kdmp/pkg/version.gitVersion=${RELEASE_VER} \
	-X github.com/portworx/kdmp/pkg/version.gitCommit=${GIT_SHA} \
	-X github.com/portworx/kdmp/pkg/version.buildDate=${BUILD_DATE}" \
	-a $(BASE_DIR)/cmd/executor/kopia

build-nfs-executor:
	@echo "Build nfs-executor"
	go build -o $(BIN)/nfsexecutor -ldflags="-s -w \
	-X github.com/portworx/kdmp/pkg/version.gitVersion=${RELEASE_VER} \
	-X github.com/portworx/kdmp/pkg/version.gitCommit=${GIT_SHA} \
	-X github.com/portworx/kdmp/pkg/version.buildDate=${BUILD_DATE}" \
	-a $(BASE_DIR)/cmd/executor/nfs

container-restic-executor:
	@echo "Build restice-executor docker image"
	docker build --tag $(RESTICEXECUTOR_DOCKER_IMAGE) -f Dockerfile.resticexecutor .

container-kopia-executor:
	@echo "Build kopia-executor docker image"
	docker build --tag $(KOPIAEXECUTOR_DOCKER_IMAGE) -f Dockerfile.kopia .

container-nfs-executor:
	@echo "Build nfs-executor docker image"
	docker build --tag $(NFSEXECUTOR_DOCKER_IMAGE) -f Dockerfile.nfs .

deploy-restic-executor:
	@echo "Deploy kdmp docker image"
	docker push $(RESTICEXECUTOR_DOCKER_IMAGE)

deploy-kopia-executor:
	@echo "Deploy kopia docker image"
	docker push $(KOPIAEXECUTOR_DOCKER_IMAGE)

deploy-nfs-executor:
	@echo "Deploy nfs docker image"
	docker push $(NFSEXECUTOR_DOCKER_IMAGE)

### pxc-exporter targets ###
build-pxc-exporter: gogenerate
	@echo "Build kdmp"
	go build -o ${BIN}/pxc-exporter -ldflags="-s -w \
	-X github.com/portworx/kdmp/pkg/version.gitVersion=${RELEASE_VER} \
	-X github.com/portworx/kdmp/pkg/version.gitCommit=${GIT_SHA} \
	-X github.com/portworx/kdmp/pkg/version.buildDate=${BUILD_DATE}" \
	-a $(BASE_DIR)/cmd/exporter


container-pxc-exporter:
deploy-pxc-exporter:
